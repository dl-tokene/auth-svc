package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"reflect"
)

type EthRPCConfiger interface {
	EthRPCConfig() *EthRPCConfig
}

type EthRPCConfig struct {
	Client *ethclient.Client `fig:"client"`
}

func NewEthRPCConfiger(getter kv.Getter) EthRPCConfiger {
	return &ethRPCConfig{
		getter: getter,
	}
}

type ethRPCConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *ethRPCConfig) EthRPCConfig() *EthRPCConfig {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "eth_rpc")
		config := EthRPCConfig{}
		err := figure.Out(&config).With(figure.BaseHooks, GethHook).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &config
	}).(*EthRPCConfig)
}

var (
	GethHook = figure.Hooks{
		"*ethclient.Client": func(value interface{}) (reflect.Value, error) {
			switch v := value.(type) {
			case string:
				client, err := ethclient.Dial(v)
				if err != nil {
					return reflect.Value{}, err
				}
				return reflect.ValueOf(client), nil
			default:
				return reflect.Value{}, errors.Errorf("unsupported conversion from %T", value)
			}
		},
	}
)

func (c EthRPCConfig) EthClient() *ethclient.Client {
	return c.Client
}
