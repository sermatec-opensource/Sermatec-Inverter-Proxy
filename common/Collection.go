package common

type InverterResponseCollection struct {
	SystemInformation      *GetSystemInformationResponse
	BatteryInformation     *GetBatteryInformationResponse
	GridAndLoadInformation *GetGridAndPVAndBackupInformationResponse
}
