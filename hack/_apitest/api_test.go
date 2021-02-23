package apitest

import (
	"os"
	"testing"

	"github.com/IdanAtias/csi-test/v2/pkg/sanity"
)

func TestMyDriver(t *testing.T) {
	config := &sanity.Config{
		TargetPath:                os.TempDir() + "/csi-target",
		StagingPath:               os.TempDir() + "/csi-staging",
		Address:                   "/tmp/e2e-csi-sanity.sock",
		TestNodeVolumeAttachLimit: true,
	}

	sanity.Test(t, config)
}
