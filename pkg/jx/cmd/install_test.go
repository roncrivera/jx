package cmd_test

import (
	"os"
	"path"
	"testing"

	"fmt"

	"github.com/jenkins-x/jx/pkg/jx/cmd"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestInstall(t *testing.T) {
	t.Parallel()
	testDir := path.Join("test_data", "install_cloud_environments_repo")
	_, err := os.Stat(testDir)
	assert.NoError(t, err)

	version, err := cmd.LoadVersionFromCloudEnvironmentsDir(testDir)
	assert.NoError(t, err)

	assert.Equal(t, "0.0.1436", version, "For Makefile in dir %s", testDir)
}

func TestGenerateProwSecret(t *testing.T) {
	fmt.Println(util.RandStringBytesMaskImprSrc(41))
}

func TestGetSafeUsername(t *testing.T) {
	t.Parallel()
	username := `Your active configuration is: [cloudshell-16392]
tutorial@bamboo-depth-206411.iam.gserviceaccount.com`
	assert.Equal(t, cmd.GetSafeUsername(username), "tutorial@bamboo-depth-206411.iam.gserviceaccount.com")

	username = `tutorial@bamboo-depth-206411.iam.gserviceaccount.com`
	assert.Equal(t, cmd.GetSafeUsername(username), "tutorial@bamboo-depth-206411.iam.gserviceaccount.com")
}
