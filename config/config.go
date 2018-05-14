package config

import (
	"bytes"
	"errors"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"sync"
)

type AppConfig struct {
}

func NewAppConfig() *AppConfig {
	return new(AppConfig)
}

func (c *AppConfig) Load(configValue interface{}, readers ...io.Reader) error {
	if len(readers) == 0 {
		return errors.New("[AppConfig][Load]empty paths")
	}

	buff := bytes.NewBuffer([]byte{})
	for _, reader := range readers {
		dat, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}
		buff.Write(dat)
	}

	if err := yaml.Unmarshal(buff.Bytes(), configValue); err != nil {
		return err
	}

	return nil
}

var instance *AppConfig
var once sync.Once

func GetAppConfig() *AppConfig {
	once.Do(func() {
		instance = NewAppConfig()
	})

	return instance
}
