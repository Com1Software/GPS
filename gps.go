package gps

import "fmt"

func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}


func GetSerialPort() string {
    port := "com1"
    return port
}
