/*

这个文件定义了一个 NormalPacker 结构体，
用于处理消息的打包和解包。具体来说，
它将消息转换为字节流进行传输，并从字节流中解析出消息


*/

package network

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

type NormalPacker struct {
	// todo 处理大端小端。数据如果从高位存放，对应的是大端。如果从低位存放，对应的是小端
	Order binary.ByteOrder
}

/* Pack 方法用于将 Message 结构体打包成字节流。具体步骤如下：
1. 创建一个字节切片 buffer，长度为消息长度（8字节）+ ID（8字节）+ 数据长度。
2. 将消息总长度写入 buffer 的前 8 个字节。
3. 将消息 ID 写入 buffer 的接下来的 8 个字节。
4. 将消息数据复制到 buffer 的剩余部分。
5. 返回打包后的字节流。
*/

// Pack | data length | id | data |
func (p *NormalPacker) Pack(message *Message) ([]byte, error) {
	buffer := make([]byte, 8+8+len(message.Data))
	p.Order.PutUint64(buffer[:8], uint64(len(buffer)))
	p.Order.PutUint64(buffer[8:16], message.Id)
	copy(buffer[16:], message.Data)
	return buffer, nil
}

/* Unpack 方法用于从字节流中解析出 Message 结构体。具体步骤如下：
1. 设置读取超时时间为 1 秒。
2. 创建一个字节切片 buffer，长度为 16 字节（消息长度 + ID）。
3. 从 reader 中读取 16 字节数据到 buffer。
4. 解析出消息总长度和消息 ID。
5. 计算数据部分的长度，并创建相应长度的字节切片 dataBuffer。
6. 从 reader 中读取数据部分到 dataBuffer。
7. 创建并返回解析出的 Message 结构体。
*/

func (p *NormalPacker) Unpack(reader io.Reader) (*Message, error) {
	err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, 8+8)
	_, err = io.ReadFull(reader, buffer)
	if err != nil {
		return nil, err
	}
	totalLen := p.Order.Uint64(buffer[:8])
	id := p.Order.Uint64(buffer[8:])
	dataLen := totalLen - 16
	dataBuffer := make([]byte, dataLen)
	_, err = io.ReadFull(reader, dataBuffer)
	if err != nil {
		return nil, err
	}
	msg := &Message{
		Id:   id,
		Data: dataBuffer,
	}
	return msg, nil
}

/*
字节流是一种数据传输格式，其中数据以连续的字节序列形式进行传输或存储。
字节流可以用于在计算机系统之间传输数据，或者在内存中存储数据。
字节流的主要特点是数据以字节为单位进行处理，
这使得它非常适合用于网络通信和文件读写等场景
*/
