# kubernetes-metrics-server
Helm chart for Kubernetes Metrics Server running in Tenant Clusters


* Installs the [kubernetes-metrics-server](https://github.com/kubernetes-incubator/metrics-server)

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

Deployment to Tenant Clusters is handled by [app-operator](https://github.com/giantswarm/app-operator).
