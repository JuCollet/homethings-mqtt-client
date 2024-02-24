package services

import (
	"encoding/json"
	"home-things/constants"
	"io"
	"net/http"
	"os"
)

type Device struct {
	DeviceId     string `json:"device_id"`
	UserDeviceId int    `json:"user_device_id"`
}

type DeviceDataType struct {
	Device
	DeviceDataTypeId int    `json:"device_data_type_id"`
	Key              string `json:"key"`
}

func getDataTypes(req *http.Request) (device []DeviceDataType, err error) {
	client := &http.Client{}
	req.Header.Set(constants.HEADER_API_KEY, os.Getenv("API_KEY"))
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	d := []DeviceDataType{}

	deviceErr := json.Unmarshal(body, &d)

	if deviceErr != nil {
		return nil, deviceErr
	}

	return d, nil
}

func GetDevicesDataTypes() (device []DeviceDataType, err error) {
	url := os.Getenv("API_BASE_URL") + "/devices-data-types"
	req, _ := http.NewRequest("GET", url, nil)
	return getDataTypes(req)
}

func GetDeviceDataTypes(deviceId int32) (device []DeviceDataType, err error) {
	url := os.Getenv("API_BASE_URL") + "/device-data-types/" + string(deviceId)

	req, _ := http.NewRequest("GET", url, nil)
	return getDataTypes(req)
}
