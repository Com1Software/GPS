package gps

import "fmt"
import "go.bug.st/serial"

func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}


func GetSerialPort() string {
    ports, err := serial.GetPortsList()
	if err != nil {
		fmt.Printf("Port Error: %v\n", port)
	}
	if len(ports) == 0 {
		fmt.Printf("Port Not Found \n")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
    return port
}
