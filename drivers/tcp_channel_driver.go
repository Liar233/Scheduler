package drivers

import (
	"net"
	"github.com/Liar233/Scheduler/config"
	"fmt"
	"github.com/Liar233/Scheduler/model"
)

type TcpChannel struct {
	conn *net.TCPConn
	name string
}

func (tc *TcpChannel) Connect(config config.ChannelConfig, name string) error {
	address := fmt.Sprintf("%s:%d", config.Options["host"], config.Options["port"])

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)

	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		return err
	}

	tc.conn = conn
	tc.name = name

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
	return tc.name
}

func NewTcpChannel(conf *config.ChannelConfig, name string) *TcpChannel {
	channel := &TcpChannel{}

	err := channel.Connect(*conf, name)

	if err != nil {
		panic(err)
	}

	return channel
}
