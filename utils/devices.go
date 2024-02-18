package utils

import "home-things/pkg/services"

func GetUniqueDevices(devices []services.DeviceDataType) []services.DeviceDataType {
	seen := make(map[string]bool)
	uniqueDevices := []services.DeviceDataType{}

	for _, d := range devices {
		if _, ok := seen[d.DeviceId]; !ok {
			seen[d.DeviceId] = true
			uniqueDevices = append(uniqueDevices, d)
		}
	}

	return uniqueDevices
}
