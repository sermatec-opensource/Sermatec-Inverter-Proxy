package common

type GetControlCabinetInformationCommand struct {
}

func NewGetControlCabinetInformationCommand() *GetControlCabinetInformationCommand {
	return &GetControlCabinetInformationCommand{}
}

func (instance *GetControlCabinetInformationCommand) GetCommand() DatagramCommand {
	return DatagramCommand{0x0b, 0x00}
}
func (instance *GetControlCabinetInformationCommand) GetData() DatagramData {
	return DatagramData{}
}
