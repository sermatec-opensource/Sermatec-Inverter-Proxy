package common

type GetSystemInformationCommand struct {
}

func NewGetSystemInformationCommand() *GetSystemInformationCommand {
	return &GetSystemInformationCommand{}
}

func (instance *GetSystemInformationCommand) GetCommand() DatagramCommand {
	return DatagramCommand{0x98, 0x00}
}
func (instance *GetSystemInformationCommand) GetData() DatagramData {
	return DatagramData{}
}
