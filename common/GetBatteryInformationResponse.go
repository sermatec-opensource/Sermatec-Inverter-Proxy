package common

import (
	"errors"
)

type GetBatteryInformationResponse struct {
	voltage               float32
	current               float32
	temperature           float32
	stateOfCharge         uint16
	stateOfHealth         uint16
	state                 uint16
	maxChargingCurrent    float32
	maxDischargingCurrent float32
}

func NewGetBatteryInformationFromDatagram(datagram Datagram) (*GetBatteryInformationResponse, error) {
	if !datagram.IsValid() {
		return nil, errors.New("datagram is invalid")
	}

	response := &GetBatteryInformationResponse{
		voltage:               ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 0, 10),
		current:               ParseInverterDatagramDataAsSignedFloat(datagram.data, 2, 10),
		temperature:           ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 4, 10),
		stateOfCharge:         ParseInverterDatagramDataAsUnsignedInt(datagram.data, 6),
		stateOfHealth:         ParseInverterDatagramDataAsUnsignedInt(datagram.data, 8),
		state:                 ParseInverterDatagramDataAsUnsignedInt(datagram.data, 10),
		maxChargingCurrent:    ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 12, 10),
		maxDischargingCurrent: ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 14, 10),
	}
	return response, nil
}
func (instance *GetBatteryInformationResponse) GetVoltage() float32 {
	return instance.voltage
}
func (instance *GetBatteryInformationResponse) GetCurrent() float32 {
	return instance.current
}
func (instance *GetBatteryInformationResponse) GetTemperature() float32 {
	return instance.temperature
}
func (instance *GetBatteryInformationResponse) GetMaxChargingCurrent() float32 {
	return instance.maxChargingCurrent
}
func (instance *GetBatteryInformationResponse) GetMaxDischargingCurrent() float32 {
	return instance.maxDischargingCurrent
}
func (instance *GetBatteryInformationResponse) GetState() uint16 {
	return instance.state
}
func (instance *GetBatteryInformationResponse) GetStateOfCharge() uint16 {
	return instance.stateOfCharge
}
func (instance *GetBatteryInformationResponse) GetStateOfHealth() uint16 {
	return instance.stateOfHealth
}
func (instance *GetBatteryInformationResponse) IsCharging() bool {
	return instance.state == 0x0011
}
func (instance *GetBatteryInformationResponse) IsDischarging() bool {
	return instance.state == 0x0022
}
func (instance *GetBatteryInformationResponse) IsStandBy() bool {
	return instance.state == 0x0033
}
