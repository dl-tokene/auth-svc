package config

import (
	"github.com/ethereum/go-ethereum/common"
	vault "github.com/hashicorp/vault/api"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"golang.org/x/net/context"
	"log"
	"reflect"
)

type VaultConfiger interface {
	VaultConfig() *VaultConfig
	RegistryConfig() *RegistryConfig
}

type VaultConfig struct {
	Endpoint        string `json:"endpoint"`
	Token           string `json:"token"`
	MountPath       string `json:"mount_path"`
	PrivateKeyPath  string `json:"private_key_path"`
	AddressRegistry string `json:"address_registry"`
	SecretPath      string `json:"secret_path"`
}

type RegistryConfig struct {
	Address common.Address `json:"address"`
}

func NewVaultConfiger(getter kv.Getter) VaultConfiger {
	return &vaultConfig{
		getter: getter,
	}
}

type vaultConfig struct {
	getter    kv.Getter
	once      comfig.Once
	vaultOnce comfig.Once
}

func (c *vaultConfig) VaultConfig() *VaultConfig {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "vault")
		config := VaultConfig{}
		err := figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &config
	}).(*VaultConfig)
}

func (c *vaultConfig) RegistryConfig() *RegistryConfig {
	rg := RegistryConfig{}
	rg.Address = c.getContractAddress()
	return &rg
}

func (c *vaultConfig) getContractAddress() common.Address {
	config := struct {
		Addresses   map[string]string `fig:"addresses"`
		ProjectName string            `fig:"projectName"`
		StartBlock  int               `json:"startBlock"`
	}{}

	err := c.getVaultSecret(c.VaultConfig().SecretPath, &config)
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress(config.Addresses["MasterContractsRegistry"])

	return address
}

func (c *vaultConfig) getVaultSecret(key string, out interface{}) error {
	vaultClient := c.vaultClient()
	secret, err := vaultClient.KVv2(c.VaultConfig().MountPath).Get(context.Background(), key)
	if err != nil {
		return errors.Wrap(err, "failed to get secret data")
	}

	return figure.
		Out(out).
		With(figure.BaseHooks, VaultHook).
		From(secret.Data).
		Please()
}

func (c *vaultConfig) vaultClient() *vault.Client {
	return c.vaultOnce.Do(func() interface{} {
		conf := vault.DefaultConfig()
		conf.Address = c.VaultConfig().Endpoint
		client, err := vault.NewClient(conf)
		if err != nil {
			log.Panicf("unable to initialize Vault client: %v", err)
		}
		token := c.VaultConfig().Token

		client.SetToken(token)

		return client
	}).(*vault.Client)
}

var VaultHook = figure.Hooks{
	"map[string]interface{}": func(value interface{}) (reflect.Value, error) {
		if value == nil {
			return reflect.Value{}, nil
		}

		var params map[string]interface{}
		switch s := value.(type) {
		case map[interface{}]interface{}:
			params = make(map[string]interface{})
			for key, value := range s {
				params[key.(string)] = value
			}
		case map[string]interface{}:
			params = s
		default:
			return reflect.Value{}, errors.New("unexpected type while figure map[string]interface{}")
		}

		return reflect.ValueOf(params), nil
	},
	"map[string]string": func(value interface{}) (reflect.Value, error) {
		if value == nil {
			return reflect.Value{}, nil
		}

		var params map[string]string
		switch s := value.(type) {
		case map[interface{}]interface{}:
			params = make(map[string]string)
			for key, value := range s {
				params[key.(string)] = value.(string)
			}
		case map[interface{}]string:
			params = make(map[string]string)
			for key, value := range s {
				params[key.(string)] = value
			}
		case map[string]interface{}:
			params = make(map[string]string)
			for key, value := range s {
				params[key] = value.(string)
			}
		case map[string]string:
			params = s

		default:
			return reflect.Value{}, errors.New("unexpected type while figure map[string]interface{}")
		}

		return reflect.ValueOf(params), nil
	},
}
