package test


import(
	"State_processor/netcode"
	"fmt"
	"github.com/yuin/gopher-lua"
	"strconv"
	"io/ioutil"

)

//NetCodeTest tests the simple netcode implementation
func NetCodeTest(in int){

	address, done1 := netcode.ArbitraryHost()
	fmt.Printf("inbound IP: %s \n", address)
	for i := 0; i < in; i++{
		go netcode.SendMessage("This is message " + strconv.Itoa(i), address)
		fmt.Printf("Received :%s\n", <-done1)
	}
	
}

//LuaGoTest tests the lua go implementation
func LuaGoTest() {
	//fmt.Print("Hello World")

	fmt.Printf("", )
	dat, err := ioutil.ReadFile("/test/luaCode.lua")
	check(err)
	//fmt.Print(string(dat))
	var runString = string(dat)

	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(runString); err != nil {
		panic(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
