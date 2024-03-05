package mqtt

import (
	"fmt"
	"home-things/pkg/services"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Init(devices []services.Device) mqtt.Client {
	options := mqtt.NewClientOptions().AddBroker(os.Getenv("MQTT_BROKER_HOST"))
	options.SetUsername(os.Getenv("MQTT_USER_NAME"))
	options.SetPassword(os.Getenv("MQTT_TOKEN"))
	options.SetDefaultPublishHandler(MessageHandler)
	options.AutoReconnect = true

	options.OnConnect = func(c mqtt.Client) {
		fmt.Println("Connected to TTN Broker")

		for _, d := range devices {
			go func(device services.Device) {

				if conn := c.Subscribe(fmt.Sprintf("v3/the-home-things@ttn/devices/%v/up", device.DeviceId), 0, nil); conn.Wait() {
					fmt.Println("Subscribed to up-messages of device", device.DeviceId)
					if conn.Error() != nil {
						fmt.Println(conn.Error())
					}
				}

			}(d)

		}

	}

	options.OnConnectionLost = func(c mqtt.Client, err error) {
		fmt.Println("Connection lost")
		fmt.Println(err.Error())
	}

	client := mqtt.NewClient(options)

	if conn := client.Connect(); conn.Wait() && conn.Error() != nil {
		panic(conn.Error())
	}

	return client
}
