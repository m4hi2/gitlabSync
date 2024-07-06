package config

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

var (
	now   = time.Now().Truncate(0)
	mt, _ = now.MarshalText()
)

var sampleConfig = &Config{
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

func TestConfig_Read(t *testing.T) {
	wantConfig := sampleConfig

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

	config := sampleConfig
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

func TestReadWriteToLocalFile(t *testing.T) {
	tmp, _ := os.CreateTemp("", "*config.json")
	// It's closed here because os.CreateTemp opens file for reading.
	_ = tmp.Close()
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Fatal(err)
		}
	}(tmp.Name())

	config := sampleConfig
	err := WriteLocalConfigFile(tmp.Name(), config)
	if err != nil {
		log.Fatalf("write config error: %v", err)
	}

	c := &Config{}

	err = ReadLocalConfigFile(tmp.Name(), c)
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}

	if !reflect.DeepEqual(config, c) {
		log.Fatalf("want config: %s \n got config: %s", config.String(), c.String())
	}
}
