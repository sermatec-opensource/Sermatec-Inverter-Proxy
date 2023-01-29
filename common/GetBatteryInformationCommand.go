package common

type GetBatteryInformationCommand struct {
}

func NewGetBatteryInformationCommand() *GetBatteryInformationCommand {
	return &GetBatteryInformationCommand{}
}
func (instance *GetBatteryInformationCommand) GetCommand() DatagramCommand {
	return DatagramCommand{0x0a, 0x00}
}
func (instance *GetBatteryInformationCommand) GetData() DatagramData {
	return DatagramData{}
}
