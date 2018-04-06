package main

import (
	"State_processor/LanguagesAndHelpers"
	"State_processor/netcode"
	//"State_processor/test"
	"fmt"
	"os"
	"strings"
	"sync"
)

//test
//another test

const (
	differentInfoDelimiter   string = "¡"
	nameTextDelimiter        string = "‽"
	controllerArgumentString        = "controller"
	processorArgumentString         = "processor"
)

//stores a list of all programs. Stored in the format TastoysUser-PluginId
var program map[string]string
var programMutex = &sync.Mutex{}
var language map[string]string
var languageMutex = &sync.Mutex{}

func main() {
	//test.NetCodeTest(100)
	//test.LuaGoTest()
	argParse()
}

type fidomsg struct {
	PluginID       int      `json:"plugin_id"`
	PluginType     int      `json:"plugin_type"`
	SourcePlatform string   `json:"source_platform"`
	SourceChannel  string   `json:"source_channel"`
	SourceUser     string   `json:"source_user"`
	TastoysUser    string   `json:"tastoys_user"`
	Timestamp      string   `json:"timestamp"`
	Data           struct { //I think that this needs to be a string
	} `json:"data"`
}

func argParse() {
	args := os.Args

	switch strings.ToLower(args[1]) {
	case controllerArgumentString:
		controller()
	case processorArgumentString:
		processor(args[2])
	}
}

func controller() {

}

func processor(controllerAddress string) {
	//create and send to the controller the receiver for the

	commandAddress, inputChannel := netcode.ArbitraryHost()                     //create listening channel at string commandAddress
	commandAddressString := "newProcessor" + nameTextDelimiter + commandAddress //create the command that tells the controller where to send to
	netcode.SendMessage(commandAddressString, controllerAddress)                //send to the controller at controller address the commandString
	var commandString string                                                    //create the string that the commands are going to go into.
	var commands []string                                                       //create a string array for the commands.
	for {
		commandString = <-inputChannel                                                    //read from the channel the next command string.
		commands = textSplitter(commandString, differentInfoDelimiter, nameTextDelimiter) //split the commands. normally only one comand

		for i := 0; i < len(commands); i += 2 {
			switch strings.ToLower(commands[i]) {
			case "newjob":
				newJob(commands[i+1]) //add the job code (second command)
			case "runjob":
				runJob(commands[i+1]) //run the code in the second argument
			case "removejob":
				removeJob(commands[i+1]) // remove from the list of active jobs the input job.
			case "listjobs":
				listJobs(commands[i+1]) //return to the input address (second arg) the list of all jobs in the processor
			case "ping":
				netcode.SendMessage("pong", commands[i+1]) //respond to a ping message with a pong.
			default:
			}
		}
	}
}

//splits the text provided according to a primary and secondairy delimiter.
func textSplitter(text, primaryDelim, secondairyDelim string) []string {

	pairs := strings.Split(text, primaryDelim)
	var split []string
	for _, v := range pairs {
		fmt.Printf("%s\n", v)
		split = append(split, strings.SplitN(v, secondairyDelim, 2)...)
	}
	return split
}

//adds a job string to the list of jobs.
func newJob(input string) {
	langFetch := "data:language" // get from the fido the language
	lang := LanguagesAndHelpers.GetJSONPiece(input, langFetch)
	codeFetch := "data:code"
	code := LanguagesAndHelpers.GetJSONPiece(input, codeFetch)
	id := LanguagesAndHelpers.GetJSONPiece(input, "TastoysUser")
	user := LanguagesAndHelpers.GetJSONPiece(input, "PluginID")
	codeID:=user + "-" + id
	languageMutex.Lock()
	language[codeID] = lang
	languageMutex.Unlock()
	programMutex.Lock()
	program[codeID] = code
	programMutex.Unlock()


	
}

//run a job based on the given info.
func runJob(input string) {
	id := LanguagesAndHelpers.GetJSONPiece(input, "TastoysUser")
	user := LanguagesAndHelpers.GetJSONPiece(input, "PluginID")
	codeID:=user + "-" + id
	switch language[codeID]{
	case "LUA" :
		LanguagesAndHelpers.RunLuaCode(program[codeID])
	}
}

func removeJob(input string) {

	id := LanguagesAndHelpers.GetJSONPiece(input, "TastoysUser")
	user := LanguagesAndHelpers.GetJSONPiece(input, "PluginID")
	codeID:=user + "-" + id

	languageMutex.Lock()
	delete(language, codeID)
	languageMutex.Unlock()
	programMutex.Lock()
	delete(program, codeID)
	programMutex.Unlock()

}

func listJobs(input string) {

	response := ""
	for id := range(program){
		response += ":"+id
	}

	addressFetch := "date:address"
	responseAddress := LanguagesAndHelpers.GetJSONPiece(input, addressFetch)
	netcode.SendMessage(response, responseAddress)
}
