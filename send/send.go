package Sender

import (
	"log"
	"net"

	"github.com/sijiramakun/seapick/utils"
	"github.com/urfave/cli/v2"
)

var TYPE = "tcp"
var receiverAddr string
var fileAddr string

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "receiver",
		Aliases:     []string{"r"},
		Usage:       "Address of the receiver",
		Destination: &receiverAddr,
	},
	&cli.StringFlag{
		Name:        "file",
		Aliases:     []string{"f"},
		Usage:       "Address of the file to send",
		Destination: &fileAddr,
	},
}

func Send(ctx *cli.Context) {

	if receiverAddr == "" {
		log.Fatalf("Receiver address is missing, input the --receiver flag")
	}
	tcpServer, err := net.ResolveTCPAddr(TYPE, receiverAddr)
	utils.CheckError(err)

	conn, err := net.DialTCP(TYPE, nil, tcpServer)

	if fileAddr == "" || !isFile(fileAddr) {
		log.Fatalf("File address is missing, input the --receiver flag")
	}

	sendFIle(fileAddr, conn)

	received := make([]byte, 1024)

	_, errr := conn.Read(received)

	utils.CheckError(errr)

	println(received)

}
