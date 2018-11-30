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
	config *config.ChannelConfig
}

func (tc *TcpChannel) Init(config *config.ChannelConfig, name string) {
	tc.name = name
	tc.config = config
}

func (tc *TcpChannel) Connect() error {
	address := fmt.Sprintf("%s:%d", tc.config.Options["host"], tc.config.Options["port"])

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)

	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		return err
	}

	tc.conn = conn

	return nil
}

func (tc *TcpChannel) Fire(e *model.Event) error {
	err := tc.Connect()

	if err != nil {
		return err
	}

	_, err = tc.conn.Write([]byte(e.Payload + "\n"))

	if err != nil {
		return err
	}

	tc.conn.Close()

	return nil
}

func (tc *TcpChannel) Name() string {
	return tc.name
}

func NewTcpChannel(conf *config.ChannelConfig, name string) *TcpChannel {
	channel := &TcpChannel{}

	channel.Init(conf, name)

	return channel
}
