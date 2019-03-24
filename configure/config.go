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

type Config struct {
	ReqParam ReqParam `json:"req_param"`
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
