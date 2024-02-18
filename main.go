package main

import (
	"home-things/internal"
	"home-things/internal/mqtt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	internal.FetchDevicesDataTypes()

	devices := internal.GetDevices()
	client := mqtt.Init(devices)
	mqtt.ListenForDevices(client)

	select {}

}
