// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/active_directory"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/administrators"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/alert_watchers"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/alerts"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/api_clients"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/apps"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/array_connections"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/arrays"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/audits"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/authorization"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/certificates"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/connections"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/controllers"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/directories"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/directory_exports"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/directory_services"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/directory_snapshots"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/dns"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/drives"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/file_systems"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/hardware"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/host_groups"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/hosts"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/kmip"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/maintenance_windows"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/network_interfaces"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/offloads"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/pod_replica_links"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/pods"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/policies"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/ports"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/protection_group_snapshots"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/protection_groups"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/remote_pods"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/remote_protection_group_snapshots"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/remote_protection_groups"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/remote_volume_snapshots"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/sessions"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/smi_s"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/smtp"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/snmp_agents"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/snmp_managers"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/software"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/subnets"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/support"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/syslog"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/volume_groups"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/volume_snapshots"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/volumes"
)

// Default flasharray HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new flasharray HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Flasharray {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new flasharray HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Flasharray {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new flasharray client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Flasharray {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Flasharray)
	cli.Transport = transport
	cli.ActiveDirectory = active_directory.New(transport, formats)
	cli.Administrators = administrators.New(transport, formats)
	cli.AlertWatchers = alert_watchers.New(transport, formats)
	cli.Alerts = alerts.New(transport, formats)
	cli.APIClients = api_clients.New(transport, formats)
	cli.Apps = apps.New(transport, formats)
	cli.ArrayConnections = array_connections.New(transport, formats)
	cli.Arrays = arrays.New(transport, formats)
	cli.Audits = audits.New(transport, formats)
	cli.Authorization = authorization.New(transport, formats)
	cli.Certificates = certificates.New(transport, formats)
	cli.Connections = connections.New(transport, formats)
	cli.Controllers = controllers.New(transport, formats)
	cli.Directories = directories.New(transport, formats)
	cli.DirectoryExports = directory_exports.New(transport, formats)
	cli.DirectoryServices = directory_services.New(transport, formats)
	cli.DirectorySnapshots = directory_snapshots.New(transport, formats)
	cli.DNS = dns.New(transport, formats)
	cli.Drives = drives.New(transport, formats)
	cli.FileSystems = file_systems.New(transport, formats)
	cli.Hardware = hardware.New(transport, formats)
	cli.HostGroups = host_groups.New(transport, formats)
	cli.Hosts = hosts.New(transport, formats)
	cli.KMIP = kmip.New(transport, formats)
	cli.MaintenanceWindows = maintenance_windows.New(transport, formats)
	cli.NetworkInterfaces = network_interfaces.New(transport, formats)
	cli.Offloads = offloads.New(transport, formats)
	cli.PodReplicaLinks = pod_replica_links.New(transport, formats)
	cli.Pods = pods.New(transport, formats)
	cli.Policies = policies.New(transport, formats)
	cli.Ports = ports.New(transport, formats)
	cli.ProtectionGroupSnapshots = protection_group_snapshots.New(transport, formats)
	cli.ProtectionGroups = protection_groups.New(transport, formats)
	cli.RemotePods = remote_pods.New(transport, formats)
	cli.RemoteProtectionGroupSnapshots = remote_protection_group_snapshots.New(transport, formats)
	cli.RemoteProtectionGroups = remote_protection_groups.New(transport, formats)
	cli.RemoteVolumeSnapshots = remote_volume_snapshots.New(transport, formats)
	cli.Sessions = sessions.New(transport, formats)
	cli.SMIs = smi_s.New(transport, formats)
	cli.SMTP = smtp.New(transport, formats)
	cli.SNMPAgents = snmp_agents.New(transport, formats)
	cli.SNMPManagers = snmp_managers.New(transport, formats)
	cli.Software = software.New(transport, formats)
	cli.Subnets = subnets.New(transport, formats)
	cli.Support = support.New(transport, formats)
	cli.Syslog = syslog.New(transport, formats)
	cli.VolumeGroups = volume_groups.New(transport, formats)
	cli.VolumeSnapshots = volume_snapshots.New(transport, formats)
	cli.Volumes = volumes.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Flasharray is a client for flasharray
