package common

type GetLoadInformationCommand struct {
}

func (instance *GetLoadInformationCommand) GetCommand() DatagramCommand {
	return DatagramCommand{0x0d, 0x00}
}
func (instance *GetLoadInformationCommand) GetData() DatagramData {
	return DatagramData{}
}
