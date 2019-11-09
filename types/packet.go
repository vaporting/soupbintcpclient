package types

import (
	"encoding/binary"

	"errors"
)

// IPacket defines the packet interface
type IPacket interface {
	Bytes() []byte
}

// Packet is used to store each packet info.
type Packet struct {
	len   uint16
	pType string
	text  []byte

	// interface
	IPacket
}

// Bytes returns bytes with all info in pkt
func (pkt *Packet) Bytes() []byte {
	pktBytes := make([]byte, 2+1+len(pkt.text))
	temp := make([]byte, 2)
	binary.BigEndian.PutUint16(temp, pkt.len)
	copy(pktBytes[:2], temp)
	pktBytes[2] = pkt.pType[0]
	copy(pktBytes[3:], pkt.text)
	return pktBytes
}

// NewPacket creates Packet
func NewPacket(pType string, text []byte) (*Packet, error) {
	sum := len(pType) + len(text)
	if sum > 0xffff {
		return nil, errors.New("size overflow")
	}
	return &Packet{len: uint16(sum), pType: pType, text: text}, nil
}
