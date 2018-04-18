// Code generated by go generate; DO NOT EDIT.
// Generated at 2018-02-13 17:51:30.579035757 +0100 CET m=+17.606651348
// For version v1.5.0, base URL https://dl.itch.ovh/libc7zip
package formulas

import "github.com/itchio/boar/szextractor/types"

var ByOsArch = types.DepSpecMap{"linux-386": types.DepSpec{Entries: []types.DepEntry{types.DepEntry{Name: "7z.so", Size: 2291392, Hashes: []types.DepHash{types.DepHash{Algo: "sha1", Value: "5a3313d9d69ce3fe648022a7aa588b0c5195de65"}, types.DepHash{Algo: "sha256", Value: "1cc4e611b4cc91de144fe12e9fbc5242f775d37574c6eda261e3d6fd888c9c2b"}}}, types.DepEntry{Name: "libc7zip.so", Size: 173832, Hashes: []types.DepHash{types.DepHash{Algo: "sha256", Value: "67e0f16f9068a8c3a6a5b359d22b3cf820e9c42c103b73a986e12ae54c49f5a3"}, types.DepHash{Algo: "sha1", Value: "b8582cec8504b0d4a10bc90a6d4ca19416406881"}}}}, Sources: []string{"https://dl.itch.ovh/libc7zip/linux-386/v1.5.0/libc7zip.zip"}}, "darwin-amd64": types.DepSpec{Entries: []types.DepEntry{types.DepEntry{Name: "7z.so", Size: 2276972, Hashes: []types.DepHash{types.DepHash{Algo: "sha1", Value: "c8ab6bea9e01218f3a47d8786ef25e5e76e5eedb"}, types.DepHash{Algo: "sha256", Value: "a2fb72a26b28257b6c738b0ef27b8f2044ed4fb23a4125386f659e19b8e3828b"}}}, types.DepEntry{Name: "libc7zip.dylib", Size: 160008, Hashes: []types.DepHash{types.DepHash{Algo: "sha256", Value: "441cf2026e6223d52d71af631cb7fbb76ed6c386d36670fbff5c48014e9d76ea"}, types.DepHash{Algo: "sha1", Value: "e3550c11b590328079d61cdebdc9efab22dfef58"}}}}, Sources: []string{"https://dl.itch.ovh/libc7zip/darwin-amd64/v1.5.0/libc7zip.zip"}}, "windows-386": types.DepSpec{Entries: []types.DepEntry{types.DepEntry{Name: "7z.dll", Size: 1079408, Hashes: []types.DepHash{types.DepHash{Algo: "sha1", Value: "16bf27695d927e6142c7d05dc78980cb327f36b8"}, types.DepHash{Algo: "sha256", Value: "71e68d9913afcd008a9c76d6efdce6aaf1b8a3d16d2652df484b4bbe1da76fbb"}}}, types.DepEntry{Name: "c7zip.dll", Size: 397928, Hashes: []types.DepHash{types.DepHash{Algo: "sha1", Value: "15443cd2724e7d3af843bb9cf2d5624d570fe847"}, types.DepHash{Algo: "sha256", Value: "85ae5b6c8f8635044f3e5168b52359938b3a8c179c80bc3640f34f3f4ef75c7b"}}}}, Sources: []string{"https://dl.itch.ovh/libc7zip/windows-386/v1.5.0/libc7zip.zip"}}, "windows-amd64": types.DepSpec{Entries: []types.DepEntry{types.DepEntry{Name: "7z.dll", Size: 1614952, Hashes: []types.DepHash{types.DepHash{Algo: "sha1", Value: "1ee14b9f4aac98925911a2aaf8aae7de1f787592"}, types.DepHash{Algo: "sha256", Value: "5dd34626363aa9d30e7533f3df0fbb00209c63718ee388d8234822621cbd95e1"}}}, types.DepEntry{Name: "c7zip.dll", Size: 539760, Hashes: []types.DepHash{types.DepHash{Algo: "sha1", Value: "05fb56690f8092a486a9e5a1ec0b152e8ed6676f"}, types.DepHash{Algo: "sha256", Value: "c0c3cce72cff3f6b68c8f52337083d2bb4f0ad5de7554ba2763f201346632536"}}}}, Sources: []string{"https://dl.itch.ovh/libc7zip/windows-amd64/v1.5.0/libc7zip.zip"}}, "linux-amd64": types.DepSpec{Entries: []types.DepEntry{types.DepEntry{Name: "7z.so", Size: 2273248, Hashes: []types.DepHash{types.DepHash{Algo: "sha1", Value: "ff47097a16402a6144e868daff4416e6ab8c2584"}, types.DepHash{Algo: "sha256", Value: "23f949fdbcb3cb0722539d728824138c8b5048577d9de350ceb17da0071b3a99"}}}, types.DepEntry{Name: "libc7zip.so", Size: 193152, Hashes: []types.DepHash{types.DepHash{Algo: "sha256", Value: "fc5e3d95adb21efe666b6b5cc0ff46594302e5efe7704b386a117790a6979b6e"}, types.DepHash{Algo: "sha1", Value: "f93a8dda18a527273e7f828fc7ef69fea18cb67d"}}}}, Sources: []string{"https://dl.itch.ovh/libc7zip/linux-amd64/v1.5.0/libc7zip.zip"}}}