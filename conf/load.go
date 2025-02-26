package conf

import "path"

var config *Config

func load() error {
	baseDir, err := FindBaseDir()
	if err != nil {
		return err
	}

	config, err = ReadConfig(path.Join(*baseDir, configName))
	if err != nil {
		return err
	}
	return nil
}

func Init() error {
	return load()
}

func Get() (*Config, error) {
	if config == nil {
		if err := load(); err != nil {
			return nil, err
		}
	}
	return config, nil
}

func ShouldGet() *Config {
	if config == nil {
		panic("config is not loaded")
	}
	return config
}
