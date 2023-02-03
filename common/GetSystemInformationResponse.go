package common

import (
	"errors"
)

type GetSystemInformationResponse struct {
	pcuVersion                int16
	batteryManufacturerNumber int16
	modelCode                 int16
	productSerial             string
	productSerialLN           string
}

func NewGetSystemInformationResponseFromDatagram(datagram Datagram) (*GetSystemInformationResponse, error) {
	if !datagram.IsValid() {
		return nil, errors.New("datagram is invalid")
	}

	response := &GetSystemInformationResponse{
		pcuVersion:                ParseInverterDatagramDataAsSignedInt(datagram.data, 0),
		batteryManufacturerNumber: ParseInverterDatagramDataAsSignedInt(datagram.data, 2),
		modelCode:                 ParseInverterDatagramDataAsSignedInt(datagram.data, 4),
		productSerial:             ParseInverterDatagramDataAsString(datagram.data, 6, 26),
		productSerialLN:           ParseInverterDatagramDataAsString(datagram.data, 32, 18),
	}
	return response, nil
}
func (instance *GetSystemInformationResponse) GetPCUVersion() int16 {
	return instance.pcuVersion
}
func (instance *GetSystemInformationResponse) GetBatteryManufacturerNumber() int16 {
	return instance.batteryManufacturerNumber
}
func (instance *GetSystemInformationResponse) GetModelCode() int16 {
	return instance.modelCode
}
func (instance *GetSystemInformationResponse) GetProductSerial() string {
	return instance.productSerial
}
func (instance *GetSystemInformationResponse) GetProductSerialLN() string {
	return instance.productSerialLN
}
