package common

import (
	"errors"
)

type GetControlCabinetInformationResponse struct {
	firstPVVoltage             float32
	firstPVCurrent             float32
	firstPVPower               uint16
	secondPVVoltage            float32
	secondPVCurrent            float32
	secondPVPower              uint16
	inverterPhaseAVoltage      float32
	inverterPhaseACurrent      float32
	gridPhaseAVoltage          float32
	gridLineABVoltage          float32
	gridPhaseACurrent          float32
	inverterPhaseBVoltage      float32
	inverterPhaseBCurrent      float32
	gridPhaseBVoltage          float32
	gridLineBCVoltage          float32
	gridPhaseBCurrent          float32
	inverterPhaseCVoltage      float32
	inverterPhaseCCurrent      float32
	gridPhaseCVoltage          float32
	gridLineCAVoltage          float32
	gridPhaseCCurrent          float32
	gridFrequency              float32
	powerFactor                float32
	gridActivePower            int16
	gridReactivePower          int16
	systemApparentPower        int16
	batteryCurrent             float32
	batteryVoltage             float32
	busDCPositiveVoltage       float32
	busDCNegativeVoltage       float32
	busDCBilateralVoltage      float32
	busDCPower                 int16
	internalTemperature        float32
	backupBusDCPositiveVoltage float32
	backupBusDCNegativeVoltage float32
	deviceTypeCode             uint16
	dspHighVersion             uint16
	dspLowVersion              uint16
	parallelAddress            uint16
	workEfficiency             uint16
	battery1Current            float32
	battery2Current            float32
	moduleA1Temperature        float32
	moduleB1Temperature        float32
	moduleC1Temperature        float32
	loadPhaseAVoltage          float32
	loadPhaseBVoltage          float32
	loadPhaseCVoltage          float32
	loadFrequency              float32
	loadPhaseACurrent          float32
	loadPhaseBCurrent          float32
	loadPhaseCCurrent          float32
	loadPowerFactor            float32
	loadActivePower            int16
	loadReactivePower          int16
	loadApparentPower          int16
}

func NewGetControlCabinetInformationFromDatagram(datagram Datagram) (*GetControlCabinetInformationResponse, error) {
	if !datagram.IsValid() {
		return nil, errors.New("datagram is invalid")
	}

	response := &GetControlCabinetInformationResponse{
		firstPVVoltage:             ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 0, 10),
		firstPVCurrent:             ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 2, 10),
		firstPVPower:               ParseInverterDatagramDataAsUnsignedInt(datagram.data, 4),
		secondPVVoltage:            ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 6, 10),
		secondPVCurrent:            ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 8, 10),
		secondPVPower:              ParseInverterDatagramDataAsUnsignedInt(datagram.data, 10),
		inverterPhaseAVoltage:      ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 12, 10),
		inverterPhaseACurrent:      ParseInverterDatagramDataAsSignedFloat(datagram.data, 14, 10),
		gridPhaseAVoltage:          ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 16, 10),
		gridLineABVoltage:          ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 18, 10),
		gridPhaseACurrent:          ParseInverterDatagramDataAsSignedFloat(datagram.data, 20, 10),
		inverterPhaseBVoltage:      ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 22, 10),
		inverterPhaseBCurrent:      ParseInverterDatagramDataAsSignedFloat(datagram.data, 24, 10),
		gridPhaseBVoltage:          ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 26, 10),
		gridLineBCVoltage:          ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 28, 10),
		gridPhaseBCurrent:          ParseInverterDatagramDataAsSignedFloat(datagram.data, 30, 10),
		inverterPhaseCVoltage:      ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 32, 10),
		inverterPhaseCCurrent:      ParseInverterDatagramDataAsSignedFloat(datagram.data, 34, 10),
		gridPhaseCVoltage:          ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 36, 10),
		gridLineCAVoltage:          ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 38, 10),
		gridPhaseCCurrent:          ParseInverterDatagramDataAsSignedFloat(datagram.data, 40, 10),
		gridFrequency:              ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 42, 100),
		powerFactor:                ParseInverterDatagramDataAsSignedFloat(datagram.data, 44, 1000),
		gridActivePower:            ParseInverterDatagramDataAsSignedInt(datagram.data, 46),
		gridReactivePower:          ParseInverterDatagramDataAsSignedInt(datagram.data, 48),
		systemApparentPower:        ParseInverterDatagramDataAsSignedInt(datagram.data, 50),
		batteryCurrent:             ParseInverterDatagramDataAsSignedFloat(datagram.data, 52, 10),
		batteryVoltage:             ParseInverterDatagramDataAsSignedFloat(datagram.data, 54, 10),
		busDCPositiveVoltage:       ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 56, 10),
		busDCNegativeVoltage:       ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 58, 10),
		busDCBilateralVoltage:      ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 60, 10),
		busDCPower:                 ParseInverterDatagramDataAsSignedInt(datagram.data, 62),
		internalTemperature:        ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 64, 10),
		backupBusDCPositiveVoltage: ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 66, 10),
		backupBusDCNegativeVoltage: ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 68, 10),
		deviceTypeCode:             ParseInverterDatagramDataAsUnsignedInt(datagram.data, 70),
		dspHighVersion:             ParseInverterDatagramDataAsUnsignedInt(datagram.data, 72),
		dspLowVersion:              ParseInverterDatagramDataAsUnsignedInt(datagram.data, 74),
		parallelAddress:            ParseInverterDatagramDataAsUnsignedInt(datagram.data, 76),
		workEfficiency:             ParseInverterDatagramDataAsUnsignedInt(datagram.data, 78),
		battery1Current:            ParseInverterDatagramDataAsSignedFloat(datagram.data, 80, 10),
		battery2Current:            ParseInverterDatagramDataAsSignedFloat(datagram.data, 82, 10),
		moduleA1Temperature:        ParseInverterDatagramDataAsSignedFloat(datagram.data, 84, 10),
		moduleB1Temperature:        ParseInverterDatagramDataAsSignedFloat(datagram.data, 86, 10),
		moduleC1Temperature:        ParseInverterDatagramDataAsSignedFloat(datagram.data, 88, 10),
		loadPhaseAVoltage:          ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 90, 10),
		loadPhaseBVoltage:          ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 92, 10),
		loadPhaseCVoltage:          ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 94, 10),
		loadFrequency:              ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 96, 100),
		loadPhaseACurrent:          ParseInverterDatagramDataAsSignedFloat(datagram.data, 98, 10),
		loadPhaseBCurrent:          ParseInverterDatagramDataAsSignedFloat(datagram.data, 100, 10),
		loadPhaseCCurrent:          ParseInverterDatagramDataAsSignedFloat(datagram.data, 102, 10),
		loadPowerFactor:            ParseInverterDatagramDataAsSignedFloat(datagram.data, 104, 1000),
		loadActivePower:            ParseInverterDatagramDataAsSignedInt(datagram.data, 106),
		loadReactivePower:          ParseInverterDatagramDataAsSignedInt(datagram.data, 108),
		loadApparentPower:          ParseInverterDatagramDataAsSignedInt(datagram.data, 110),
	}

	return response, nil
}

