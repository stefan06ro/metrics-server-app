[![CircleCI](https://circleci.com/gh/giantswarm/metrics-server-app.svg?style=svg)](https://circleci.com/gh/giantswarm/metrics-server-app)

# kubernetes-metrics-server

Helm chart for Kubernetes Metrics Server running in Tenant Clusters

* Installs the [kubernetes-metrics-server].

## Release Process

* Ensure CHANGELOG.md is up to date.
* Create a new GitHub release which will push the tarball to the [default-catalog].
* Update [cluster-operator] with the new version.

[app-operator]: https://github.com/giantswarm/app-operator
[cluster-operator]: https://github.com/giantswarm/app-operator
[default-catalog]: https://github.com/giantswarm/default-catalog
[default-test-catalog]: https://github.com/giantswarm/default-test-catalog

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/metrics-server-app.git
$ cd metrics-server-app
$ helm install metrics-server-app/helm/metrics-server-app
```

Provide a custom `values.yaml`:

```bash
$ helm install metrics-server-app -f values.yaml
```

 ## Release Process

* Ensure CHANGELOG.md is up to date.
* Create a new GitHub release which will push the tarball to the [default-catalog].
* Update [cluster-operator] with the new version.

[app-operator]: https://github.com/giantswarm/app-operator
[cluster-operator]: https://github.com/giantswarm/app-operator
[default-catalog]: https://github.com/giantswarm/default-catalog
[default-test-catalog]: https://github.com/giantswarm/default-test-catalog
[kubernetes-metrics-server]: https://github.com/kubernetes-incubator/metrics-server
