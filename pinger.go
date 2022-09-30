package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var hostName string
	fmt.Println("Enter the ip addres: ")
	fmt.Scanln(&hostName)
	portNum := "80"
	seconds := 5
	timeOut := time.Duration(seconds) * time.Second

	conn, err := net.DialTimeout("tcp", hostName+":"+portNum, timeOut)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Connection established between %s and localhost with time out of %d seconds.\n", hostName, int64(seconds))
	fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
	fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())

}