package Receiver

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/sijiramakun/seapick/utils"
)

func handleIncomingRequests(conn net.Conn) {
	println("Received a request: " + conn.RemoteAddr().String())
	headerBuffer := make([]byte, 1024)

	_, err := conn.Read(headerBuffer)
	utils.CheckError(err)

	var name string
	var reps uint32

	if headerBuffer[0] == byte(1) && headerBuffer[1023] == byte(0) {
		reps = binary.BigEndian.Uint32(headerBuffer[1:5])
		lengthOfName := binary.BigEndian.Uint32(headerBuffer[5:9])
		name = string(headerBuffer[9 : 9+lengthOfName])
	} else {
		log.Fatal("Invalid header")
	}

	conn.Write([]byte("Header Received"))

	dataBuffer := make([]byte, 1024)

	errr := os.MkdirAll("./received/", os.ModePerm) //create folder first
	utils.CheckError(errr)

	file, err := os.Create("./received/" + name) //create file next
	utils.CheckError(err)

	fmt.Println(name, reps, "PINPOINT THISSSSSSSSSSSSSSSSSSSSS")

	for i := 0; i < int(reps); i++ {
		_, err := conn.Read(dataBuffer)
		utils.CheckError(err)

		if dataBuffer[0] == byte(0) && dataBuffer[1023] == byte(1) {
			segmentNumber := dataBuffer[1:5]
			fmt.Printf("Segment Number: %d\n", binary.BigEndian.Uint32(segmentNumber))
			length := binary.BigEndian.Uint32(dataBuffer[5:9])
			fmt.Printf("File Data: %s\n", hex.EncodeToString(dataBuffer[9:9+length]))
			file.Write(dataBuffer[9 : 9+length])
		} else {
			log.Fatal("Invalid Segment")
		}

		conn.Write([]byte("Segment Received"))

	}

	time := time.Now().UTC().Format("Monday, 02-Jan-06 15:04:05 MST")

	conn.Write([]byte(time))

	file.Close()
	conn.Close()

	return
}
