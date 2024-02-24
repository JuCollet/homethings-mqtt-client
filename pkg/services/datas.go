package services

import (
	"bytes"
	"encoding/json"
	"home-things/constants"
	"io"
	"net/http"
	"os"
)

type Data struct {
	UserDeviceID     int32   `json:"user_device_id"`
	DeviceDataTypeId int32   `json:"device_data_type_id"`
	Value            float32 `json:"value"`
}

func CreateData(data Data) (err error) {
	url := os.Getenv("API_BASE_URL") + "/data"
	body, err := json.Marshal(data)

	if err != nil {
		return err
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(constants.HEADER_API_KEY, os.Getenv("API_KEY"))
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	_, respErr := io.ReadAll(res.Body)

	return respErr

}
