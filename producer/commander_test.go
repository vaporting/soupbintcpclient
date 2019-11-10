package producer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"soupbintcpclient/types"

	"sort"
)

func TestCommanderPushMsg(t *testing.T) {
	c := make(chan []byte)
	cmdr := NewCommander(c)
	pkt, _ := types.NewPacket(types.PktTDbg, []byte("1000"))
	msg1 := &types.Message{Msec: 1000, Pkt: pkt}
	cmdr.PushMsg(msg1) // SUT
	pkt, _ = types.NewPacket(types.PktTDbg, []byte("0"))
	msg2 := &types.Message{Msec: 0, Pkt: pkt}
	cmdr.PushMsg(msg2) // SUT

	// Assertion
	assert.Equal(t, 2, len(cmdr.msgQ))
	assert.Equal(t, cmdr.msgQ[0].Msec, msg1.Msec)
	assert.Equal(t, cmdr.msgQ[0].Pkt, msg1.Pkt)
	assert.Equal(t, cmdr.msgQ[1].Msec, msg2.Msec)
	assert.Equal(t, cmdr.msgQ[1].Pkt, msg2.Pkt)
}

func TestCommanderOrder(t *testing.T) {
	c := make(chan []byte)
	cmdr := NewCommander(c)
	expMq := types.MessageQueue{}
	pkt, _ := types.NewPacket(types.PktTDbg, []byte("1000"))
	msg := &types.Message{Msec: 1000, Pkt: pkt}
	expMq = append(expMq, msg)
	cmdr.PushMsg(msg)
	pkt, _ = types.NewPacket(types.PktTDbg, []byte("0"))
	msg = &types.Message{Msec: 0, Pkt: pkt}
	expMq = append(expMq, msg)
	cmdr.PushMsg(msg)
	pkt, _ = types.NewPacket(types.PktTDbg, []byte("1500"))
	msg = &types.Message{Msec: 1500, Pkt: pkt}
	expMq = append(expMq, msg)
	cmdr.PushMsg(msg)
	pkt, _ = types.NewPacket(types.PktTDbg, []byte("3000"))
	msg = &types.Message{Msec: 3000, Pkt: pkt}
	expMq = append(expMq, msg)
	cmdr.PushMsg(msg)

	// SUT
	go cmdr.Order()

	sort.Stable(expMq)

	// Assertion
	i := 0
	for i < len(expMq) {
		temp := <-c
		assert.Equal(t, expMq[i].Pkt.Bytes(), temp)
		i++
	}
}
