package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	goonvif "github.com/xrayalex/onvif"
	"github.com/xrayalex/onvif/device"
	"github.com/xrayalex/onvif/gosoap"
	"github.com/xrayalex/onvif/xsd/onvif"
)

const (
	login    = "admin"
	password = "Supervisor"
)

func readResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.13.14:80",
		Username:   login,
		Password:   password,
		HttpClient: new(http.Client),
	})
	if err != nil {
		panic(err)
	}

	//Preparing commands
	UserLevel := onvif.UserLevel("User")
	systemDateAndTyme := device.GetSystemDateAndTime{}
	getCapabilities := device.GetCapabilities{Category: []onvif.CapabilityCategory{"All"}}
	createUser := device.CreateUsers{
		User: []onvif.UserRequest{
			{
				Username:  "TestUser",
				Password:  "TestPassword",
				UserLevel: &UserLevel,
			},
		},
	}

	//Commands execution
	systemDateAndTymeResponse, err := dev.CallMethod(systemDateAndTyme)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(readResponse(systemDateAndTymeResponse))
	}
	getCapabilitiesResponse, err := dev.CallMethod(getCapabilities)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(readResponse(getCapabilitiesResponse))
	}
	createUserResponse, err := dev.CallMethod(createUser)
	if err != nil {
		log.Println(err)
	} else {
		/*
			You could use https://github.com/xrayalex/onvif/gosoap for pretty printing response
		*/
		fmt.Println(gosoap.SoapMessage(readResponse(createUserResponse)).StringIndent())
	}

}
