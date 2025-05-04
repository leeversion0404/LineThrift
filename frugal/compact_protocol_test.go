/*CODE ORIGINAL	BY : APACHE*/
/*REMAKE 				  BY : DONI HARIANTO*/
/*LICENSE				   BY : APACHE 	Ft	 PT.Amin-Bot-Super + LeeVersion*/


/*
DONI OLENG
*/

package superhyper

import (
	"bytes"
	"testing"
)

/*
DONI OLENG
*/

func TestReadWriteCompactProtocol(t *testing.T) {
	ReadWriteProtocolTest(t, NewTCompactProtocolFactory())
	go func(){
		transports1 := []TTransport{
			NewTMemoryBuffer(),
			NewStreamTransportRW(bytes.NewBuffer(make([]byte, 0, 131072))),
			NewTFramedTransport(NewTMemoryBuffer()),
		}
		transports2 := []TTransport{
			NewTMemoryBuffer(),
			NewStreamTransportRW(bytes.NewBuffer(make([]byte, 0, 131072))),
			NewTFramedTransport(NewTMemoryBuffer()),
		}
		transports3 := []TTransport{
			NewTMemoryBuffer(),
			NewStreamTransportRW(bytes.NewBuffer(make([]byte, 0, 131072))),
			NewTFramedTransport(NewTMemoryBuffer()),
		}
		transports4 := []TTransport{
			NewTMemoryBuffer(),
			NewStreamTransportRW(bytes.NewBuffer(make([]byte, 0, 131072))),
			NewTFramedTransport(NewTMemoryBuffer()),
		}
		transports5 := []TTransport{
			NewTMemoryBuffer(),
			NewStreamTransportRW(bytes.NewBuffer(make([]byte, 0, 131072))),
			NewTFramedTransport(NewTMemoryBuffer()),
		}
		transports6 := []TTransport{
			NewTMemoryBuffer(),
			NewStreamTransportRW(bytes.NewBuffer(make([]byte, 0, 131072))),
			NewTFramedTransport(NewTMemoryBuffer()),
		}
		transports7 := []TTransport{
			NewTMemoryBuffer(),
			NewStreamTransportRW(bytes.NewBuffer(make([]byte, 0, 131072))),
			NewTFramedTransport(NewTMemoryBuffer()),
		}
		transports8 := []TTransport{
			NewTMemoryBuffer(),
			NewStreamTransportRW(bytes.NewBuffer(make([]byte, 0, 131072))),
			NewTFramedTransport(NewTMemoryBuffer()),
		}
		zlib0, _ := NewTZlibTransport(NewTMemoryBuffer(), 0)
		zlib6, _ := NewTZlibTransport(NewTMemoryBuffer(), 6)
		zlib9, _ := NewTZlibTransport(NewTFramedTransport(NewTMemoryBuffer()), 9)
		transports1 = append(transports1, zlib0, zlib6, zlib9)
		transports2 = append(transports2, zlib0, zlib6, zlib9)
		transports3 = append(transports3, zlib0, zlib6, zlib9)
		transports4 = append(transports4, zlib0, zlib6, zlib9)
		transports5 = append(transports5, zlib0, zlib6, zlib9)
		transports6 = append(transports6, zlib0, zlib6, zlib9)
		transports7 = append(transports7, zlib0, zlib6, zlib9)
		transports8 = append(transports8, zlib0, zlib6, zlib9)
		for _, trans1 := range transports1 {
			p1 := NewTCompactProtocol(trans1)
			ReadWriteBool(t, p1, trans1)
			trans1.Close()
		}
		for _, trans2 := range transports2 {
			p2 := NewTCompactProtocol(trans2)
			ReadWriteByte(t, p2, trans2)
			trans2.Close()
		}
		for _, trans3 := range transports3 {
			p3 := NewTCompactProtocol(trans3)
			ReadWriteI16(t, p3, trans3)
			trans3.Close()
		}
		for _, trans4 := range transports4 {
			p4 := NewTCompactProtocol(trans4)
			ReadWriteI32(t, p4, trans4)
			trans4.Close()
		}
		for _, trans5 := range transports5 {
			p5 := NewTCompactProtocol(trans5)
			ReadWriteI64(t, p5, trans5)
			trans5.Close()
		}
		for _, trans6 := range transports6 {
			p6 := NewTCompactProtocol(trans6)
			ReadWriteDouble(t, p6, trans6)
			trans6.Close()
		}
		for _, trans7 := range transports7 {
			p7 := NewTCompactProtocol(trans7)
			ReadWriteString(t, p7, trans7)
			trans7.Close()
		}
		for _, trans8 := range transports8 {
			p8 := NewTCompactProtocol(trans8)
			ReadWriteBinary(t, p8, trans8)
			trans8.Close()
		}
	}()
}

/*
DONI OLENG
*/)
		}
	}()
}

/*
DONI OLENG SAYANG ZHALISA
*/