# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.4.0] - 2021-08-13

## [1.3.0] - 2021-05-10

### Added

- Added new configuration value `extraArgs`.

## [1.2.2] - 2021-03-26

### Changed

- Set docker.io as the default registry

## [1.2.1] - 2020-12-10

- Push app to control plane catalogs

## [1.2.0] - 2020-12-10

### Changed

- Updated metrics-server version to 0.4.1.

## [1.1.1] - 2020-07-23

### Changed

- Updated metrics-server version to 0.3.6.
- Updated architect-orb to 0.10.0.

## [1.1.0] - 2020-06-17

### Changed

- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.

## [v1.0.0] 2020-01-03

### Changed

- Updated manifests for Kubernetes 1.16.

## [v0.4.1]

### Added

- Template image registry.

## [v0.4.0]

### Changed

- Migrated to be deployed via an app CR not a chartconfig CR.

## [v0.3.4]

### Changed

- Change priority class to `giantswarm-critical`.

## [v0.3.3]

### Changed

- Upgrade metric server to the latest version [0.3.3](https://github.com/kubernetes-incubator/metrics-server/releases/tag/v0.3.3).

## [v0.3.2]

### Changed

- Upgrade metric server to the latest version [0.3.2](https://github.com/kubernetes-incubator/metrics-server/releases/tag/v0.3.2). Notable changes there:

    - Ignore pods not in Running phase

    - Do not skip unready nodes

    - Add kubelet-certificate-authority flag

[Unreleased]: https://github.com/giantswarm/metrics-server-app/compare/v1.4.0...HEAD
[1.4.0]: https://github.com/giantswarm/metrics-server-app/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/giantswarm/metrics-server-app/compare/v1.2.2...v1.3.0
[1.2.2]: https://github.com/giantswarm/metrics-server-app/compare/v1.2.1...v1.2.2
[1.2.1]: https://github.com/giantswarm/metrics-server-app/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/giantswarm/metrics-server-app/compare/v1.1.1...v1.2.0
[1.1.1]: https://github.com/giantswarm/metrics-server-app/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/giantswarm/metrics-server-app/compare/v1.0.0...v1.1.0
[v1.0.0]: https://github.com/giantswarm/metrics-server-app/pull/11
[v0.3.2]: https://github.com/giantswarm/kubernetes-metrics-server/pull/12
