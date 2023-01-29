package common

import (
	"errors"
)

type GetGridAndPVAndBackupInformationResponse struct {
	firstPVVoltage        float32
	firstPVCurrent        float32
	firstPVPower          uint16
	secondPVVoltage       float32
	secondPVCurrent       float32
	secondPVPower         uint16
	inverterPhaseAVoltage float32
	inverterPhaseACurrent float32
	gridPhaseAVoltage     float32
	gridLineABVoltage     float32
	gridPhaseACurrent     float32
	inverterPhaseBVoltage float32
	inverterPhaseBCurrent float32
	gridPhaseBVoltage     float32
	gridLineBCVoltage     float32
	gridPhaseBCurrent     float32
	inverterPhaseCVoltage float32
	inverterPhaseCCurrent float32
	gridPhaseCVoltage     float32
	gridLineCAVoltage     float32
	gridPhaseCCurrent     float32
	gridFrequency         float32
	powerFactor           float32
	gridActivePower       int16
	gridReactivePower     int16
	gridApparentPower     int16
	deviceTypeCode        uint16
	dspHighVersion        uint16
	dspLowVersion         uint16
	loadPhaseAVoltage     float32
	loadPhaseBVoltage     float32
	loadPhaseCVoltage     float32
	loadFrequency         float32
	loadPhaseACurrent     float32
	loadPhaseBCurrent     float32
	loadPhaseCCurrent     float32
	loadPowerFactor       float32
	loadActivePower       int16
	loadReactivePower     int16
	loadApparentPower     int16
}

func NewGetGridAndPVAndBackupInformationFromDatagram(datagram Datagram) (*GetGridAndPVAndBackupInformationResponse, error) {
	if !datagram.IsValid() {
		return nil, errors.New("datagram is invalid")
	}

	response := &GetGridAndPVAndBackupInformationResponse{
		firstPVVoltage:        ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 0, 10),
		firstPVCurrent:        ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 2, 10),
		firstPVPower:          ParseInverterDatagramDataAsUnsignedInt(datagram.data, 4),
		secondPVVoltage:       ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 6, 10),
		secondPVCurrent:       ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 8, 10),
		secondPVPower:         ParseInverterDatagramDataAsUnsignedInt(datagram.data, 10),
		inverterPhaseAVoltage: ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 12, 10),
		inverterPhaseACurrent: ParseInverterDatagramDataAsSignedFloat(datagram.data, 14, 10),
		gridPhaseAVoltage:     ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 16, 10),
		gridLineABVoltage:     ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 18, 10),
		gridPhaseACurrent:     ParseInverterDatagramDataAsSignedFloat(datagram.data, 20, 10),
		inverterPhaseBVoltage: ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 22, 10),
		inverterPhaseBCurrent: ParseInverterDatagramDataAsSignedFloat(datagram.data, 24, 10),
		gridPhaseBVoltage:     ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 26, 10),
		gridLineBCVoltage:     ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 28, 10),
		gridPhaseBCurrent:     ParseInverterDatagramDataAsSignedFloat(datagram.data, 30, 10),
		inverterPhaseCVoltage: ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 32, 10),
		inverterPhaseCCurrent: ParseInverterDatagramDataAsSignedFloat(datagram.data, 34, 10),
		gridPhaseCVoltage:     ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 36, 10),
		gridLineCAVoltage:     ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 38, 10),
		gridPhaseCCurrent:     ParseInverterDatagramDataAsSignedFloat(datagram.data, 40, 10),
		gridFrequency:         ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 42, 100),
		powerFactor:           ParseInverterDatagramDataAsSignedFloat(datagram.data, 44, 1000),
		gridActivePower:       ParseInverterDatagramDataAsSignedInt(datagram.data, 46),
		gridReactivePower:     ParseInverterDatagramDataAsSignedInt(datagram.data, 48),
		gridApparentPower:     ParseInverterDatagramDataAsSignedInt(datagram.data, 50),
		deviceTypeCode:        ParseInverterDatagramDataAsUnsignedInt(datagram.data, 70),
		dspHighVersion:        ParseInverterDatagramDataAsUnsignedInt(datagram.data, 72),
		dspLowVersion:         ParseInverterDatagramDataAsUnsignedInt(datagram.data, 74),
		loadPhaseAVoltage:     ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 90, 10),
		loadPhaseBVoltage:     ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 92, 10),
		loadPhaseCVoltage:     ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 94, 10),
		loadFrequency:         ParseInverterDatagramDataAsUnsignedFloat(datagram.data, 96, 100),
		loadPhaseACurrent:     ParseInverterDatagramDataAsSignedFloat(datagram.data, 98, 10),
		loadPhaseBCurrent:     ParseInverterDatagramDataAsSignedFloat(datagram.data, 100, 10),
		loadPhaseCCurrent:     ParseInverterDatagramDataAsSignedFloat(datagram.data, 102, 10),
		loadPowerFactor:       ParseInverterDatagramDataAsSignedFloat(datagram.data, 104, 1000),
		loadActivePower:       ParseInverterDatagramDataAsSignedInt(datagram.data, 106),
		loadReactivePower:     ParseInverterDatagramDataAsSignedInt(datagram.data, 108),
		loadApparentPower:     ParseInverterDatagramDataAsSignedInt(datagram.data, 110),
	}

	return response, nil
}

