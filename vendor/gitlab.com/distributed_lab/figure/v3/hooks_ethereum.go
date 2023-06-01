package figure

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"reflect"
)

var (
	// EthereumHooks set of default hooks for common types of Ethereum
	EthereumHooks = Hooks{
		"common.Address": func(value interface{}) (reflect.Value, error) {
			switch v := value.(type) {
			case string:
				if !common.IsHexAddress(v) {
					// provide value does not look like valid address
					return reflect.Value{}, errors.New("invalid address")
				}
				return reflect.ValueOf(common.HexToAddress(v)), nil
			default:
				return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
			}
		},
		"*ecdsa.PrivateKey": func(raw interface{}) (reflect.Value, error) {
			switch value := raw.(type) {
			case string:
				kp, err := crypto.HexToECDSA(value)
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to init keypair")
				}
				return reflect.ValueOf(kp), nil
			default:
				return reflect.Value{}, fmt.Errorf("cant init keypair from type: %T", value)
			}
		},
	}
)
