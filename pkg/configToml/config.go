package configToml

import (
	"os"

	"github.com/BurntSushi/toml"
)

func Get(path string, config interface{}) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	_, err := toml.DecodeFile(path, config)
	if err != nil {
		return err
	}
	return nil
}
