package test


import(
	"State_processor/netcode"
	"fmt"

	"strconv"
)

//NetCode tests netcode
func NetCode(in int){

	address, done1 := netcode.ArbitraryHost()
	fmt.Printf("inbound IP: %s \n", address)
	for i := 0; i < in; i++{
		go netcode.SendMessage("This is message " + strconv.Itoa(i), address)
		fmt.Printf("Received :%s\n", <-done1)
	}
	
}