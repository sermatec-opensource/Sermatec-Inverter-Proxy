package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/onzeway/Sermatec-Inverter-Proxy/client"
	"github.com/onzeway/Sermatec-Inverter-Proxy/common"
	"log"
	"os"
	"strconv"
	"sync"
)

var inverterLastReceivedData *common.InverterResponseCollection

func logInverterCollectionData(collection common.InverterResponseCollection) {
	if collection.SystemInformation != nil {
		log.Printf("Inverter PCU Version: %s", collection.SystemInformation.GetPCUVersion())
		log.Printf("Inverter Serial ID: %s", collection.SystemInformation.GetSerialId())
	}

	if collection.BatteryInformation != nil {
		log.Printf("Battery Voltage: %.2f", collection.BatteryInformation.GetVoltage())
		log.Printf("Battery Current: %.2f", collection.BatteryInformation.GetCurrent())
		log.Printf("Battery Temperature: %.2f", collection.BatteryInformation.GetTemperature())
		log.Printf("Battery Max Charging Current: %.2f", collection.BatteryInformation.GetMaxChargingCurrent())
		log.Printf("Battery Max Discharging Current: %.2f", collection.BatteryInformation.GetMaxDischargingCurrent())
		log.Printf("Battery State of Charge: %d", collection.BatteryInformation.GetStateOfCharge())
		log.Printf("Battery State of Health: %d", collection.BatteryInformation.GetStateOfHealth())
		log.Printf("Battery State: %d", collection.BatteryInformation.GetState())
		log.Printf("Battery StandBy: %t", collection.BatteryInformation.IsStandBy())
		log.Printf("Battery Charging: %t", collection.BatteryInformation.IsCharging())
		log.Printf("Battery Discharging: %t", collection.BatteryInformation.IsDischarging())
	}

	if collection.GridAndLoadInformation != nil {
		log.Printf("Device type: %d", collection.GridAndLoadInformation.GetDeviceType())
		log.Printf("DSP Version: %d.%d", collection.GridAndLoadInformation.GetDSPHighVersion(), collection.GridAndLoadInformation.GetDSPLowVersion())

		log.Printf("PV1 Voltage: %.2f V", collection.GridAndLoadInformation.GetPV1Voltage())
		log.Printf("PV1 Current: %.2f A", collection.GridAndLoadInformation.GetPV1Current())
		log.Printf("PV1 Power: %d W", collection.GridAndLoadInformation.GetPV1Power())
		log.Printf("PV2 Voltage: %.2f V", collection.GridAndLoadInformation.GetPV2Voltage())
		log.Printf("PV2 Current: %.2f A", collection.GridAndLoadInformation.GetPV2Current())
		log.Printf("PV2 Power: %d W", collection.GridAndLoadInformation.GetPV2Power())

		log.Printf("Grid Phase A Voltage: %.2f V", collection.GridAndLoadInformation.GetGridPhaseAVoltage())
		log.Printf("Grid Phase A Current: %.2f A", collection.GridAndLoadInformation.GetGridPhaseACurrent())
		log.Printf("Grid Phase B Voltage: %.2f V", collection.GridAndLoadInformation.GetGridPhaseBVoltage())
		log.Printf("Grid Phase B Current: %.2f A", collection.GridAndLoadInformation.GetGridPhaseBCurrent())
		log.Printf("Grid Phase C Voltage: %.2f V", collection.GridAndLoadInformation.GetGridPhaseCVoltage())
		log.Printf("Grid Phase C Current: %.2f A", collection.GridAndLoadInformation.GetGridPhaseCCurrent())

		log.Printf("Grid Frequency: %.2f Hz", collection.GridAndLoadInformation.GetGridFrequency())

		log.Printf("Grid Line AB Voltage: %.2f V", collection.GridAndLoadInformation.GetGridLineABVoltage())
		log.Printf("Grid Line BC Voltage: %.2f V", collection.GridAndLoadInformation.GetGridLineBCVoltage())
		log.Printf("Grid Line CA Voltage: %.2f V", collection.GridAndLoadInformation.GetGridLineCAVoltage())

		log.Printf("Load Phase A Voltage: %.2f V", collection.GridAndLoadInformation.GetLoadPhaseAVoltage())
		log.Printf("Load Phase A Current: %.2f A", collection.GridAndLoadInformation.GetLoadPhaseACurrent())
		log.Printf("Load Phase B Voltage: %.2f V", collection.GridAndLoadInformation.GetLoadPhaseBVoltage())
		log.Printf("Load Phase B Current: %.2f A", collection.GridAndLoadInformation.GetLoadPhaseBCurrent())
		log.Printf("Load Phase C Voltage: %.2f V", collection.GridAndLoadInformation.GetLoadPhaseCVoltage())
		log.Printf("Load Phase C Current: %.2f A", collection.GridAndLoadInformation.GetLoadPhaseCCurrent())

		log.Printf("Load Frequency: %.2f Hz", collection.GridAndLoadInformation.GetLoadFrequency())

		log.Printf("Grid Active Power : %d W", collection.GridAndLoadInformation.GetGridActivePower())
		log.Printf("Grid Reactive Power : %d Var", collection.GridAndLoadInformation.GetGridReactivePower())
		log.Printf("Grid Apparent Power : %d VA", collection.GridAndLoadInformation.GetGridApparentPower())
		log.Printf("Load Power Factor : %.3f", collection.GridAndLoadInformation.GetLoadPowerFactor())
		log.Printf("Load Active Power : %d W", collection.GridAndLoadInformation.GetLoadActivePower())
		log.Printf("Load Reactive Power : %d Var", collection.GridAndLoadInformation.GetLoadReactivePower())
		log.Printf("Load Apparent Power : %d VA", collection.GridAndLoadInformation.GetLoadApparentPower())
	}
}

