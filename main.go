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
	mqtt.Init(devices)

	select {}

}
