// +build realfa

package faclient_test

import (
	"fmt"
	"testing"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient"
	//"purity.purestorage.com/pkg/faclient"
)

func TestWithRealArray(t *testing.T) {
	// hostname := "gs64-13.dev.purestorage.com"
	// apiToken := "fbab6e1f-5262-b7ac-8dbb-d52a869f8fad"
	hostname := "vm-yizhao"
	apiToken := "f2532610-629c-870d-c45c-999555510c14"

	fa := faclient.New("test", hostname, apiToken)
	fmt.Printf("Supported versions are %v", fa.SupportedVersions)

	if err := getPgroup(fa); err != nil {
		t.Errorf("Failed to get pgroup %v", err)
	}

	/*
		if err := testVolumeAndSnapshot(t, fa); err != nil {
			t.Errorf("Volume test failed %v", err)
		}
	*/
}

/*
func testVolumeAndSnapshot(t *testing.T, fa *FlashArrayClient) error {
	// Assumes volume exist
	volume := "faclient-test"

	// copy to a new volume
	toVolume := "faclient-test-copy"
	if err := fa.CreateVolumeFromSnap(volume, toVolume, true); err != nil {
		t.Errorf("Failed to create volume %v", err)
		return err
	}

	sendSnap := ""
	if resp, err := fa.SendSnapshot([]string{toVolume}, "vm-yizhao2"); err != nil {
		t.Errorf("Failed to send volume %v", err)
		return err
	} else {
		sendSnap = (*resp)[0].Name
	}

	if _, err := fa.GetSnapshotTransfer([]string{sendSnap}); err != nil {
		t.Errorf("Failed to get volume transfer %v", err)
		return err
	}

	if err := fa.DeleteVolumeSnapshots([]string{sendSnap}); err != nil {
		t.Errorf("Failed to delete volume snapshot %v", err)
		return err
	}

	if err := fa.DeleteVolumes([]string{toVolume}); err != nil {
		t.Errorf("Failed to delete volume %v", err)
		return err
	}
	return nil
}
*/

func getPgroup(fa *faclient.FlashArrayClient) error {
	/*
		if _, err := fa.GetPgroupList(true); err != nil {
			return err
		}
	*/
	if _, err := fa.GetPgroupList24(); err != nil {
		return err
	}
	return nil
}
