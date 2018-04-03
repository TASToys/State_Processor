package test

//This file is the file used for testing and learning purposes. Eventually it will become testing only
//However, until such a time as I (deef0000dragon1) am comfortable with the language, I will continue
//to use it as a playground of sorts to test different aspects of golang and the state_processor

import (
	"State_processor/netcode"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/yuin/gopher-lua"

	"log"
)

//AllTests runs all tests in test.go
func AllTests() {
	fmt.Println("\n\n*****************STARTING TESTS*****************")
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
func NetCodeTest(in int) {

	address, done1 := netcode.ArbitraryHost()
	fmt.Printf("inbound IP: %s \n", address)
	for i := 0; i < in; i++ {
		go netcode.SendMessage("This is message "+strconv.Itoa(i), address)
		fmt.Printf("Received: %s\n", <-done1)
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

	//create Vm, and run the VM code.
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

//LuaLimitedEnvironmentPassedInStringTest is designed to test the limiting of the lua environment.
func LuaLimitedEnvironmentPassedInStringTest() {
	dat, err := ioutil.ReadFile("./test/luaCode.lua")
	check(err)
	//fmt.Print(string(dat))
	var runString = string(dat)
	TrueRunCode(runString)

}

/*tested but failing code. Kept for posterity shold I come back to examine how to get the out of lua sandboxing working.
func doScriptInSandbox(L *lua.LState, script string) error {
	io := L.GetGlobal("io").(*lua.LTable)
	orgopen := io.RawGetH(lua.LString("open"))
	defer io.RawSetH(lua.LString("open"), orgopen)
	sandBoxFunc := L.NewFunction(func(L *lua.LState) int {
		L.RaiseError("can not call in a sandbox environment.")
		return 0
	})
	io.RawSetH(lua.LString("open"), sandBoxFunc)
	err := L.DoString(script)
	return err
}

func runcode() {
	L := lua.NewState()

	fmt.Printf("L State:%v\n\n\n", ((L.Env).Len()))
	script := `
      local fp = assert(io.open("test.txt"))
      fp:close()
    `
	if err := doScriptInSandbox(L, script); err != nil {
		fmt.Println(err.Error())
	}
	if err := L.DoString(script); err != nil {
		fmt.Println(err.Error())
	}
}
*/

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
	} else {
		fmt.Print("Passed the test\n")
	}

}

//TrueRunCode is designed to take in code into inputCode and run it in the predetermined sandboxed format.
//inputCode string is the code to be run by the function.
func TrueRunCode(inputCode string) {
	start := `
	local env = {
		ipairs = ipairs,
		next = next,
		pairs = pairs,
		pcall = pcall,
		tonumber = tonumber,
		tostring = tostring,
		type = type,
		unpack = unpack,,
		print = print,
		string = { byte = string.byte, char = string.char, find = string.find, 
			format = string.format, gmatch = string.gmatch, gsub = string.gsub, 
			len = string.len, lower = string.lower, match = string.match, 
			rep = string.rep, reverse = string.reverse, sub = string.sub, 
			upper = string.upper },
		table = { insert = table.insert, maxn = table.maxn, remove = table.remove, 
			sort = table.sort },
		math = { abs = math.abs, acos = math.acos, asin = math.asin, 
			atan = math.atan, atan2 = math.atan2, ceil = math.ceil, cos = math.cos, 
			cosh = math.cosh, deg = math.deg, exp = math.exp, floor = math.floor, 
			fmod = math.fmod, frexp = math.frexp, huge = math.huge, 
			ldexp = math.ldexp, log = math.log, log10 = math.log10, max = math.max, 
			min = math.min, modf = math.modf, pi = math.pi, pow = math.pow, 
			rad = math.rad, random = math.random, sin = math.sin, sinh = math.sinh, 
			sqrt = math.sqrt, tan = math.tan, tanh = math.tanh },
		os = { clock = os.clock, difftime = os.difftime, time = os.time },
	}

	--may not add codoutines. 
	--coroutine = { create = coroutine.create, resume = coroutine.resume, 
	--	running = coroutine.running, status = coroutine.status, 
	--	wrap = coroutine.wrap },

	local function run(untrusted_code) --run function
  	if untrusted_code:byte(1) == 27 then return nil, "binary bytecode prohibited" end
  	local untrusted_function, message = loadstring(untrusted_code)
  	if not untrusted_function then return nil, message end
  	setfenv(untrusted_function, env) --set the environment
  	return pcall(untrusted_function) --run the code using the set environment
	end
	
	run [[`

	//It would be wise to confirm against http://lua-users.org/wiki/SandBoxes what to add and remove. 
	runString := start + inputCode + "]]"
	L := lua.NewState()
	if err := L.DoString(runString); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Print("Passed the test\n")
	}

}
