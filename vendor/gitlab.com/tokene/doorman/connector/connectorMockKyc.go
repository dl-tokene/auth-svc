package connector

import (
	"net/http"
	"strings"

	"gitlab.com/tokene/doorman/resources"
)

type ConnectorMockKyc struct {
	ServiceUrl string
}

func NewConnectorMockKyc(serviceUrl string) ConnectorI {
	return ConnectorMockKyc{
		ServiceUrl: serviceUrl,
	}
}

func (c ConnectorMockKyc) GenerateJwtPair(address string, purpose string) (resources.JwtPairResponse, error) {
	return resources.JwtPairResponse{}, nil
}

func (c ConnectorMockKyc) ValidateJwt(token string) (string, error) {
	return strings.ToLower("0x750Bd531CEA1f68418DDF2373193CfbD86A69058"), nil
}

func (c ConnectorMockKyc) RefreshJwt(refreshToken string) (resources.JwtPairResponse, error) {
	return resources.JwtPairResponse{}, nil
}
func (c ConnectorMockKyc) GetAuthToken(r *http.Request) (string, error) {
	return "token", nil
}
func (c ConnectorMockKyc) CheckPermission(owner string, token string) error {
	return nil
}
func (c ConnectorMockKyc) CheckPermissionID(id, resource, token string) error {
	return nil
}
func (c ConnectorMockKyc) CheckPurpose(token string) (string, error) {
	return "session", nil
}
