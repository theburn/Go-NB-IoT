package device

import (
	"Go-NB-IoT/client"
	"Go-NB-IoT/configure"
	log "Go-NB-IoT/logging"
	"encoding/json"
	"fmt"
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

func (d *DeviceCredentials) RegisterDevice(c *client.NBHttpClient) error {

	reqParam := client.RequestParam{}
	reqParam.URL = configure.NBIoTConfig.ReqParam.IoTHost +
		registerDeviceURI + configure.NBIoTConfig.ReqParam.AppID

	reqParam.Method = "POST"
	reqParam.ContentType = "application/json"

	var err error

	if reqParam.ReqBody, err = json.Marshal(d); err != nil {
		log.Error("json Marshal Failed, ", d, err)
		return err
	}

	err = c.Request(&reqParam)

	fmt.Println(string(reqParam.RespBody))

	return err
}
