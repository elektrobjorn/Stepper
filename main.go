package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.bug.st/serial"
)

func main() {
	var selectedPort string

	// Get list of open COM ports
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	// Select COM port
	myApp := app.New()
	myWindow := myApp.NewWindow("Select COM port")
	fmt.Printf("myWindow done\n")

	radio := widget.NewRadioGroup(ports, func(selectedPort string) {
		fmt.Printf("Radio set to %v\n", selectedPort)
	})
	fmt.Printf("Radio done\n")
	// Open COM port
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	fmt.Printf("Mode done\n")
	fmt.Printf("selectedPort: %v\n", selectedPort)
	port, err := serial.Open(selectedPort, mode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Open done\n")
	log.Println("Selected", selectedPort)
	defer port.Close()

	myWindow.SetContent(container.NewVBox(radio))
	myWindow.ShowAndRun()
}