func (instance *GetGridAndPVAndBackupInformationResponse) GetDeviceType() uint16 {
	return instance.deviceTypeCode
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetDSPHighVersion() uint16 {
	return instance.dspHighVersion
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetDSPLowVersion() uint16 {
	return instance.dspLowVersion
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetPV1Voltage() float32 {
	return instance.firstPVVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetPV2Voltage() float32 {
	return instance.secondPVVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetPV1Current() float32 {
	return instance.firstPVCurrent
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetPV2Current() float32 {
	return instance.secondPVCurrent
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetPV1Power() uint16 {
	return instance.firstPVPower
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetPV2Power() uint16 {
	return instance.secondPVPower
}

func (instance *GetGridAndPVAndBackupInformationResponse) GetInverterPhaseAVoltage() float32 {
	return instance.inverterPhaseAVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetInverterPhaseACurrent() float32 {
	return instance.inverterPhaseACurrent
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridPhaseAVoltage() float32 {
	return instance.gridPhaseAVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridPhaseACurrent() float32 {
	return instance.gridPhaseACurrent
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridLineABVoltage() float32 {
	return instance.gridLineABVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetInverterPhaseBVoltage() float32 {
	return instance.inverterPhaseBVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetInverterPhaseBCurrent() float32 {
	return instance.inverterPhaseBCurrent
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridPhaseBVoltage() float32 {
	return instance.gridPhaseBVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridPhaseBCurrent() float32 {
	return instance.gridPhaseBCurrent
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridLineBCVoltage() float32 {
	return instance.gridLineBCVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetInverterPhaseCVoltage() float32 {
	return instance.inverterPhaseCVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetInverterPhaseCCurrent() float32 {
	return instance.inverterPhaseCCurrent
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridPhaseCVoltage() float32 {
	return instance.gridPhaseCVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridPhaseCCurrent() float32 {
	return instance.gridPhaseCCurrent
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridLineCAVoltage() float32 {
	return instance.gridLineCAVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridFrequency() float32 {
	return instance.gridFrequency
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridActivePower() int16 {
	return instance.gridActivePower
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridReactivePower() int16 {
	return instance.gridReactivePower
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetGridApparentPower() int16 {
	return instance.gridApparentPower
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetLoadPhaseAVoltage() float32 {
	return instance.loadPhaseAVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetLoadPhaseACurrent() float32 {
	return instance.loadPhaseACurrent
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetLoadPhaseBVoltage() float32 {
	return instance.loadPhaseBVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetLoadPhaseBCurrent() float32 {
	return instance.loadPhaseBCurrent
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetLoadPhaseCVoltage() float32 {
	return instance.loadPhaseCVoltage
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetLoadPhaseCCurrent() float32 {
	return instance.loadPhaseCCurrent
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetLoadFrequency() float32 {
	return instance.loadFrequency
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetLoadPowerFactor() float32 {
	return instance.loadPowerFactor
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetLoadActivePower() int16 {
	return instance.loadActivePower
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetLoadReactivePower() int16 {
	return instance.loadReactivePower
}
func (instance *GetGridAndPVAndBackupInformationResponse) GetLoadApparentPower() int16 {
	return instance.loadApparentPower
}