type Flasharray struct {
	ActiveDirectory active_directory.ClientService

	Administrators administrators.ClientService

	AlertWatchers alert_watchers.ClientService

	Alerts alerts.ClientService

	APIClients api_clients.ClientService

	Apps apps.ClientService

	ArrayConnections array_connections.ClientService

	Arrays arrays.ClientService

	Audits audits.ClientService

	Authorization authorization.ClientService

	Certificates certificates.ClientService

	Connections connections.ClientService

	Controllers controllers.ClientService

	Directories directories.ClientService

	DirectoryExports directory_exports.ClientService

	DirectoryServices directory_services.ClientService

	DirectorySnapshots directory_snapshots.ClientService

	DNS dns.ClientService

	Drives drives.ClientService

	FileSystems file_systems.ClientService

	Hardware hardware.ClientService

	HostGroups host_groups.ClientService

	Hosts hosts.ClientService

	KMIP kmip.ClientService

	MaintenanceWindows maintenance_windows.ClientService

	NetworkInterfaces network_interfaces.ClientService

	Offloads offloads.ClientService

	PodReplicaLinks pod_replica_links.ClientService

	Pods pods.ClientService

	Policies policies.ClientService

	Ports ports.ClientService

	ProtectionGroupSnapshots protection_group_snapshots.ClientService

	ProtectionGroups protection_groups.ClientService

	RemotePods remote_pods.ClientService

	RemoteProtectionGroupSnapshots remote_protection_group_snapshots.ClientService

	RemoteProtectionGroups remote_protection_groups.ClientService

	RemoteVolumeSnapshots remote_volume_snapshots.ClientService

	Sessions sessions.ClientService

	SMIs smi_s.ClientService

	SMTP smtp.ClientService

	SNMPAgents snmp_agents.ClientService

	SNMPManagers snmp_managers.ClientService

	Software software.ClientService

	Subnets subnets.ClientService

	Support support.ClientService

	Syslog syslog.ClientService

	VolumeGroups volume_groups.ClientService

	VolumeSnapshots volume_snapshots.ClientService

	Volumes volumes.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Flasharray) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.ActiveDirectory.SetTransport(transport)
	c.Administrators.SetTransport(transport)
	c.AlertWatchers.SetTransport(transport)
	c.Alerts.SetTransport(transport)
	c.APIClients.SetTransport(transport)
	c.Apps.SetTransport(transport)
	c.ArrayConnections.SetTransport(transport)
	c.Arrays.SetTransport(transport)
	c.Audits.SetTransport(transport)
	c.Authorization.SetTransport(transport)
	c.Certificates.SetTransport(transport)
	c.Connections.SetTransport(transport)
	c.Controllers.SetTransport(transport)
	c.Directories.SetTransport(transport)
	c.DirectoryExports.SetTransport(transport)
	c.DirectoryServices.SetTransport(transport)
	c.DirectorySnapshots.SetTransport(transport)
	c.DNS.SetTransport(transport)
	c.Drives.SetTransport(transport)
	c.FileSystems.SetTransport(transport)
	c.Hardware.SetTransport(transport)
	c.HostGroups.SetTransport(transport)
	c.Hosts.SetTransport(transport)
	c.KMIP.SetTransport(transport)
	c.MaintenanceWindows.SetTransport(transport)
	c.NetworkInterfaces.SetTransport(transport)
	c.Offloads.SetTransport(transport)
	c.PodReplicaLinks.SetTransport(transport)
	c.Pods.SetTransport(transport)
	c.Policies.SetTransport(transport)
	c.Ports.SetTransport(transport)
	c.ProtectionGroupSnapshots.SetTransport(transport)
	c.ProtectionGroups.SetTransport(transport)
	c.RemotePods.SetTransport(transport)
	c.RemoteProtectionGroupSnapshots.SetTransport(transport)
	c.RemoteProtectionGroups.SetTransport(transport)
	c.RemoteVolumeSnapshots.SetTransport(transport)
	c.Sessions.SetTransport(transport)
	c.SMIs.SetTransport(transport)
	c.SMTP.SetTransport(transport)
	c.SNMPAgents.SetTransport(transport)
	c.SNMPManagers.SetTransport(transport)
	c.Software.SetTransport(transport)
	c.Subnets.SetTransport(transport)
	c.Support.SetTransport(transport)
	c.Syslog.SetTransport(transport)
	c.VolumeGroups.SetTransport(transport)
	c.VolumeSnapshots.SetTransport(transport)
	c.Volumes.SetTransport(transport)
}
