package targets

import (
	"github.com/metalkube/kni-installer/pkg/asset"
	"github.com/metalkube/kni-installer/pkg/asset/cluster"
	"github.com/metalkube/kni-installer/pkg/asset/ignition/bootstrap"
	"github.com/metalkube/kni-installer/pkg/asset/ignition/machine"
	"github.com/metalkube/kni-installer/pkg/asset/installconfig"
	"github.com/metalkube/kni-installer/pkg/asset/kubeconfig"
	"github.com/metalkube/kni-installer/pkg/asset/machines"
	"github.com/metalkube/kni-installer/pkg/asset/manifests"
	"github.com/metalkube/kni-installer/pkg/asset/templates/content/bootkube"
	"github.com/metalkube/kni-installer/pkg/asset/templates/content/openshift"
	"github.com/metalkube/kni-installer/pkg/asset/tls"
)

var (
	// InstallConfig are the install-config targeted assets.
	InstallConfig = []asset.WritableAsset{
		&installconfig.InstallConfig{},
	}

	// Manifests are the manifests targeted assets.
	Manifests = []asset.WritableAsset{
		&machines.Master{},
		&manifests.Manifests{},
		&manifests.Openshift{},
	}

	// ManifestTemplates are the manifest-templates targeted assets.
	ManifestTemplates = []asset.WritableAsset{
		&bootkube.KubeCloudConfig{},
		&bootkube.MachineConfigServerTLSSecret{},
		&bootkube.Pull{},
		&bootkube.CVOOverrides{},
		&bootkube.HostEtcdServiceEndpointsKubeSystem{},
		&bootkube.KubeSystemConfigmapEtcdServingCA{},
		&bootkube.KubeSystemConfigmapRootCA{},
		&bootkube.KubeSystemSecretEtcdClient{},
		&bootkube.OpenshiftMachineConfigOperator{},
		&bootkube.EtcdServiceKubeSystem{},
		&bootkube.HostEtcdServiceKubeSystem{},
		&bootkube.OpenshiftConfigSecretEtcdMetricsClient{},
		&bootkube.OpenshiftConfigConfigmapEtcdMetricsServingCA{},
		&openshift.BindingDiscovery{},
		&openshift.CloudCredsSecret{},
		&openshift.KubeadminPasswordSecret{},
		&openshift.RoleCloudCredsSecretReader{},
	}

	// IgnitionConfigs are the ignition-configs targeted assets.
	IgnitionConfigs = []asset.WritableAsset{
		&kubeconfig.AdminClient{},
		&machine.Master{},
		&machine.Worker{},
		&bootstrap.Bootstrap{},
		&cluster.Metadata{},
	}

	// Cluster are the cluster targeted assets.
	Cluster = []asset.WritableAsset{
		&cluster.TerraformVariables{},
		&kubeconfig.AdminClient{},
		&tls.JournalCertKey{},
		&cluster.Metadata{},
		&cluster.Cluster{},
	}
)
