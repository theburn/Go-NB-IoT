package client

import (
	"github.com/theburn/Go-NB-IoT/configure"
	log "github.com/theburn/Go-NB-IoT/logging"
	"crypto/tls"
	"encoding/json"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

type NBHttpClient struct {
	client   *fasthttp.Client
	createAt time.Time
	authInfo loginResponse
	rwLock   sync.RWMutex
}

type loginResponse struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int64  `json:"expiresIn"`
	Scope        string `json:"scope"`
}

type ReqRespParam struct {
	URL         string
	Method      string `default:"POST"`
	ContentType string `default:"application/json"`
	ReqBody     []byte
	RespBody    []byte
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
		MaxIdleConnDuration: 60 * time.Second,
	}

	return c, nil

}

func (c *NBHttpClient) Login() error {
	args := fasthttp.AcquireArgs()
	args.Add("appId", configure.NBIoTConfig.ReqParam.AppID)
	args.Add("secret", configure.NBIoTConfig.ReqParam.Secret)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(configure.NBIoTConfig.ReqParam.IoTHost + loginURI)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.SetBody(args.QueryString())

	resp := fasthttp.AcquireResponse()

	if err := c.client.Do(req, resp); err != nil {
		log.Error("Login is Failed!", err)
		return err
	}

	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	if err := json.Unmarshal(resp.Body(), &c.authInfo); err != nil {
		log.Errorf("resp body unmarshal failed! %s, %s", string(resp.Body()), err)
		return err
	}

	c.createAt = time.Now()

	log.Info("Login Successed!")

	return nil

}

func (c *NBHttpClient) RefreshToken() error {

	c.rwLock.RLock()
	jsonArgs := map[string]string{
		"appId":        configure.NBIoTConfig.ReqParam.AppID,
		"secret":       configure.NBIoTConfig.ReqParam.Secret,
		"refreshToken": c.authInfo.RefreshToken,
	}
	c.rwLock.RUnlock()

	reqBody, _ := json.Marshal(jsonArgs)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(configure.NBIoTConfig.ReqParam.IoTHost + refreshTokenURI)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBody(reqBody)

	resp := fasthttp.AcquireResponse()

	if err := c.client.Do(req, resp); err != nil {
		log.Error("Refresh is Failed!", err)
		return err
	}

	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	if err := json.Unmarshal(resp.Body(), &c.authInfo); err != nil {
		log.Errorf("resp body unmarshal failed! %s, %s", string(resp.Body()), err)
		return err
	}

	c.createAt = time.Now()

	log.Info("RefreshToken Successed!")

	return nil

}

func (c *NBHttpClient) tokenIsExpire() bool {

	// if the time of token is  less than 5 min.  refresh it.
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()

	if int64(time.Now().Sub(c.createAt).Seconds())+300 > c.authInfo.ExpiresIn {
		return true
	}

	return false
}

func (c *NBHttpClient) Request(reqRespParam *ReqRespParam) error {

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(reqRespParam.URL)
	req.Header.SetMethod(reqRespParam.Method)
	req.Header.SetContentType(reqRespParam.ContentType)
	req.SetBody(reqRespParam.ReqBody)

	req.Header.Add("app_key", configure.NBIoTConfig.ReqParam.AppID)
	req.Header.Add("Authorization", "Bearer "+c.authInfo.AccessToken)

	resp := fasthttp.AcquireResponse()

	if c.tokenIsExpire() {
		c.RefreshToken()
	}

	if err := c.client.Do(req, resp); err != nil {
		log.Error("Request is Failed!", err)
		return err
	}

	reqRespParam.RespBody = resp.Body()
	return nil
}
