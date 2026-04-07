package main

import (
	"context"
	"log"

	"github.com/bougou/go-ipmi"
)

func main() {
	// a := app.New()
	// w := a.NewWindow("Hello World!")
	// w.SetContent(widget.NewLabel("Hello World!"))
	// w.ShowAndRun()


	client, err := ipmi.NewClient(Hostname, Port, Username, Password)
	client.WithInterface(ipmi.InterfaceLanplus)
	ctx := context.Background()
	if err != nil {
		panic(err)
	}

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Unable to stablish connection: %v", err)
	}

	info, err := client.GetDeviceID(ctx)
	if err != nil {
		log.Fatalf("Unable to get device ID: %v", err)
	}

	log.Printf("Device ID: %d\n", info.DeviceID)
	log.Printf("FirmwareVision: %s\n", info.FirmwareVersionStr())
	log.Printf("Manufacturer ID: %d\n", info.ManufacturerID)
	log.Printf("IPMIVersion: %d.%d\n", info.MajorIPMIVersion, info.MinorIPMIVersion)

	sdrs, err := client.GetSDRs(ctx)
	if err != nil {
		log.Fatalf("Unable to get SDRs: %v", err)
	}
	for _, sdr := range sdrs {
		// switch sdr.SensorName() {
		// case "FAN1":
			log.Printf(sdr.String())
		// }
	}
}
