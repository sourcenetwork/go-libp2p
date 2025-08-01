package autonat

import (
	"github.com/libp2p/go-libp2p/core/network"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

var _ network.Notifiee = (*AmbientAutoNAT)(nil)

// Listen is part of the network.Notifiee interface
func (as *AmbientAutoNAT) Listen(_ network.Network, _ ma.Multiaddr) {}

// ListenClose is part of the network.Notifiee interface
func (as *AmbientAutoNAT) ListenClose(_ network.Network, _ ma.Multiaddr) {}

// Connected is part of the network.Notifiee interface
func (as *AmbientAutoNAT) Connected(_ network.Network, c network.Conn) {
	if c.Stat().Direction == network.DirInbound &&
		manet.IsPublicAddr(c.RemoteMultiaddr()) {
		select {
		case as.inboundConn <- c:
		default:
		}
	}
}

// Disconnected is part of the network.Notifiee interface
func (as *AmbientAutoNAT) Disconnected(_ network.Network, _ network.Conn) {}
