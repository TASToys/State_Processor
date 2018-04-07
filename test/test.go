package test

//This file is the file used for testing and learning purposes. Eventually it will become testing only
//However, until such a time as I (deef0000dragon1) am comfortable with the language, I will continue
//to use it as a playground of sorts to test different aspects of golang and the state_processor

import (
	"State_processor/LanguagesAndHelpers"
	"State_processor/netcode"
	"fmt"
	"github.com/yuin/gopher-lua"
	"io/ioutil"
	"strconv"
	"sync"
//	"sync/atomic"
)

//AllTests runs all tests in test.go
func AllTests() {
	fmt.Println("\n\n*****************STARTING SIMPLE TESTS*****************")
	fmt.Println("\n***NETCODE TESTS***")
	NetCodeTest(3)
	fmt.Println("\n***LUA GO TESTS***")
	LuaGoTest()
	fmt.Println("\n***LUA LIMITED ENVIRONMENT TESTS***")
	LuaLimitedEnvironmentTest()
	fmt.Println("\n***PASSED IN STRING TESTS***")
	LuaLimitedEnvironmentPassedInStringTest()
	fmt.Println("\n***GLOBAL MAP TESTS***")
	GlobalMapTest()
	fmt.Println("\n***JSON PARSE TESTS***")
	JSONParseTest()
	fmt.Print("******************ENDING SIMPLE TESTS******************\n\n\n")

}

//CoreTests runs all core tests as necessary for core State_Processor functionality
func CoreTests(){
	fmt.Println("\n\n******************STARTING CORE TESTS******************")
	fmt.Println("The following tests and actions are the initialization") 
	fmt.Println("and running of a tastoys State_processor component.")
	fmt.Printf("Included in the test are the following\n\n")
	fmt.Println("Starting of State_processor controller")
	fmt.Println("Starting of State_processor processor")
	fmt.Println("Loading a LUA Job by controller")
	fmt.Println("Routing a LUA Job from controller to processor")
	fmt.Println("adding a LUA Job by Processor")
	fmt.Println("Running a LUA Job by Processor")
	fmt.Print("******************ENDING CORE TESTS******************\n\n\n")
}

var globalMap map[string]string

//GlobalMapTest tests the global map.
func GlobalMapTest() {
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



var sum = 0
var done = 0

//MutexTest tests mutexes
func mutexTest() {
	var mutex = &sync.Mutex{}
	num := 10000
	for i := 0; i < num; i++ {

		go func() {
			total := 0
			for i := 0; i < num; i++ {
				total++
			}
			mutex.Lock()
				sum+=total
				done++
			mutex.Unlock()
			//fmt.Print("done")
		}()

	}
	
	for done < num{

	}
	fmt.Printf("Total %d\n", sum)

}

//NetCodeTest tests the simple netcode implementation
func NetCodeTest(in int) {

	address, done1 := netcode.ArbitraryHost()
	fmt.Printf("inbound IP: %s \n", address)
	for i := 0; i < in; i++ {
		go netcode.SendMessage("This is message "+strconv.Itoa(i), address)
		fmt.Printf("Received: %s\n", <-done1)
	}

}

//error checker that pacnis on an error
func check(err error) {
	if err != nil {
		panic(err)
	}
}

//goal: write json parsing code that can pull out specific data from arbitrary json.

//GetJSONPieceTest tests getting a part of arbitrary json
func GetJSONPieceTest() {
	data, err := ioutil.ReadFile("./test/example.json")
	check(err)
	loc := "1:owner:login"
	LanguagesAndHelpers.GetJSONPiece(string(data), loc) //should return "deef0000dragon1"
}

//LuaGoTest tests the lua go implementation
func LuaGoTest() {
	//fmt.Print("Hello World")

	files, err := ioutil.ReadDir("./test")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	dat, err := ioutil.ReadFile("./test/luaCode.lua")
	if err != nil {
		panic(err)
	}
	//fmt.Print(string(dat))
	var runString = string(dat)

	//create Vm, and run the VM code.
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(runString); err != nil {
		panic(err)
	}
}

//LuaLimitedEnvironmentPassedInStringTest is designed to test the limiting of the lua environment.
func LuaLimitedEnvironmentPassedInStringTest() {
	dat, err := ioutil.ReadFile("./test/luaCode.lua")
	if err != nil {
		panic(err)
	} //fmt.Print(string(dat))
	var runString = string(dat)
	LanguagesAndHelpers.RunLuaCode(runString)

}

//LuaLimitedEnvironmentTest tests basic lua limited environment code using print etc.
func LuaLimitedEnvironmentTest() {
	L := lua.NewState()
	script := `print("Hello World!") --test working print
	
	local env = {print=print, math={cos=math.cos}} --set environmental variables
	--setting math=math does give access to all math functions
	--however, using the above syntax, specific functions can be selected. 


	local function run(untrusted_code) --run function
  	if untrusted_code:byte(1) == 27 then return nil, "binary bytecode prohibited" end
  	local untrusted_function, message = loadstring(untrusted_code)
  	if not untrusted_function then return nil, message end
  	setfenv(untrusted_function, env) --set the environment
  	return pcall(untrusted_function) --run the code using the set environment
	end
	
	print("Should Print Test") 
	run [[print("test")]]
	print("Should Not print the sin of 5")
	run [[print(math.sin(5))]]
	print("should print the cos of 5")
	run [[print(math.cos(5))]]
	print("done")

	`

	if err := L.DoString(script); err != nil {
		fmt.Println(err.Error())
	}
}

//JSONParseTest tests to make sure that the JSON parser is parsing the input JSON correctally.
func JSONParseTest() {

	data, err := ioutil.ReadFile("./test/example.json")
	if err != nil {
		panic(err)
	} //fmt.Print(string(dat))
	var runString = string(data)

	searchString1 := "stuff:data:1:owner:id"
	searchString2 := "stuff:data:0:name"

	out1 := LanguagesAndHelpers.GetJSONPiece(runString, searchString1)
	out2 := LanguagesAndHelpers.GetJSONPiece(runString, searchString2)

	fmt.Printf("Should Be \n3053654 CE2801-Embedded-Systems-1\nIS\n%s %s\n", out1, out2)
}
