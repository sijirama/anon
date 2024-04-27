package Receiver

import (
	"fmt"
	"net"

	//tea "github.com/charmbracelet/bubbletea"
	"github.com/sijiramakun/seapick/components"
	"github.com/sijiramakun/seapick/utils"
)

func Receive() {

	ip := getLocalIP()
	port := findAvailablePort()
	addr := ip + ":" + port

	listener, err := net.Listen("tcp", addr)
	utils.CheckError(err)

	defer listener.Close()

	//INFO: spinner component
	message := fmt.Sprintf("Server is Listening on %s", addr)
	components.Spinner(message)

	for {
		conn, err := listener.Accept()
		utils.CheckError(err)
		go handleIncomingRequests(conn)
	}

}
