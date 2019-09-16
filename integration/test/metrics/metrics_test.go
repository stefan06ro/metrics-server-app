// +build k8srequired

package metrics

import (
	"context"
	"testing"
	"time"

	"k8s.io/helm/pkg/helm"

	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
)

const (
	longMaxInterval    = 2 * time.Minute
	metricsAPIEndpoint = "/apis/metrics.k8s.io/v1beta1"
	shortMaxInterval   = 5 * time.Second
)

// TestMetrics ensures that deployed metrics-server chart exposes node-metrics
// via Kubernetes API extension.
func TestMetrics(t *testing.T) {
	ctx := context.Background()

	// Install chart and wait for deployed status
	err := ms.Test(ctx)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	// Check metrics availability
	err = checkMetricsAvailability(ctx)
	if err != nil {
		t.Fatalf("could not get metrics: %v", err)
	}

	// Delete release
	err = helmClient.DeleteRelease(ctx, chartName, helm.DeletePurge(true))
	if err != nil {
		t.Fatalf("failed to teardown resource: %v", err)
	}
}

func checkMetricsAvailability(ctx context.Context) error {
	var err error

	restClient := k8sClients.K8sClient().CoreV1().RESTClient()

	l.LogCtx(ctx, "level", "debug", "message", "waiting for the metrics become available")

	o := func() error {

		_, err := restClient.Get().RequestURI(metricsAPIEndpoint).Stream()
		if err != nil {
			return err
		}

		return nil
	}
	b := backoff.NewConstant(longMaxInterval, shortMaxInterval)
	n := backoff.NewNotifier(l, ctx)

	err = backoff.RetryNotify(o, b, n)
	if err != nil {
		return microerror.Mask(err)
	}

	l.LogCtx(ctx, "level", "debug", "message", "successfully retrieved metrics from metrics server")

	return nil
}
