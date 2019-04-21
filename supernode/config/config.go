package config

import (
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/dragonflyoss/Dragonfly/common/util"
)

// NewConfig create a instant with default values.
func NewConfig() *Config {
	return &Config{
		BaseProperties: NewBaseProperties(),
	}
}

// Config contains all configuration of supernode.
type Config struct {
	*BaseProperties `yaml:"base"`
	Plugins         map[PluginType][]*PluginProperties `yaml:"plugins"`
	Storages        map[string]interface{}             `yaml:"storages"`
}

// Load loads config properties from the giving file.
func (c *Config) Load(path string) error {
	return util.LoadYaml(path, c)
}

func (c *Config) String() string {
	if out, err := yaml.Marshal(c); err == nil {
		return string(out)
	}
	return ""
}

// NewBaseProperties create a instant with default values.
func NewBaseProperties() *BaseProperties {
	home := filepath.Join(string(filepath.Separator), "home", "admin", "supernode")
	return &BaseProperties{
		ListenPort:              8002,
		HomeDir:                 home,
		SchedulerCorePoolSize:   10,
		DownloadPath:            filepath.Join(home, "repo", "download"),
		PeerUpLimit:             5,
		PeerDownLimit:           5,
		EliminationLimit:        5,
		FailureCountLimit:       5,
		LinkLimit:               20,
		SystemReservedBandwidth: 20,
		MaxBandwidth:            200,
		EnableProfiler:          false,
		Debug:                   false,
	}
}

// BaseProperties contains all basic properties of supernode.
type BaseProperties struct {
	// ListenPort is the port supernode server listens on.
	// default:
	ListenPort int `yaml:"listenPort"`

	// HomeDir is working directory of supernode.
	// default: /home/admin/supernode
	HomeDir string `yaml:"homeDir"`

	// the core pool size of ScheduledExecutorService.
	// When a request to start a download task, supernode will construct a thread concurrent pool
	// to download pieces of source file and write to specified storage.
	// Note: source file downloading is into pieces via range attribute set in HTTP header.
	// default: 10
	SchedulerCorePoolSize int `yaml:"schedulerCorePoolSize"`

	// DownloadPath specifies the path where to store downloaded files from source address.
	// This path can be set beyond BaseDir, such as taking advantage of a different disk from BaseDir's.
	// default: $BaseDir/downloads
	DownloadPath string `yaml:"downloadPath"`

	// PeerUpLimit is the upload limit of a peer. When dfget starts to play a role of peer,
	// it can only stand PeerUpLimit upload tasks from other peers.
	// default: 5
	PeerUpLimit int `yaml:"peerUpLimit"`

	// PeerDownLimit is the download limit of a peer. When a peer starts to download a file/image,
	// it will download file/image in the form of pieces. PeerDownLimit mean that a peer can only
	// stand starting PeerDownLimit concurrent downloading tasks.
	// default: 4
	PeerDownLimit int `yaml:"peerDownLimit"`

	// When dfget node starts to play a role of peer, it will provide services for other peers
	// to pull pieces. If it runs into an issue when providing services for a peer, its self failure
	// increases by 1. When the failure limit reaches EliminationLimit, the peer will isolate itself
	// as a unhealthy state. Then this dfget will be no longer called by other peers.
	// default: 5
	EliminationLimit int `yaml:"eliminationLimit"`

	// FailureCountLimit is the failure count limit set in supernode for dfget client.
	// When a dfget client takes part in the peer network constructed by supernode,
	// supernode will command the peer to start distribution task.
	// When dfget client fails to finish distribution task, the failure count of client
	// increases by 1. When failure count of client reaches to FailureCountLimit(default 5),
	// dfget client will be moved to blacklist of supernode to stop playing as a peer.
	// default: 5
	FailureCountLimit int `yaml:"failureCountLimit"`

	// LinkLimit is set for supernode to limit every piece download network speed (unit: MB/s).
	// default: 20
	LinkLimit int `yaml:"linkLimit"`

	// SystemReservedBandwidth is the network bandwidth reserved for system software.
	// unit: MB/s
	// default: 20
	SystemReservedBandwidth int `yaml:"systemReservedBandwidth"`

	// MaxBandwidth is the network bandwidth that supernode can use.
	// unit: MB/s
	// default: 200
	MaxBandwidth int `yaml:"maxBandwidth"`

	// Whether to enable profiler
	// default: false
	EnableProfiler bool `yaml:"enableProfiler"`

	// Whether to open DEBUG level
	// default: false
	Debug bool `yaml:"debug"`
}
