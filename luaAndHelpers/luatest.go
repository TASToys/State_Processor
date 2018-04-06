package luaAndHelpers

import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

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

//RunLuaCode is designed to take in lua code into inputCode and run it in the predetermined sandboxed format.
//inputCode string is the code to be run by the function.
func RunLuaCode(inputCode string) {
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
	}

}
