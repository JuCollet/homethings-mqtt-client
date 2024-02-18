package internal

import (
	"home-things/pkg/services"
	"home-things/utils"
	"slices"
)

var DevicesDataTypes []services.DeviceDataType

func FetchDevicesDataTypes() error {
	devicesDataTypes, err := services.GetDevicesDataTypes()

	if err != nil {
		panic(err)
	}

	DevicesDataTypes = devicesDataTypes

	return nil
}

func GetDevicesDataTypes() []services.DeviceDataType {
	return DevicesDataTypes
}

func AddDevicesDataTypes(devicesDataType services.DeviceDataType) {
	idx := slices.IndexFunc(DevicesDataTypes, func(d services.DeviceDataType) bool {
		return d.DeviceId == devicesDataType.DeviceId && d.Key == devicesDataType.Key
	})

	if idx != -1 {
		return
	}

	DevicesDataTypes = append(DevicesDataTypes, devicesDataType)
}

func GetDevices() []services.Device {
	uniqueDataTypes := utils.GetUniqueDevices(DevicesDataTypes)
	devices := []services.Device{}

	for _, v := range uniqueDataTypes {
		devices = append(devices, services.Device{
			DeviceId:     v.DeviceId,
			UserDeviceId: v.UserDeviceId,
		})
	}

	return devices
}
