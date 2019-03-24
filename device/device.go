package device

import (
	"Go-NB-IoT/client"
	"Go-NB-IoT/configure"
	log "Go-NB-IoT/logging"
	"encoding/json"
)

type DeviceCredentials struct {
	VerifyCode string `json:"verifyCode"`
	NodeId     string `json:"nodeId"`
	EndUserId  string `json:"endUserId"`
	Timeout    int    `json:"timeout"`
	IsSecure   bool   `json:"isSecure"`
}

type DataConfigDTO struct {
	DataAgingTime int `json:"dataAgingTime"` // 0 - 90 Day
}

type DeviceConfigDTO struct {
	DataConfig DataConfigDTO `json:"dataConfig"`
}

type DeviceProfile struct {
	Name             string          `json:"name"`
	EndUser          string          `json:"endUser"`
	Mute             string          `json:"mute"` // TRUE , FALSE
	ManufacturerID   string          `json:"manufacturerId"`
	ManufacturerName string          `json:"manufacturerName"`
	DeviceType       string          `json:"deviceType"`
	Model            string          `json:"model"`
	Location         string          `json:"location"`
	ProtocolType     string          `json:"protocolType"`
	DeviceConfig     DeviceConfigDTO `json:"deviceConfig"`
	Region           string          `json:"region"`
	Organization     string          `json:"organization"`
	TimeZone         string          `json:"timezone"`
	IsSecure         bool            `json:"isSecure"`
	Psk              string          `json:"psk"`
}

type DeviceIdInfo struct {
	DeviceId   string `json:"deviceId"`
	VerifyCode string `json:"verifyCode"`
	Timeout    int    `json:"timeout"`
	Psk        string `json:"psk"`
}

func (d *DeviceCredentials) RegisterDevice(c *client.NBHttpClient) (*DeviceIdInfo, error) {

	reqRespParam := client.ReqRespParam{}
	reqRespParam.URL = configure.NBIoTConfig.ReqParam.IoTHost +
		registerDeviceURI + configure.NBIoTConfig.ReqParam.AppID

	reqRespParam.Method = "POST"
	reqRespParam.ContentType = "application/json"

	var err error

	if reqRespParam.ReqBody, err = json.Marshal(d); err != nil {
		log.Error("json Marshal Failed, ", d, err)
		return nil, err
	}

	if err = c.Request(&reqRespParam); err != nil {
		log.Error("Request error!", err)
		return nil, err
	}

	var deviceIdInfo DeviceIdInfo

	if err = json.Unmarshal(reqRespParam.RespBody, &deviceIdInfo); err != nil {
		log.Error("json DeviceIdInfo Unmarshal Failed, ", reqRespParam.RespBody, err)
		return nil, err
	}

	return &deviceIdInfo, err
}
