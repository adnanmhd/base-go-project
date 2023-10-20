package service

import (
	"base-go-project/util"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type ThirdPartyServiceConfigImpl struct {
	client *resty.Client
	config util.ThirdPartyServiceConfig
}

func New(client *resty.Client, conf util.ThirdPartyServiceConfig) ThirdPartyServiceConfigImpl {
	return ThirdPartyServiceConfigImpl{
		client: client,
		config: conf,
	}
}

func (d ThirdPartyServiceConfigImpl) GetToken() (*AuthResponse, error) {
	authResponse := new(AuthResponse)
	resp, err := d.client.
		R().
		SetHeaders(map[string]string{
			"Content-Type":  "application/x-www-form-urlencoded",
			"cache-control": "no-cache",
			"Accept":        "application/json",
		}).
		SetFormData(map[string]string{
			"client_id":     "",
			"client_secret": "",
			"grant_type":    "",
			"username":      "",
			"password":      "",
		}).
		SetResult(authResponse).
		//SetError(Error{}).
		Post("URL")

	if resp.IsError() {
		log.Errorf("Error get token, http status code: %d, response body: %s", resp.RawResponse.StatusCode, string(resp.Body()))
		err = errors.New(fmt.Sprintf("error %d", resp.RawResponse.StatusCode))
	}

	return authResponse, err
}

func (d ThirdPartyServiceConfigImpl) send(url, token string, reqByte []byte) (*Payload, error) {
	log.Info("request to digital-service: ", string(reqByte))

	responseData := new(Payload)

	// send
	req := d.client.R().SetBody(reqByte).SetResult(responseData).SetAuthToken(token)
	resp, err := req.Execute(resty.MethodPost, url)

	if resp.IsError() {
		log.Errorf("Error hit to digital-service, http status code: %d, response body: %s", resp.RawResponse.StatusCode, string(resp.Body()))
		err = errors.New(fmt.Sprintf("error %d", resp.RawResponse.StatusCode))
	}

	log.Info("response from digital-service: ", string(resp.Body()))
	return responseData, err
}
