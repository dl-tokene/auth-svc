package helpers

import (
	"errors"
	"regexp"
)

var AddressRegexp = regexp.MustCompile("^(0x)?[0-9a-fA-F]{40}$")
var Purposes = []string{"session"}

func ValidatePurposes(value interface{}) error {
	str, _ := value.(string)
	for _, purpose := range Purposes {
		if purpose == str {
			return nil
		}
	}
	return errors.New("invalid purpose")
}
