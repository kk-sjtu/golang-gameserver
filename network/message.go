package network

type Message struct {
	Id   uint64
	Data []byte // 可能需要序列化
}
