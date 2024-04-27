package Sender

import (
	"fmt"
	"log"
	"net"

	"github.com/sijiramakun/seapick/utils"
	"github.com/urfave/cli/v2"
)

var TYPE = "tcp"
var recieverAddr string
var fileAddr string

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "reciever",
		Aliases:     []string{"r"},
		Usage:       "Address of the reciever",
		Destination: &recieverAddr,
	},
	&cli.StringFlag{
		Name:        "file",
		Aliases:     []string{"f"},
		Usage:       "Address of the file to send",
		Destination: &fileAddr,
	},
}

func Send(ctx *cli.Context) {
	fmt.Println(recieverAddr, fileAddr)
	if recieverAddr == "" {
		log.Fatalf("Receiver address is missing, input the --reciever flag")
	}
	tcpServer, err := net.ResolveTCPAddr(TYPE, recieverAddr)
	utils.CheckError(err)

	conn, err := net.DialTCP(TYPE, nil, tcpServer)

	if fileAddr == "" || !isFile(fileAddr) {
		log.Fatalf("File address is missing, input the --reciever flag")
	}

	sendFIle(fileAddr, conn)

	received := make([]byte, 1024)

	_, errr := conn.Read(received)

	utils.CheckError(errr)

	println(received)

}
