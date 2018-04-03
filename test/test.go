package test

//This file is the file used for testing and learning purposes. Eventually it will become testing only
//However, until such a time as I (deef0000dragon1) am comfortable with the language, I will continue
//to use it as a playground of sorts to test different aspects of golang and the state_processor

import(
	"State_processor/netcode"
	"fmt"
	"github.com/yuin/gopher-lua"
	"strconv"
	"io/ioutil"
//	"os"
//	"path/filepath"

	"log"
)

//AllTests runs all tests in test.go
func AllTests(){
	fmt.Println("\n\n*****************STARTING TESTS*****************")
	NetCodeTest(3)
	fmt.Println()
	LuaGoTest()
	fmt.Println()
	GlobalMapTest()
	fmt.Print("******************ENDING TESTS******************\n\n\n")
}

	var globalMap map[string]string

//GlobalMapTest tests the global map.
func GlobalMapTest(){
	globalMap = make(map[string]string)
	globalMap["test1"] = "test1Value"
	globalMap["test2"] = "test2Value"
	globalMap["test3"] = "test3Value"
	fmt.Printf("%v\n", globalMap)
	globalMap["test1"] = "altered"
	globalMap["test2"] = "altered"
	globalMap["test3"] = "altered"
	fmt.Printf("%v\n", globalMap)

}

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

	files, err := ioutil.ReadDir("./test")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
            fmt.Println(f.Name())
	}
	dat, err := ioutil.ReadFile("./test/luaCode.lua")
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


