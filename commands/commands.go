package commands

import (
	"encoding/json"

	"github.com/theburn/Go-NB-IoT/client"
	"github.com/theburn/Go-NB-IoT/configure"
	log "github.com/theburn/Go-NB-IoT/logging"
)

type CommandDTO struct {
	ServiceId string          `json:"serviceId"`
	Method    string          `json:"method"`
	Paras     json.RawMessage `json:"paras"`
}

type DeviceCreateCmdReq struct {
	DeviceID string     `json:"deviceId"`
	Command  CommandDTO `json:"command"`
	//CallbackUrl   string     `json:"callbackUrl"`
	ExpireTime    int `json:"expireTime"`
	MaxRetransmit int `json:"maxRetransmit"`
}

type DeviceCreateCmdResp struct {
	CommandId          string     `json:"commandId"`
	AppId              string     `json:"appId"`
	DeviceId           string     `json:"deviceId"`
	Command            CommandDTO `json:"command"`
	CallbackUrl        string     `json:"callbackUrl"`
	ExpireTime         int        `json:"expireTime"`
	Status             string     `json:"status"`
	CreationTime       string     `json:"creationTime"`
	ExecuteTime        string     `json:"executeTime"`
	PlatformIssuedTime string     `json:"platformIssuedTime"`
	DeliveredTime      string     `json:"deliveredTime"`
	IssuedTimes        int        `json:"issuedTimes"`
	MaxRetransmit      int        `json:"maxRetransmit"`
}

func (s *DeviceCreateCmdReq) DeviceCreateCmdBusiness(c *client.NBHttpClient) (*DeviceCreateCmdResp, error) {

	reqRespParam := client.ReqRespParam{}
	reqRespParam.URL = configure.NBIoTConfig.ReqParam.IoTHost + createDeviceCmdURI

	reqRespParam.Method = "POST"
	reqRespParam.ContentType = "application/json"

	var err error

	if reqRespParam.ReqBody, err = json.Marshal(s); err != nil {
		log.Error("json Marshal Failed, ", s, err)
		return nil, err
	}

	if err = c.Request(&reqRespParam); err != nil {
		log.Error("Request error!", err)
		return nil, err
	}

	log.Debugf("reqBody: %+v", string(reqRespParam.ReqBody))

	var cmdResp DeviceCreateCmdResp

	if err = json.Unmarshal(reqRespParam.RespBody, &cmdResp); err != nil {
		log.Error("json SubscriptionsBusinessResp Unmarshal Failed, ", reqRespParam.RespBody, err)
		return nil, err
	}

	log.Debugf("respBody: %v", string(reqRespParam.RespBody))
	return &cmdResp, err
}
