# Changelog

## [0.2.0](https://github.com/mowies/lifecycle-controller/compare/v0.3.0...v0.2.0) (2022-11-02)


### ⚠ BREAKING CHANGES

* Rename to lifecycle toolkit (#286)
* **operator:** Modified behavior of KeptnAppVersion and KeptnWorkloadInstance to support pre and post deployment evaluation checks with Prometheus montoring
* **operator:** now the namespaces have to be annotated/labeled with keptn.sh/lifecycle-controller=enabled when the lifecycle controller should be used
* **operator:** Implementation of the KeptnApp CRD and Controller. This modifies the behaviour of the KeptnWorkloadInstance and Keptn MutatingWebhook

### Features

* Add Dashboards for Applications and Workloads ([#219](https://github.com/mowies/lifecycle-controller/issues/219)) ([48589e2](https://github.com/mowies/lifecycle-controller/commit/48589e2a521df0ff7c607a9fb74f47c06f81d3bf))
* Add scheduler with annotations ([#31](https://github.com/mowies/lifecycle-controller/issues/31)) ([9e29019](https://github.com/mowies/lifecycle-controller/commit/9e29019c098fd4f1d5e36500bd2c7ef410421aa8))
* Added tutorial for setting up observability example ([#145](https://github.com/mowies/lifecycle-controller/issues/145)) ([28f5a9c](https://github.com/mowies/lifecycle-controller/commit/28f5a9c24d031694e2066318bc85ae6e79dfd095))
* Bootstrap Service CR and controller ([#21](https://github.com/mowies/lifecycle-controller/issues/21)) ([c714ecc](https://github.com/mowies/lifecycle-controller/commit/c714eccc3b9c4d1309036fc9d193da3154b4cac5))
* First draft of a scheduler ([#19](https://github.com/mowies/lifecycle-controller/issues/19)) ([1884c86](https://github.com/mowies/lifecycle-controller/commit/1884c8678a681ed322a0ef2ea07fad3e24e01237))
* first podtatohead sample deployment manifests ([#45](https://github.com/mowies/lifecycle-controller/issues/45)) ([3e92d27](https://github.com/mowies/lifecycle-controller/commit/3e92d277ebf1a9063ebcf80f05ebe62958e45cbb))
* First Version of Function Execution ([#35](https://github.com/mowies/lifecycle-controller/issues/35)) ([f6badfd](https://github.com/mowies/lifecycle-controller/commit/f6badfd19f9f0b15c04364be7b03f524c920a015))
* initial version of function runtime ([#26](https://github.com/mowies/lifecycle-controller/issues/26)) ([c8800ee](https://github.com/mowies/lifecycle-controller/commit/c8800ee352b5d0d5eccd7338cd4fa6a3ae7d2efa))
* Inject keptn-scheduler when resource contains Keptn annotations ([#18](https://github.com/mowies/lifecycle-controller/issues/18)) ([4530e86](https://github.com/mowies/lifecycle-controller/commit/4530e8602beb4fc923b767eb586e44752f725400))
* **lfc-scheduler:** Move from Helm to Kustomize ([#53](https://github.com/mowies/lifecycle-controller/issues/53)) ([d7ba5f3](https://github.com/mowies/lifecycle-controller/commit/d7ba5f35f1b32451f833d9fd53079b4162837bde))
* **main:** Make LFC development environment installable with one command ([#138](https://github.com/mowies/lifecycle-controller/issues/138)) ([832ca37](https://github.com/mowies/lifecycle-controller/commit/832ca37d5a19297a63e17a8d367c126af37275c4))
* Namespace keptn-lifecycle-controller-system should never call webhook ([#192](https://github.com/mowies/lifecycle-controller/issues/192)) ([913a9ff](https://github.com/mowies/lifecycle-controller/commit/913a9ffd62f93aa7831b35e29853afff6213a0c9))
* **operator:** Add additional metrics for Deployment duration and interval ([#220](https://github.com/mowies/lifecycle-controller/issues/220)) ([71383c0](https://github.com/mowies/lifecycle-controller/commit/71383c0680cd17bec96b01155376cff683034d24))
* **operator:** Add commit hash, buildtime, buildversion to OTel resource attributes ([#121](https://github.com/mowies/lifecycle-controller/issues/121)) ([5a2ef61](https://github.com/mowies/lifecycle-controller/commit/5a2ef61b965472cfe850672d04b4361f5d48ca0d))
* **operator:** add fallback behavior when no keptn annotations are set ([#171](https://github.com/mowies/lifecycle-controller/issues/171)) ([b6cc674](https://github.com/mowies/lifecycle-controller/commit/b6cc674adb787615fc79dbbc5b10668c367e4736))
* **operator:** Add information about current phase in workloadinstances and appversions ([#200](https://github.com/mowies/lifecycle-controller/issues/200)) ([55fa4e9](https://github.com/mowies/lifecycle-controller/commit/55fa4e97c62aec7bd1a45f85d47cfaca48f3dd8f))
* **operator:** Add KeptnApplication controller ([#137](https://github.com/mowies/lifecycle-controller/issues/137)) ([271f5a8](https://github.com/mowies/lifecycle-controller/commit/271f5a830f216c9f827457d8a391c25d56aed2e3))
* **operator:** Add separate trace for Deployment ([#222](https://github.com/mowies/lifecycle-controller/issues/222)) ([6966e3d](https://github.com/mowies/lifecycle-controller/commit/6966e3d467e058471f15e90159ed749490bc30b2))
* **operator:** Add Spans for handling webhook requests and inject TraceContext ([#115](https://github.com/mowies/lifecycle-controller/issues/115)) ([812f2c5](https://github.com/mowies/lifecycle-controller/commit/812f2c5d49314617cb9c7532262e15edecd9f078))
* **operator:** Add support for OTel collector ([#139](https://github.com/mowies/lifecycle-controller/issues/139)) ([ac3f0d2](https://github.com/mowies/lifecycle-controller/commit/ac3f0d222f43abff7f35f1eb8de5ec80ff7dd8dc))
* **operator:** Added metrics ([#55](https://github.com/mowies/lifecycle-controller/issues/55)) ([f8a3cee](https://github.com/mowies/lifecycle-controller/commit/f8a3ceea6d1628750e7c3a7c9cd3372642bd0611))
* **operator:** Added minimal context information ([#170](https://github.com/mowies/lifecycle-controller/issues/170)) ([eebe420](https://github.com/mowies/lifecycle-controller/commit/eebe4200aac74a7c2cbc73720d1d9ac6a0c1fc72))
* **operator:** Allow pre- and post-deployment tasks as labels or annotations ([#181](https://github.com/mowies/lifecycle-controller/issues/181)) ([4241fe7](https://github.com/mowies/lifecycle-controller/commit/4241fe7cfab91aa6d38309eacf5712436a6e8327))
* **operator:** Bootstrap evaluation CRD from app ([#184](https://github.com/mowies/lifecycle-controller/issues/184)) ([74c3dbc](https://github.com/mowies/lifecycle-controller/commit/74c3dbc7b6d78d8ca7eafbac50abb8c3473701eb))
* **operator:** Bootstrap evaluation CRD from WorkloadInstance ([#188](https://github.com/mowies/lifecycle-controller/issues/188)) ([95e206b](https://github.com/mowies/lifecycle-controller/commit/95e206b4165b0277f5acbc67fc78a8e28f06741b))
* **operator:** Bootstrap KeptnEvaluationProvider and KeptnEvaluation Definition CRDs ([#165](https://github.com/mowies/lifecycle-controller/issues/165)) ([03d2346](https://github.com/mowies/lifecycle-controller/commit/03d234610fd8ef9f21e756450c7f503cb236f302))
* **operator:** Fix phase naming ([#197](https://github.com/mowies/lifecycle-controller/issues/197)) ([3739127](https://github.com/mowies/lifecycle-controller/commit/3739127d2794d75c489a6af04acf57b82920ca46))
* **operator:** Improve state and phase information ([#211](https://github.com/mowies/lifecycle-controller/issues/211)) ([6982074](https://github.com/mowies/lifecycle-controller/commit/6982074cae4e8147c4643aae821c284614d542b3))
* **operator:** Introduce KeptnEvaluation Controller + CRD ([#168](https://github.com/mowies/lifecycle-controller/issues/168)) ([1ce044a](https://github.com/mowies/lifecycle-controller/commit/1ce044a3470f815597d725d424a5491f828f2c4c))
* **operator:** Introduce OTel tracing for Task controller ([#128](https://github.com/mowies/lifecycle-controller/issues/128)) ([0baf7a9](https://github.com/mowies/lifecycle-controller/commit/0baf7a9d8058877247bc264eb6fdb645b0a77a60))
* **operator:** Introduce OTel tracing for Workload controller ([#125](https://github.com/mowies/lifecycle-controller/issues/125)) ([bc03709](https://github.com/mowies/lifecycle-controller/commit/bc03709b744d61ad966b5fba9f70dbeaffa10119))
* **operator:** Introduce OTel tracing for WorkloadInstance controller ([#131](https://github.com/mowies/lifecycle-controller/issues/131)) ([a195614](https://github.com/mowies/lifecycle-controller/commit/a1956141fe80e5b1afd79fb33198313e1dbff7fa))
* **operator:** Introduce Prometheus evaluation ([#183](https://github.com/mowies/lifecycle-controller/issues/183)) ([c2ab773](https://github.com/mowies/lifecycle-controller/commit/c2ab7733291928eaea5c38287c63e45d12754ba1))
* **operator:** namespace should be annotated when the lifecycle controller is used ([#178](https://github.com/mowies/lifecycle-controller/issues/178)) ([fa8b875](https://github.com/mowies/lifecycle-controller/commit/fa8b8758ebb5a29064f255a66d9066a863bf0944))
* **operator:** Use Async Gauges for active KLC Entities ([#206](https://github.com/mowies/lifecycle-controller/issues/206)) ([9d61ab2](https://github.com/mowies/lifecycle-controller/commit/9d61ab2664d5f3339ed5af4e1303eacf2fc89dec))
* sample function for deno runtime ([#27](https://github.com/mowies/lifecycle-controller/issues/27)) ([2501e46](https://github.com/mowies/lifecycle-controller/commit/2501e46a18dfc4ab436669fa7c42c570abad5a52))
* **scheduler:** Add OTel Resource Attributes ([#147](https://github.com/mowies/lifecycle-controller/issues/147)) ([b952156](https://github.com/mowies/lifecycle-controller/commit/b9521568e95e7855ee4fef5d55559376e2d398d9))
* **scheduler:** Add support for OTel collector ([#146](https://github.com/mowies/lifecycle-controller/issues/146)) ([9fd210d](https://github.com/mowies/lifecycle-controller/commit/9fd210d0355e5d17316f5daa8a8e289a03755d46))
* **scheduler:** Add tracing support ([#129](https://github.com/mowies/lifecycle-controller/issues/129)) ([60651d1](https://github.com/mowies/lifecycle-controller/commit/60651d15c78f9e0aa786d4dd4836c9ae828b14f3))
* **scheduler:** Background check for pod status in permit plugin ([#124](https://github.com/mowies/lifecycle-controller/issues/124)) ([97ceef6](https://github.com/mowies/lifecycle-controller/commit/97ceef6938603e315c4e1c8d2bb697aabc3dd7f8))
* **scheduler:** Disable gRPC logs when creating OTLP exporter ([#151](https://github.com/mowies/lifecycle-controller/issues/151)) ([d0f69b9](https://github.com/mowies/lifecycle-controller/commit/d0f69b9509543a5a11f22e8940a71018509ba048))
* Sign container images with sigstore/cosign ([#21](https://github.com/mowies/lifecycle-controller/issues/21)) ([66459fa](https://github.com/mowies/lifecycle-controller/commit/66459fa5395631a264c72314f2b47250120a4098))
* substitute event task ([#43](https://github.com/mowies/lifecycle-controller/issues/43)) ([3644a7d](https://github.com/mowies/lifecycle-controller/commit/3644a7d9a0d4a565a9d857348a63ed91d8cb8102))
* Switch to distroless-base image ([#46](https://github.com/mowies/lifecycle-controller/issues/46)) ([0a735b2](https://github.com/mowies/lifecycle-controller/commit/0a735b2ca22a02ca42faf7d091741d39e0f5a547))
* Webhook creates Service, Service creates ServiceRun, ServiceRun creates Event ([#30](https://github.com/mowies/lifecycle-controller/issues/30)) ([5ae58c3](https://github.com/mowies/lifecycle-controller/commit/5ae58c33abe965e79bb405e74c0f308f1220d4ee))


### Bug Fixes

* Added namespace to task definition for podtato head example ([#72](https://github.com/mowies/lifecycle-controller/issues/72)) ([7081f27](https://github.com/mowies/lifecycle-controller/commit/7081f2772aee5abec840a58c7ab700603e84cf52))
* Fix CODEOWNERS syntax ([0be5197](https://github.com/mowies/lifecycle-controller/commit/0be5197c19ea3066d28fe8e97f274efff21f66ff))
* fixed namespace in scheduler kustomization ([#63](https://github.com/mowies/lifecycle-controller/issues/63)) ([237bf4f](https://github.com/mowies/lifecycle-controller/commit/237bf4f480161f48aa0c4b5f2afbff433447d2a8))
* Missed error ([#76](https://github.com/mowies/lifecycle-controller/issues/76)) ([a59aa15](https://github.com/mowies/lifecycle-controller/commit/a59aa1552795bce15e39195af235fd42d1448e61))
* **operator:** Fix nil pointer exception in case of app not being found ([#233](https://github.com/mowies/lifecycle-controller/issues/233)) ([de9a016](https://github.com/mowies/lifecycle-controller/commit/de9a01654d7b54809932ef973860ede59f541310))
* **operator:** Fixed starting deployments, when no corresponding app-version is available ([#210](https://github.com/mowies/lifecycle-controller/issues/210)) ([3efa13e](https://github.com/mowies/lifecycle-controller/commit/3efa13e72b900a11a7dd4f65e0fbaae02211a6e9))
* **operator:** Get desired amount of replicas from upper level resource ([#89](https://github.com/mowies/lifecycle-controller/issues/89)) ([6767832](https://github.com/mowies/lifecycle-controller/commit/67678327c2531c25ea0cdb6f1b805365ae454719))
* **operator:** Update workload if spec changes ([#90](https://github.com/mowies/lifecycle-controller/issues/90)) ([ec01ad2](https://github.com/mowies/lifecycle-controller/commit/ec01ad2ccd04f0c4e6f9ba47e01c5bada128aa3b))
* **operator:** Update workload instance controller, add example ([#102](https://github.com/mowies/lifecycle-controller/issues/102)) ([e679c10](https://github.com/mowies/lifecycle-controller/commit/e679c1070f0130bd2d6616bf1856956e64dc0bac))
* **operator:** Use pointer receiver for SpanHandler methods to ensure span map is populated; thread safety via mutex ([#288](https://github.com/mowies/lifecycle-controller/issues/288)) ([a127a42](https://github.com/mowies/lifecycle-controller/commit/a127a42717068a43c60b4cc30abd56bc1478669c))
* query jobs before creating ([#79](https://github.com/mowies/lifecycle-controller/issues/79)) ([47f82b8](https://github.com/mowies/lifecycle-controller/commit/47f82b891d9d20ade2928faae307009e5c96ae22))
* scheduler config plugin configuration ([#68](https://github.com/mowies/lifecycle-controller/issues/68)) ([4c4e3c6](https://github.com/mowies/lifecycle-controller/commit/4c4e3c60a0e11267dc69ea7d8470555e3ee4f91e))
* **scheduler:** Create new context when starting background routine for pod checks ([#148](https://github.com/mowies/lifecycle-controller/issues/148)) ([543ca87](https://github.com/mowies/lifecycle-controller/commit/543ca876b27d90cb906ddb2643112a62dc923f56))
* **scheduler:** Fix the status the scheduler is acting on (preDeploymentEvaluationStatus) ([#226](https://github.com/mowies/lifecycle-controller/issues/226)) ([1a0dd92](https://github.com/mowies/lifecycle-controller/commit/1a0dd929930eb078070fb84b9bab0133ef4bccd9))
* **scheduler:** Ignoring OTel error logs ([#150](https://github.com/mowies/lifecycle-controller/issues/150)) ([0be89a5](https://github.com/mowies/lifecycle-controller/commit/0be89a56445a0356275f040dedad8fc8716a0fdd))
* Typo in observability example ([#248](https://github.com/mowies/lifecycle-controller/issues/248)) ([2f6be5f](https://github.com/mowies/lifecycle-controller/commit/2f6be5fe091951231dde005b3b9c99dcf07cab87))


### CI

* Prepare release ([#110](https://github.com/mowies/lifecycle-controller/issues/110)) ([9d7644b](https://github.com/mowies/lifecycle-controller/commit/9d7644b718e29bd37da398d89dc8b51997667358))


### Build

* Prepare release ([#149](https://github.com/mowies/lifecycle-controller/issues/149)) ([5be4504](https://github.com/mowies/lifecycle-controller/commit/5be4504e365b1c89ffc3069871a3f0fc0ecc7482))


### Docs

* Add documentation for OTel collector as pre-requisite ([#185](https://github.com/mowies/lifecycle-controller/issues/185)) ([bc3900c](https://github.com/mowies/lifecycle-controller/commit/bc3900ca64f6c7a0ef22ab94a9665aac17a83372))
* Add example for ArgoCD ([#179](https://github.com/mowies/lifecycle-controller/issues/179)) ([daf622d](https://github.com/mowies/lifecycle-controller/commit/daf622d47068f70539eb5819bc81dfe72e1b105c))
* Add flux example ([#187](https://github.com/mowies/lifecycle-controller/issues/187)) ([02cceb3](https://github.com/mowies/lifecycle-controller/commit/02cceb37d64c52a12d0779f015cf488b4ad3729f))
* Add GitHub Discussions badge and the App Lifecycle WG reference ([#42](https://github.com/mowies/lifecycle-controller/issues/42)) ([844a5f4](https://github.com/mowies/lifecycle-controller/commit/844a5f4270c6576d4bac0a08611299d4cba08560))
* Add more documentation  ([#77](https://github.com/mowies/lifecycle-controller/issues/77)) ([078080d](https://github.com/mowies/lifecycle-controller/commit/078080daf871fca50ee17f147d7d883042a7c05d))
* Add proper version badge in readme ([#114](https://github.com/mowies/lifecycle-controller/issues/114)) ([e4add2d](https://github.com/mowies/lifecycle-controller/commit/e4add2de2340f160fe30bd0cd6831107339b175e))
* Add temporary sub-project logo to the repository ([#207](https://github.com/mowies/lifecycle-controller/issues/207)) ([3708cb3](https://github.com/mowies/lifecycle-controller/commit/3708cb31dca6d8fb179bf8e46aa422ced3b877ff))
* Fix name of keptnappversions ([#215](https://github.com/mowies/lifecycle-controller/issues/215)) ([d6e3e2c](https://github.com/mowies/lifecycle-controller/commit/d6e3e2c2859ee1882902c570b7564a999f479f47))
* Improve installation steps ([#154](https://github.com/mowies/lifecycle-controller/issues/154)) ([d183e4f](https://github.com/mowies/lifecycle-controller/commit/d183e4f6b3102e426b9e29d0648cdf0c4c7cc19e))
* Improve podtato example with HTTP service lookup ([#113](https://github.com/mowies/lifecycle-controller/issues/113)) ([81b1236](https://github.com/mowies/lifecycle-controller/commit/81b1236dcff7bd37afd0e39f11638fe01406c7c4))
* Update manifest name in readme ([#111](https://github.com/mowies/lifecycle-controller/issues/111)) ([e51dbbc](https://github.com/mowies/lifecycle-controller/commit/e51dbbc0198f734fb3905b280bc1ff2e0b24d39e))
* Update the repository links in README after the org migration ([#208](https://github.com/mowies/lifecycle-controller/issues/208)) ([a1ac506](https://github.com/mowies/lifecycle-controller/commit/a1ac5060d909e9fbe0d7874aaee20af06805f033))


### Dependency Updates

* update actions/checkout action to v3 ([#282](https://github.com/mowies/lifecycle-controller/issues/282)) ([99eae9c](https://github.com/mowies/lifecycle-controller/commit/99eae9ce94ebc34ce876bbb5c1d19954f83e36d1))
* update golang docker tag to v1.18.8 ([#275](https://github.com/mowies/lifecycle-controller/issues/275)) ([c510824](https://github.com/mowies/lifecycle-controller/commit/c51082481338edc7405d42baaf15139cb35b51b9))


### Other

* Add [@mowies](https://github.com/mowies), [@philipp-hinteregger](https://github.com/philipp-hinteregger), and [@bacherfl](https://github.com/bacherfl) as owners ([#51](https://github.com/mowies/lifecycle-controller/issues/51)) ([cda7d49](https://github.com/mowies/lifecycle-controller/commit/cda7d49a0b78459923682e0dc95fe0f354c4ad4e))
* Add Evaluation instructions ([#190](https://github.com/mowies/lifecycle-controller/issues/190)) ([6717b89](https://github.com/mowies/lifecycle-controller/commit/6717b8931496be4235c3945390be53633ccb9e43))
* Add example Grafana dashboard to observability example ([#199](https://github.com/mowies/lifecycle-controller/issues/199)) ([9c20600](https://github.com/mowies/lifecycle-controller/commit/9c20600f8a5dd3149f040cf2253cd4b787cc08d3))
* added helpers in Makefile for local development ([#75](https://github.com/mowies/lifecycle-controller/issues/75)) ([3944ad7](https://github.com/mowies/lifecycle-controller/commit/3944ad783b638a0f3bbf35fb5397a6fc5bb5065c))
* Added readme with make instructions for the scheduler ([#38](https://github.com/mowies/lifecycle-controller/issues/38)) ([6049eec](https://github.com/mowies/lifecycle-controller/commit/6049eecff8399be237c029b8f334ab0491a69cd9))
* Added scheduler description into the README ([#52](https://github.com/mowies/lifecycle-controller/issues/52)) ([d7ce656](https://github.com/mowies/lifecycle-controller/commit/d7ce6567f3e4876111c59b2bc54ff786e7474787))
* Adjust conventional commit configuration ([#103](https://github.com/mowies/lifecycle-controller/issues/103)) ([1d89b46](https://github.com/mowies/lifecycle-controller/commit/1d89b46050cb019614eb36549691e3164797a238))
* Clean up release-please manifest file ([#108](https://github.com/mowies/lifecycle-controller/issues/108)) ([43ce011](https://github.com/mowies/lifecycle-controller/commit/43ce0116e4173c5aa2d659deb3eca3b9460f474b))
* **deps:** Update dependencies ([#265](https://github.com/mowies/lifecycle-controller/issues/265)) ([7a87bb8](https://github.com/mowies/lifecycle-controller/commit/7a87bb87b697b3052cc4e4cdded3f22cff641ccb))
* Fix OWNERS file ([#73](https://github.com/mowies/lifecycle-controller/issues/73)) ([2d7f412](https://github.com/mowies/lifecycle-controller/commit/2d7f4125b81d4e25f3e87ff1d3edcff4e44f9629))
* Introduce failing observability example for podtatohead ([#204](https://github.com/mowies/lifecycle-controller/issues/204)) ([f29910d](https://github.com/mowies/lifecycle-controller/commit/f29910d2feb8931cb990794899bea275d47ab7b2))
* **main:** release 0.1.0 ([#62](https://github.com/mowies/lifecycle-controller/issues/62)) ([5fa7733](https://github.com/mowies/lifecycle-controller/commit/5fa7733d5dbdcd38c1ab93d6469d730f630cb42d))
* **main:** release 0.1.0 ([#85](https://github.com/mowies/lifecycle-controller/issues/85)) ([ead06ed](https://github.com/mowies/lifecycle-controller/commit/ead06edf248d6883b180adf002c22619e0e36b4b))
* **main:** release 0.2.0 ([#120](https://github.com/mowies/lifecycle-controller/issues/120)) ([ec2359a](https://github.com/mowies/lifecycle-controller/commit/ec2359a7c318eb7caceee1164f5fdffa431555ee))
* **main:** release 0.3.0 ([#164](https://github.com/mowies/lifecycle-controller/issues/164)) ([0489374](https://github.com/mowies/lifecycle-controller/commit/048937400b7daffbf801388d7a6487e13f73c7bd))
* **operator:** Add workload and app version to KeptnTask ([#201](https://github.com/mowies/lifecycle-controller/issues/201)) ([fde0c67](https://github.com/mowies/lifecycle-controller/commit/fde0c67a4dd0b01006d1e0f6b0a240307c07bca4))
* **operator:** Support Progressing state in every phase + refactoring + speed improvements ([#236](https://github.com/mowies/lifecycle-controller/issues/236)) ([af1da5d](https://github.com/mowies/lifecycle-controller/commit/af1da5d938ce46a3dd6970a467842b01db09c33d))
* release 0.1.0 ([4c46a42](https://github.com/mowies/lifecycle-controller/commit/4c46a4297c540b9da30c5a373624d4b8e8a88231))
* release 0.1.0 ([afa8493](https://github.com/mowies/lifecycle-controller/commit/afa849324fa422352ed61faa7f0dc75d74c3c25d))
* Remove obsolete Event CRD manifests ([#83](https://github.com/mowies/lifecycle-controller/issues/83)) ([b6c83fe](https://github.com/mowies/lifecycle-controller/commit/b6c83fe5aba6f79d2faf4cb467a19eb48ad617b3))
* Rename to lifecycle toolkit ([#286](https://github.com/mowies/lifecycle-controller/issues/286)) ([9177c76](https://github.com/mowies/lifecycle-controller/commit/9177c76535b1b9dad9dc64c2d34e5e92819fcd2c))
* renamed resources according to diagram ([#41](https://github.com/mowies/lifecycle-controller/issues/41)) ([3ea81da](https://github.com/mowies/lifecycle-controller/commit/3ea81da93cc997927882cf931387322d4ca7b37c))
* **scheduler:** Rename lfc-scheduler to scheduler ([#91](https://github.com/mowies/lifecycle-controller/issues/91)) ([1d98a2d](https://github.com/mowies/lifecycle-controller/commit/1d98a2d5712cebe1e81599b35ed31bf40a86e080))
* set initial release version ([#58](https://github.com/mowies/lifecycle-controller/issues/58)) ([277f826](https://github.com/mowies/lifecycle-controller/commit/277f826a5096685e7a93c69321c4a2f7ac4a5116))
* set initial release version to release 0.1.0 ([#56](https://github.com/mowies/lifecycle-controller/issues/56)) ([b2e2984](https://github.com/mowies/lifecycle-controller/commit/b2e2984dedc78effd2595234f0d68b54b3d4bf74))
* Update dependencies and fixes ([#281](https://github.com/mowies/lifecycle-controller/issues/281)) ([5f5eda9](https://github.com/mowies/lifecycle-controller/commit/5f5eda9c599f421db0c7c94f9f5432945fabea3c))
* Update README wording ([#23](https://github.com/mowies/lifecycle-controller/issues/23)) ([628fd9f](https://github.com/mowies/lifecycle-controller/commit/628fd9fe74832d4aedd82418bd98078f616529d7))
* Update repo URL everywhere ([#216](https://github.com/mowies/lifecycle-controller/issues/216)) ([33d494c](https://github.com/mowies/lifecycle-controller/commit/33d494c537ea055d61e6a32d63c7812e0af90575))
* Updated Prometheus Network policy for granting access from lifecycle controller namespace ([#191](https://github.com/mowies/lifecycle-controller/issues/191)) ([bd77527](https://github.com/mowies/lifecycle-controller/commit/bd775276ad1324278c4bc3c82a9c0352d02bcece))
* Updated scheduler readme and developer instructions ([#123](https://github.com/mowies/lifecycle-controller/issues/123)) ([9bd5d14](https://github.com/mowies/lifecycle-controller/commit/9bd5d1461cdeeca851b6ccb78ee7e6ff0b500c1c))

## [0.3.0](https://github.com/keptn/lifecycle-toolkit/compare/v0.2.0...v0.3.0) (2022-10-20)


### ⚠ BREAKING CHANGES

* **operator:** Modified behavior of KeptnAppVersion and KeptnWorkloadInstance to support pre and post deployment evaluation checks with Prometheus montoring
* **operator:** now the namespaces have to be annotated/labeled with keptn.sh/lifecycle-toolkit=enabled when the lifecycle controller should be used
* **operator:** Implementation of the KeptnApp CRD and Controller. This modifies the behaviour of the KeptnWorkloadInstance and Keptn MutatingWebhook

### Features

* Namespace keptn-lifecycle-toolkit-system should never call webhook ([#192](https://github.com/keptn/lifecycle-toolkit/issues/192)) ([913a9ff](https://github.com/keptn/lifecycle-toolkit/commit/913a9ffd62f93aa7831b35e29853afff6213a0c9))
* **operator:** add fallback behavior when no keptn annotations are set ([#171](https://github.com/keptn/lifecycle-toolkit/issues/171)) ([b6cc674](https://github.com/keptn/lifecycle-toolkit/commit/b6cc674adb787615fc79dbbc5b10668c367e4736))
* **operator:** Add KeptnApplication controller ([#137](https://github.com/keptn/lifecycle-toolkit/issues/137)) ([271f5a8](https://github.com/keptn/lifecycle-toolkit/commit/271f5a830f216c9f827457d8a391c25d56aed2e3))
* **operator:** Added minimal context information ([#170](https://github.com/keptn/lifecycle-toolkit/issues/170)) ([eebe420](https://github.com/keptn/lifecycle-toolkit/commit/eebe4200aac74a7c2cbc73720d1d9ac6a0c1fc72))
* **operator:** Allow pre- and post-deployment tasks as labels or annotations ([#181](https://github.com/keptn/lifecycle-toolkit/issues/181)) ([4241fe7](https://github.com/keptn/lifecycle-toolkit/commit/4241fe7cfab91aa6d38309eacf5712436a6e8327))
* **operator:** Bootstrap evaluation CRD from app ([#184](https://github.com/keptn/lifecycle-toolkit/issues/184)) ([74c3dbc](https://github.com/keptn/lifecycle-toolkit/commit/74c3dbc7b6d78d8ca7eafbac50abb8c3473701eb))
* **operator:** Bootstrap evaluation CRD from WorkloadInstance ([#188](https://github.com/keptn/lifecycle-toolkit/issues/188)) ([95e206b](https://github.com/keptn/lifecycle-toolkit/commit/95e206b4165b0277f5acbc67fc78a8e28f06741b))
* **operator:** Bootstrap KeptnEvaluationProvider and KeptnEvaluation Definition CRDs ([#165](https://github.com/keptn/lifecycle-toolkit/issues/165)) ([03d2346](https://github.com/keptn/lifecycle-toolkit/commit/03d234610fd8ef9f21e756450c7f503cb236f302))
* **operator:** Fix phase naming ([#197](https://github.com/keptn/lifecycle-toolkit/issues/197)) ([3739127](https://github.com/keptn/lifecycle-toolkit/commit/3739127d2794d75c489a6af04acf57b82920ca46))
* **operator:** Introduce KeptnEvaluation Controller + CRD ([#168](https://github.com/keptn/lifecycle-toolkit/issues/168)) ([1ce044a](https://github.com/keptn/lifecycle-toolkit/commit/1ce044a3470f815597d725d424a5491f828f2c4c))
* **operator:** Introduce Prometheus evaluation ([#183](https://github.com/keptn/lifecycle-toolkit/issues/183)) ([c2ab773](https://github.com/keptn/lifecycle-toolkit/commit/c2ab7733291928eaea5c38287c63e45d12754ba1))
* **operator:** namespace should be annotated when the lifecycle controller is used ([#178](https://github.com/keptn/lifecycle-toolkit/issues/178)) ([fa8b875](https://github.com/keptn/lifecycle-toolkit/commit/fa8b8758ebb5a29064f255a66d9066a863bf0944))


### Docs

* Add documentation for OTel collector as pre-requisite ([#185](https://github.com/keptn/lifecycle-toolkit/issues/185)) ([bc3900c](https://github.com/keptn/lifecycle-toolkit/commit/bc3900ca64f6c7a0ef22ab94a9665aac17a83372))
* Add example for ArgoCD ([#179](https://github.com/keptn/lifecycle-toolkit/issues/179)) ([daf622d](https://github.com/keptn/lifecycle-toolkit/commit/daf622d47068f70539eb5819bc81dfe72e1b105c))
* Add flux example ([#187](https://github.com/keptn/lifecycle-toolkit/issues/187)) ([02cceb3](https://github.com/keptn/lifecycle-toolkit/commit/02cceb37d64c52a12d0779f015cf488b4ad3729f))
* Improve installation steps ([#154](https://github.com/keptn/lifecycle-toolkit/issues/154)) ([d183e4f](https://github.com/keptn/lifecycle-toolkit/commit/d183e4f6b3102e426b9e29d0648cdf0c4c7cc19e))


### Other

* Add Evaluation instructions ([#190](https://github.com/keptn/lifecycle-toolkit/issues/190)) ([6717b89](https://github.com/keptn/lifecycle-toolkit/commit/6717b8931496be4235c3945390be53633ccb9e43))
* Add example Grafana dashboard to observability example ([#199](https://github.com/keptn/lifecycle-toolkit/issues/199)) ([9c20600](https://github.com/keptn/lifecycle-toolkit/commit/9c20600f8a5dd3149f040cf2253cd4b787cc08d3))
* Updated Prometheus Network policy for granting access from lifecycle controller namespace ([#191](https://github.com/keptn/lifecycle-toolkit/issues/191)) ([bd77527](https://github.com/keptn/lifecycle-toolkit/commit/bd775276ad1324278c4bc3c82a9c0352d02bcece))

## [0.2.0](https://github.com/keptn/lifecycle-toolkit/compare/v0.1.0...v0.2.0) (2022-10-12)


### Features

* Added tutorial for setting up observability example ([#145](https://github.com/keptn/lifecycle-toolkit/issues/145)) ([28f5a9c](https://github.com/keptn/lifecycle-toolkit/commit/28f5a9c24d031694e2066318bc85ae6e79dfd095))
* **main:** Make LFC development environment installable with one command ([#138](https://github.com/keptn/lifecycle-toolkit/issues/138)) ([832ca37](https://github.com/keptn/lifecycle-toolkit/commit/832ca37d5a19297a63e17a8d367c126af37275c4))
* **operator:** Add commit hash, buildtime, buildversion to OTel resource attributes ([#121](https://github.com/keptn/lifecycle-toolkit/issues/121)) ([5a2ef61](https://github.com/keptn/lifecycle-toolkit/commit/5a2ef61b965472cfe850672d04b4361f5d48ca0d))
* **operator:** Add Spans for handling webhook requests and inject TraceContext ([#115](https://github.com/keptn/lifecycle-toolkit/issues/115)) ([812f2c5](https://github.com/keptn/lifecycle-toolkit/commit/812f2c5d49314617cb9c7532262e15edecd9f078))
* **operator:** Add support for OTel collector ([#139](https://github.com/keptn/lifecycle-toolkit/issues/139)) ([ac3f0d2](https://github.com/keptn/lifecycle-toolkit/commit/ac3f0d222f43abff7f35f1eb8de5ec80ff7dd8dc))
* **operator:** Added metrics ([#55](https://github.com/keptn/lifecycle-toolkit/issues/55)) ([f8a3cee](https://github.com/keptn/lifecycle-toolkit/commit/f8a3ceea6d1628750e7c3a7c9cd3372642bd0611))
* **operator:** Introduce OTel tracing for Task controller ([#128](https://github.com/keptn/lifecycle-toolkit/issues/128)) ([0baf7a9](https://github.com/keptn/lifecycle-toolkit/commit/0baf7a9d8058877247bc264eb6fdb645b0a77a60))
* **operator:** Introduce OTel tracing for Workload controller ([#125](https://github.com/keptn/lifecycle-toolkit/issues/125)) ([bc03709](https://github.com/keptn/lifecycle-toolkit/commit/bc03709b744d61ad966b5fba9f70dbeaffa10119))
* **operator:** Introduce OTel tracing for WorkloadInstance controller ([#131](https://github.com/keptn/lifecycle-toolkit/issues/131)) ([a195614](https://github.com/keptn/lifecycle-toolkit/commit/a1956141fe80e5b1afd79fb33198313e1dbff7fa))
* **scheduler:** Add OTel Resource Attributes ([#147](https://github.com/keptn/lifecycle-toolkit/issues/147)) ([b952156](https://github.com/keptn/lifecycle-toolkit/commit/b9521568e95e7855ee4fef5d55559376e2d398d9))
* **scheduler:** Add support for OTel collector ([#146](https://github.com/keptn/lifecycle-toolkit/issues/146)) ([9fd210d](https://github.com/keptn/lifecycle-toolkit/commit/9fd210d0355e5d17316f5daa8a8e289a03755d46))
* **scheduler:** Add tracing support ([#129](https://github.com/keptn/lifecycle-toolkit/issues/129)) ([60651d1](https://github.com/keptn/lifecycle-toolkit/commit/60651d15c78f9e0aa786d4dd4836c9ae828b14f3))
* **scheduler:** Background check for pod status in permit plugin ([#124](https://github.com/keptn/lifecycle-toolkit/issues/124)) ([97ceef6](https://github.com/keptn/lifecycle-toolkit/commit/97ceef6938603e315c4e1c8d2bb697aabc3dd7f8))
* **scheduler:** Disable gRPC logs when creating OTLP exporter ([#151](https://github.com/keptn/lifecycle-toolkit/issues/151)) ([d0f69b9](https://github.com/keptn/lifecycle-toolkit/commit/d0f69b9509543a5a11f22e8940a71018509ba048))


### Bug Fixes

* **scheduler:** Create new context when starting background routine for pod checks ([#148](https://github.com/keptn/lifecycle-toolkit/issues/148)) ([543ca87](https://github.com/keptn/lifecycle-toolkit/commit/543ca876b27d90cb906ddb2643112a62dc923f56))
* **scheduler:** Ignoring OTel error logs ([#150](https://github.com/keptn/lifecycle-toolkit/issues/150)) ([0be89a5](https://github.com/keptn/lifecycle-toolkit/commit/0be89a56445a0356275f040dedad8fc8716a0fdd))


### Docs

* Add proper version badge in readme ([#114](https://github.com/keptn/lifecycle-toolkit/issues/114)) ([e4add2d](https://github.com/keptn/lifecycle-toolkit/commit/e4add2de2340f160fe30bd0cd6831107339b175e))
* Improve podtato example with HTTP service lookup ([#113](https://github.com/keptn/lifecycle-toolkit/issues/113)) ([81b1236](https://github.com/keptn/lifecycle-toolkit/commit/81b1236dcff7bd37afd0e39f11638fe01406c7c4))
* Update manifest name in readme ([#111](https://github.com/keptn/lifecycle-toolkit/issues/111)) ([e51dbbc](https://github.com/keptn/lifecycle-toolkit/commit/e51dbbc0198f734fb3905b280bc1ff2e0b24d39e))


### Other

* Updated scheduler readme and developer instructions ([#123](https://github.com/keptn/lifecycle-toolkit/issues/123)) ([9bd5d14](https://github.com/keptn/lifecycle-toolkit/commit/9bd5d1461cdeeca851b6ccb78ee7e6ff0b500c1c))


### Build

* Prepare release ([#149](https://github.com/keptn/lifecycle-toolkit/issues/149)) ([5be4504](https://github.com/keptn/lifecycle-toolkit/commit/5be4504e365b1c89ffc3069871a3f0fc0ecc7482))

## 0.1.0 (2022-10-04)


### Features

* Add scheduler with annotations ([#31](https://github.com/keptn/lifecycle-toolkit/issues/31)) ([9e29019](https://github.com/keptn/lifecycle-toolkit/commit/9e29019c098fd4f1d5e36500bd2c7ef410421aa8))
* Bootstrap Service CR and controller ([#21](https://github.com/keptn/lifecycle-toolkit/issues/21)) ([c714ecc](https://github.com/keptn/lifecycle-toolkit/commit/c714eccc3b9c4d1309036fc9d193da3154b4cac5))
* First draft of a scheduler ([#19](https://github.com/keptn/lifecycle-toolkit/issues/19)) ([1884c86](https://github.com/keptn/lifecycle-toolkit/commit/1884c8678a681ed322a0ef2ea07fad3e24e01237))
* first podtatohead sample deployment manifests ([#45](https://github.com/keptn/lifecycle-toolkit/issues/45)) ([3e92d27](https://github.com/keptn/lifecycle-toolkit/commit/3e92d277ebf1a9063ebcf80f05ebe62958e45cbb))
* First Version of Function Execution ([#35](https://github.com/keptn/lifecycle-toolkit/issues/35)) ([f6badfd](https://github.com/keptn/lifecycle-toolkit/commit/f6badfd19f9f0b15c04364be7b03f524c920a015))
* initial version of function runtime ([#26](https://github.com/keptn/lifecycle-toolkit/issues/26)) ([c8800ee](https://github.com/keptn/lifecycle-toolkit/commit/c8800ee352b5d0d5eccd7338cd4fa6a3ae7d2efa))
* Inject keptn-scheduler when resource contains Keptn annotations ([#18](https://github.com/keptn/lifecycle-toolkit/issues/18)) ([4530e86](https://github.com/keptn/lifecycle-toolkit/commit/4530e8602beb4fc923b767eb586e44752f725400))
* **lfc-scheduler:** Move from Helm to Kustomize ([#53](https://github.com/keptn/lifecycle-toolkit/issues/53)) ([d7ba5f3](https://github.com/keptn/lifecycle-toolkit/commit/d7ba5f35f1b32451f833d9fd53079b4162837bde))
* sample function for deno runtime ([#27](https://github.com/keptn/lifecycle-toolkit/issues/27)) ([2501e46](https://github.com/keptn/lifecycle-toolkit/commit/2501e46a18dfc4ab436669fa7c42c570abad5a52))
* substitute event task ([#43](https://github.com/keptn/lifecycle-toolkit/issues/43)) ([3644a7d](https://github.com/keptn/lifecycle-toolkit/commit/3644a7d9a0d4a565a9d857348a63ed91d8cb8102))
* Switch to distroless-base image ([#46](https://github.com/keptn/lifecycle-toolkit/issues/46)) ([0a735b2](https://github.com/keptn/lifecycle-toolkit/commit/0a735b2ca22a02ca42faf7d091741d39e0f5a547))
* Webhook creates Service, Service creates ServiceRun, ServiceRun creates Event ([#30](https://github.com/keptn/lifecycle-toolkit/issues/30)) ([5ae58c3](https://github.com/keptn/lifecycle-toolkit/commit/5ae58c33abe965e79bb405e74c0f308f1220d4ee))


### Bug Fixes

* Added namespace to task definition for podtato head example ([#72](https://github.com/keptn/lifecycle-toolkit/issues/72)) ([7081f27](https://github.com/keptn/lifecycle-toolkit/commit/7081f2772aee5abec840a58c7ab700603e84cf52))
* Fix CODEOWNERS syntax ([0be5197](https://github.com/keptn/lifecycle-toolkit/commit/0be5197c19ea3066d28fe8e97f274efff21f66ff))
* fixed namespace in scheduler kustomization ([#63](https://github.com/keptn/lifecycle-toolkit/issues/63)) ([237bf4f](https://github.com/keptn/lifecycle-toolkit/commit/237bf4f480161f48aa0c4b5f2afbff433447d2a8))
* Missed error ([#76](https://github.com/keptn/lifecycle-toolkit/issues/76)) ([a59aa15](https://github.com/keptn/lifecycle-toolkit/commit/a59aa1552795bce15e39195af235fd42d1448e61))
* **operator:** Get desired amount of replicas from upper level resource ([#89](https://github.com/keptn/lifecycle-toolkit/issues/89)) ([6767832](https://github.com/keptn/lifecycle-toolkit/commit/67678327c2531c25ea0cdb6f1b805365ae454719))
* **operator:** Update workload if spec changes ([#90](https://github.com/keptn/lifecycle-toolkit/issues/90)) ([ec01ad2](https://github.com/keptn/lifecycle-toolkit/commit/ec01ad2ccd04f0c4e6f9ba47e01c5bada128aa3b))
* **operator:** Update workload instance controller, add example ([#102](https://github.com/keptn/lifecycle-toolkit/issues/102)) ([e679c10](https://github.com/keptn/lifecycle-toolkit/commit/e679c1070f0130bd2d6616bf1856956e64dc0bac))
* query jobs before creating ([#79](https://github.com/keptn/lifecycle-toolkit/issues/79)) ([47f82b8](https://github.com/keptn/lifecycle-toolkit/commit/47f82b891d9d20ade2928faae307009e5c96ae22))
* scheduler config plugin configuration ([#68](https://github.com/keptn/lifecycle-toolkit/issues/68)) ([4c4e3c6](https://github.com/keptn/lifecycle-toolkit/commit/4c4e3c60a0e11267dc69ea7d8470555e3ee4f91e))


### Miscellaneous Chores

* release 0.1.0 ([4c46a42](https://github.com/keptn/lifecycle-toolkit/commit/4c46a4297c540b9da30c5a373624d4b8e8a88231))
* release 0.1.0 ([afa8493](https://github.com/keptn/lifecycle-toolkit/commit/afa849324fa422352ed61faa7f0dc75d74c3c25d))


### Continuous Integration

* Prepare release ([#110](https://github.com/keptn/lifecycle-toolkit/issues/110)) ([9d7644b](https://github.com/keptn/lifecycle-toolkit/commit/9d7644b718e29bd37da398d89dc8b51997667358))

## Changelog
