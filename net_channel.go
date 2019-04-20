packge "netChannel"

import (
	"net"
	"time"
)

type NetChannel struct {
	chnnelID int
}

func (nchan *NetChannel) Read(b []byte) (n int, err error) {

}

func (nchan *NetChannel) Write(b []byte) (n int, err error) {

}

func (nchan *NetChannel) Close() error {

}

func (nchan *NetChannel) LocalAddr() net.Addr {

}

func (nchan *NetChannel) RemoteAddr() net.Addr {

}

func (nchan *NetChannel) SetDeadLine(t time.Time) error {

}

func (nchan *NetChannel) SetReadDeadLine(t time.Time) error {

}

func (nchan *NetChannel) SetWriteDeadLine(t time.TIme) error {

}