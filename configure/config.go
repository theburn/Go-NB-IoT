package configure

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/theburn/Go-NB-IoT/logging"
)

/* from gwuhaolin/livego configure */

/*
{
    "server_param" {
        "listen_port":  "9880",
		"static_path": "/usr/local/github.com/theburn/Go-NB-IoT/static"
    },
    "req_param" : {
        "cert_file": "certs/server.crt",
        "key_file": "certs/server.key",
        "iot_host": "https://180.101.147.89:8743",
        "app_id": "<AppID>",
        "secret": "<SECRET>"
    },
	"amqp_param" :{
		"amqp_url":"amqp://user:pass@127.0.0.1:5672/"
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

type ServerParam struct {
	ListenPort string `json:"listen_port"`
	StaticPath string `json:"static_path"`
}

type Config struct {
	ReqParam    ReqParam    `json:"req_param"`
	ServerParam ServerParam `json:"server_param"`
}

var NBIoTConfig Config

func LoadConfig(configfilename string) error {
	log.Infof("starting load configure file(%s)......", configfilename)
	data, err := ioutil.ReadFile(configfilename)
	if err != nil {
		log.Errorf("ReadFile %s error:%v", configfilename, err)
		return err
	}

	log.Debugf("loadconfig: \r\n%s", string(data))

	err = json.Unmarshal(data, &NBIoTConfig)
	if err != nil {
		log.Errorf("json.Unmarshal error:%v", err)
		return err
	}
	log.Debugf("get config json data:%v", NBIoTConfig)
	return nil
}
