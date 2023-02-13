package common

import (
	"errors"
)

type totalPV struct {
	dailyPVPowerGeneration    float32
	totalPVPowerGeneration    int32
	dailyLoadPowerConsumption float32
	totalLoadPowerConsumption int32
	dailyMoneySaving          float32
	totalMoneySaving          float32
	dayPVPower                []float32
	monthPVPower              []int16
	yearPVPower               []int32
	historyPVPower            []int32
}
type totalGrid struct {
	dailyGridConsumption float32
	totalGridConsumption int32
}

type GetTotalPowerDataResponse struct {
	totalPV
	totalGrid
}

func (instance *GetTotalPowerDataResponse) GetDailyPVPowerGeneration() float32 {
	return instance.dailyPVPowerGeneration
}

func (instance *GetTotalPowerDataResponse) GetTotalPVPowerGeneration() int32 {
	return instance.totalPVPowerGeneration
}

func (instance *GetTotalPowerDataResponse) GetDailyLoadPowerConsumption() float32 {
	return instance.dailyLoadPowerConsumption
}

func (instance *GetTotalPowerDataResponse) GetTotalLoadPowerConsumption() int32 {
	return instance.totalLoadPowerConsumption
}

func (instance *GetTotalPowerDataResponse) GetDailyMoneySaving() float32 {
	return instance.dailyMoneySaving
}

func (instance *GetTotalPowerDataResponse) GetTotalMoneySaving() float32 {
	return instance.totalMoneySaving
}

func (instance *GetTotalPowerDataResponse) GetDayPVPower() []float32 {
	return instance.dayPVPower
}

func (instance *GetTotalPowerDataResponse) GetMonthPVPower() []int16 {
	return instance.monthPVPower
}

func (instance *GetTotalPowerDataResponse) GetYearPVPower() []int32 {
	return instance.yearPVPower
}

func (instance *GetTotalPowerDataResponse) GetHistoryPVPower() []int32 {
	return instance.historyPVPower
}

func (instance *GetTotalPowerDataResponse) GetDailyGridConsumption() float32 {
	return instance.dailyGridConsumption
}
func (instance *GetTotalPowerDataResponse) GetTotalGridConsumption() int32 {
	return instance.totalGridConsumption
}

func NewGetTotalPowerDataFromDatagram(datagram Datagram) (*GetTotalPowerDataResponse, error) {
	if !datagram.IsValid() {
		return nil, errors.New("datagram is invalid")
	}

	response := &GetTotalPowerDataResponse{
		totalPV: totalPV{
			dailyPVPowerGeneration:    ParseInverterDatagramDataAsSignedFloat(datagram.data, 0, 10),
			totalPVPowerGeneration:    ParseInverterDatagramDataAsSignedLong(datagram.data, 2),
			dailyLoadPowerConsumption: ParseInverterDatagramDataAsSignedFloat(datagram.data, 6, 10),
			totalLoadPowerConsumption: ParseInverterDatagramDataAsSignedLong(datagram.data, 8),
			dailyMoneySaving:          ParseInverterDatagramDataAsSignedFloat(datagram.data, 12, 10),
			totalMoneySaving:          ParseInverterDatagramDataAsSignedLongFloat(datagram.data, 14, 10),
			dayPVPower:                ParseInverterDatagramDataAsRepeatedSignedFloat(datagram.data, 18, 10, 48),
			monthPVPower:              ParseInverterDatagramDataAsRepeatedSignedInt(datagram.data, 114, 31),
			yearPVPower:               ParseInverterDatagramDataAsRepeatedSignedLong(datagram.data, 176, 12),
			historyPVPower:            ParseInverterDatagramDataAsRepeatedSignedLong(datagram.data, 224, 5),
		},
		totalGrid: totalGrid{
			dailyGridConsumption: ParseInverterDatagramDataAsSignedFloat(datagram.data, 244, 10),
			totalGridConsumption: ParseInverterDatagramDataAsSignedLong(datagram.data, 246),
		},
	}
	return response, nil
}
