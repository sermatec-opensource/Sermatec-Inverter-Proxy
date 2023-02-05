package common

import (
	"bytes"
	"encoding/binary"
	"log"
)

type InverterCommand interface {
	GetCommand() DatagramCommand
	GetData() DatagramData
}

type DatagramCommand = [2]byte
type DatagramData = []byte

type Datagram struct {
	preamble [2]byte
	source   byte
	target   byte
	command  DatagramCommand
	length   byte
	data     DatagramData
	checksum byte
	end      byte
}

func NewRequestDatagram(inverterCommand InverterCommand) *Datagram {
	datagram := &Datagram{
		preamble: [2]byte{0xfe, 0x55},
		source:   0x64,
		target:   0x14,
		command:  inverterCommand.GetCommand(),
		length:   byte(len(inverterCommand.GetData())),
		data:     inverterCommand.GetData(),
		checksum: 0x00,
		end:      0xae,
	}
	datagram.checksum = calcChecksum(datagram)
	return datagram
}

func GetDatagramFromRawData(rawData []byte) *Datagram {
	datagram := Datagram{
		preamble: [2]byte{rawData[0], rawData[1]},
		source:   rawData[2],
		target:   rawData[3],
		command:  DatagramCommand{rawData[4], rawData[5]},
		length:   rawData[6],
		data:     DatagramData{},
		checksum: 0x00,
		end:      0x00,
	}

	var messageLength uint8 = rawData[6]
	dataMaxIndex := 7 + messageLength
	if messageLength > 0 {
		datagram.data = rawData[7:dataMaxIndex]
	}
	datagram.checksum = rawData[dataMaxIndex]
	dataMaxIndex++
	datagram.end = rawData[dataMaxIndex]

	log.Printf("datagram Message (hex): %x", datagram.data)

	return &datagram
}

func ParseInverterDatagramDataAsUnsignedFloat(data []byte, firstByteAddress uint16, factor uint16) float32 {
	return float32(ParseInverterDatagramDataAsUnsignedInt(data, firstByteAddress)) / float32(factor)
}
func ParseInverterDatagramDataAsSignedFloat(data []byte, firstByteAddress uint16, factor uint16) float32 {
	return float32(ParseInverterDatagramDataAsSignedInt(data, firstByteAddress)) / float32(factor)
}
func ParseInverterDatagramDataAsSignedLongFloat(data []byte, firstByteAddress uint16, factor uint16) float32 {
	return float32(ParseInverterDatagramDataAsSignedLong(data, firstByteAddress)) / float32(factor)
}

func ParseInverterDatagramDataAsUnsignedInt(data []byte, firstByteAddress uint16) uint16 {
	endAddress := firstByteAddress + 2
	return binary.BigEndian.Uint16(data[firstByteAddress:endAddress])
}

func ParseInverterDatagramDataAsSignedLong(data []byte, firstByteAddress uint16) int32 {
	endAddress := firstByteAddress + 4
	var value int32
	err := binary.Read(bytes.NewReader(data[firstByteAddress:endAddress]), binary.BigEndian, &value)
	if err != nil {
		log.Printf("An error has occurred when parsing data as signed long: %v", err)
		return 0
	}
	return value
}

func ParseInverterDatagramDataAsRepeatedSignedFloat(data []byte, firstByteAddress uint16, factor uint16, repeat int) []float32 {
	var value []float32
	var currentFirstByteAddress uint16
	for i := 0; i < repeat; i++ {
		currentFirstByteAddress = firstByteAddress + 2*uint16(i)
		value = append(value, ParseInverterDatagramDataAsSignedFloat(data, currentFirstByteAddress, factor))
	}
	return value
}

func ParseInverterDatagramDataAsRepeatedSignedInt(data []byte, firstByteAddress uint16, repeat int) []int16 {
	var value []int16
	var currentFirstByteAddress uint16
	for i := 0; i < repeat; i++ {
		currentFirstByteAddress = firstByteAddress + 2*uint16(i)
		value = append(value, ParseInverterDatagramDataAsSignedInt(data, currentFirstByteAddress))
	}
	return value
}

func ParseInverterDatagramDataAsRepeatedSignedLong(data []byte, firstByteAddress uint16, repeat int) []int32 {
	var value []int32
	var currentFirstByteAddress uint16
	for i := 0; i < repeat; i++ {
		currentFirstByteAddress = firstByteAddress + 4*uint16(i)
		value = append(value, ParseInverterDatagramDataAsSignedLong(data, currentFirstByteAddress))
	}
	return value
}

func ParseInverterDatagramDataAsString(data []byte, firstByteAddress uint16, strLen uint16) string {
	endAddress := firstByteAddress + strLen
	return string(data[firstByteAddress:endAddress])
}

func ParseInverterDatagramDataAsSignedInt(data []byte, firstByteAddress uint16) int16 {
	endAddress := firstByteAddress + 2
	var value int16
	err := binary.Read(bytes.NewReader(data[firstByteAddress:endAddress]), binary.BigEndian, &value)
	if err != nil {
		log.Printf("An error has occurred when parsing data as signed int: %v", err)
		return 0
	}
	return value
}

func calcChecksum(datagram *Datagram) byte {
	var checkSumValue byte = 0x0f
	var valuesToCompare []byte
	valuesToCompare = append(valuesToCompare, datagram.preamble[:]...)
	valuesToCompare = append(valuesToCompare, datagram.source, datagram.target)
	valuesToCompare = append(valuesToCompare, datagram.command[:]...)
	valuesToCompare = append(valuesToCompare, datagram.length)
	valuesToCompare = append(valuesToCompare, datagram.data[:]...)
	for _, value := range valuesToCompare {
		checkSumValue = checkSumValue ^ value
	}
	return checkSumValue
}

func (instance *Datagram) IsValid() bool {
	if instance.end != 0xae {
		return false
	}
	expectedChecksum := calcChecksum(instance)
	return expectedChecksum == instance.checksum
}

func (instance *Datagram) GetCommand() DatagramCommand {
	return instance.command
}

func (instance *Datagram) ToBytes() []byte {
	var query []byte
	query = append(query, instance.preamble[:]...)
	query = append(query, instance.source, instance.target)
	query = append(query, instance.command[:]...)
	query = append(query, instance.length)
	query = append(query, instance.data[:]...)
	query = append(query, instance.checksum, instance.end)
	return query
}
