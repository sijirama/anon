package Receiver

import (
	"encoding/binary"
	"log"
	"net"
	"os"
	"time"

	"github.com/sijiramakun/seapick/utils"
)

func handleIncomingRequests(conn net.Conn) {
	println("Received a request: " + conn.RemoteAddr().String())
	headerBuffer := make([]byte, utils.HeaderSize)

	_, err := conn.Read(headerBuffer)
	utils.CheckError(err)

	var name string
	var reps uint32

	if headerBuffer[0] == byte(1) && headerBuffer[utils.HeaderSize-1] == byte(0) {
		reps = binary.BigEndian.Uint32(headerBuffer[1:5])
		lengthOfName := binary.BigEndian.Uint32(headerBuffer[5:9])
		name = string(headerBuffer[9 : 9+lengthOfName])
	} else {
		log.Fatal("Invalid header")
	}

	conn.Write([]byte("Header Received"))

	dataBuffer := make([]byte, utils.DataSize)

	errr := os.MkdirAll("./received/", os.ModePerm) //create folder first
	utils.CheckError(errr)

	file, err := os.Create("./received/" + name) //create file next
	utils.CheckError(err)

	println("Sending ", utils.DataSize, " bytes per chunck.")
	for i := 0; i < int(reps); i++ {
		_, err := conn.Read(dataBuffer)
		utils.CheckError(err)

		//remove this crap
		// if dataBuffer[0] == byte(0) {
		// 	length := binary.BigEndian.Uint32(dataBuffer[5:9])
		// 	segmentNumber := dataBuffer[1:5]
		// 	seg := fmt.Sprintf("%d", binary.BigEndian.Uint32(segmentNumber))
		// 	println(length, seg, len(dataBuffer))
		// }

		//fmt.Println(name, "Number of reps", reps, " Rep number", i, "PINPOINT THISSSSSSSSSSSSSSSSSSSSS")
		if dataBuffer[0] == byte(0) && dataBuffer[utils.DataSize-1] == byte(1) {

			//segmentNumber := dataBuffer[1:5]
			//fmt.Printf("Segment Number: %d\n", binary.BigEndian.Uint32(segmentNumber))
			//fmt.Printf("File Data: %s\n", hex.EncodeToString(dataBuffer[9:9+length]))

			length := binary.BigEndian.Uint32(dataBuffer[5:9])
			file.Write(dataBuffer[9 : 9+length])
		} else {
			log.Fatal("Invalid Segment")
			dataBuffer = []byte{0}

		}

		conn.Write([]byte("Segment Received"))

	}

	time := time.Now().UTC().Format("Monday, 02-Jan-06 15:04:05 MST")

	conn.Write([]byte(time))

	file.Close()
	conn.Close()

	return
}
