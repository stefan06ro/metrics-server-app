# kubernetes-metrics-server
Helm chart for Kubernetes Metrics Server running in Tenant Clusters


* Installs the [kubernetes-metrics-server](https://github.com/kubernetes-incubator/metrics-server)

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/kubernetes-metrics-server.git
$ cd kubernetes-metrics-server
$ helm install kubernetes-metrics-server/helm/kubernetes-metrics-server-chart
```

Provide a custom `values.yaml`:

```bash
$ helm install kubernetes-metrics-server-chart -f values.yaml
```

Deployment to Tenant Clusters is handled by [chart-operator](https://github.com/giantswarm/chart-operator).
