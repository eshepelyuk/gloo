package kubernetes

import (
	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/pkg/errors"

	"github.com/solo-io/gloo-api/pkg/api/types/v1"
	"github.com/solo-io/gloo/pkg/bootstrap"
	"github.com/solo-io/gloo/pkg/endpointdiscovery"
	"github.com/solo-io/gloo/pkg/plugin"
)

func init() {
	plugin.Register(&Plugin{}, createEndpointDiscovery)
}

func createEndpointDiscovery(opts bootstrap.Options) (endpointdiscovery.Interface, error) {
	kubeConfig := opts.KubeOptions.KubeConfig
	masterUrl := opts.KubeOptions.MasterURL
	syncFrequency := opts.ConfigWatcherOptions.SyncFrequency
	disc, err := NewEndpointDiscovery(kubeConfig, masterUrl, syncFrequency)
	if err != nil {
		return nil, err
	}
	return disc, err
}

type Plugin struct{}

const (
	// define Upstream type name
	UpstreamTypeKube = "kubernetes"
)

func (p *Plugin) GetDependencies(_ *v1.Config) *plugin.Dependencies {
	return nil
}

func (p *Plugin) ProcessUpstream(_ *plugin.UpstreamPluginParams, in *v1.Upstream, out *envoyapi.Cluster) error {
	if in.Type != UpstreamTypeKube {
		return nil
	}
	// decode does validation for us
	if _, err := DecodeUpstreamSpec(in.Spec); err != nil {
		return errors.Wrap(err, "invalid kubernetes upstream spec")
	}

	// just configure the cluster to use EDS:ADS and call it a day
	out.Type = envoyapi.Cluster_EDS
	out.EdsClusterConfig = &envoyapi.Cluster_EdsClusterConfig{
		EdsConfig: &envoycore.ConfigSource{
			ConfigSourceSpecifier: &envoycore.ConfigSource_Ads{
				Ads: &envoycore.AggregatedConfigSource{},
			},
		},
	}
	return nil
}
