package Sender

/*
the header model
Start (all 1s) - 1 byte
Reps (number of segments) - 4 bytes
Lengthofname - 4 bytes
Name - `lengthofname` bytes
End (all 0s) - 1 byte;

the segment model
Start (all 0s) - 1 byte
Segment number - 4 bytes
Lengthofdata - 4 bytes
Data - `lengthofdata` bytes
End (all 1s) - 1 byte

*/

import (
	"encoding/binary"
	"fmt"
	"github.com/sijiramakun/seapick/utils"
	"net"
	"os"
)

var metadatasize = 10
var headerSize = utils.HeaderSize - metadatasize //1014
var DataSize = utils.DataSize - metadatasize

func sendFIle(path string, conn *net.TCPConn) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)

	utils.CheckError(err)

	header := prepareFIleMetaData(file)

	dataBuffer := make([]byte, DataSize)

	headerBuffer := []byte{1}

	segmentBuffer := []byte{0}

	temp := make([]byte, 4) // Temporary buffer for uint32

	received := make([]byte, 100)

	for i := 0; i < int(header.reps); i++ {
		//n, _ := file.ReadAt(dataBuffer, int64(i*1014))

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

		n, _ := file.ReadAt(dataBuffer, int64(i*DataSize))

		//rep number
		binary.BigEndian.PutUint32(temp, uint32(i))
		segmentBuffer = append(segmentBuffer, temp...)

		//fmt.Println("length of data is ", n)

		//length of data
		binary.BigEndian.PutUint32(temp, uint32(n))
		segmentBuffer = append(segmentBuffer, temp...)

		//data
		segmentBuffer = append(segmentBuffer, dataBuffer...)

		segmentBuffer = append(segmentBuffer, 1)

		// fmt.Println("segment buffer of data is ", segmentBuffer[len(segmentBuffer)-1])
		// fmt.Println("length of segment buffer of data is ", len(segmentBuffer))

		_, err = conn.Write(segmentBuffer)

		utils.CheckError(err)

		_, err = conn.Read(received)
		//fmt.Println(string(received))

		utils.CheckError(err)

		//reset segment buffer
		segmentBuffer = []byte{0}

	}

	fmt.Println("File sent successfully")

}
