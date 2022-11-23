package config

import (
	"time"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type ServiceConfiger interface {
	ServiceConfig() *ServiceConfig
}

type ServiceConfig struct {
	TokenKey               string        `fig:"token_key,required"`
	TokenExpireTime        time.Duration `fig:"token_expire_time,required"`
	RefreshTokenExpireTime time.Duration `fig:"refresh_token_expire_time,required"`
}

func NewServiceConfiger(getter kv.Getter) ServiceConfiger {
	return &serviceConfig{
		getter: getter,
	}
}

type serviceConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *serviceConfig) ServiceConfig() *ServiceConfig {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "service")
		config := ServiceConfig{}
		err := figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &config
	}).(*ServiceConfig)
}
