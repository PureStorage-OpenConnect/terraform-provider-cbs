package faclient

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/client"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/client/authorization"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/client/protection_groups"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/client/remote_volume_snapshots"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/client/volume_snapshots"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/client/volumes"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

type FlashArrayClient struct {
	mgmtEndpoint string
	apiToken     string
	HTTPClient   *http.Client
	//UserAgent    string

	transport         *httptransport.Runtime
	SupportedVersions []string

	// clients of different versions
	fa24 *client.Flasharray
}

func (f *FlashArrayClient) set24Client() {
	f.fa24 = client.New(f.transport, strfmt.Default)
	// get supported version
	resp, err := f.fa24.Authorization.GetAPIAPIVersion(nil)
	fmt.Printf("Response %v, Error %v\n,", resp, err)
	if err != nil {
		fmt.Println("Ignore error of fetching versions")
	}
	f.SupportedVersions = resp.GetPayload().Version
}

func (f *FlashArrayClient) setTransport() {
	f.transport = httptransport.NewWithClient(f.mgmtEndpoint, "", []string{"https"}, f.HTTPClient)
	f.transport.SetDebug(true)
}

func New(serviceName string, mgmtEndpoint string, apiToken string) *FlashArrayClient {
	// jar, _ := cookiejar.New(nil)
	f := &FlashArrayClient{
		mgmtEndpoint: mgmtEndpoint,
		apiToken:     apiToken,
		HTTPClient: &http.Client{
			// http.Client doesn't set the default Timeout,
			// so it will be blocked forever without Timeout setting
			Timeout: time.Second * 30,
			// Jar:     jar, // Not setting cookiejar because it will return same response even though array state changes
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		//UserAgent: restclient.BuildUserAgent(appName, appVersion),
	}
	f.setTransport()
	f.set24Client()
	return f
}

// configureAuth will start a new session
// TODO: only do it when session expires
func (f *FlashArrayClient) configureAuth() error {
	fmt.Println("Create Auth Client")
	authClient := f.fa24.Authorization
	// Issue /login request, create a session, getting xAuthToken
	var xAuthToken string
	param := authorization.NewPostAPI24LoginParams()
	param.SetAPIToken(&f.apiToken)
	resp, err := authClient.PostAPI24Login(param)
	if err != nil {
		fmt.Printf("Response %v, Error %v, xAuthToken %v\n", resp, err, xAuthToken)
		return err
	} else {
		xAuthToken = resp.XAuthToken
	}
	// Use the x-auth-token as default authentication
	apiKeyQueryAuth := httptransport.APIKeyAuth("x-auth-token", "header", xAuthToken)
	f.transport.DefaultAuthentication = apiKeyQueryAuth
	return nil
}

// PgroupInfo contains the configuration of a protection group
type PgroupInfo struct {
	Name   string `json:"name"`
	Source string `json:"source"`
	// Targets []PgroupTargetInfo `json:"targets"`
	// Volumes []string           `json:"volumes"`

	ReplicationFrequency int64 `json:"replicate_frequency"`
	ReplicationEnabled   bool  `json:"replicate_enabled"`
}

func (f *FlashArrayClient) GetPgroupList24() (*[]PgroupInfo, error) {
	f.configureAuth()

	getPgParam := protection_groups.NewGetAPI24ProtectionGroupsParams()
	resp, err := f.fa24.ProtectionGroups.GetAPI24ProtectionGroups(getPgParam)
	if err != nil {
		fmt.Printf("Response %v, Error %v\n", resp, err)
		return nil, err
	}

	respPgroups := resp.GetPayload().Items

	pgroupInfoReturn := make([]PgroupInfo, len(respPgroups))
	for i, pg := range respPgroups {
		pgroupInfoReturn[i] = PgroupInfo{
			Name:                 pg.Name,
			Source:               pg.Source.Name,
			ReplicationFrequency: pg.ReplicationSchedule.Frequency,
			ReplicationEnabled:   pg.ReplicationSchedule.Enabled,
		}
	}

	// Need separate request for getting targets
	if resp2, err2 := f.fa24.ProtectionGroups.GetAPI24ProtectionGroupsTargets(nil); true {
		fmt.Printf("Response %v, Error %v\n", resp2, err2)
	}

	// Need separate request for getting volumes
	if resp2, err2 := f.fa24.ProtectionGroups.GetAPI24ProtectionGroupsVolumes(nil); true {
		fmt.Printf("Response %v, Error %v\n", resp2, err2)
	}
	return &pgroupInfoReturn, nil
}

// The result of the snapshot that is being sent
type RemoteVolumeSnapshot struct {
	Name       string // The name of the snapshot that is being sent
	SourceName string // The name of the volume that sent snapshot belongs
}

// SendSnapshot sends snapshots to the "to" FA
// It returns the info about the snapshot that is being sent
func (f *FlashArrayClient) SendSnapshot(snapshots []string, to string) (*[]RemoteVolumeSnapshot, error) {
	f.configureAuth()

	params := remote_volume_snapshots.NewPostAPI24RemoteVolumeSnapshotsParams()
	params.SetSourceNames(snapshots)
	params.SetOn(&to)

	resp, err := f.fa24.RemoteVolumeSnapshots.PostAPI24RemoteVolumeSnapshots(params)
	if err != nil {
		fmt.Printf("Response %v, Error %v\n", resp, err)
		return nil, err
	}

	items := resp.GetPayload().Items

	remoteVolumeSnapshots := make([]RemoteVolumeSnapshot, len(items))
	for i, pg := range items {
		remoteVolumeSnapshots[i] = RemoteVolumeSnapshot{
			Name:       pg.Name,
			SourceName: pg.Source.Name,
		}
	}

	return &remoteVolumeSnapshots, nil
}

type SnapshotTransferStats struct {
	Name                 string
	DataTransferred      int64
	PhysicalBytesWritten int64
	Progress             float32
	CompletionTime       int64 // Non-zero value being snapshot is completed
}

// GetSnapshotTransfer returns transfer stats of given snapshots
func (f *FlashArrayClient) GetSnapshotTransfer(snapshots []string) (*[]SnapshotTransferStats, error) {
	f.configureAuth()

	params := volume_snapshots.NewGetAPI24VolumeSnapshotsTransferParams()
	params.SetNames(snapshots)
	resp, err := f.fa24.VolumeSnapshots.GetAPI24VolumeSnapshotsTransfer(params)
	if err != nil {
		fmt.Printf("Response %v, Error %v\n", resp, err)
		return nil, err
	}
	items := resp.GetPayload().Items
	result := make([]SnapshotTransferStats, len(items))
	for i, item := range items {
		result[i] = SnapshotTransferStats{
			Name:                 item.Name,
			DataTransferred:      item.DataTransferred,
			PhysicalBytesWritten: item.PhysicalBytesWritten,
			Progress:             item.Progress,
			CompletionTime:       item.Completed,
		}
	}

	return &result, err
}

func (f *FlashArrayClient) destroyVolumeSnapshots(snapshots []string) error {
	f.configureAuth()

	params := volume_snapshots.NewPatchAPI24VolumeSnapshotsParams()
	params.SetNames(snapshots)
	replicationSnapshot := true
	params.SetReplicationSnapshot(&replicationSnapshot)

	destroyPatch := models.VolumeSnapshotPatch{
		models.DestroyedPatchPost{
			Destroyed: true,
		},
		models.NewName{},
	}
	params.SetVolumeSnapshot(&destroyPatch)
	_, err := f.fa24.VolumeSnapshots.PatchAPI24VolumeSnapshots(params)
	if err != nil {
		return err
	}
	return nil
}

func (f *FlashArrayClient) eradicateVolumeSnapshots(snapshots []string) error {
	f.configureAuth()

	params := volume_snapshots.NewDeleteAPI24VolumeSnapshotsParams()
	params.SetNames(snapshots)
	replicationSnapshot := true
	params.SetReplicationSnapshot(&replicationSnapshot)
	_, err := f.fa24.VolumeSnapshots.DeleteAPI24VolumeSnapshots(params)
	return err
}

// DeleteVolumeSnapshots destroys and eradicates the given snapshot list
func (f *FlashArrayClient) DeleteVolumeSnapshots(snapshots []string) error {
	// Destroy and then eradicate snapshots
	if err := f.destroyVolumeSnapshots(snapshots); err != nil {
		fmt.Printf("Failed to destroy snapshots %v\n", err)
		// Proceed to eradicate anyway
	}
	if err := f.eradicateVolumeSnapshots(snapshots); err != nil {
		return err
	}
	return nil
}

func (f *FlashArrayClient) destroyVolumes(volumeNames []string) error {
	f.configureAuth()

	destroyParams := volumes.NewPatchAPI24VolumesParams()
	destroyParams.SetNames(volumeNames)

	destroyPatch := models.VolumePatch{}
	destroyPatch.Destroyed = true
	destroyParams.SetVolume(&destroyPatch)
	_, err := f.fa24.Volumes.PatchAPI24Volumes(destroyParams)
	if err != nil {
		return err
	}
	return nil
}

func (f *FlashArrayClient) eradicateVolumes(volumeNames []string) error {
	f.configureAuth()

	params := volumes.NewDeleteAPI24VolumesParams()
	params.SetNames(volumeNames)
	_, err := f.fa24.Volumes.DeleteAPI24Volumes(params)
	return err
}

// DeleteVolumes destroys and eradicates the given volume list
func (f *FlashArrayClient) DeleteVolumes(volumes []string) error {
	if err := f.destroyVolumes(volumes); err != nil {
		fmt.Printf("Failed to destroy snapshots %v\n", err)
		// Proceed to eradicate anyway
	}
	if err := f.eradicateVolumes(volumes); err != nil {
		return err
	}
	return nil
}
