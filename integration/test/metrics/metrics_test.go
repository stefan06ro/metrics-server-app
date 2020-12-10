// +build k8srequired

package metrics

import (
	"context"
	"testing"
	"time"

	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	metricsAPIEndpoint = "/apis/metrics.k8s.io/v1beta1"
)

// TestMetrics ensures that deployed metrics-server chart exposes node-metrics
// via Kubernetes API extension.
func TestMetrics(t *testing.T) {
	ctx := context.Background()
	var err error

	// Check metrics-server deployment is ready.
	err = checkReadyDeployment(ctx)
	if err != nil {
		t.Fatalf("could not get metrics: %v", err)
	}

	// Check metrics availability.
	err = checkMetricsAvailability(ctx)
	if err != nil {
		t.Fatalf("could not get metrics: %v", err)
	}
}

func checkMetricsAvailability(ctx context.Context) error {
	var err error

	restClient := appTest.K8sClient().CoreV1().RESTClient()

	l.LogCtx(ctx, "level", "debug", "message", "waiting for the metrics become available")

	o := func() error {
		_, err := restClient.Get().RequestURI(metricsAPIEndpoint).Stream(context.TODO())
		if err != nil {
			return err
		}

		return nil
	}
	b := backoff.NewConstant(2*time.Minute, 5*time.Second)
	n := backoff.NewNotifier(l, ctx)

	err = backoff.RetryNotify(o, b, n)
	if err != nil {
		return microerror.Mask(err)
	}

	l.LogCtx(ctx, "level", "debug", "message", "successfully retrieved metrics from metrics server")

	return nil
}

func checkReadyDeployment(ctx context.Context) error {
	var err error

	l.LogCtx(ctx, "level", "debug", "message", "waiting for ready deployment")

	o := func() error {
		deploy, err := appTest.K8sClient().AppsV1().Deployments(metav1.NamespaceSystem).Get(ctx, app, metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			return microerror.Maskf(executionFailedError, "deployment %#q in %#q not found", app, metav1.NamespaceSystem)
		} else if err != nil {
			return microerror.Mask(err)
		}

		if deploy.Status.ReadyReplicas != *deploy.Spec.Replicas {
			return microerror.Maskf(executionFailedError, "deployment %#q want %d replicas %d ready", app, *deploy.Spec.Replicas, deploy.Status.ReadyReplicas)
		}

		return nil
	}
	b := backoff.NewConstant(2*time.Minute, 5*time.Second)
	n := backoff.NewNotifier(l, ctx)

	err = backoff.RetryNotify(o, b, n)
	if err != nil {
		return microerror.Mask(err)
	}

	l.LogCtx(ctx, "level", "debug", "message", "deployment is ready")

	return nil
}
