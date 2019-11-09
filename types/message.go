package types

// Message stores msec and packet
type Message struct {
	Msec int64 // Millisecond
	Pkt  *Packet
}

// MessageQueue stores messages
type MessageQueue []*Message

func (mq MessageQueue) Len() int {
	return len(mq)
}

func (mq MessageQueue) Less(i, j int) bool {
	return mq[i].Msec < mq[j].Msec
}

func (mq MessageQueue) Swap(i, j int) {
	mq[i], mq[j] = mq[j], mq[i]
}
