// +build k8srequired

package metrics

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/giantswarm/apprclient"
	e2esetup "github.com/giantswarm/e2esetup/chart"
	"github.com/giantswarm/e2esetup/chart/env"
	"github.com/giantswarm/e2esetup/k8s"
	"github.com/giantswarm/e2etests/managedservices"
	"github.com/giantswarm/helmclient"
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/afero"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/kubernetes-metrics-server/integration/templates"
)

const (
	testName = "metrics"

	metricsServerName = "metrics-server"
	chartName         = "kubernetes-metrics-server"
)

var (
	a          *apprclient.Client
	helmClient *helmclient.Client
	k8sSetup   *k8s.Setup
	k8sClients *k8sclient.Clients
	l          micrologger.Logger
	ms         *managedservices.ManagedServices
)

func init() {
	var err error

	{
		c := micrologger.Config{}
		l, err = micrologger.New(c)
		if err != nil {
			panic(err.Error())
		}
	}

	{
		c := apprclient.Config{
			Fs:     afero.NewOsFs(),
			Logger: l,

			Address:      "https://quay.io",
			Organization: "giantswarm",
		}
		a, err = apprclient.New(c)
		if err != nil {
			panic(err.Error())
		}
	}

	{
		c := k8sclient.ClientsConfig{
			Logger: l,

			KubeConfigPath: env.KubeConfigPath(),
		}
		k8sClients, err = k8sclient.NewClients(c)
		if err != nil {
			panic(err.Error())
		}
	}

	{
		c := k8s.SetupConfig{
			Logger: l,

			Clients: k8sClients,
		}
		k8sSetup, err = k8s.NewSetup(c)
		if err != nil {
			panic(err.Error())
		}
	}

	{
		c := helmclient.Config{
			Logger:          l,
			K8sClient:       k8sClients.K8sClient(),
			RestConfig:      k8sClients.RestConfig(),
			TillerNamespace: "giantswarm",
		}
		helmClient, err = helmclient.New(c)
		if err != nil {
			panic(err.Error())
		}
	}

	{
		c := managedservices.Config{
			ApprClient: a,
			Clients:    k8sClients,
			HelmClient: helmClient,
			Logger:     l,

			ChartConfig: managedservices.ChartConfig{
				ChannelName:     fmt.Sprintf("%s-%s", env.CircleSHA(), testName),
				ChartName:       chartName,
				ChartValues:     templates.MetricsServerValues,
				Namespace:       metav1.NamespaceSystem,
				RunReleaseTests: false,
			},
			ChartResources: managedservices.ChartResources{
				Deployments: []managedservices.Deployment{
					{
						Name:      metricsServerName,
						Namespace: metav1.NamespaceSystem,
						DeploymentLabels: map[string]string{
							"giantswarm.io/service-type": "managed",
							"app":                        metricsServerName,
						},
						MatchLabels: map[string]string{
							"app": metricsServerName,
						},
						PodLabels: map[string]string{
							"giantswarm.io/service-type": "managed",
							"app":                        metricsServerName,
						},
						Replicas: 1,
					},
				},
			},
		}
		ms, err = managedservices.New(c)
		if err != nil {
			panic(err.Error())
		}
	}
}

// TestMain allows us to have common setup and teardown steps that are run
// once for all the tests https://golang.org/pkg/testing/#hdr-Main.
func TestMain(m *testing.M) {
	ctx := context.Background()

	{
		c := e2esetup.Config{
			HelmClient: helmClient,
			Setup:      k8sSetup,
		}

		v, err := e2esetup.Setup(ctx, m, c)
		if err != nil {
			l.LogCtx(ctx, "level", "error", "message", "e2e test failed", "stack", fmt.Sprintf("%#v\n", err))
		}

		os.Exit(v)
	}
}
