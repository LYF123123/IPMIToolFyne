package main

import (
	"IPMITOOLFYNE/config"
	"context"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bougou/go-ipmi"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World!")
	w.Resize(fyne.NewSize(500, 500))

	message := widget.NewLabel("System Initialization")
	button_connect := widget.NewButton("Connect", func() {
		clientConfig := config.GetClientConfig()
		client, err := ipmi.NewClient(clientConfig.Host, clientConfig.Port, clientConfig.User, clientConfig.Pass)
		client.WithInterface(ipmi.InterfaceLanplus)
		ctx := context.Background()
		if err != nil {
			panic(err)
		}

		if err := client.Connect(ctx); err != nil {
			message.SetText("Unable to stablish connection!")
			log.Fatalf("Unable to stablish connection: %v", err)
		}
		message.SetText("Connected to the BMC Device!")
	})
	w.SetContent(container.NewVBox(message, button_connect))
	w.ShowAndRun()

	// clientConfig := config.GetClientConfig()
	// client, err := ipmi.NewClient(clientConfig.Host, clientConfig.Port, clientConfig.User, clientConfig.Pass)
	// client.WithInterface(ipmi.InterfaceLanplus)
	// ctx := context.Background()
	// if err != nil {
	// 	panic(err)
	// }

	// if err := client.Connect(ctx); err != nil {
	// 	log.Fatalf("Unable to stablish connection: %v", err)
	// }

	// info, err := client.GetDeviceID(ctx)
	// if err != nil {
	// 	log.Fatalf("Unable to get device ID: %v", err)
	// }

	// log.Printf("Device ID: %d\n", info.DeviceID)
	// log.Printf("FirmwareVision: %s\n", info.FirmwareVersionStr())
	// log.Printf("Manufacturer ID: %d\n", info.ManufacturerID)
	// log.Printf("IPMIVersion: %d.%d\n", info.MajorIPMIVersion, info.MinorIPMIVersion)

	// sdrs, err := client.GetSDRs(ctx)
	// if err != nil {
	// 	log.Fatalf("Unable to get SDRs: %v", err)
	// }
	// for _, sdr := range sdrs {
	// 	// switch sdr.SensorName() {
	// 	// case "FAN1":
	// 	log.Printf(sdr.String())
	// 	// }
	// }
}
