package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/as/kafka"
)

func s(str string) kafka.Str {
	return kafka.Str{
		N:    uint16(len(str)),
		Data: []byte(str),
	}
}

func main() {
	txbuf := new(bytes.Buffer)
	buf := new(bytes.Buffer)
	txhdr := kafka.TX{
		ApiKey: uint16(kafka.TxProd),
		ApiVer: 1,
	}
	txhdr.WriteBinary(buf)
	fmt.Printf("%# x\n", buf)
	hr := &kafka.ProdTX{
		NeedAcks: 0,
		Timeout:  18348483,
		N:        1,
		Topics: []kafka.ProdTXTopic{{
			Name: s("t1"),
			Parts: []kafka.ProdTXPart{{
				ID:         1,
				N:          5,
				MessageSet: []byte("hello"), // this should be a record with the right headers but it breaks our other app
			}},
		}},
	}
	fmt.Printf("%#v\n", hr)
	hr.WriteBinary(txbuf)
	fmt.Printf("%#v\n", hr)
	binary.Write(buf, binary.BigEndian, txbuf.Len())
	io.Copy(buf, txbuf)

	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#x\n", buf.Bytes())
	io.Copy(conn, buf)
	io.Copy(os.Stdin, conn)
}

func heartbeat() {
	txbuf := new(bytes.Buffer)
	buf := new(bytes.Buffer)
	txhdr := kafka.TX{
		ApiKey: uint16(kafka.TxHeart),
		ApiVer: 1,
	}
	txhdr.WriteBinary(buf)
	fmt.Printf("%# x\n", buf)
	hr := &kafka.HeartTX{Group: s("alpha"), Gen: 1, Member: s("go")}
	fmt.Printf("%#v\n", hr)
	hr.WriteBinary(txbuf)
	fmt.Printf("%#v\n", hr)
	binary.Write(buf, binary.BigEndian, txbuf.Len())
	io.Copy(buf, txbuf)

	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#x\n", buf.Bytes())
	io.Copy(conn, buf)
	io.Copy(os.Stdin, conn)
}
