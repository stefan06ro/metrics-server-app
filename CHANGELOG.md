# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [v0.3.4]

### Changed

- Change prioty class from to `giantswarm-critical`.

## [v0.3.3]

### Changed

- Upgrade metric server to the latest version [0.3.3](https://github.com/kubernetes-incubator/metrics-server/releases/tag/v0.3.3).

## [v0.3.2]

### Changed

- Upgrade metric server to the latest version [0.3.2](https://github.com/kubernetes-incubator/metrics-server/releases/tag/v0.3.2). Notable changes there:

    - Ignore pods not in Running phase

    - Do not skip unready nodes

    - Add kubelet-certificate-authority flag

[0.3.2]: https://github.com/giantswarm/kubernetes-metrics-server/pull/12
