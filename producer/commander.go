package producer

import (
	"soupbintcp/types"

	"sort"

	"time"
)

// NewCommander creates commander
func NewCommander(c chan []byte) *Commander {
	return &Commander{tube: c}
}

// Commander is used to order with messages
type Commander struct {
	msgQ types.MessageQueue
	tube chan []byte
}

// PushMsg pushs the msg to messages slice
func (cmdr *Commander) PushMsg(msg *types.Message) {
	cmdr.msgQ = append(cmdr.msgQ, msg)
}

// Order orders the msg at right time
func (cmdr *Commander) Order() {
	sort.Stable(cmdr.msgQ)
	i := 0
	baseT := time.Now().UnixNano() / int64(time.Millisecond)
	for i < len(cmdr.msgQ) {
		periodT := (time.Now().UnixNano() / int64(time.Millisecond)) - baseT
		if cmdr.msgQ[i].Msec <= periodT {
			cmdr.tube <- cmdr.msgQ[i].Pkt.Bytes()
			i++
		}
	}
}
