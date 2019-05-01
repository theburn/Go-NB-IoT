package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/theburn/Go-NB-IoT/client"
	"github.com/theburn/Go-NB-IoT/configure"
	"github.com/theburn/Go-NB-IoT/device"
	log "github.com/theburn/Go-NB-IoT/logging"
)

var (
	version        = "v1.0"
	configfilename = flag.String("cfgfile", "conf/config.json", "live configure filename")
	loglevel       = flag.String("loglevel", "info", "log level")
	logfile        = flag.String("logfile", "logs/go-nb-iot.log", "log file path")
)

func init() {
	flag.Parse()
	log.SetOutputByName(*logfile)
	log.SetRotateByDay()
	log.SetLevelByString(*loglevel)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("go-nb-iot panic: ", r)
			time.Sleep(1 * time.Second)
		}
	}()
	err := configure.LoadConfig(*configfilename)
	if err != nil {
		return
	}

	log.SetMaxLogDay(7)

	// output system info
	log.Info("-----------------START----------------")
	log.Info("start go-nb-iot: ", version)

	d := device.DeviceCredentials{
		VerifyCode: "359369080878768",
		NodeId:     "359369080878768",
		EndUserId:  "112233445",
		Timeout:    0,
		IsSecure:   false,
	}

	dataConfig := device.DataConfigDTO{
		DataAgingTime: 7,
	}

	deviceConfig := device.DeviceConfigDTO{
		DataConfig: dataConfig,
	}

	p := device.DeviceProfile{
		Name:             "SZDTestDevice001",
		EndUser:          "SZD",
		Mute:             "FALSE",
		ManufacturerID:   "SZDTS001",
		ManufacturerName: "SZD",
		DeviceType:       "Water",
		Model:            "SZDTSDevice",
		Location:         "Shanghai",
		ProtocolType:     "CoAP",
		DeviceConfig:     deviceConfig,
		Region:           "Shanghai",
		Organization:     "SZDTest",
		TimeZone:         "Asia/Shanghai",
		IsSecure:         false,
		//Psk:              "00aaAA11bbBB",
	}

	if c, err := client.NewNBHttpClient(); err != nil {
		log.Error("New Client error", err)
	} else {
		c.Login()
		d, _ := d.RegisterDevice(c)
		fmt.Printf("%v\n", d)

		d.ModifyDeviceInfo(c, p)

	}

}
