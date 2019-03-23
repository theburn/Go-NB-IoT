package configure

import (
	log "Go-NB-IoT/logging"
	"encoding/json"
	"io/ioutil"
)

/* from gwuhaolin/livego configure */

/*
{
"req_param" : {
    "cert_file": "certs/server.crt",  # from IoT-platform
    "key_file": "certs/server.key",   # from IoT-platform
    "iot_host": "https://180.101.147.89:8743",
    "app_id": "<APP ID>",
    "secret": "<SECRET>"
    },
"profile_info": {
    "manufacturer_name": "NAME",
    "manufacturer_id": "ID",
    "end_user_id": "END_USER_ID",
    "location": "Unknown",
    "device_type": "DEV_TYPE",
    "device_model": "DEV_MODEL"
    }
}
*/

type ReqParam struct {
	CertFile string `json:"cert_file"`
	KeyFile  string `json:"key_file"`
	IoTHost  string `json:"iot_host"`
	AppID    string `json:"app_id"`
	Secret   string `json:"secret"`
}

type ProfileInfo struct {
	ManufacturerName string `json:"manufacturer_name"`
	ManufacturerID   string `json:"manufacturer_id"`
	EndUserID        string `json:"end_user_id"`
	Location         string `json:"location"`
	DeviceType       string `json:"device_type"`
	DeviceModel      string `json:"device_model"`
}

type Config struct {
	ReqParam    ReqParam    `json:"req_param"`
	ProfileInfo ProfileInfo `json:"profile_info"`
}

var NBIoTConfig Config

func LoadConfig(configfilename string) error {
	log.Infof("starting load configure file(%s)......", configfilename)
	data, err := ioutil.ReadFile(configfilename)
	if err != nil {
		log.Errorf("ReadFile %s error:%v", configfilename, err)
		return err
	}

	log.Infof("loadconfig: \r\n%s", string(data))

	err = json.Unmarshal(data, &NBIoTConfig)
	if err != nil {
		log.Errorf("json.Unmarshal error:%v", err)
		return err
	}
	log.Infof("get config json data:%v", NBIoTConfig)
	return nil
}
