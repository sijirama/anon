package Sender

import (
	"encoding/binary"
	"net"

	"fmt"
	"os"

	"github.com/sijiramakun/seapick/utils"
)

func sendFIle(path string, conn *net.TCPConn) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)

	utils.CheckError(err)

	header := prepareFIleMetaData(file)

	dataBuffer := make([]byte, 1014)

	headerBuffer := []byte{1}

	segmentBuffer := []byte{0}

	temp := make([]byte, 4) // Temporary buffer for uint32

	received := make([]byte, 100)

	for i := 0; i < int(header.reps); i++ {
		n, _ := file.ReadAt(dataBuffer, int64(i*1014))

		if i == 0 { //send the header in the first request

			//number of segments
			binary.BigEndian.PutUint32(temp, header.reps)
			headerBuffer = append(headerBuffer, temp...)

			//length of name
			binary.BigEndian.PutUint32(temp, uint32(len(header.name)))
			headerBuffer = append(headerBuffer, temp...)

			//name of file
			headerBuffer = append(headerBuffer, []byte(header.name)...)

			headerBuffer = append(headerBuffer, 0)

			_, err := conn.Write(headerBuffer)

			utils.CheckError(err)

			_, err = conn.Read(received)

			utils.CheckError(err)

			println(string(received))

		}

		//rep number
		binary.BigEndian.PutUint32(temp, uint32(i))
		segmentBuffer = append(segmentBuffer, temp...)

		//length of data
		binary.BigEndian.PutUint32(temp, uint32(n))
		segmentBuffer = append(segmentBuffer, temp...)

		//data
		segmentBuffer = append(segmentBuffer, dataBuffer...)

		segmentBuffer = append(segmentBuffer, 1)

		_, err = conn.Write(segmentBuffer)

		utils.CheckError(err)

		_, err = conn.Read(received)
		fmt.Println(string(received))

		utils.CheckError(err)

		//reset segment buffer
		segmentBuffer = []byte{0}

	}

}
