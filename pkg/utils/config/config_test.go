package config

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestConfig_Read(t *testing.T) {
	now := time.Now().Truncate(0)
	mt, _ := now.MarshalText()
	wantConfig := &Config{
		ConfigPath:  "~/.config/glSync/config.json",
		GitlabToken: "asdf",
		Groups: []*Group{
			{
				GroupID:      0,
				GroupName:    "test",
				GroupRootDir: "/services",
				LastSyncedAt: now,
			},
		},
	}

	configSTR := fmt.Sprintf(`{ "config_path": "~/.config/glSync/config.json", "gitlab_token": "asdf", "groups": [ { "group_id": 0, "group_name": "test", "group_root_dir": "/services", "last_synced_at": "%s" } ] }`, mt)

	rdr := strings.NewReader(configSTR)

	config := &Config{}

	err := config.Read(rdr)
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}

	if !reflect.DeepEqual(wantConfig, config) {
		log.Fatalf("want config: %s \n got config: %s", wantConfig.String(), config.String())
	}
}

func TestConfig_Write(t *testing.T) {
	var buf bytes.Buffer

	now := time.Now().Truncate(0)
	mt, _ := now.MarshalText()

	config := &Config{
		ConfigPath:  "~/.config/glSync/config.json",
		GitlabToken: "asdf",
		Groups: []*Group{
			{
				GroupID:      0,
				GroupName:    "test",
				GroupRootDir: "/services",
				LastSyncedAt: now,
			},
		},
	}
	configSTR := fmt.Sprintf(`{
  "config_path": "~/.config/glSync/config.json",
  "gitlab_token": "asdf",
  "groups": [
    {
      "group_id": 0,
      "group_name": "test",
      "group_root_dir": "/services",
      "last_synced_at": "%s"
    }
  ]
}`, mt)

	err := config.Write(&buf)
	if err != nil {
		log.Fatalf("write config error: %v", err)
	}

	if buf.String() != configSTR {
		log.Fatalf("want config: %s \n got config: %s", configSTR, buf.String())
	}
}
