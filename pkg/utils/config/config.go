package config

import (
	"encoding/json"
	"io"
	"time"
)

var (
	JSONIndent = "  "
	JSONPrefix = ""
)

type PullMethod string

const (
	PullMethodHTTP PullMethod = "http"
	PullMethodSSH  PullMethod = "ssh"
)

type Config struct {
	ConfigPath      string            `json:"config_path"`
	GitLabInstances []*GitLabInstance `json:"gitlab_instances"`
}

type Group struct {
	GroupID      int       `json:"group_id"`
	GroupName    string    `json:"group_name"`
	GroupRootDir string    `json:"group_root_dir"`
	LastSyncedAt time.Time `json:"last_synced_at"`
}

type GitLabInstance struct {
	Name        string     `json:"name"`
	HostName    string     `json:"host_name"`
	PullMethod  PullMethod `json:"pull_method"`
	GitlabToken string     `json:"gitlab_token"`
	Groups      []*Group   `json:"groups"`
}

func (c *Config) Write(writer io.Writer) error {
	cb, err := json.MarshalIndent(c, JSONPrefix, JSONIndent)
	if err != nil {
		return err
	}

	_, err = writer.Write(cb)

	return err
}

func (c *Config) Read(reader io.Reader) error {
	all, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	return json.Unmarshal(all, c)
}

func (c *Config) String() string {
	cb, _ := json.MarshalIndent(c, JSONPrefix, JSONIndent)
	return string(cb)
}
