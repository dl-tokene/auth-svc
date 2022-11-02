package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Adminer interface {
	AdminsConfig() *AdminsConfig
}

type AdminsConfig struct {
	Admins []string `fig:"admins"`
}

func NewAdminer(getter kv.Getter) Adminer {
	return &adminsConfig{
		getter: getter,
	}
}

type adminsConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *adminsConfig) AdminsConfig() *AdminsConfig {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "admins")
		config := AdminsConfig{}
		err := figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &config
	}).(*AdminsConfig)
}
