package drivers

import (
	"github.com/Liar233/Scheduler/model"
	"github.com/Liar233/Scheduler/config"
	"net"
	"fmt"
)

type TcpChannel struct {
	conn *net.TCPConn
}

func (tc *TcpChannel) Connect(config *config.ChannelConfig) error {
	address := fmt.Sprintf("%s:%s", config.Options["host"], config.Options["port"])

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)

	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil ,tcpAddr)

	if err != nil {
		return err
	}

	tc.conn = conn

	return nil
}

func (tc *TcpChannel) Fire(e *model.Event) error {
	_, err := tc.conn.Write([]byte(e.Payload))
	
	if err != nil {
		return err
	}

	return nil
}

func (tc *TcpChannel) Name() string {
	return "tcp_channel"
}

