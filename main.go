package main

import (
	"log"

	"github.com/karalabe/hid"
)

func main() {
	devices := hid.Enumerate(0x1a61, 0x3650)
	if len(devices) == 0 {
		log.Fatal("Device Not FOUND")
	}

	dev, err := devices[0].Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dev.Close()

	s := buildHdiPacket(0x04, []byte{0x00}, 0)

	var b []byte
	n, err := dev.Read(b)
	log.Println(n)
	if err != nil {
		log.Fatalf("%#+v", err)
	}

	log.Println(b)

	n, err = dev.Write(s[:])
	log.Println(n)
	if err != nil {
		log.Fatalf("%#+v", err)
	}

	_, err = dev.Read(b)
	if err != nil {
		log.Fatal(err)
	}
}

func buildHdiPacket(command byte, data []byte, length int) [65]byte {
	ret := [65]byte{}
	ret[0] = 0x00
	ret[1] = 0x04
	ret[2] = byte(length)

	return ret
}
