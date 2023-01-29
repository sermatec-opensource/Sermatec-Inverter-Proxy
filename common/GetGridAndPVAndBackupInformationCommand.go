package common

type GetGridAndPVAndBackupInformationCommand struct {
}

func NewGetGridAndPVAndBackupInformationCommand() *GetGridAndPVAndBackupInformationCommand {
	return &GetGridAndPVAndBackupInformationCommand{}
}

func (instance *GetGridAndPVAndBackupInformationCommand) GetCommand() DatagramCommand {
	return DatagramCommand{0x0b, 0x00}
}
func (instance *GetGridAndPVAndBackupInformationCommand) GetData() DatagramData {
	return DatagramData{}
}
