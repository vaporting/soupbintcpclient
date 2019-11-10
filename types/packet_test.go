package types

import (
	"testing"

	"encoding/binary"

	"github.com/stretchr/testify/assert"
)

func TestNewPacket(t *testing.T) {
	text := []byte("123")

	// SUT
	pkt, _ := NewPacket(PktTDbg, text)

	// Assertion
	assert.Equal(t, uint16(1+len(text)), pkt.len)
	assert.Equal(t, PktTDbg, pkt.pType)
	assert.Equal(t, text, pkt.text)
}

func TestPakcetBytes(t *testing.T) {
	text := []byte("123")
	pkt, _ := NewPacket(PktTCliHB, text)
	// SUT
	result := pkt.Bytes()

	expRes := make([]byte, 2+1+len(text))
	temp := make([]byte, 2)
	binary.BigEndian.PutUint16(temp, pkt.len)
	copy(expRes[:2], temp)
	expRes[2] = pkt.pType[0]
	copy(expRes[3:], text)

	// Assertion
	assert.Equal(t, expRes, result)
}
