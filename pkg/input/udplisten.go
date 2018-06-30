package input

import (
	"io"
	"net"
)

// UDPListen creates a UDP listener and returns ReadCloser
func UDPListen(addr string) (io.ReadCloser, error) {
	a, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	return net.ListenUDP("udp", a)
}
