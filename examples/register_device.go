package main

import (
	"flag"
	"fmt"
	"time"

	"Go-NB-IoT/client"
	"Go-NB-IoT/configure"
	"Go-NB-IoT/device"
	log "Go-NB-IoT/logging"
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
		VerifyCode: "9af17b52-4393-426e-958b-21ba27ebcf29",
		NodeId:     "9af17b52-4393-426e-958b-21ba27ebcf29",
		EndUserId:  "112233445",
		Timeout:    0,
		IsSecure:   false,
	}

	if c, err := client.NewNBHttpClient(); err != nil {
		log.Error("New Client error", err)
	} else {
		c.Login()
		d, _ := d.RegisterDevice(c)
		fmt.Println(d)

	}

}