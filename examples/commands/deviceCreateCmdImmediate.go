package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/theburn/Go-NB-IoT/client"
	"github.com/theburn/Go-NB-IoT/commands"
	"github.com/theburn/Go-NB-IoT/configure"
	log "github.com/theburn/Go-NB-IoT/logging"
)

var (
	version        = "v1.0"
	configfilename = flag.String("cfgfile", "conf/config.json", "live configure filename")
	loglevel       = flag.String("loglevel", "info", "log level")
	logfile        = flag.String("logfile", "logs/go-nb-iot.log", "log file path")
	deviceID       = flag.String("d", "", "deviceID")
	serviceID      = flag.String("s", "", "serviceID")
	method         = flag.String("m", "", "Method")
	paras          = flag.String("p", "", "paras")
	callbackUrl    = flag.String("c", "", "CallbackUrl")
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

	cmdReq := commands.DeviceCreateCmdReq{}
	cmdReq.DeviceID = *deviceID
	cmdReq.ExpireTime = 2 * 60 * 60
	cmdReq.MaxRetransmit = 3
	cmdReq.Command = commands.CommandDTO{}
	cmdReq.Command.Method = *method
	cmdReq.Command.ServiceId = *serviceID

	err = json.Unmarshal([]byte(*paras), &cmdReq.Command.Paras)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Debugf("req:%+v", cmdReq)

	if c, err := client.NewNBHttpClient(); err != nil {
		log.Error("New Client error", err)
	} else {
		c.Login()
		resp, err := cmdReq.DeviceCreateCmdBusiness(c)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("%v\n", resp)

	}

}
