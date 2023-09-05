# Changelog

## [0.8.2](https://github.com/mowies/lifecycle-controller/compare/lifecycle-operator-v0.8.1...lifecycle-operator-v0.8.2) (2023-09-05)


### Features

* add cloud events support ([#1843](https://github.com/mowies/lifecycle-controller/issues/1843)) ([5b47120](https://github.com/mowies/lifecycle-controller/commit/5b471203e412a919903876212ac45c04f180e482))
* **metrics-operator:** add analysis controller ([#1875](https://github.com/mowies/lifecycle-controller/issues/1875)) ([017e08b](https://github.com/mowies/lifecycle-controller/commit/017e08b0a65679ca417e363f2223b7f4fef3bc55))
* **metrics-operator:** add Analysis CRD ([#1839](https://github.com/mowies/lifecycle-controller/issues/1839)) ([9521a16](https://github.com/mowies/lifecycle-controller/commit/9521a16ce4946d3169993780f2d2a4f3a75d0445))
* monorepo setup for lifecycle-operator, scheduler and runtimes ([#1857](https://github.com/mowies/lifecycle-controller/issues/1857)) ([84e243a](https://github.com/mowies/lifecycle-controller/commit/84e243a213ffba86eddd51ccc4bf4dbd61140069))


### Bug Fixes

* **lifecycle-operator:** avoid setting the overall state of an App or WorkloadInstance between state transitions ([#1871](https://github.com/mowies/lifecycle-controller/issues/1871)) ([ee0b085](https://github.com/mowies/lifecycle-controller/commit/ee0b085b05b2b9781457eba34d5d1050b3c7a604))


### Other

* rename operator folder to lifecycle-operator ([#1819](https://github.com/mowies/lifecycle-controller/issues/1819)) ([97a2d25](https://github.com/mowies/lifecycle-controller/commit/97a2d25919c0a02165dd0dc6c7c82d57ad200139))


### Docs

* fix typos and grammar issues ([#1925](https://github.com/mowies/lifecycle-controller/issues/1925)) ([5570d55](https://github.com/mowies/lifecycle-controller/commit/5570d555bfc4bbdcbfc66b2725d5352090e5b937))


### Dependency Updates

* update dependency kubernetes-sigs/controller-tools to v0.13.0 ([#1930](https://github.com/mowies/lifecycle-controller/issues/1930)) ([8b34b63](https://github.com/mowies/lifecycle-controller/commit/8b34b63404d0339633ef41ff1cf2005deae8d2b7))
* update dependency kubernetes-sigs/kustomize to v5.1.1 ([#1853](https://github.com/mowies/lifecycle-controller/issues/1853)) ([354ab3f](https://github.com/mowies/lifecycle-controller/commit/354ab3f980c2569e17a0354ece417df40317d120))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 88a54f9 ([#1794](https://github.com/mowies/lifecycle-controller/issues/1794)) ([fc976eb](https://github.com/mowies/lifecycle-controller/commit/fc976eb07ed9a5e49ed7d4ab1dbf187cee583e64))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to bb133cf ([#1963](https://github.com/mowies/lifecycle-controller/issues/1963)) ([c7697bf](https://github.com/mowies/lifecycle-controller/commit/c7697bf54d5fe18b7c62c5b11801c6c83079b0a3))
* update module github.com/onsi/gomega to v1.27.10 ([#1796](https://github.com/mowies/lifecycle-controller/issues/1796)) ([8f14bff](https://github.com/mowies/lifecycle-controller/commit/8f14bffe27485a36e0b05b770a01e357402d92f7))
* update module golang.org/x/net to v0.14.0 ([#1855](https://github.com/mowies/lifecycle-controller/issues/1855)) ([3186188](https://github.com/mowies/lifecycle-controller/commit/31861889bf7b227f489b941ac4a52db86551fcc2))
* update module google.golang.org/grpc to v1.57.0 ([#1861](https://github.com/mowies/lifecycle-controller/issues/1861)) ([fdcbdf5](https://github.com/mowies/lifecycle-controller/commit/fdcbdf50365dfd69d16c679c6814e89570a8a0e2))
* update opentelemetry-go monorepo (minor) ([#1931](https://github.com/mowies/lifecycle-controller/issues/1931)) ([a0a1a7e](https://github.com/mowies/lifecycle-controller/commit/a0a1a7e97906ab56ed85da7ab9b6d1e13c902397))


### Refactoring

* **lifecycle-operator:** eventing and telemetry ([#1844](https://github.com/mowies/lifecycle-controller/issues/1844)) ([0130576](https://github.com/mowies/lifecycle-controller/commit/0130576a17a78453019c150af849c06553d799a6))
* **lifecycle-operator:** refactor event emitter unit tests ([#1867](https://github.com/mowies/lifecycle-controller/issues/1867)) ([2558f74](https://github.com/mowies/lifecycle-controller/commit/2558f742031e4d38a8006ce9894f41bddac7cd3f))
