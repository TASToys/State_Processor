package netcode

import(
	"net"
	"log"
	"os"
	"fmt"
	"strings"
)

const(
	bufferSize int = 1024
)

//TestNetCode sends and prints 1000 messages to an arbitrarily created server



//SendMessage takes in a message and address string and sends that message to that address. 
//this code is designed to act as part of a simple 2 piece send and receive 
func SendMessage(message string, address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		//handle error
		panic("oops")
	}
	conn.Write([]byte(message))
	conn.Close()
}

//PollMessage takes in a message and address string and sends that message to that address. 
//this code is designed to act as part of a simple 2 piece send and receive 
func PollMessage(message string, address string) (response string) {
	conn, err := net.Dial("tcp", address)
	defer conn.Close()
	buf := make([]byte, bufferSize)
	if err != nil {
		// handle error
		panic("oops")
	}
	conn.Write([]byte(message))
	//fmt.Fprintf(conn, "stuff GET / HTTP/1.0\r\n\r\n")
	_, err = conn.Read(buf)
	return strings.Trim(string(buf), string(0))
}

//ArbitraryHost hosts an arbitrarily located tcp server. All data received is sent to the provided channel.
//The returned string is the string that the host is listening to in the form of xxx.xxx.xxx.xxx:port
func ArbitraryHost() (address string, output chan string){
	output = make(chan string, 1024) 
	l, err := net.Listen("tcp", getOutboundIP().String() +":0")

	if err != nil {
		fmt.Println("Error initialing listener:", err.Error())
		os.Exit(1)
	}
	address = l.Addr().String()
	go arbitraryHostListen(output, l)
	return address, output
}

func arbitraryHostListen(output chan string, l net.Listener ){
	defer l.Close()

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		
		if err != nil {
			fmt.Println("Error  d accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleInput(conn, output)
	}

}

func handleInput (conn net.Conn, output chan string) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, bufferSize)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	output <- strings.Trim(string(buf), string(0))
	
	// Send a response back to incoming connection.
	conn.Write([]byte("pong"))
	// Close the connection when you're done with it.
	conn.Close()
}

//getOutboundIP retrieves the IP of the current system as connected to the internet
//by pinging the 8.8.8.8:80 address and examining the connection created. 
//https://stackoverflow.com/a/37382208/3073088
func getOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}