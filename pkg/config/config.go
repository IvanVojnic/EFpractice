package config

/*
import (
	"EFpractic2/pkg/repository"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Storage repository.StorageConfig `yaml:"storage"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			fmt.Sprintf("error config %s", err)
		}
	})
	return instance
}
*/
