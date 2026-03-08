# Changelog

## [1.7.0](https://github.com/neverprepared/az-pim-cli/compare/v1.6.1...v1.7.0) (2026-03-08)


### Features

* --wait/--timeout polling, --all for governance, shell completion docs ([ff0f45c](https://github.com/neverprepared/az-pim-cli/commit/ff0f45c93d36679d7f8170fe8492b78e5b30707d))
* add list active, deactivate, --json output, and token commands ([d2a7d0e](https://github.com/neverprepared/az-pim-cli/commit/d2a7d0e3d8dfa7640ce2bfed20609f4f8e9b3dc0))
* add MCP server ([af8d18d](https://github.com/neverprepared/az-pim-cli/commit/af8d18d1d64160fb6ffaff42604caa28073bc352))
* add support for additional Azure environments (us gov, china) ([#111](https://github.com/neverprepared/az-pim-cli/issues/111)) ([07a83af](https://github.com/neverprepared/az-pim-cli/commit/07a83af56bcaf109045ca2dbc7442d3f1cc2c2d9))
* bulk activate --all, sorted output, README improvements ([979640d](https://github.com/neverprepared/az-pim-cli/commit/979640d2895f22eb09fff9cba4ab8a58a0514644))
* **goreleaser:** include binaries in release ([#101](https://github.com/neverprepared/az-pim-cli/issues/101)) ([aeabca2](https://github.com/neverprepared/az-pim-cli/commit/aeabca28b2777b36a291976f51724299172b30a9))
* table output, ANSI color, HTTP retry, multiple --name targets ([0d3bd07](https://github.com/neverprepared/az-pim-cli/commit/0d3bd078ab029237552d18fb5e9f541a07583d25))


### Bug Fixes

* **build/goreleaser:** remove incorrect template key usage ([#103](https://github.com/neverprepared/az-pim-cli/issues/103)) ([a1427a3](https://github.com/neverprepared/az-pim-cli/commit/a1427a304866813939279be5c5ed480932be3d6b))
* prevent panic from JWT parsing ([0ebbaa1](https://github.com/neverprepared/az-pim-cli/commit/0ebbaa1baff1dad212d45d32143a71678f107440))
* prevent panic from JWT parsing ([#96](https://github.com/neverprepared/az-pim-cli/issues/96)) ([0ebbaa1](https://github.com/neverprepared/az-pim-cli/commit/0ebbaa1baff1dad212d45d32143a71678f107440))


### Chores

* **deps:** upgrade dependencies ([#99](https://github.com/neverprepared/az-pim-cli/issues/99)) ([202b948](https://github.com/neverprepared/az-pim-cli/commit/202b9485e8003023e5bfa9c0b0f078b4b52660ae))
* **main:** release 1.10.0 ([#108](https://github.com/neverprepared/az-pim-cli/issues/108)) ([a82127c](https://github.com/neverprepared/az-pim-cli/commit/a82127cd9435821a9c74845a5702ab346a96aca7))
* **main:** release 1.11.0 ([#112](https://github.com/neverprepared/az-pim-cli/issues/112)) ([cdb9492](https://github.com/neverprepared/az-pim-cli/commit/cdb949255acfc2b3cf1a4e2a07c16087d71b5c74))
* **main:** release 1.12.0 ([d2ccf9d](https://github.com/neverprepared/az-pim-cli/commit/d2ccf9de31abb7a873e57f3155be90c4fe71015c))
* **main:** release 1.12.0 ([1dc54b6](https://github.com/neverprepared/az-pim-cli/commit/1dc54b6fa72ff01f9d4214391eada6ca0fa748fc))
* **main:** release 1.13.0 ([34f3034](https://github.com/neverprepared/az-pim-cli/commit/34f303470c261c5e26dacc8212f2140508bd3fd0))
* **main:** release 1.13.0 ([ab807b1](https://github.com/neverprepared/az-pim-cli/commit/ab807b17c9e3e37b5e4a2916873d10513a94cc60))
* **main:** release 1.6.1 ([539d7d3](https://github.com/neverprepared/az-pim-cli/commit/539d7d3eee44f22446a2a36e3f1c901533ec37d0))
* **main:** release 1.6.1 ([eefffb0](https://github.com/neverprepared/az-pim-cli/commit/eefffb073d66039594b12fd5135e21878e28d746))
* **main:** release 1.7.0 ([#92](https://github.com/neverprepared/az-pim-cli/issues/92)) ([9a46319](https://github.com/neverprepared/az-pim-cli/commit/9a4631930c447ffa2db64bce7c6e0c363b4f8c75))
* **main:** release 1.8.0 ([#100](https://github.com/neverprepared/az-pim-cli/issues/100)) ([a207941](https://github.com/neverprepared/az-pim-cli/commit/a2079415ced2e294dbb341c12bfa2f11f0b8df59))
* **main:** release 1.9.0 ([#106](https://github.com/neverprepared/az-pim-cli/issues/106)) ([4d8a7b0](https://github.com/neverprepared/az-pim-cli/commit/4d8a7b0666dcec079d40e4b3d31e239d56deeb86))
* migrate to neverprepared/az-pim-cli ([60df4cc](https://github.com/neverprepared/az-pim-cli/commit/60df4ccc38440f9d7881b258f357ee8856b8d4d6))
* replace dummy JWT with minimal fake token ([ad5bf0a](https://github.com/neverprepared/az-pim-cli/commit/ad5bf0a068166b15b52bf05be3562f08e98b5703))
* upgrade dependencies ([202b948](https://github.com/neverprepared/az-pim-cli/commit/202b9485e8003023e5bfa9c0b0f078b4b52660ae))


### Tests

* add coverage for new features ([070f5b8](https://github.com/neverprepared/az-pim-cli/commit/070f5b862e9ab0293ae5fb8b189248d4abeececa))


### Build

* fix broken goreleaser builds ([#105](https://github.com/neverprepared/az-pim-cli/issues/105)) ([506acde](https://github.com/neverprepared/az-pim-cli/commit/506acde2183d3a649e9d0d49d578f34e10bed112))
* **goreleaser:** include macos-universal binary in release archives ([#107](https://github.com/neverprepared/az-pim-cli/issues/107)) ([fb45385](https://github.com/neverprepared/az-pim-cli/commit/fb453852f51c7840c051e603b216886715fa7476))


### Continuous Integration

* add build workflow for all target platforms ([f991bd1](https://github.com/neverprepared/az-pim-cli/commit/f991bd1476cfa983c6f85e7261af40a103605a9f))
* disable Snyk jobs until SNYK_TOKEN is configured ([6ce411e](https://github.com/neverprepared/az-pim-cli/commit/6ce411e7a319236c99208f2db78154f2ba45480e))
* pin semgrep container and snyk action to specific versions ([d659971](https://github.com/neverprepared/az-pim-cli/commit/d659971d888bde381a7430172522a9e6438de45b))
* run GoReleaser after release-please creates a release ([f7cdfb6](https://github.com/neverprepared/az-pim-cli/commit/f7cdfb6a13a2e1fe102a3e671698d11e4c0f9ab8))
* set permissions for workflows ([#91](https://github.com/neverprepared/az-pim-cli/issues/91)) ([2f10183](https://github.com/neverprepared/az-pim-cli/commit/2f1018335cf34d0172d332bb8da63daba157784b))
* use GITHUB_TOKEN for release-please instead of PAT ([7bbc4a8](https://github.com/neverprepared/az-pim-cli/commit/7bbc4a8caef116cbe3a65a4d761ac58148a40963))

## [1.6.1](https://github.com/neverprepared/az-pim-cli/compare/v1.13.0...v1.6.1) (2026-03-08)


### ⚠ BREAKING CHANGES

* use proper terms for 'azure resources' type ([#59](https://github.com/neverprepared/az-pim-cli/issues/59))

### Features

* --wait/--timeout polling, --all for governance, shell completion docs ([ff0f45c](https://github.com/neverprepared/az-pim-cli/commit/ff0f45c93d36679d7f8170fe8492b78e5b30707d))
* activate roles ([90d7eca](https://github.com/neverprepared/az-pim-cli/commit/90d7ecaae6349b99568ae476b72a994ddc31b5c3))
* add 'version' command ([#30](https://github.com/neverprepared/az-pim-cli/issues/30)) ([7a74290](https://github.com/neverprepared/az-pim-cli/commit/7a74290953295fe0b49f36197d2ba90f0f4897e5))
* add list active, deactivate, --json output, and token commands ([d2a7d0e](https://github.com/neverprepared/az-pim-cli/commit/d2a7d0e3d8dfa7640ce2bfed20609f4f8e9b3dc0))
* add MCP server ([af8d18d](https://github.com/neverprepared/az-pim-cli/commit/af8d18d1d64160fb6ffaff42604caa28073bc352))
* add reason to activate command ([#4](https://github.com/neverprepared/az-pim-cli/issues/4)) ([842c1af](https://github.com/neverprepared/az-pim-cli/commit/842c1af530956039312d05a8adf429db4710ce5d))
* add support for additional Azure environments (us gov, china) ([#111](https://github.com/neverprepared/az-pim-cli/issues/111)) ([07a83af](https://github.com/neverprepared/az-pim-cli/commit/07a83af56bcaf109045ca2dbc7442d3f1cc2c2d9))
* add support for setting start-time ([#81](https://github.com/neverprepared/az-pim-cli/issues/81)) ([fb57862](https://github.com/neverprepared/az-pim-cli/commit/fb57862d164e7833b2274da9705f2b71e5e36fd8))
* add support for specifying 'ticket number' and 'ticket system' ([#56](https://github.com/neverprepared/az-pim-cli/issues/56)) ([955a3dd](https://github.com/neverprepared/az-pim-cli/commit/955a3dd2c7eaab249bad0f6eb0f87c6edd36a5d1))
* bulk activate --all, sorted output, README improvements ([979640d](https://github.com/neverprepared/az-pim-cli/commit/979640d2895f22eb09fff9cba4ab8a58a0514644))
* check for various request status types ([#14](https://github.com/neverprepared/az-pim-cli/issues/14)) ([c2b2167](https://github.com/neverprepared/az-pim-cli/commit/c2b2167178d3300b6d73e1033840385214b065a5))
* dry-run for 'activate' ([#22](https://github.com/neverprepared/az-pim-cli/issues/22)) ([b462889](https://github.com/neverprepared/az-pim-cli/commit/b462889bf0d488150ef80b5bdf048a820f45438b))
* **goreleaser:** include binaries in release ([#101](https://github.com/neverprepared/az-pim-cli/issues/101)) ([aeabca2](https://github.com/neverprepared/az-pim-cli/commit/aeabca28b2777b36a291976f51724299172b30a9))
* improved error messages and logging ([#68](https://github.com/neverprepared/az-pim-cli/issues/68)) ([40e3cc3](https://github.com/neverprepared/az-pim-cli/commit/40e3cc3ec45f25f13df2e0e501025640166e0d0a))
* list eligible roles ([32a97d9](https://github.com/neverprepared/az-pim-cli/commit/32a97d9106bb3d8b9b9fb6cac89cb8279cf97eb2))
* Support for Entra roles ([#61](https://github.com/neverprepared/az-pim-cli/issues/61)) ([9568654](https://github.com/neverprepared/az-pim-cli/commit/9568654dbbe37442b8207dd4a3b3176de2fe4b48))
* support for PIM Entra groups ([#16](https://github.com/neverprepared/az-pim-cli/issues/16)) ([f01d98a](https://github.com/neverprepared/az-pim-cli/commit/f01d98a474ca4a166be5c9eeaa6480dc35c56494))
* support new Azure Entra ID PIM API ([#6](https://github.com/neverprepared/az-pim-cli/issues/6)) ([f2902f4](https://github.com/neverprepared/az-pim-cli/commit/f2902f4d3ca6d78c680adacaf9f316502a61814c))
* table output, ANSI color, HTTP retry, multiple --name targets ([0d3bd07](https://github.com/neverprepared/az-pim-cli/commit/0d3bd078ab029237552d18fb5e9f541a07583d25))
* use az-cli for auth ([5502b46](https://github.com/neverprepared/az-pim-cli/commit/5502b464de8abd1dbf0fa89a4193713bc98275b9))
* use proper terms for 'azure resources' type ([#59](https://github.com/neverprepared/az-pim-cli/issues/59)) ([0c17c92](https://github.com/neverprepared/az-pim-cli/commit/0c17c923221c8def5b36cddc0cc1e705afcf6df0))


### Bug Fixes

* **activate:** Role selection on `activate` selects incorrect role ([#8](https://github.com/neverprepared/az-pim-cli/issues/8)) ([c64fad2](https://github.com/neverprepared/az-pim-cli/commit/c64fad2f144955df20fbc7f567cdbd768df6a0ef))
* **build/goreleaser:** remove incorrect template key usage ([#103](https://github.com/neverprepared/az-pim-cli/issues/103)) ([a1427a3](https://github.com/neverprepared/az-pim-cli/commit/a1427a304866813939279be5c5ed480932be3d6b))
* fix casing role on activate ([#3](https://github.com/neverprepared/az-pim-cli/issues/3)) ([ac69cbc](https://github.com/neverprepared/az-pim-cli/commit/ac69cbc11bd9d25a6530758070d5bd8ef02bb4c1))
* **pim-client:** resolve invalid logic for building a request ([#76](https://github.com/neverprepared/az-pim-cli/issues/76)) ([287679b](https://github.com/neverprepared/az-pim-cli/commit/287679b4cf540ccabc550cb310b44bc4df795fb8))
* **pim-client:** resolve invalid logic for building a request dynamically ([287679b](https://github.com/neverprepared/az-pim-cli/commit/287679b4cf540ccabc550cb310b44bc4df795fb8))
* prevent panic from JWT parsing ([0ebbaa1](https://github.com/neverprepared/az-pim-cli/commit/0ebbaa1baff1dad212d45d32143a71678f107440))
* prevent panic from JWT parsing ([#96](https://github.com/neverprepared/az-pim-cli/issues/96)) ([0ebbaa1](https://github.com/neverprepared/az-pim-cli/commit/0ebbaa1baff1dad212d45d32143a71678f107440))
* use exact matching for the role selection ([#12](https://github.com/neverprepared/az-pim-cli/issues/12)) ([9d9cfd5](https://github.com/neverprepared/az-pim-cli/commit/9d9cfd5146f5a4e255d43a21eb88aca21a1dfd1f))


### Documentation

* **github:** add project guidelines ([#31](https://github.com/neverprepared/az-pim-cli/issues/31)) ([d63b3c2](https://github.com/neverprepared/az-pim-cli/commit/d63b3c24f32d3d871ece525097a9b7024d367f79))
* initial docs ([ffa6c25](https://github.com/neverprepared/az-pim-cli/commit/ffa6c250cd91a9d29c1a79c47b41d89eddd44e44))


### Code Refactoring

* create interface for azure client ([#72](https://github.com/neverprepared/az-pim-cli/issues/72)) ([00d1d70](https://github.com/neverprepared/az-pim-cli/commit/00d1d7007547c35dc552e877e5af11a0cea6559e))


### Chores

* add FUNDING file ([#74](https://github.com/neverprepared/az-pim-cli/issues/74)) ([3bb73bc](https://github.com/neverprepared/az-pim-cli/commit/3bb73bcffa2a0f08e0487902c926515e359c5b04))
* added LICENSE ([#25](https://github.com/neverprepared/az-pim-cli/issues/25)) ([68dd63b](https://github.com/neverprepared/az-pim-cli/commit/68dd63ba044f171466110b4e6eeae0cced61d222))
* bootstrap releases for path: . ([#34](https://github.com/neverprepared/az-pim-cli/issues/34)) ([1f156a4](https://github.com/neverprepared/az-pim-cli/commit/1f156a48df289ef3d7404c651a8b3b8eaeb716d5))
* create release v1.6.1 ([#89](https://github.com/neverprepared/az-pim-cli/issues/89)) ([837ea97](https://github.com/neverprepared/az-pim-cli/commit/837ea97a44b9eab04a71ee32dba773e7140ddbc1))
* **deps:** bump golang.org/x/crypto from 0.24.0 to 0.31.0 in the go_modules group ([#79](https://github.com/neverprepared/az-pim-cli/issues/79)) ([0d2d9b4](https://github.com/neverprepared/az-pim-cli/commit/0d2d9b496a8b5daf0221ee456331116ab2a94351))
* **deps:** bump golang.org/x/crypto in the go_modules group ([0d2d9b4](https://github.com/neverprepared/az-pim-cli/commit/0d2d9b496a8b5daf0221ee456331116ab2a94351))
* **deps:** bump golang.org/x/net from 0.22.0 to 0.23.0 ([#11](https://github.com/neverprepared/az-pim-cli/issues/11)) ([3049152](https://github.com/neverprepared/az-pim-cli/commit/30491525d1b82a965762e9837e751bb4128fa8c1))
* **deps:** bump golang.org/x/net from 0.26.0 to 0.33.0 in the go_modules group ([#84](https://github.com/neverprepared/az-pim-cli/issues/84)) ([6c2191d](https://github.com/neverprepared/az-pim-cli/commit/6c2191df22a3d646a5a0bf12ee3351e74fc4a40c))
* **deps:** bump golang.org/x/net in the go_modules group ([6c2191d](https://github.com/neverprepared/az-pim-cli/commit/6c2191df22a3d646a5a0bf12ee3351e74fc4a40c))
* **deps:** upgrade dependencies ([#99](https://github.com/neverprepared/az-pim-cli/issues/99)) ([202b948](https://github.com/neverprepared/az-pim-cli/commit/202b9485e8003023e5bfa9c0b0f078b4b52660ae))
* **docs:** cleanup ([e737a78](https://github.com/neverprepared/az-pim-cli/commit/e737a782b9b29d9ed91d5fdf44d141516f27da07))
* ensure consistent flag names ([#18](https://github.com/neverprepared/az-pim-cli/issues/18)) ([e2ae673](https://github.com/neverprepared/az-pim-cli/commit/e2ae673d83f15a19225c90ac88f4ff64b2599653))
* ignores ([0bdb13d](https://github.com/neverprepared/az-pim-cli/commit/0bdb13dabf796abd611f55fa9c1b147407ab0ea5))
* **main:** release 1.0.1 ([#49](https://github.com/neverprepared/az-pim-cli/issues/49)) ([51cdbd0](https://github.com/neverprepared/az-pim-cli/commit/51cdbd08a1030a448edda79be7147bcb176d3f2e))
* **main:** release 1.1.0 ([#57](https://github.com/neverprepared/az-pim-cli/issues/57)) ([9fb01cd](https://github.com/neverprepared/az-pim-cli/commit/9fb01cd1d6f588d4fee3f3aca4ddb848ee46a332))
* **main:** release 1.10.0 ([#108](https://github.com/neverprepared/az-pim-cli/issues/108)) ([a82127c](https://github.com/neverprepared/az-pim-cli/commit/a82127cd9435821a9c74845a5702ab346a96aca7))
* **main:** release 1.11.0 ([#112](https://github.com/neverprepared/az-pim-cli/issues/112)) ([cdb9492](https://github.com/neverprepared/az-pim-cli/commit/cdb949255acfc2b3cf1a4e2a07c16087d71b5c74))
* **main:** release 1.12.0 ([d2ccf9d](https://github.com/neverprepared/az-pim-cli/commit/d2ccf9de31abb7a873e57f3155be90c4fe71015c))
* **main:** release 1.12.0 ([1dc54b6](https://github.com/neverprepared/az-pim-cli/commit/1dc54b6fa72ff01f9d4214391eada6ca0fa748fc))
* **main:** release 1.13.0 ([34f3034](https://github.com/neverprepared/az-pim-cli/commit/34f303470c261c5e26dacc8212f2140508bd3fd0))
* **main:** release 1.13.0 ([ab807b1](https://github.com/neverprepared/az-pim-cli/commit/ab807b17c9e3e37b5e4a2916873d10513a94cc60))
* **main:** release 1.2.0 ([#63](https://github.com/neverprepared/az-pim-cli/issues/63)) ([89707d9](https://github.com/neverprepared/az-pim-cli/commit/89707d926ef6dea86dbc5dc2a9faf42e227b6a8c))
* **main:** release 1.3.0 ([#65](https://github.com/neverprepared/az-pim-cli/issues/65)) ([a4f8461](https://github.com/neverprepared/az-pim-cli/commit/a4f8461e1f90695c6d733666d99afb8e212e6470))
* **main:** release 1.4.0 ([#70](https://github.com/neverprepared/az-pim-cli/issues/70)) ([38817b6](https://github.com/neverprepared/az-pim-cli/commit/38817b6d8fbb0da5306beb8ff1ca8bf13cdec197))
* **main:** release 1.5.0 ([#77](https://github.com/neverprepared/az-pim-cli/issues/77)) ([eda5f4b](https://github.com/neverprepared/az-pim-cli/commit/eda5f4bda5f6458d57ed31e550212d9562977309))
* **main:** release 1.6.0 ([#85](https://github.com/neverprepared/az-pim-cli/issues/85)) ([e23a5d4](https://github.com/neverprepared/az-pim-cli/commit/e23a5d46a7ee68edb7e5848fe3328a16a63fe1f5))
* **main:** release 1.6.1 ([#90](https://github.com/neverprepared/az-pim-cli/issues/90)) ([38dd958](https://github.com/neverprepared/az-pim-cli/commit/38dd958d14e5fada70ce8ac3cea385378aa60111))
* **main:** release 1.7.0 ([#92](https://github.com/neverprepared/az-pim-cli/issues/92)) ([9a46319](https://github.com/neverprepared/az-pim-cli/commit/9a4631930c447ffa2db64bce7c6e0c363b4f8c75))
* **main:** release 1.8.0 ([#100](https://github.com/neverprepared/az-pim-cli/issues/100)) ([a207941](https://github.com/neverprepared/az-pim-cli/commit/a2079415ced2e294dbb341c12bfa2f11f0b8df59))
* **main:** release 1.9.0 ([#106](https://github.com/neverprepared/az-pim-cli/issues/106)) ([4d8a7b0](https://github.com/neverprepared/az-pim-cli/commit/4d8a7b0666dcec079d40e4b3d31e239d56deeb86))
* migrate to neverprepared/az-pim-cli ([60df4cc](https://github.com/neverprepared/az-pim-cli/commit/60df4ccc38440f9d7881b258f357ee8856b8d4d6))
* replace deprecated ioutil package (SA1019) ([765f753](https://github.com/neverprepared/az-pim-cli/commit/765f753c573221505ac9cad7fbf2efa6f5dcf5f3))
* replace dummy JWT with minimal fake token ([ad5bf0a](https://github.com/neverprepared/az-pim-cli/commit/ad5bf0a068166b15b52bf05be3562f08e98b5703))
* update dependencies ([#10](https://github.com/neverprepared/az-pim-cli/issues/10)) ([66a8963](https://github.com/neverprepared/az-pim-cli/commit/66a89636a6c6a6428044da19f86415ca98cce457))
* update dependencies ([#28](https://github.com/neverprepared/az-pim-cli/issues/28)) ([d111506](https://github.com/neverprepared/az-pim-cli/commit/d111506717481b9b2958c3337400941903a080af))
* upgrade dependencies ([202b948](https://github.com/neverprepared/az-pim-cli/commit/202b9485e8003023e5bfa9c0b0f078b4b52660ae))
* upgrade dependencies ([#88](https://github.com/neverprepared/az-pim-cli/issues/88)) ([2649bf7](https://github.com/neverprepared/az-pim-cli/commit/2649bf722686ada8c89cec0509dd5a659c461c01))


### Tests

* add coverage for new features ([070f5b8](https://github.com/neverprepared/az-pim-cli/commit/070f5b862e9ab0293ae5fb8b189248d4abeececa))
* add unit tests for the client and utils ([#73](https://github.com/neverprepared/az-pim-cli/issues/73)) ([33f8811](https://github.com/neverprepared/az-pim-cli/commit/33f88114d88b8ed4170d47bc5a8cdf0ab9a2c350))


### Linting

* resolve 'empty branch' linter error (SA9003) ([3a5c36c](https://github.com/neverprepared/az-pim-cli/commit/3a5c36c314b83261d527f95cbce9ed97f7057b53))
* resolve 'unused' linter error ([0afba95](https://github.com/neverprepared/az-pim-cli/commit/0afba95e2ad8af4db98fbe77aa83aaeb44518cc8))


### Build

* change versioning strategy ([#64](https://github.com/neverprepared/az-pim-cli/issues/64)) ([913af01](https://github.com/neverprepared/az-pim-cli/commit/913af01d814188841bb1984a16a2a07fc7219a43))
* fix broken goreleaser builds ([#105](https://github.com/neverprepared/az-pim-cli/issues/105)) ([506acde](https://github.com/neverprepared/az-pim-cli/commit/506acde2183d3a649e9d0d49d578f34e10bed112))
* **goreleaser:** include macos-universal binary in release archives ([#107](https://github.com/neverprepared/az-pim-cli/issues/107)) ([fb45385](https://github.com/neverprepared/az-pim-cli/commit/fb453852f51c7840c051e603b216886715fa7476))


### Continuous Integration

* add build workflow for all target platforms ([f991bd1](https://github.com/neverprepared/az-pim-cli/commit/f991bd1476cfa983c6f85e7261af40a103605a9f))
* add golangci-lint config ([c9aa6bf](https://github.com/neverprepared/az-pim-cli/commit/c9aa6bfcbdb141fd1b91e1bc4137862b81951c91))
* add release-please workflow ([#36](https://github.com/neverprepared/az-pim-cli/issues/36)) ([cddf23d](https://github.com/neverprepared/az-pim-cli/commit/cddf23dc93b8516007ced7e11e78869f9351efc8))
* add workflow triggered by merge to main ([#37](https://github.com/neverprepared/az-pim-cli/issues/37)) ([4adb696](https://github.com/neverprepared/az-pim-cli/commit/4adb696c5511730585e93e30042e141f6f9d77bf))
* added semgrep workflow ([#29](https://github.com/neverprepared/az-pim-cli/issues/29)) ([292c7c4](https://github.com/neverprepared/az-pim-cli/commit/292c7c4f84b6b3b4249c5e526d85939a39d976a8))
* added snyk workflow ([#26](https://github.com/neverprepared/az-pim-cli/issues/26)) ([2bb20e2](https://github.com/neverprepared/az-pim-cli/commit/2bb20e21611afcaf808585886b48508eea3bde03))
* config for goreleaser ([113ab32](https://github.com/neverprepared/az-pim-cli/commit/113ab3270b711a78bb3a364283a5e7492353b51c))
* disable Snyk jobs until SNYK_TOKEN is configured ([6ce411e](https://github.com/neverprepared/az-pim-cli/commit/6ce411e7a319236c99208f2db78154f2ba45480e))
* **fix:** set permissions for scheduled workflow ([#33](https://github.com/neverprepared/az-pim-cli/issues/33)) ([b69157a](https://github.com/neverprepared/az-pim-cli/commit/b69157aae5bd5aae6a2fad51ee8aecbacea88544))
* pin semgrep container and snyk action to specific versions ([d659971](https://github.com/neverprepared/az-pim-cli/commit/d659971d888bde381a7430172522a9e6438de45b))
* **pre-commit:** add golangci-lint 'full' hook ([d00bd26](https://github.com/neverprepared/az-pim-cli/commit/d00bd26e716f9cde5206057e5baaa985b42ae7d8))
* **pre-commit:** added conventional-pre-commit hook ([3280f37](https://github.com/neverprepared/az-pim-cli/commit/3280f3777c6e1e2de231faa141e5f37b4cbc3e88))
* **pre-commit:** update hooks ([7e5e1cc](https://github.com/neverprepared/az-pim-cli/commit/7e5e1cc80a31b9485d4fceb404a3ce0e4e8a8668))
* run GoReleaser after release-please creates a release ([f7cdfb6](https://github.com/neverprepared/az-pim-cli/commit/f7cdfb6a13a2e1fe102a3e671698d11e4c0f9ab8))
* set permissions for workflows ([#91](https://github.com/neverprepared/az-pim-cli/issues/91)) ([2f10183](https://github.com/neverprepared/az-pim-cli/commit/2f1018335cf34d0172d332bb8da63daba157784b))
* use GITHUB_TOKEN for release-please instead of PAT ([7bbc4a8](https://github.com/neverprepared/az-pim-cli/commit/7bbc4a8caef116cbe3a65a4d761ac58148a40963))
* workflow to check for conventional-commits ([9e95133](https://github.com/neverprepared/az-pim-cli/commit/9e951338401542d8729874c1c31d36654232b44c))
* workflow to run on new tags (release) ([c39224f](https://github.com/neverprepared/az-pim-cli/commit/c39224f504aa0ec60d646c054eef2351333a93ab))
* workflow to run on pull requests ([9ab9c31](https://github.com/neverprepared/az-pim-cli/commit/9ab9c31f874b7c02844d913190ec4cc9823173b7))
* workflow to run pre-commit ([29d9aab](https://github.com/neverprepared/az-pim-cli/commit/29d9aab34aee4143865b1c44bd64c9a359e0202e))

## [1.13.0](https://github.com/mindmorass/az-pim-cli/compare/v1.12.0...v1.13.0) (2026-03-08)


### Continuous Integration

* run GoReleaser after release-please creates a release ([9ff7e6e](https://github.com/mindmorass/az-pim-cli/commit/9ff7e6e08bf9e89c8187769441603f879459536d))

## [1.12.0](https://github.com/mindmorass/az-pim-cli/compare/v1.11.0...v1.12.0) (2026-03-07)


### Continuous Integration

* add build workflow for all target platforms ([4e79802](https://github.com/mindmorass/az-pim-cli/commit/4e79802069c79d3a830583a942216e8d17675de4))
* disable Snyk jobs until SNYK_TOKEN is configured ([ebccfe0](https://github.com/mindmorass/az-pim-cli/commit/ebccfe0256d5c4c3d8fea534bd60380998da24a4))
* use GITHUB_TOKEN for release-please instead of PAT ([11d0de9](https://github.com/mindmorass/az-pim-cli/commit/11d0de9df2e2f4dc461d19306630094388aa0970))

## [1.11.0](https://github.com/netr0m/az-pim-cli/compare/v1.10.0...v1.11.0) (2026-01-01)


### Features

* add support for additional Azure environments (us gov, china) ([#111](https://github.com/netr0m/az-pim-cli/issues/111)) ([b180576](https://github.com/netr0m/az-pim-cli/commit/b1805761e280d6c0db39cd30a51d2b35cd0e7685))

## [1.10.0](https://github.com/netr0m/az-pim-cli/compare/v1.9.0...v1.10.0) (2025-10-03)


### Build

* **goreleaser:** include macos-universal binary in release archives ([#107](https://github.com/netr0m/az-pim-cli/issues/107)) ([6e63b46](https://github.com/netr0m/az-pim-cli/commit/6e63b4636f28c8d99cf36d61a1d759fa98ac5021))

## [1.9.0](https://github.com/netr0m/az-pim-cli/compare/v1.8.0...v1.9.0) (2025-10-03)


### Features

* **goreleaser:** include binaries in release ([#101](https://github.com/netr0m/az-pim-cli/issues/101)) ([808cdf1](https://github.com/netr0m/az-pim-cli/commit/808cdf12ddd1b95f2e96b83b325683190d05f68f))


### Bug Fixes

* **build/goreleaser:** remove incorrect template key usage ([#103](https://github.com/netr0m/az-pim-cli/issues/103)) ([ef49cb2](https://github.com/netr0m/az-pim-cli/commit/ef49cb2a81b455215559202298b126e99d97182c))


### Build

* fix broken goreleaser builds ([#105](https://github.com/netr0m/az-pim-cli/issues/105)) ([46cb5b5](https://github.com/netr0m/az-pim-cli/commit/46cb5b526787e3d299fbd08dcbb76dea0f03bcd6))

## [1.8.0](https://github.com/netr0m/az-pim-cli/compare/v1.7.0...v1.8.0) (2025-10-03)


### Chores

* **deps:** upgrade dependencies ([#99](https://github.com/netr0m/az-pim-cli/issues/99)) ([f4cf17c](https://github.com/netr0m/az-pim-cli/commit/f4cf17c62213f3273b2c9760cec9cffcffb39a9a))
* upgrade dependencies ([f4cf17c](https://github.com/netr0m/az-pim-cli/commit/f4cf17c62213f3273b2c9760cec9cffcffb39a9a))

## [1.7.0](https://github.com/netr0m/az-pim-cli/compare/v1.6.1...v1.7.0) (2025-07-27)


### Bug Fixes

* prevent panic from JWT parsing ([01a09ca](https://github.com/netr0m/az-pim-cli/commit/01a09ca5aa46e437136f264f9bc412a6dc34a86b))
* prevent panic from JWT parsing ([#96](https://github.com/netr0m/az-pim-cli/issues/96)) ([01a09ca](https://github.com/netr0m/az-pim-cli/commit/01a09ca5aa46e437136f264f9bc412a6dc34a86b))


### Continuous Integration

* set permissions for workflows ([#91](https://github.com/netr0m/az-pim-cli/issues/91)) ([e4bb4d7](https://github.com/netr0m/az-pim-cli/commit/e4bb4d7617a0561ae2fad3fb00c1e12d1548d5fc))

## [1.6.1](https://github.com/netr0m/az-pim-cli/compare/v1.6.0...v1.6.1) (2025-05-30)


### Chores

* create release v1.6.1 ([#89](https://github.com/netr0m/az-pim-cli/issues/89)) ([4b0316c](https://github.com/netr0m/az-pim-cli/commit/4b0316cbbfb5091b9fb301bb901a39c1bfd58d91))
* upgrade dependencies ([#88](https://github.com/netr0m/az-pim-cli/issues/88)) ([b70f388](https://github.com/netr0m/az-pim-cli/commit/b70f38815df5cfaaa4093a3c35440131376b0ecf))

## [1.6.0](https://github.com/netr0m/az-pim-cli/compare/v1.5.0...v1.6.0) (2025-02-21)


### Features

* add support for setting start-time ([#81](https://github.com/netr0m/az-pim-cli/issues/81)) ([ee8a4a9](https://github.com/netr0m/az-pim-cli/commit/ee8a4a914be91c7ef3e7d84da3cdcd66b8e31fe9))

## [1.5.0](https://github.com/netr0m/az-pim-cli/compare/v1.4.0...v1.5.0) (2024-11-21)


### Bug Fixes

* **pim-client:** resolve invalid logic for building a request ([#76](https://github.com/netr0m/az-pim-cli/issues/76)) ([ece6a96](https://github.com/netr0m/az-pim-cli/commit/ece6a96be07f771ce9308f47750ff41c2c4676d8))
* **pim-client:** resolve invalid logic for building a request dynamically ([ece6a96](https://github.com/netr0m/az-pim-cli/commit/ece6a96be07f771ce9308f47750ff41c2c4676d8))

## [1.4.0](https://github.com/netr0m/az-pim-cli/compare/v1.3.0...v1.4.0) (2024-11-05)


### Features

* improved error messages and logging ([#68](https://github.com/netr0m/az-pim-cli/issues/68)) ([bbeea03](https://github.com/netr0m/az-pim-cli/commit/bbeea03b138d28653cc667954cd56cc25a9d9fa5))


### Code Refactoring

* create interface for azure client ([#72](https://github.com/netr0m/az-pim-cli/issues/72)) ([7391369](https://github.com/netr0m/az-pim-cli/commit/7391369453d3d24dd17e024e48100260d68da4da))

## [1.3.0](https://github.com/netr0m/az-pim-cli/compare/v1.2.0...v1.3.0) (2024-10-21)


### Features

* Support for Entra roles ([#61](https://github.com/netr0m/az-pim-cli/issues/61)) ([dd9ed19](https://github.com/netr0m/az-pim-cli/commit/dd9ed193c7bee3a85ad3cc62ada4bc2630378393))

## [1.2.0](https://github.com/netr0m/az-pim-cli/compare/v1.1.0...v1.2.0) (2024-10-21)


### ⚠ BREAKING CHANGES

* use proper terms for 'azure resources' type ([#59](https://github.com/netr0m/az-pim-cli/issues/59))

### Features

* use proper terms for 'azure resources' type ([#59](https://github.com/netr0m/az-pim-cli/issues/59)) ([6411902](https://github.com/netr0m/az-pim-cli/commit/641190289f99d2599d7dd789c5c3ea10845746ae))

## [1.1.0](https://github.com/netr0m/az-pim-cli/compare/v1.0.1...v1.1.0) (2024-09-13)


### Features

* add support for specifying 'ticket number' and 'ticket system' ([#56](https://github.com/netr0m/az-pim-cli/issues/56)) ([a62c52f](https://github.com/netr0m/az-pim-cli/commit/a62c52ff158a018d46598fa6c631ebc020c52d53))

## 1.0.1 (2024-07-01)


### Features

* activate roles ([7cdb3be](https://github.com/netr0m/az-pim-cli/commit/7cdb3be77fe393028096d066192a6c1631b3ac3d))
* add 'version' command ([#30](https://github.com/netr0m/az-pim-cli/issues/30)) ([e24a15f](https://github.com/netr0m/az-pim-cli/commit/e24a15f6fb1aa020e6e7191080c3b56363eac355))
* add reason to activate command ([#4](https://github.com/netr0m/az-pim-cli/issues/4)) ([8b43135](https://github.com/netr0m/az-pim-cli/commit/8b4313595e4b534c304619c973d42e2c8e8b1d35))
* check for various request status types ([#14](https://github.com/netr0m/az-pim-cli/issues/14)) ([57e4472](https://github.com/netr0m/az-pim-cli/commit/57e447247280dc092cc2b9ee817a53b599b47ae9))
* dry-run for 'activate' ([#22](https://github.com/netr0m/az-pim-cli/issues/22)) ([05c4095](https://github.com/netr0m/az-pim-cli/commit/05c40956017909a14f3015f2de10c4a5e43303e2))
* list eligible roles ([eb3e15a](https://github.com/netr0m/az-pim-cli/commit/eb3e15ae475d065613c1cb816dc6082e9d008c76))
* support for PIM Entra groups ([#16](https://github.com/netr0m/az-pim-cli/issues/16)) ([6fddc87](https://github.com/netr0m/az-pim-cli/commit/6fddc870a990bc6065b8dd053544fc141421428f))
* support new Azure Entra ID PIM API ([#6](https://github.com/netr0m/az-pim-cli/issues/6)) ([700323c](https://github.com/netr0m/az-pim-cli/commit/700323cc0c90674f8d1b8fd9db6db96933e15bbc))
* use az-cli for auth ([95e7553](https://github.com/netr0m/az-pim-cli/commit/95e7553cd7142b0ba35f7054f4762b23764804d3))


### Bug Fixes

* **activate:** Role selection on `activate` selects incorrect role ([#8](https://github.com/netr0m/az-pim-cli/issues/8)) ([6cb1079](https://github.com/netr0m/az-pim-cli/commit/6cb1079b62cabf219232c9e829198d70b4b122e8))
* fix casing role on activate ([#3](https://github.com/netr0m/az-pim-cli/issues/3)) ([9d92cff](https://github.com/netr0m/az-pim-cli/commit/9d92cff54a4515eb44e6226c623fe8f59cf9817c))
* use exact matching for the role selection ([#12](https://github.com/netr0m/az-pim-cli/issues/12)) ([0bf37e6](https://github.com/netr0m/az-pim-cli/commit/0bf37e6db2e648179442326c0b101328e4fd7e82))


### Documentation

* **github:** add project guidelines ([#31](https://github.com/netr0m/az-pim-cli/issues/31)) ([5fc195b](https://github.com/netr0m/az-pim-cli/commit/5fc195bda5e78fd66b0fc996b3259d380b40f102))
* initial docs ([c315b5c](https://github.com/netr0m/az-pim-cli/commit/c315b5c44dab5102e8a7678c09e3c81d35f87a09))


### Continuous Integration

* add release-please workflow ([#36](https://github.com/netr0m/az-pim-cli/issues/36)) ([67f357d](https://github.com/netr0m/az-pim-cli/commit/67f357d1dfb1a2bc981ad257085757e59d934b90))
* add workflow triggered by merge to main ([#37](https://github.com/netr0m/az-pim-cli/issues/37)) ([4b24cf9](https://github.com/netr0m/az-pim-cli/commit/4b24cf90b8a58a5a71c36347149418b233fa038b))
