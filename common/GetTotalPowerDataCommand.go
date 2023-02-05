package common

type GetTotalPowerDataCommand struct {
}

func NewGetTotalPowerDataCommand() *GetSystemInformationCommand {
	return &GetSystemInformationCommand{}
}
func (instance *GetTotalPowerDataCommand) GetCommand() DatagramCommand {
	return DatagramCommand{0x99, 0x00}
}
func (instance *GetTotalPowerDataCommand) GetData() DatagramData {
	return DatagramData{}
}
