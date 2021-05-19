package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("vim-go")
}

func checkAlive(ip string) bool {
	alive := false
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, "22"), time.Second*30)
	if err == nil {
		alive = true
	}
	return alive
}
