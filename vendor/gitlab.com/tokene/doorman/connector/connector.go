package connector

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	"gitlab.com/tokene/doorman/resources"
)

type Connector struct {
	ServiceUrl string
	Client     *http.Client
}

func NewConnector(serviceUrl string) ConnectorI {
	return Connector{
		ServiceUrl: serviceUrl,
		Client: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}
func (c Connector) DoAuthRequest(method string, url string, token string, body interface{}, expectedStatus int) (*http.Response, error) {
	postBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	req.Header.Set("Authorization", "Bearer "+token)

	response, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	if expectedStatus != response.StatusCode {
		response.Body.Close()
		return nil, errors.New("Bad status")
	}
	return response, nil
}
func (c Connector) DoAuthRequestWithDecode(method string, url string, token string, body interface{}, decodeModel interface{}, expectedStatus int) error {
	response, err := c.DoAuthRequest(method, url, token, body, expectedStatus)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(decodeModel); err != nil {
		return errors.Wrap(err, "failed to unmarshal")
	}

	return nil
}

func (c Connector) GenerateJwtPair(address string, purpose string) (resources.JwtPairResponse, error) {
	model := resources.JwtPairResponse{}
	err := c.DoAuthRequestWithDecode("POST", c.ServiceUrl+"/get_token_pair", "", NewClaimsModel(address, purpose), &model, http.StatusOK)
	return model, err
}

func (c Connector) ValidateJwt(token string) (string, error) {
	model := resources.JwtValidationResponse{}
	err := c.DoAuthRequestWithDecode("GET", c.ServiceUrl+"/validate_token", token, nil, &model, http.StatusOK)
	return model.Data.Attributes.EthAddress, err
}
func (c Connector) RefreshJwt(refreshToken string) (resources.JwtPairResponse, error) {
	model := resources.JwtPairResponse{}
	err := c.DoAuthRequestWithDecode("GET", c.ServiceUrl+"/refresh_token", refreshToken, nil, &model, http.StatusOK)
	return model, err
}
func (c Connector) GetAuthToken(r *http.Request) (string, error) {
	return helpers.Authenticate(r)
}

func (c Connector) CheckPermission(owner string, token string) error {
	response, err := c.DoAuthRequest("GET", c.ServiceUrl+"/check_permission/"+owner, token, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	response.Body.Close()
	return nil
}
func (c Connector) CheckPurpose(token string) (string, error) {
	model := resources.Purpose{}
	err := c.DoAuthRequestWithDecode("GET", c.ServiceUrl+"/check_purpose", token, nil, &model, http.StatusOK)
	if err != nil {
		return "", err
	}
	return model.Type, nil
}
