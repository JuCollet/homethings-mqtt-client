package mqtt

import (
	"encoding/json"
	"home-things/internal"
	"home-things/pkg/services"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"golang.org/x/exp/slices"
)

type Payload struct {
	EndDeviceIds struct {
		DeviceId string `json:"device_id"`
	} `json:"end_device_ids"`
	Message struct {
		DecodedPayload map[string]interface{} `json:"decoded_payload"`
	} `json:"uplink_message"`
}

func MessageHandler(client mqtt.Client, msg mqtt.Message) {
	var payload Payload

	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		return

	}

	devicesDataTypes := internal.GetDevicesDataTypes()

	for key, value := range payload.Message.DecodedPayload {
		idx := slices.IndexFunc(devicesDataTypes, func(d services.DeviceDataType) bool {
			return d.DeviceId == payload.EndDeviceIds.DeviceId && d.Key == key
		})

		if idx == -1 {
			continue
		}

		deviceDataType := devicesDataTypes[idx]
		var val float64

		switch v := value.(type) {
		case float64:
			val = v
		case int:
			val = float64(v)
		}

		services.CreateData(services.Data{
			UserDeviceID:     int32(deviceDataType.UserDeviceId),
			DeviceDataTypeId: int32(deviceDataType.DeviceDataTypeId),
			Value:            float32(val),
		})

	}

}
