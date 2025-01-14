package api

import (
	"fmt"
	"io"
	"net"
)

type (
	addr    string
	network string
)

type Dialer struct {
	Addr    addr
	Network network
}

func New(network network, port addr) Dialer {
	return Dialer{
		port,
		network,
	}
}

func (d Dialer) Dial() ([]byte,error) {
	con,conErr:=net.Dial(string(d.Network),string(d.Addr))

	if conErr!=nil{
		return nil,fmt.Errorf("error during connecting to network \n %v",conErr)
	}

	defer con.Close()

	data,dataErr:=	io.ReadAll(con)

	if dataErr!=nil{
		return nil, fmt.Errorf("error during reading message %v",dataErr)
	}

	return data,nil
}
