package superhyper

import (
	"fmt"
	"testing"
)

var (
	RandProxies       = RandomChoice([]string{
		"donioleng877:donioleng877@139.162.109.153:3128",
		"donioleng877:donioleng877@172.105.228.157:3128",
		"donioleng877:donioleng877@139.162.89.197:3128",
		"donioleng877:donioleng877@139.162.79.115:3128",
		"donioleng877:donioleng877@172.104.80.121:3128",
		"donioleng877:donioleng877@139.162.123.80:3128",
		"donioleng877:donioleng877@172.105.200.233:3128",
		"donioleng877:donioleng877@139.162.82.12:3128",
		"donioleng877:donioleng877@172.104.76.227:3128",
		"donioleng877:donioleng877@139.162.97.160:3128",
		"donioleng877:donioleng877@139.162.97.23:3128",
		"donioleng877:donioleng877@139.162.99.125:3128",
		"donioleng877:donioleng877@139.162.99.143:3128",
		"donioleng877:donioleng877@172.105.229.236:3128",
	})
)

func RandomChoice(slice []string) string {
	rand.Seed(time.Now().UTC().UnixNano())
    return slice[rand.Intn(len(slice))]
}

func TestSocketIsntListeningAfterInterrupt(t *testing.T) {
	addr := fmt.Sprintf("%s", RandProxies)
	socket := CreateServerSocket(t, addr)
	socket.Listen()
	socket.Interrupt()
	newSocket := CreateServerSocket(t, addr)
	err := newSocket.Listen()
	defer newSocket.Interrupt()
	if err != nil {
		t.Fatalf("Failed to rebinds: %s", err)
	}
}

func TestSocketConcurrency(t *testing.T) {
	addr := fmt.Sprintf("%s", RandProxies)
	socket := CreateServerSocket(t, addr)
	go func() { socket.Listen() }()
	go func() { socket.Interrupt() }()
}

func CreateServerSocket(t *testing.T, addr string) *TServerSocket {
	socket, err := NewTServerSocket(addr)
	if err != nil {
		t.Fatalf("Failed to create server socket: %s", err)
	}
	return socket
}
