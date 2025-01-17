package boar

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/itchio/dash"

	"github.com/itchio/boar/szextractor"
	"github.com/itchio/boar/szextractor/xzsource"
	"github.com/itchio/savior/bzip2source"
	"github.com/itchio/savior/gzipsource"
	"github.com/itchio/savior/seeksource"

	"github.com/itchio/savior/tarextractor"
	"github.com/itchio/savior/zipextractor"
	"github.com/pkg/errors"

	"github.com/itchio/headway/state"
	"github.com/itchio/httpkit/eos"
	"github.com/itchio/savior"
)

type Strategy int

const (
	StrategyNone Strategy = 0

	StrategyZip Strategy = 100
	// linux binaries for example - they might be MojoSetup installers
	// (.zip files), or they might not be.
	StrategyZipUnsure Strategy = 101

	StrategyTar    Strategy = 200
	StrategyTarGz  Strategy = 201
	StrategyTarBz2 Strategy = 202
	StrategyTarXz  Strategy = 203

	StrategySevenZip Strategy = 300
	// .exe files for example - might be self-extracting
	// archives 7-zip can handle, or they might not.
	StrategySevenZipUnsure Strategy = 301

	// .dmg files can only be properly extracted on macOS.
	// 7-zip struggles with ISO9660 disk images for example,
	// and doesn't support APFS yet (as of 18.05)
	StrategyDmg Strategy = 400

	// .rar files we do *not* want to open while probing
	StrategyRar Strategy = 500
)

func (as Strategy) String() string {
	switch as {
	case StrategyZip, StrategyZipUnsure:
		return "zip"
	case StrategyTar:
		return "tar"
	case StrategyTarGz:
		return "tar.gz"
	case StrategyTarBz2:
		return "tar.bz2"
	case StrategyTarXz:
		return "tar.xz"
	case StrategySevenZip, StrategySevenZipUnsure:
		return "7-zip"
	default:
		return "<no strategy>"
	}
}

type EntriesLister interface {
	Entries() []*savior.Entry
}

type Info struct {
	Strategy    Strategy
	Features    savior.ExtractorFeatures
	Format      string
	PostExtract []string
}

func (ai *Info) String() string {
	res := ""
	res += fmt.Sprintf("%s (via %s)", ai.Format, ai.Strategy)
	res += fmt.Sprintf(", %s", ai.Features)
	return res
}

// Probe attempts to determine the type of an archive
// Returns (nil, nil) if it is not a recognized archive type
// Returns (nil, non-nil) if it IS a recognized archive type, but something's
// wrong with it.
// Returns (non-nil, nil) if it is a recognized archive type and we
// are confident we can extract it correctly.
func Probe(params ProbeParams) (*Info, error) {
	var strategy Strategy
	file := params.File
	consumer := params.Consumer
	ext := getExt(file, consumer)

	if params.Candidate != nil && params.Candidate.Flavor == dash.FlavorNativeLinux {
		// might be a mojosetup installer - if not, we won't know what to do with it
		strategy = StrategyZipUnsure
	} else {
		strategy = getStrategy(ext)
	}

	if strategy == StrategyNone {
		return nil, nil
	}

	info := &Info{
		Strategy: strategy,
	}

	{
		checkEarlyExit := true

		switch info.Strategy {
		case StrategySevenZip:
			info.Features = szextractor.FeaturesByExtension(ext)
		default:
			checkEarlyExit = false
		}

		if checkEarlyExit && !info.Features.RandomAccess {
			// no random access means the format isn't "htfs-friendly",
			// ie. enumerating entries might end up downloading the whole
			// file. since we're just probing right now, let's just return
			// what little info we have.
			return info, nil
		}
	}

	// now actually try to open it
	ex, err := info.GetExtractor(file, consumer)
	if err != nil {
		switch strategy {
		case StrategySevenZipUnsure:
			// we didn't know that one until we try, so it's just
			// not a recognized archive format
			consumer.Warnf("Tried opening archive with 7-zip but we got: %v", err)
			consumer.Warnf("Ignoring...")
			return nil, nil
		case StrategyZipUnsure:
			// we didn't know that one until we try, so it's just
			// not a recognized archive format
			consumer.Warnf("Tried opening as a zip but we got: %v", err)
			consumer.Warnf("Ignoring...")
			return nil, nil
		default:
			return nil, errors.Wrap(err, "opening archive")
		}
	}

	if szex, ok := ex.(szextractor.SzExtractor); ok {
		// this codepath runs when we did not have a .zip, .tar.gz etc. file
		// extension, but 7-zip detected such a file anyway. in this case, we prefer
		// "native" decompressors, implemented in pure Go.
		info.Format = szex.GetFormat()
		useNativeDecompressor := true
		switch info.Format {
		case "gzip":
			info.Strategy = StrategyTarGz
		case "bzip2":
			info.Strategy = StrategyTarBz2
		case "xz":
			info.Strategy = StrategyTarXz
		case "tar":
			info.Strategy = StrategyTar
		case "zip":
			info.Strategy = StrategyZip
		default:
			useNativeDecompressor = false
		}

		if useNativeDecompressor {
			ex, err = info.GetExtractor(file, consumer)
			if err != nil {
				return nil, errors.Wrap(err, "getting extractor for file")
			}

			info.Format = info.Strategy.String()
		}
	} else {
		info.Format = info.Strategy.String()
	}
	info.Features = ex.Features()

	var entries []*savior.Entry
	// only try listing entries if we have random access support,
	// otherwise, we risk downloading the whole archive just to list entries.
	if info.Features.RandomAccess {
		if el, ok := ex.(EntriesLister); ok {
			entries = el.Entries()
			if params.OnEntries != nil {
				params.OnEntries(entries)
			}
		}
	}

	return info, nil
}

