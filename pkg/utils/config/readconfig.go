package config

import "os"

func ReadLocalConfigFile(path string, config *Config) error {
	file, err := os.Open(path)
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}(file)

	if err != nil {
		return err
	}

	return config.Read(file)
}