func main() {
	dotEnvFile := flag.String("dot-env-file", "/run/secrets/.env", "path to DotEnv file")
	clientEnabled := flag.Bool("client", false, "client enabled")
	mockServerEnabled := flag.Bool("mock-server", false, "mock server mode")
	proxyEnabled := flag.Bool("proxy", false, "proxy server enabled")
	webUIEnabled := flag.Bool("ui", false, "enabling web ui")
	help := flag.Bool("h", false, "showing the command help")

	flag.Parse()

	if *help {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	log.Println("## Options provided ##")
	log.Printf("==> DotEnv file path : %s", *dotEnvFile)
	log.Printf("==> Client Enabled : %t", *clientEnabled)
	log.Printf("==> Mock Server Enabled : %t", *mockServerEnabled)
	log.Printf("==> Proxy Enabled : %t", *proxyEnabled)
	log.Printf("==> WebUI Enabled : %t", *webUIEnabled)

	err := godotenv.Load(*dotEnvFile)
	if err != nil {
		log.Panicf("Error loading %s file", *dotEnvFile)
	}

	log.Println("## EnvVar Provided ##")
	log.Printf("==> INVERTER_HOST : %s", os.Getenv("INVERTER_HOST"))
	log.Printf("==> INVERTER_PORT : %s", os.Getenv("INVERTER_PORT"))
	log.Printf("==> CLIENT_POLLING_INTERVAL : %s", os.Getenv("CLIENT_POLLING_INTERVAL"))

	inverterDataCollectionChan := make(chan common.InverterResponseCollection)
	var inverterClientWaitingGroup sync.WaitGroup

	if *clientEnabled {
		port, err := strconv.Atoi(os.Getenv("INVERTER_PORT"))
		if err != nil {
			panic(err)
		}
		pollingInterval, err := strconv.Atoi(os.Getenv("CLIENT_POLLING_INTERVAL"))
		if err != nil {
			panic(err)
		}
		inverterClient := client.NewClient(os.Getenv("INVERTER_HOST"), port)
		inverterClientWaitingGroup.Add(1)
		go inverterClient.GetData(pollingInterval, inverterDataCollectionChan, &inverterClientWaitingGroup)

		go func() {
			for responseCollection := range inverterDataCollectionChan {
				inverterLastReceivedData = &responseCollection
				logInverterCollectionData(responseCollection)
			}
		}()
	}

	inverterClientWaitingGroup.Wait()
	os.Exit(0)
}
