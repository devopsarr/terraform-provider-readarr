# Changelog

## [2.0.0](https://github.com/devopsarr/terraform-provider-readarr/compare/v1.5.0...v2.0.0) (2023-05-24)


### ⚠ BREAKING CHANGES

* remove deprecated indexer rarbg

### Features

* add custom format condition data source ([4d5e9dd](https://github.com/devopsarr/terraform-provider-readarr/commit/4d5e9dda84681f0b123147704e77b1cc29b28821))
* add custom format condition release group data source ([c58e211](https://github.com/devopsarr/terraform-provider-readarr/commit/c58e2113d26c45030375a3d1d219ad9d2ba09e44))
* add custom format condition release title data source ([bf71eb8](https://github.com/devopsarr/terraform-provider-readarr/commit/bf71eb8445d322107f2e6db56ea14a7fdaa76d35))
* add custom format condition size data source ([23e5365](https://github.com/devopsarr/terraform-provider-readarr/commit/23e5365ce3df66ea24e599bc2b0f18823d52591b))
* add custom format data source ([a6ee007](https://github.com/devopsarr/terraform-provider-readarr/commit/a6ee0070df71f88b421b71cc91eef51dea911afc))
* add custom format resource ([11911e7](https://github.com/devopsarr/terraform-provider-readarr/commit/11911e71e3c8ac35acc3c7d7a2a269e3881bf558))
* add custom formats data source ([eda00b1](https://github.com/devopsarr/terraform-provider-readarr/commit/eda00b108871f367a8ce16f7472abfd29db3f94a))
* add format field in quality profile ([f6d56c8](https://github.com/devopsarr/terraform-provider-readarr/commit/f6d56c8a3850ac9cffd936ca4015ce85d3ff4e03))
* add notification on app update flag ([faad53f](https://github.com/devopsarr/terraform-provider-readarr/commit/faad53fb98b35fa6fa5646845241605cc34d190b))
* move notification flags to optional ([8f34637](https://github.com/devopsarr/terraform-provider-readarr/commit/8f346377daff4af24aefa70ef320b0a39e10ac77))
* remove deprecated indexer rarbg ([560dbba](https://github.com/devopsarr/terraform-provider-readarr/commit/560dbba1d3bbdc015aec1595db86695f0c44faac))
* update release profile fields ([2a9dc59](https://github.com/devopsarr/terraform-provider-readarr/commit/2a9dc597d803f4204b3c647d27189fb750505311))

## [1.5.0](https://github.com/devopsarr/terraform-provider-readarr/compare/v1.4.0...v1.5.0) (2023-03-09)


### Features

* add author data source ([aca7806](https://github.com/devopsarr/terraform-provider-readarr/commit/aca7806ee3b7feee93b10cee2de4e57287462163))
* add author resource ([2817b66](https://github.com/devopsarr/terraform-provider-readarr/commit/2817b660a0366045837d186f32bbdf5f8e6ddb34))
* add authors data source ([5de5d21](https://github.com/devopsarr/terraform-provider-readarr/commit/5de5d21d06fba9c9669d716eed31ed05ea849fe6))
* add delay profile data source ([cda5337](https://github.com/devopsarr/terraform-provider-readarr/commit/cda5337af590710f76b787584a72b737e3b2029f))
* add delay profile resource ([21f64cc](https://github.com/devopsarr/terraform-provider-readarr/commit/21f64cc5c6ec5f76fb2a57a4410a6b179c82efb3))
* add delay profiles data source ([76a0ad2](https://github.com/devopsarr/terraform-provider-readarr/commit/76a0ad2ebc0bf51d43ec9e2c4891316cc98491eb))
* add download client aria2 resource ([89b6ba4](https://github.com/devopsarr/terraform-provider-readarr/commit/89b6ba4fe0a5ab7a8cdf5d9932f622b43303abdc))
* add download client deluge resource ([095e76d](https://github.com/devopsarr/terraform-provider-readarr/commit/095e76d483dcac158aef6d6e08c51cc8e3d6caa6))
* add download client flood resource ([49707e1](https://github.com/devopsarr/terraform-provider-readarr/commit/49707e175c023f507c95be357d1b293adc6e447f))
* add download client hadouken resource ([6a7ce07](https://github.com/devopsarr/terraform-provider-readarr/commit/6a7ce077f0ed05249a25ec926428e6c56dd242e0))
* add download client nzbget resource ([668b205](https://github.com/devopsarr/terraform-provider-readarr/commit/668b2059631c8fe5b10b130ad3a734be564fde54))
* add download client nzbvortex resource ([3c9929e](https://github.com/devopsarr/terraform-provider-readarr/commit/3c9929e8514238484ebfb88c3dd08dc6d0dcb9fd))
* add download client pneumatic resource ([8cc4a88](https://github.com/devopsarr/terraform-provider-readarr/commit/8cc4a8828dd7dca6f09a42b4fc21f62cf819fd10))
* add download client qbittorrent resource ([bd95a51](https://github.com/devopsarr/terraform-provider-readarr/commit/bd95a51eabed49fcab83c3d257b9136149739e86))
* add download client rtorrent resource ([54da221](https://github.com/devopsarr/terraform-provider-readarr/commit/54da2218c5c60d8287133e92a5652f4b731abc56))
* add download client sabnzb resource ([757f6ad](https://github.com/devopsarr/terraform-provider-readarr/commit/757f6adeaa10f16f86ebea44c91d4cac9d6efceb))
* add download client torrent blackhole resource ([6c5d03e](https://github.com/devopsarr/terraform-provider-readarr/commit/6c5d03ebababdd78500aa8e786d9482f8f95ebb4))
* add download client torrent download station resource ([f1112fa](https://github.com/devopsarr/terraform-provider-readarr/commit/f1112fa55db65a825f712643d5ad7c6f2f8a5a05))
* add download client usenet blackhole resource ([5570218](https://github.com/devopsarr/terraform-provider-readarr/commit/55702188221792f502fd1f30877ccf6ab21c32e9))
* add download client usenet download station resource ([7444693](https://github.com/devopsarr/terraform-provider-readarr/commit/744469315020fd2b9b5c156dfca0d77f31f7a0ed))
* add download client utorrent resource ([64e6a93](https://github.com/devopsarr/terraform-provider-readarr/commit/64e6a931c843bd79642aa828286618612bfdbf7f))
* add download client vuze resource ([65bdaee](https://github.com/devopsarr/terraform-provider-readarr/commit/65bdaee181e6a87b3bee491ae05874ad4e778a5b))
* add import list data source ([f036664](https://github.com/devopsarr/terraform-provider-readarr/commit/f036664e8f2337bb75208be85e80cf12cbc3191a))
* add import list exclusion data source ([e6cf8ad](https://github.com/devopsarr/terraform-provider-readarr/commit/e6cf8adff2244490e554117e8df8c0056d34c76f))
* add import list exclusion resource ([565daa4](https://github.com/devopsarr/terraform-provider-readarr/commit/565daa4c912b7b6b72b6e4048faff3219a521335))
* add import list exclusions data source ([0e51a5f](https://github.com/devopsarr/terraform-provider-readarr/commit/0e51a5f53de41f994365cad9fb6490324cfce374))
* add import list goodreads bookshelf resource ([bcf7db5](https://github.com/devopsarr/terraform-provider-readarr/commit/bcf7db551a215759371fa0f573290adb3f3529c8))
* add import list goodreads list resource ([09d363b](https://github.com/devopsarr/terraform-provider-readarr/commit/09d363b9bf4d6e6fdab7c7eedd2737dfae4785e4))
* add import list goodreads owned books resource ([9d892fa](https://github.com/devopsarr/terraform-provider-readarr/commit/9d892faf58cd8677cd2fcf126775afb4bda86695))
* add import list goodreads series resource ([7f54c45](https://github.com/devopsarr/terraform-provider-readarr/commit/7f54c45b759661d2c1eeee6791a22ca7a28b2ab6))
* add import list lazy librarian resource ([7893626](https://github.com/devopsarr/terraform-provider-readarr/commit/78936260c9f4975a7a8a532d9bd24aa2ddedb3c0))
* add import list readarr resource ([31066d1](https://github.com/devopsarr/terraform-provider-readarr/commit/31066d104bc55915984f6259731aee213e0f35b3))
* add import list resource ([f92feff](https://github.com/devopsarr/terraform-provider-readarr/commit/f92feffcc2cf7f7e6c6ff6312ed46bc19736b9ec))
* add import lists data source ([d66b5c6](https://github.com/devopsarr/terraform-provider-readarr/commit/d66b5c6e4486e0725191bf2616b20bb3dd207349))
* add indexer config data source ([6db342e](https://github.com/devopsarr/terraform-provider-readarr/commit/6db342e5d56141c973068d0b2d82f779d35db1c7))
* add indexer config resource ([634135e](https://github.com/devopsarr/terraform-provider-readarr/commit/634135e5ac011ff747f300036a884468b8dcd6e6))
* add indexer data source ([cebc9ff](https://github.com/devopsarr/terraform-provider-readarr/commit/cebc9ffdfd72dd392098ecfe46eb20ba0cb053af))
* add indexer filelist resource ([d7f92e5](https://github.com/devopsarr/terraform-provider-readarr/commit/d7f92e58f54eecc0593e09160984fb552a24237a))
* add indexer gazelle resource ([85ee185](https://github.com/devopsarr/terraform-provider-readarr/commit/85ee18561c4fd65b41e13c329fba08e5c37afdd4))
* add indexer iptorrents resource ([454b484](https://github.com/devopsarr/terraform-provider-readarr/commit/454b48431c688a3237f8e291c1f1f9a40cfe3ba4))
* add indexer newznab resource ([a9c9a53](https://github.com/devopsarr/terraform-provider-readarr/commit/a9c9a53596cde304216b9a7fb7b67ccbf8d71535))
* add indexer nyaa resource ([dfdc098](https://github.com/devopsarr/terraform-provider-readarr/commit/dfdc0980fc8c36e7525d8903e60360eedc278919))
* add indexer rarbg resource ([c060cd1](https://github.com/devopsarr/terraform-provider-readarr/commit/c060cd15461a95dab342f8d59b045fe747d4c7f2))
* add indexer resource ([13be362](https://github.com/devopsarr/terraform-provider-readarr/commit/13be362d5d37cb1cb6dfe2457aca19691572cdfe))
* add indexer torrent rss resource ([cd9d468](https://github.com/devopsarr/terraform-provider-readarr/commit/cd9d468c51d201c97ee7319f28add16664129e06))
* add indexer torrentleech resource ([0efd6be](https://github.com/devopsarr/terraform-provider-readarr/commit/0efd6be6978668008c23928ab8d54ac0cdf8dedd))
* add indexer torznab resource ([2123e04](https://github.com/devopsarr/terraform-provider-readarr/commit/2123e04b6175914ece856e4f0172f2aa81997073))
* add indexers data source ([1269703](https://github.com/devopsarr/terraform-provider-readarr/commit/1269703424fd1ebb18d22faabbc076c6a27d3b09))
* add media management data source ([087422a](https://github.com/devopsarr/terraform-provider-readarr/commit/087422a7ee8b858eb38347330bb1b313f8ed402f))
* add media management resource ([5bb44b4](https://github.com/devopsarr/terraform-provider-readarr/commit/5bb44b45f79914e8609fb91e037fa462923c8195))
* add metadata config data source ([66c23f6](https://github.com/devopsarr/terraform-provider-readarr/commit/66c23f66f8355f444462d22548d14009466546ba))
* add metadata config resource ([2e026ad](https://github.com/devopsarr/terraform-provider-readarr/commit/2e026ad6d10476f5e2629793b6a7cbdfd5baa964))
* add metadata profile data source ([4cd6fdf](https://github.com/devopsarr/terraform-provider-readarr/commit/4cd6fdf7107f83ff655072167fb95ab27b1af6e4))
* add metadata profile resource ([a5b5048](https://github.com/devopsarr/terraform-provider-readarr/commit/a5b5048b47fe2620f986123a236551d0687eb616))
* add metadata profiles data source ([65edd82](https://github.com/devopsarr/terraform-provider-readarr/commit/65edd8269a2ef2db3bec7f8eaab04ae4c4164a9e))
* add naming data source ([601ec91](https://github.com/devopsarr/terraform-provider-readarr/commit/601ec9111a864c2cf5373e5c07380c16a3d42d77))
* add naming resource ([95c2e5d](https://github.com/devopsarr/terraform-provider-readarr/commit/95c2e5d3634f466d2fcc973602ca277149200225))
* add notification boxcar resource ([3176f4d](https://github.com/devopsarr/terraform-provider-readarr/commit/3176f4deca3c59955c44ad1a65743e10e6b88886))
* add notification discord resource ([9590d57](https://github.com/devopsarr/terraform-provider-readarr/commit/9590d57b53487a8456d94dea82a3e4e6e14687ea))
* add notification email resource ([89b059b](https://github.com/devopsarr/terraform-provider-readarr/commit/89b059b31dced580ea3df31f339043c87d23889c))
* add notification goodreads bookshelves resource ([e160958](https://github.com/devopsarr/terraform-provider-readarr/commit/e1609587c5b24472e0cb35a18990628821229e69))
* add notification goodreads owned books resource ([1efe13b](https://github.com/devopsarr/terraform-provider-readarr/commit/1efe13b5fe8965d99b9e5ee55adf0bc397e07df3))
* add notification gotify resource ([c259e63](https://github.com/devopsarr/terraform-provider-readarr/commit/c259e6361b908320e3cb43c75bb233c4a1b0d815))
* add notification join resource ([b82c651](https://github.com/devopsarr/terraform-provider-readarr/commit/b82c651fbc84f0329ddae4e9088db08e77fcfc67))
* add notification kavita resource ([0b794df](https://github.com/devopsarr/terraform-provider-readarr/commit/0b794dfa218c08788b0017662b94dc104abd775f))
* add notification mailgun resource ([0387316](https://github.com/devopsarr/terraform-provider-readarr/commit/0387316625dc3c265f988d7cbcf8650295703674))
* add notification notifiarr resource ([ae8faf8](https://github.com/devopsarr/terraform-provider-readarr/commit/ae8faf87a9ea601f2bcf21a44dcb37bef4a90dd1))
* add notification ntfy resource ([9be491a](https://github.com/devopsarr/terraform-provider-readarr/commit/9be491abf75f6044e0071463ed24a55fc4f48a98))
* add notification prowl resource ([33b7323](https://github.com/devopsarr/terraform-provider-readarr/commit/33b7323f41c4a845eb3a63f900a06bf20504747a))
* add notification pushbullet resource ([670fd54](https://github.com/devopsarr/terraform-provider-readarr/commit/670fd54210b863dec7896f09cd7c6d92e2e16872))
* add notification pushover resource ([83cb199](https://github.com/devopsarr/terraform-provider-readarr/commit/83cb1998dbb2467b017e152047467881f19eeb82))
* add notification sendgrid resource ([ff1235d](https://github.com/devopsarr/terraform-provider-readarr/commit/ff1235d9ecde8bf63ab6cce413a69de8d9d476cd))
* add notification slack resource ([5ff6aba](https://github.com/devopsarr/terraform-provider-readarr/commit/5ff6aba1e52846cf67fac6acf9a7d11cd8911597))
* add notification subsonic resource ([537ef33](https://github.com/devopsarr/terraform-provider-readarr/commit/537ef3356b9d0f25a19a933fe4867330095199b3))
* add notification synology indexer resource ([6deb2d5](https://github.com/devopsarr/terraform-provider-readarr/commit/6deb2d52d729e2e9cd81c79d91004e1e9d8f6765))
* add notification telegram resource ([49f9c39](https://github.com/devopsarr/terraform-provider-readarr/commit/49f9c391c690b155672a19f554716c46dce72007))
* add notification twitter resource ([f01fc6d](https://github.com/devopsarr/terraform-provider-readarr/commit/f01fc6dc504033fc2fed6e87b96137533fbfe6c1))
* add quality data source ([673b7ba](https://github.com/devopsarr/terraform-provider-readarr/commit/673b7ba56d0603857a8225ea601ea9df119b2121))
* add quality definition data source ([90ce163](https://github.com/devopsarr/terraform-provider-readarr/commit/90ce1639ce4eadb18a62cb2089a9999622bc15a4))
* add quality definition resource ([adc24e0](https://github.com/devopsarr/terraform-provider-readarr/commit/adc24e0583aa696c76e0a1254c3695c2444fe285))
* add quality definitions data source ([8628784](https://github.com/devopsarr/terraform-provider-readarr/commit/86287840fc83e9d630787f21491a801bebb48382))
* add quality profile data source ([8da976b](https://github.com/devopsarr/terraform-provider-readarr/commit/8da976bac64b5ab0756b8b0d7d81793750700305))
* add quality profile resource ([1c14427](https://github.com/devopsarr/terraform-provider-readarr/commit/1c1442762a32856ed34701a2ffd0e7a48b105ea9))
* add quality profiles data source ([7068756](https://github.com/devopsarr/terraform-provider-readarr/commit/70687561010ba27438f8f0fb7495443d49b8141e))
* add release profile data source ([6827cb8](https://github.com/devopsarr/terraform-provider-readarr/commit/6827cb8173b6d42030ff5a1ebfc43f560b990b8a))
* add release profile resource ([68ca6ad](https://github.com/devopsarr/terraform-provider-readarr/commit/68ca6adc28f9f81aa9f1a314d04c5fdde1c7f15d))
* add release profiles data source ([6b15ec5](https://github.com/devopsarr/terraform-provider-readarr/commit/6b15ec55b19f03fbdedfcb733399d460f0acb96c))
* add root folder data source ([df388c7](https://github.com/devopsarr/terraform-provider-readarr/commit/df388c7f0e173c762d80a21dcb8b5a622b93cda1))
* add root folder resource ([6152a56](https://github.com/devopsarr/terraform-provider-readarr/commit/6152a56199651451207a0eeaefe535112fbb606b))
* add root folders data source ([99abdd5](https://github.com/devopsarr/terraform-provider-readarr/commit/99abdd5d05ac26b39d576743b8eb7e61447114e7))


### Bug Fixes

* data source read from config instead of state ([f8e833c](https://github.com/devopsarr/terraform-provider-readarr/commit/f8e833cbcf4ea4816235da796dd3206ce850876e))
* download client wrong parameters ([33ff6ae](https://github.com/devopsarr/terraform-provider-readarr/commit/33ff6ae0d48c1ec4569d885ef3e8ae1d1cc4a1d6))
* notification correct field types ([fd985b6](https://github.com/devopsarr/terraform-provider-readarr/commit/fd985b6a9881f781f05d2d3b4d0222d16ded7155))
* notification flags on sdk ([b5a723a](https://github.com/devopsarr/terraform-provider-readarr/commit/b5a723a5d35aedff8ee732c42f0588b8fe94790b))
* rename book related download client fields ([a4085c7](https://github.com/devopsarr/terraform-provider-readarr/commit/a4085c746f40947e16d9e03b6132ee1537c6438b))
* update music references to book ([46d32fd](https://github.com/devopsarr/terraform-provider-readarr/commit/46d32fdf8f21dda41f945bce7ec2da3506c1dba6))

## [1.4.0](https://github.com/devopsarr/terraform-provider-readarr/compare/v1.3.0...v1.4.0) (2022-11-17)


### Features

* add download client config datasource ([309d41d](https://github.com/devopsarr/terraform-provider-readarr/commit/309d41d0f17b6186fc27265ee92a5a5a868574f8))
* add download client config resource ([9d53e74](https://github.com/devopsarr/terraform-provider-readarr/commit/9d53e7484bd1ebdd3e7baa4fc98a24a87d580cf9))
* add download client datasource ([900feba](https://github.com/devopsarr/terraform-provider-readarr/commit/900feba8048e1b8b6737ca4003f138d4a8ede04a))
* add download client resource ([cafcb43](https://github.com/devopsarr/terraform-provider-readarr/commit/cafcb4392338d24a71c777cf0200b48979781314))
* add download client transmission resource ([04f277d](https://github.com/devopsarr/terraform-provider-readarr/commit/04f277d324b9ef96c2ef691008a5c15264e9dd32))
* add download clients datasource ([b196538](https://github.com/devopsarr/terraform-provider-readarr/commit/b19653864f1162b080d083eeea26026b8852d555))
* add remote path mapping datasource ([64ab078](https://github.com/devopsarr/terraform-provider-readarr/commit/64ab078f7dcb7fab86e4167d48d3e71b2c438d90))
* add remote path mapping resource ([0c11750](https://github.com/devopsarr/terraform-provider-readarr/commit/0c117501ac966ca0c313c95c4a3b624fef1f1b4a))
* add remote path mappings datasource ([8cd16d6](https://github.com/devopsarr/terraform-provider-readarr/commit/8cd16d69b393e9219f99d7a53534200913d4366c))


### Bug Fixes

* remove on_application_update notification flag ([4cfb1b8](https://github.com/devopsarr/terraform-provider-readarr/commit/4cfb1b82a154ddca0c4e17985682f2b8d42ebc7b))

## [1.3.0](https://github.com/devopsarr/terraform-provider-readarr/compare/v1.2.1...v1.3.0) (2022-11-15)


### Features

* add notification resource ([2ebbb68](https://github.com/devopsarr/terraform-provider-readarr/commit/2ebbb682070a338bf360e0d0502f5e4ca5c03a60))
* add notification webhook resource ([cbba1dd](https://github.com/devopsarr/terraform-provider-readarr/commit/cbba1dd4fd844d479e6e957509e75131dc8ed0d3))
* add notifications data source ([03e41f9](https://github.com/devopsarr/terraform-provider-readarr/commit/03e41f9ae07e20872658c652118f9714e8157dda))
* add notiication datasource ([fb0abdb](https://github.com/devopsarr/terraform-provider-readarr/commit/fb0abdbd96225714c42aa115e8797e52569de077))

## [1.2.1](https://github.com/devopsarr/terraform-provider-readarr/compare/v1.2.0...v1.2.1) (2022-09-06)


### Bug Fixes

* categories rendering ([a32f718](https://github.com/devopsarr/terraform-provider-readarr/commit/a32f718463a5c40f7c445a9673b73f239f234311))

## [1.2.0](https://github.com/devopsarr/terraform-provider-readarr/compare/v1.1.0...v1.2.0) (2022-09-06)


### Features

* add system status datasource ([e208594](https://github.com/devopsarr/terraform-provider-readarr/commit/e208594062f1b2e5454fd0b15e149e18c29f7e56))
* add tag datasource ([f3fc1d9](https://github.com/devopsarr/terraform-provider-readarr/commit/f3fc1d9da59640cafe1bed84b473e6e905ac8c92))

## [1.1.0](https://github.com/devopsarr/terraform-provider-readarr/compare/v1.0.0...v1.1.0) (2022-08-29)


### Features

* add validators ([535541e](https://github.com/devopsarr/terraform-provider-readarr/commit/535541e4c9f1273af47a3bb7f0f1bd18979c605a))


### Bug Fixes

* remove set parameter for framework 0.9.0 ([47d31ee](https://github.com/devopsarr/terraform-provider-readarr/commit/47d31ee2c95cb6576204790deb48452c4df2f12f))
* repo reference ([d214aae](https://github.com/devopsarr/terraform-provider-readarr/commit/d214aae942cb0576605c9126af427f6ac9d4b5dc))

## 1.0.0 (2022-03-15)


### Features

* first configuration ([d2a3b9d](https://github.com/devopsarr/terraform-provider-readarr/commit/d2a3b9d4beb87d202a3b4e541c2581f62c32fc20))
