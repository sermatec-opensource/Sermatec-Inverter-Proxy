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
		log.Println("=================== System Information ====================")
		log.Printf("Inverter PCU Version: %d", collection.SystemInformation.GetPCUVersion())
		log.Printf("Inverter Battery Manufatcurer code: %d", collection.SystemInformation.GetBatteryManufacturerNumber())
		log.Printf("Inverter Model Code: %d", collection.SystemInformation.GetModelCode())
		log.Printf("Inverter Product Serial: %s", collection.SystemInformation.GetProductSerial())
		log.Printf("Inverter Product Serial LN: %s", collection.SystemInformation.GetProductSerialLN())
	}

	if collection.BatteryInformation != nil {
		log.Println("=================== Battery Information ====================")
		log.Printf("Battery Voltage: %.2f V", collection.BatteryInformation.GetVoltage())
		log.Printf("Battery Current: %.2f A", collection.BatteryInformation.GetCurrent())
		log.Printf("Battery Temperature: %.2f °C", collection.BatteryInformation.GetTemperature())
		log.Printf("Battery Max Charging Current: %.2f A", collection.BatteryInformation.GetMaxChargingCurrent())
		log.Printf("Battery Max Discharging Current: %.2f A", collection.BatteryInformation.GetMaxDischargingCurrent())
		log.Printf("Battery State of Charge: %d", collection.BatteryInformation.GetStateOfCharge())
		log.Printf("Battery State of Health: %d", collection.BatteryInformation.GetStateOfHealth())
		log.Printf("Battery State: %d", collection.BatteryInformation.GetState())
		log.Printf("Battery StandBy: %t", collection.BatteryInformation.IsStandBy())
		log.Printf("Battery Charging: %t", collection.BatteryInformation.IsCharging())
		log.Printf("Battery Discharging: %t", collection.BatteryInformation.IsDischarging())
		log.Printf("Battery Charge Cut-Off Voltage: %.2f V", collection.BatteryInformation.GetChargeCutOffVoltage())
		log.Printf("Battery Discharge Cut-Off Voltage: %.2f V", collection.BatteryInformation.GetDischargeCutOffVoltage())
		log.Printf("Battery Charge / Discharge Times: %d", collection.BatteryInformation.GetChargeDischargeTimes())
		log.Printf("Battery Pressure: %d", collection.BatteryInformation.GetPressure())
		log.Printf("Battery Warning: %d", collection.BatteryInformation.GetWarning())
		log.Printf("Battery Error: %d", collection.BatteryInformation.GetError())
		log.Printf("Battery Communication Status: %d", collection.BatteryInformation.GetComStatus())
	}

	if collection.ControlCabinetInformation != nil {
		log.Println("=================== Control Cabinet Information ====================")
		log.Printf("Device type: %d", collection.ControlCabinetInformation.GetDeviceType())
		log.Printf("DSP High Version: %d", collection.ControlCabinetInformation.GetDSPHighVersion())
		log.Printf("DSP Low Version: %d", collection.ControlCabinetInformation.GetDSPLowVersion())
		log.Printf("Inverter Internal Temperature : %.2f °C", collection.ControlCabinetInformation.GetInternalTemperature())
		log.Printf("Module A1 Temperature : %.2f °C", collection.ControlCabinetInformation.GetModuleA1Temperature())
		log.Printf("Module B1 Temperature : %.2f °C", collection.ControlCabinetInformation.GetModuleB1Temperature())
		log.Printf("Module C1 Temperature : %.2f °C", collection.ControlCabinetInformation.GetModuleC1Temperature())
		log.Printf("Parallel Address : %d", collection.ControlCabinetInformation.GetParallelAddress())

		log.Printf("PV1 Voltage: %.2f V", collection.ControlCabinetInformation.GetPV1Voltage())
		log.Printf("PV1 Current: %.2f A", collection.ControlCabinetInformation.GetPV1Current())
		log.Printf("PV1 Power: %d W", collection.ControlCabinetInformation.GetPV1Power())
		log.Printf("PV2 Voltage: %.2f V", collection.ControlCabinetInformation.GetPV2Voltage())
		log.Printf("PV2 Current: %.2f A", collection.ControlCabinetInformation.GetPV2Current())
		log.Printf("PV2 Power: %d W", collection.ControlCabinetInformation.GetPV2Power())

		log.Printf("Grid Phase A Voltage: %.2f V", collection.ControlCabinetInformation.GetGridPhaseAVoltage())
		log.Printf("Grid Phase A Current: %.2f A", collection.ControlCabinetInformation.GetGridPhaseACurrent())
		log.Printf("Grid Phase B Voltage: %.2f V", collection.ControlCabinetInformation.GetGridPhaseBVoltage())
		log.Printf("Grid Phase B Current: %.2f A", collection.ControlCabinetInformation.GetGridPhaseBCurrent())
		log.Printf("Grid Phase C Voltage: %.2f V", collection.ControlCabinetInformation.GetGridPhaseCVoltage())
		log.Printf("Grid Phase C Current: %.2f A", collection.ControlCabinetInformation.GetGridPhaseCCurrent())

		log.Printf("Grid Frequency: %.2f Hz", collection.ControlCabinetInformation.GetGridFrequency())
		log.Printf("Power Factor: %.3f", collection.ControlCabinetInformation.GetPowerFactor())

		log.Printf("Grid Line AB Voltage: %.2f V", collection.ControlCabinetInformation.GetGridLineABVoltage())
		log.Printf("Grid Line BC Voltage: %.2f V", collection.ControlCabinetInformation.GetGridLineBCVoltage())
		log.Printf("Grid Line CA Voltage: %.2f V", collection.ControlCabinetInformation.GetGridLineCAVoltage())

		log.Printf("Load Phase A Voltage: %.2f V", collection.ControlCabinetInformation.GetLoadPhaseAVoltage())
		log.Printf("Load Phase A Current: %.2f A", collection.ControlCabinetInformation.GetLoadPhaseACurrent())
		log.Printf("Load Phase B Voltage: %.2f V", collection.ControlCabinetInformation.GetLoadPhaseBVoltage())
		log.Printf("Load Phase B Current: %.2f A", collection.ControlCabinetInformation.GetLoadPhaseBCurrent())
		log.Printf("Load Phase C Voltage: %.2f V", collection.ControlCabinetInformation.GetLoadPhaseCVoltage())
		log.Printf("Load Phase C Current: %.2f A", collection.ControlCabinetInformation.GetLoadPhaseCCurrent())

		log.Printf("Load Frequency: %.2f Hz", collection.ControlCabinetInformation.GetLoadFrequency())

		log.Printf("Grid Active Power : %d W", collection.ControlCabinetInformation.GetGridActivePower())
		log.Printf("Grid Reactive Power : %d Var", collection.ControlCabinetInformation.GetGridReactivePower())
		log.Printf("System Apparent Power : %d VA", collection.ControlCabinetInformation.GetSystemApparentPower())

		log.Printf("Load Power Factor : %.3f", collection.ControlCabinetInformation.GetLoadPowerFactor())
		log.Printf("Load Active Power : %d W", collection.ControlCabinetInformation.GetLoadActivePower())
		log.Printf("Load Reactive Power : %d Var", collection.ControlCabinetInformation.GetLoadReactivePower())
		log.Printf("Load Apparent Power : %d VA", collection.ControlCabinetInformation.GetLoadApparentPower())

		log.Printf("Battery Voltage : %.2f V", collection.ControlCabinetInformation.GetBatteryVoltage())
		log.Printf("Battery Current : %.2f A", collection.ControlCabinetInformation.GetBatteryCurrent())
		log.Printf("Battery #1 Current : %.2f A", collection.ControlCabinetInformation.GetBatteryNo1Current())
		log.Printf("Battery #2 Current : %.2f A", collection.ControlCabinetInformation.GetBatteryNo2Current())

		log.Printf("Bus DC Positive Voltage : %.2f V", collection.ControlCabinetInformation.GetBusDCPositiveVoltage())
		log.Printf("Bus DC Negative Voltage : %.2f V", collection.ControlCabinetInformation.GetBusDCNegativeVoltage())
		log.Printf("Bus DC Bilateral Voltage : %.2f V", collection.ControlCabinetInformation.GetBusDCBilateralVoltage())
		log.Printf("Bus DC Power : %d W", collection.ControlCabinetInformation.GetBusDCPower())

		log.Printf("Backup Bus DC Positive Voltage : %.2f V", collection.ControlCabinetInformation.GetBackupBusDCPositiveVoltage())
		log.Printf("Backup Bus DC Negative Voltage : %.2f V", collection.ControlCabinetInformation.GetBackupBusDCNegativeVoltage())

		log.Printf("Work Efficiency : %d", collection.ControlCabinetInformation.GetWorkEfficiency())
	}

	if collection.TotalPowerData != nil {
		log.Println("=================== Total Power Data ====================")
		log.Printf("PV Daily Power generation : %.2f W", collection.TotalPowerData.GetDailyPVPowerGeneration())
		log.Printf("PV Total Power generation : %d W", collection.TotalPowerData.GetTotalPVPowerGeneration())
		log.Printf("Load Daily Power consumption : %.2f W", collection.TotalPowerData.GetDailyLoadPowerConsumption())
		log.Printf("Load Total Power consumption : %d W", collection.TotalPowerData.GetTotalLoadPowerConsumption())
		log.Printf("Money Daily saving : %.2f", collection.TotalPowerData.GetDailyMoneySaving())
		log.Printf("Money Total saving : %.2f", collection.TotalPowerData.GetTotalMoneySaving())
		log.Printf("Datapoints for PV Power Day (48 points) : %v", collection.TotalPowerData.GetDayPVPower())
		log.Printf("Datapoints for PV Power Month (31 points) : %v", collection.TotalPowerData.GetMonthPVPower())
		log.Printf("Datapoints for PV Power Year (12 points) : %v", collection.TotalPowerData.GetYearPVPower())
		log.Printf("Datapoints for PV Power History (5 points) : %v", collection.TotalPowerData.GetHistoryPVPower())
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

	if *proxyEnabled {
	}

	inverterClientWaitingGroup.Wait()
	os.Exit(0)
}
