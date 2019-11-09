package net

import (
	"net"

	"soupbintcp/types"

	"time"
)

// NewKeepAliveTCPConn creates KeepAliveTCPConn
func NewKeepAliveTCPConn(addr string, port string) (*KeepAliveTCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr+":"+port)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	hbPkt, _ := types.NewPacket(types.PktTCliHB, []byte(""))
	return &KeepAliveTCPConn{conn: conn, tube: make(chan []byte), hbPkt: *hbPkt}, err

}

// KeepAliveTCPConn is a keep alive tcp connection
type KeepAliveTCPConn struct {
	conn  *net.TCPConn
	tube  chan []byte
	hbPkt types.Packet
}

// Start start sending message to Server
func (conn *KeepAliveTCPConn) Start() error {
	baseT := time.Now().UnixNano() / int64(time.Millisecond)
	var err error = nil
	for {
		select {
		case msg := <-conn.tube:
			_, err := conn.conn.Write(msg)
			if err != nil {
				break
			}
			baseT = time.Now().UnixNano() / int64(time.Millisecond)
		default:
			curT := time.Now().UnixNano() / int64(time.Millisecond)
			if curT-baseT > 999 {
				_, err := conn.conn.Write(conn.hbPkt.Bytes())
				if err != nil {
					break
				}
				baseT = time.Now().UnixNano() / int64(time.Millisecond)
			}
		}
	}
	return err
}

// Close closes the connection
func (conn *KeepAliveTCPConn) Close() error {
	return conn.conn.Close()
}

// GetTube returns feeding channel
func (conn *KeepAliveTCPConn) GetTube() chan []byte {
	return conn.tube
}
