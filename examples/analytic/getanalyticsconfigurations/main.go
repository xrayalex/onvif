package main

import (
	"github.com/xrayalex/onvif"
	"github.com/xrayalex/onvif/media2"
	"io/ioutil"
	"log"
)

func main() {
	dev, err := onvif.NewDevice(onvif.DeviceParams{
		Xaddr:    "192.168.12.148", // BOSCH
		Username: "administrator",
		Password: "Password1!",
		AuthMode: onvif.UsernameTokenAuth,
	})
	if err != nil {
		log.Fatalln("fail to new device:", err)
	}

	res, err := dev.CallMethod(media2.GetAnalyticsConfigurations{})
	if err != nil {
		log.Fatalln("fail to CallMethod:", err)
	}
	bs, _ := ioutil.ReadAll(res.Body)

	log.Printf(">> Result: %+v \n %s", res.StatusCode, bs)
}
