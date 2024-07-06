package config

import "os"

func WriteLocalConfigFile(path string, config *Config) error {
	file, err := os.Create(path)
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}(file)

	if err != nil {
		return err
	}

	err = config.Write(file)
	if err != nil {
		return err
	}

	return nil
}
