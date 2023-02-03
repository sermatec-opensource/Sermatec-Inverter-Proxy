package client

import (
	"errors"
	"fmt"
	"github.com/onzeway/Sermatec-Inverter-Proxy/common"
	"log"
	"net"
	"sync"
	"time"
)

type Client struct {
	host string
	port int
}

func NewClient(host string, port int) *Client {
	return &Client{
		host: host,
		port: port,
	}
}

func (instance Client) send(query []byte) ([]byte, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", instance.host, instance.port))
	if err != nil {
		log.Printf("ResolveTCPAddr failed: %s", err.Error())
		return nil, err
	}

	dialer := net.Dialer{Timeout: 5 * time.Second}
	conn, err := dialer.Dial("tcp", tcpAddr.String())
	if err != nil {
		log.Printf("Dial failed: %s", err.Error())
		return nil, err
	}

	_, err = conn.Write(query)
	if err != nil {
		log.Printf("Write to server failed: %s", err.Error())
		return nil, err
	}

	log.Printf("write to server (hex): %x", query)

	rawData := make([]byte, 1024)
	_, err = conn.Read(rawData)
	if err != nil {
		log.Printf("Write to server failed: %s", err.Error())
		return nil, err
	}
	_ = conn.Close()

	log.Printf("reply from server (hex): %x", rawData)

	return rawData, nil
}

func (instance Client) getSystemInformation() (*common.GetSystemInformationResponse, error) {
	request := common.NewRequestDatagram(common.NewGetSystemInformationCommand())
	rawData, err := instance.send(request.ToBytes())
	if err != nil {
		return nil, err
	}
	responseDatagram := common.GetDatagramFromRawData(rawData)
	if responseDatagram.GetCommand() != request.GetCommand() {
		return nil, errors.New("response does not match with request command, probably an error from inverter")
	}
	response, err := common.NewGetSystemInformationResponseFromDatagram(*responseDatagram)
	if err != nil {
		return nil, err
	}

	return response, nil
}
func (instance Client) getBatteryInformation() (*common.GetBatteryInformationResponse, error) {
	request := common.NewRequestDatagram(common.NewGetBatteryInformationCommand())
	rawData, err := instance.send(request.ToBytes())
	if err != nil {
		return nil, err
	}
	responseDatagram := common.GetDatagramFromRawData(rawData)
	if responseDatagram.GetCommand() != request.GetCommand() {
		return nil, errors.New("response does not match with request command, probably an error from inverter")
	}
	response, err := common.NewGetBatteryInformationFromDatagram(*responseDatagram)
	if err != nil {
		return nil, err
	}

	return response, nil
}
func (instance Client) getControlCabinetInformation() (*common.GetControlCabinetInformationResponse, error) {
	request := common.NewRequestDatagram(common.NewGetControlCabinetInformationCommand())
	rawData, err := instance.send(request.ToBytes())
	if err != nil {
		return nil, err
	}
	responseDatagram := common.GetDatagramFromRawData(rawData)
	if responseDatagram.GetCommand() != request.GetCommand() {
		return nil, errors.New("response does not match with request command, probably an error from inverter")
	}
	response, err := common.NewGetControlCabinetInformationFromDatagram(*responseDatagram)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (instance Client) GetData(pollingTime int, informationChan chan common.InverterResponseCollection, waitingGroup *sync.WaitGroup) {
	defer waitingGroup.Done()
	defer close(informationChan)
	for {
		systemInformation, err := instance.getSystemInformation()
		if err != nil {
			systemInformation = nil
			log.Printf("An error has occurred %v", err)
		}

		batteryInformation, err := instance.getBatteryInformation()
		if err != nil {
			batteryInformation = nil
			log.Printf("An error has occurred %v", err)
		}

		gridAndLoadInformation, err := instance.getControlCabinetInformation()
		if err != nil {
			gridAndLoadInformation = nil
			log.Printf("An error has occurred %v", err)
		}

		informationChan <- common.InverterResponseCollection{
			SystemInformation:         systemInformation,
			BatteryInformation:        batteryInformation,
			ControlCabinetInformation: gridAndLoadInformation,
		}
		time.Sleep(time.Duration(pollingTime) * time.Second)
	}
}
