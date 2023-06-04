package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"net"
	"time"
)

type NTPPacket struct {
    Settings       uint8  
    Stratum        uint8  
    Poll           int8   
    Precision      int8   
    RootDelay      uint32 
    RootDispersion uint32 
    ReferenceID    uint32 
    RefTimeSec     uint32 
    RefTimeFrac    uint32 
    OrigTimeSec    uint32 
    OrigTimeFrac   uint32 
    RxTimeSec      uint32 
    RxTimeFrac     uint32 
    TxTimeSec      uint32 
    TxTimeFrac     uint32 
}

func connectToNTPServer(){
	const host = "2.ph.pool.ntp.org:123"

	conn, err := net.Dial("udp", host)
	if err != nil {
		log.Fatal("NTP Server Connection failed : ", err)
	}

	defer conn.Close()

	if err := conn.SetDeadline(
	time.Now().Add(10 * time.Second)); err != nil {
		log.Fatal("Failed to set connection deadline : ", err)
	}

	req := &NTPPacket{Settings: 0x1B}

	if err := binary.Write(conn, binary.BigEndian, req); err != nil {
		log.Fatalf("Failed to send request to NTP Server: %v", err)
	}

	res := &NTPPacket{}
	if err := binary.Read(conn, binary.BigEndian, res); err != nil {
		log.Fatalf("Failed to read response from NTP Server: %v", err)
	}

	year1900 := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	year1970 := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

	secondsDiff := year1970.Unix() - year1900.Unix()
	absoluteDiff := math.Abs(float64(secondsDiff))

	fmt.Println("Seconds passed from 1900 to 1970", absoluteDiff)

	secs := float64(res.TxTimeSec) - absoluteDiff

	fmt.Printf("Time from NTP Server %v\n", time.Unix(int64(secs), 0))
}

func main(){
	connectToNTPServer()
}