func (instance *GetControlCabinetInformationResponse) GetDeviceType() uint16 {
	return instance.deviceTypeCode
}
func (instance *GetControlCabinetInformationResponse) GetDSPHighVersion() uint16 {
	return instance.dspHighVersion
}
func (instance *GetControlCabinetInformationResponse) GetDSPLowVersion() uint16 {
	return instance.dspLowVersion
}
func (instance *GetControlCabinetInformationResponse) GetPV1Voltage() float32 {
	return instance.firstPVVoltage
}
func (instance *GetControlCabinetInformationResponse) GetPV2Voltage() float32 {
	return instance.secondPVVoltage
}
func (instance *GetControlCabinetInformationResponse) GetPV1Current() float32 {
	return instance.firstPVCurrent
}
func (instance *GetControlCabinetInformationResponse) GetPV2Current() float32 {
	return instance.secondPVCurrent
}
func (instance *GetControlCabinetInformationResponse) GetPV1Power() uint16 {
	return instance.firstPVPower
}
func (instance *GetControlCabinetInformationResponse) GetPV2Power() uint16 {
	return instance.secondPVPower
}

func (instance *GetControlCabinetInformationResponse) GetInverterPhaseAVoltage() float32 {
	return instance.inverterPhaseAVoltage
}
func (instance *GetControlCabinetInformationResponse) GetInverterPhaseACurrent() float32 {
	return instance.inverterPhaseACurrent
}
func (instance *GetControlCabinetInformationResponse) GetGridPhaseAVoltage() float32 {
	return instance.gridPhaseAVoltage
}
func (instance *GetControlCabinetInformationResponse) GetGridPhaseACurrent() float32 {
	return instance.gridPhaseACurrent
}
func (instance *GetControlCabinetInformationResponse) GetGridLineABVoltage() float32 {
	return instance.gridLineABVoltage
}
func (instance *GetControlCabinetInformationResponse) GetInverterPhaseBVoltage() float32 {
	return instance.inverterPhaseBVoltage
}
func (instance *GetControlCabinetInformationResponse) GetInverterPhaseBCurrent() float32 {
	return instance.inverterPhaseBCurrent
}
func (instance *GetControlCabinetInformationResponse) GetGridPhaseBVoltage() float32 {
	return instance.gridPhaseBVoltage
}
func (instance *GetControlCabinetInformationResponse) GetGridPhaseBCurrent() float32 {
	return instance.gridPhaseBCurrent
}
func (instance *GetControlCabinetInformationResponse) GetGridLineBCVoltage() float32 {
	return instance.gridLineBCVoltage
}
func (instance *GetControlCabinetInformationResponse) GetInverterPhaseCVoltage() float32 {
	return instance.inverterPhaseCVoltage
}
func (instance *GetControlCabinetInformationResponse) GetInverterPhaseCCurrent() float32 {
	return instance.inverterPhaseCCurrent
}
func (instance *GetControlCabinetInformationResponse) GetGridPhaseCVoltage() float32 {
	return instance.gridPhaseCVoltage
}
func (instance *GetControlCabinetInformationResponse) GetGridPhaseCCurrent() float32 {
	return instance.gridPhaseCCurrent
}
func (instance *GetControlCabinetInformationResponse) GetGridLineCAVoltage() float32 {
	return instance.gridLineCAVoltage
}
func (instance *GetControlCabinetInformationResponse) GetGridFrequency() float32 {
	return instance.gridFrequency
}
func (instance *GetControlCabinetInformationResponse) GetGridActivePower() int16 {
	return instance.gridActivePower
}
func (instance *GetControlCabinetInformationResponse) GetGridReactivePower() int16 {
	return instance.gridReactivePower
}
func (instance *GetControlCabinetInformationResponse) GetSystemApparentPower() int16 {
	return instance.systemApparentPower
}
func (instance *GetControlCabinetInformationResponse) GetLoadPhaseAVoltage() float32 {
	return instance.loadPhaseAVoltage
}
func (instance *GetControlCabinetInformationResponse) GetLoadPhaseACurrent() float32 {
	return instance.loadPhaseACurrent
}
func (instance *GetControlCabinetInformationResponse) GetLoadPhaseBVoltage() float32 {
	return instance.loadPhaseBVoltage
}
func (instance *GetControlCabinetInformationResponse) GetLoadPhaseBCurrent() float32 {
	return instance.loadPhaseBCurrent
}
func (instance *GetControlCabinetInformationResponse) GetLoadPhaseCVoltage() float32 {
	return instance.loadPhaseCVoltage
}
func (instance *GetControlCabinetInformationResponse) GetLoadPhaseCCurrent() float32 {
	return instance.loadPhaseCCurrent
}
func (instance *GetControlCabinetInformationResponse) GetLoadFrequency() float32 {
	return instance.loadFrequency
}
func (instance *GetControlCabinetInformationResponse) GetLoadPowerFactor() float32 {
	return instance.loadPowerFactor
}
func (instance *GetControlCabinetInformationResponse) GetLoadActivePower() int16 {
	return instance.loadActivePower
}
func (instance *GetControlCabinetInformationResponse) GetLoadReactivePower() int16 {
	return instance.loadReactivePower
}
func (instance *GetControlCabinetInformationResponse) GetLoadApparentPower() int16 {
	return instance.loadApparentPower
}
func (instance *GetControlCabinetInformationResponse) GetPowerFactor() float32 {
	return instance.powerFactor
}
func (instance *GetControlCabinetInformationResponse) GetBatteryVoltage() float32 {
	return instance.batteryVoltage
}
func (instance *GetControlCabinetInformationResponse) GetBatteryCurrent() float32 {
	return instance.batteryCurrent
}
func (instance *GetControlCabinetInformationResponse) GetBatteryNo1Current() float32 {
	return instance.battery1Current
}
func (instance *GetControlCabinetInformationResponse) GetBatteryNo2Current() float32 {
	return instance.battery2Current
}
func (instance *GetControlCabinetInformationResponse) GetBusDCPower() int16 {
	return instance.busDCPower
}
func (instance *GetControlCabinetInformationResponse) GetInternalTemperature() float32 {
	return instance.internalTemperature
}
func (instance *GetControlCabinetInformationResponse) GetParallelAddress() uint16 {
	return instance.parallelAddress
}
func (instance *GetControlCabinetInformationResponse) GetBusDCPositiveVoltage() float32 {
	return instance.busDCPositiveVoltage
}
func (instance *GetControlCabinetInformationResponse) GetBusDCNegativeVoltage() float32 {
	return instance.busDCNegativeVoltage
}
func (instance *GetControlCabinetInformationResponse) GetBusDCBilateralVoltage() float32 {
	return instance.busDCBilateralVoltage
}
func (instance *GetControlCabinetInformationResponse) GetBackupBusDCNegativeVoltage() float32 {
	return instance.backupBusDCNegativeVoltage
}
func (instance *GetControlCabinetInformationResponse) GetBackupBusDCPositiveVoltage() float32 {
	return instance.backupBusDCPositiveVoltage
}
func (instance *GetControlCabinetInformationResponse) GetWorkEfficiency() uint16 {
	return instance.workEfficiency
}
func (instance *GetControlCabinetInformationResponse) GetModuleA1Temperature() float32 {
	return instance.moduleA1Temperature
}
func (instance *GetControlCabinetInformationResponse) GetModuleB1Temperature() float32 {
	return instance.moduleB1Temperature
}
func (instance *GetControlCabinetInformationResponse) GetModuleC1Temperature() float32 {
	return instance.moduleC1Temperature
}
