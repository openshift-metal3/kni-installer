package baremetal

import (
	libvirt "github.com/libvirt/libvirt-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/metalkube/kni-installer/pkg/destroy"
	"github.com/metalkube/kni-installer/pkg/types"
)

// ClusterUninstaller holds the various options for the cluster we want to delete.
type ClusterUninstaller struct {
	LibvirtURI string
	Logger     logrus.FieldLogger
}

// Run is the entrypoint to start the uninstall process.
func (o *ClusterUninstaller) Run() error {
	o.Logger.Debug("Deleting bare metal resources")

	// FIXME: close the connection
	_, err := libvirt.NewConnect(o.LibvirtURI)
	if err != nil {
		return errors.Wrap(err, "failed to connect to Libvirt daemon")
	}

	o.Logger.Debug("FIXME: delete resources!")

	return nil
}

// New returns bare metal Uninstaller from ClusterMetadata.
func New(logger logrus.FieldLogger, metadata *types.ClusterMetadata) (destroy.Destroyer, error) {
	return &ClusterUninstaller{
		LibvirtURI: metadata.ClusterPlatformMetadata.BareMetal.URI,
		Logger:     logger,
	}, nil
}
