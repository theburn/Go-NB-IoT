package client

import (
	"Go-NB-IoT/common"
	"Go-NB-IoT/configure"
	log "Go-NB-IoT/logging"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

type NBHttpClient struct {
	client   *fasthttp.Client
	authInfo loginResponse
}

type loginResponse struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int64  `json:"expiresIn"`
	Scope        string `json:"scope"`
}

func NewNBHttpClient() (*NBHttpClient, error) {
	cert, err := tls.LoadX509KeyPair(configure.NBIoTConfig.ReqParam.CertFile,
		configure.NBIoTConfig.ReqParam.KeyFile)
	if err != nil {
		return nil, err
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	tlsConfig.BuildNameToCertificate()

	c := &NBHttpClient{}
	c.client = &fasthttp.Client{
		TLSConfig:           tlsConfig,
		MaxConnsPerHost:     200,
		MaxIdleConnDuration: 120 * time.Second,
	}

	return c, nil

}

func (c *NBHttpClient) Login() error {
	args := fasthttp.AcquireArgs()
	args.Add("appId", configure.NBIoTConfig.ReqParam.AppID)
	args.Add("secret", configure.NBIoTConfig.ReqParam.Secret)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(configure.NBIoTConfig.ReqParam.IoTHost + common.LoginURI)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.SetBody(args.QueryString())

	resp := fasthttp.AcquireResponse()

	if err := c.client.Do(req, resp); err != nil {
		log.Error("Login is Failed!", err)
		return err
	}

	if err := json.Unmarshal(resp.Body(), &c.authInfo); err != nil {
		log.Errorf("resp body unmarshal failed! %s, %s", string(resp.Body()), err)
		return err
	}

	fmt.Println(c.authInfo)

	return nil

}
