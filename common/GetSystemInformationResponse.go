package common

import (
	"encoding/binary"
	"errors"
	"strconv"
	"strings"
)

type GetSystemInformationResponse struct {
	pcuVersion string
	serialId   string
}

func NewGetSystemInformationResponseFromDatagram(datagram Datagram) (*GetSystemInformationResponse, error) {
	if !datagram.IsValid() {
		return nil, errors.New("datagram is invalid")
	}

	c := binary.BigEndian
	pcuVersion := strconv.Itoa(int(c.Uint16(datagram.data[0:2])))
	nullIndex := strings.Index(string(datagram.data[6:]), "\x00") + 6
	serialId := string(datagram.data[6:nullIndex])

	response := &GetSystemInformationResponse{
		pcuVersion: pcuVersion,
		serialId:   serialId,
	}
	return response, nil
}
func (instance *GetSystemInformationResponse) GetPCUVersion() string {
	return instance.pcuVersion
}
func (instance *GetSystemInformationResponse) GetSerialId() string {
	return instance.serialId
}
