package mqtt

import (
	"fmt"
	"home-things/internal"
	"home-things/pkg/services"
	"home-things/utils"
	"slices"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ListenForDevices(client mqtt.Client) {

	topic := "devices"

	handler := func(message *kafka.Message) {
		userDeviceId, _ := strconv.Atoi(string(message.Value))

		userDeviceDataTypes, err := services.GetDeviceDataTypes(int32(userDeviceId))

		if err != nil {
			return
		}

		userDeviceDeviceId := userDeviceDataTypes[0].DeviceId
		devices := internal.GetDevicesDataTypes()

		idx := slices.IndexFunc(devices, func(d services.DeviceDataType) bool {
			return d.DeviceId == userDeviceDeviceId
		})

		if idx != -1 {
			return
		}

		if conn := client.Subscribe(fmt.Sprintf("v3/the-home-things@ttn/devices/%v/up", userDeviceDeviceId), 0, nil); conn.Wait() {
			fmt.Println("Subscribed to up-messages of device", userDeviceDeviceId)
			if conn.Error() != nil {
				fmt.Println(conn.Error())
			}
		}

		for _, newDeviceDataType := range userDeviceDataTypes {
			internal.AddDevicesDataTypes(newDeviceDataType)
		}

	}

	go utils.Subscribe(topic, handler)
}