func getExt(file eos.File, consumer *state.Consumer) string {
	stats, err := file.Stat()
	if err != nil {
		consumer.Warnf("archive: Could not stat file, going with blank extension")
		return ""
	}

	lowerName := strings.ToLower(stats.Name())
	ext := filepath.Ext(lowerName)
	if strings.HasSuffix(lowerName, ".tar"+ext) {
		ext = ".tar" + ext
	}
	return ext
}

func getStrategy(ext string) Strategy {
	switch ext {
	case ".zip":
		return StrategyZip
	case ".tar":
		return StrategyTar
	case ".tar.gz":
		return StrategyTarGz
	case ".tar.bz2":
		return StrategyTarBz2
	case ".tar.xz":
		return StrategyTarXz
	case ".7z":
		return StrategySevenZip
	case ".exe":
		return StrategySevenZipUnsure
	}

	return StrategySevenZipUnsure
}

func (ai *Info) GetExtractor(file eos.File, consumer *state.Consumer) (savior.Extractor, error) {
	switch ai.Strategy {
	case StrategyZip, StrategyZipUnsure:
		stats, err := file.Stat()
		if err != nil {
			return nil, errors.Wrap(err, "stat'ing file to open as zip archive")
		}

		ex, err := zipextractor.New(file, stats.Size())
		if err != nil {
			return nil, errors.Wrap(err, "creating zip extractor")
		}
		return ex, nil
	case StrategyTar:
		return tarextractor.New(seeksource.FromFile(file)), nil
	case StrategyTarGz:
		return tarextractor.New(gzipsource.New(seeksource.FromFile(file))), nil
	case StrategyTarBz2:
		return tarextractor.New(bzip2source.New(seeksource.FromFile(file))), nil
	case StrategyTarXz:
		xs, err := xzsource.New(file, consumer)
		if err != nil {
			return nil, errors.Wrap(err, "creating xz extractor")
		}
		return tarextractor.New(xs), nil
	case StrategySevenZip, StrategySevenZipUnsure:
		szex, err := szextractor.New(file, consumer)
		if err != nil {
			return nil, errors.Wrap(err, "creating 7-zip extractor")
		}

		// apply blacklist
		switch szex.GetFormat() {
		// cf. https://github.com/itchio/itch/issues/1700
		case "elf":
			// won't extract ELF executables
			return nil, errors.New("refusing to extract ELF file")
		case "pe":
			// won't extract PE executables
			return nil, errors.New("refusing to extract PE file")
		default:
			return szex, nil
		}
	}

	return nil, fmt.Errorf("unknown Strategy %d", ai.Strategy)
}
