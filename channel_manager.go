package netChannel

import (
	"net"
)

type ChannelManager struct {
	conn net.Conn
	channels [int]NetChannel 
}

func (cm *ChannelManager) OpenChannel() {

}

func (cm *ChannelManager) CloseChannel() {

}

func (cm *ChannelManager) GetChannel() {

}

func (cm *ChannelManager) Close() {

}