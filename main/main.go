package main

import (
	"fmt"
	"os"
	"State_processor/test"
	"State_processor/netcode"
	"strings"
	"encoding/json"
)





const(
	differentInfoDelimiter string = "¡"
	nameTextDelimiter string = ":"
	controllerArgumentString = "controller"
	processorArgumentString = "processor"
)


func main() {
	test.NetCode(100)
}


type fidomsg struct {
    PluginID		int		`json:"plugin_id"`
    PluginType		int		`json:"plugin_type"`
    SourcePlatform	string	`json:"source_platform"`
	SourceChannel	string	`json:"source_channel"`
	SourceUser		string	`json:"source_user"`
	TastoysUser		string	`json:"tastoys_user"`
    Timestamp		string	`json:"timestamp"`
    Data           struct { //I think that this needs to be a string
    } `json:"data"`
}

func argParse(){
	args := os.Args

	switch strings.ToLower(args[1]) {
	case controllerArgumentString:
		controller()
	case processorArgumentString:
		processor(args[2])
	}
}



func controller(){


}



func processor(controllerAddress string){
	//create and send to the controller the receiver for the 
	commandAddress, inputChannel := netcode.ArbitraryHost()
	commandAddressString := "newProcessor" + nameTextDelimiter + commandAddress
	netcode.SendMessage(commandAddressString, controllerAddress)
	var commandString string
	var commands []string
	for{
		commandString = <- inputChannel
		commands = textSplitter(commandString, differentInfoDelimiter, nameTextDelimiter)

		switch strings.ToLower(commands[0]) {
		case "newjob"   : newJob(commands[1:])
		case "runjob"   : runJob(commands[1:])
		case "removejob": removeJob(commands[1:])
		case "listjobs" : listJobs(commands[1:])
		case "ping"     :
			netcode.SendMessage("pong", commands[1]) 
		default:
		
		
		}
	}
}

func textSplitter(text, primaryDelim, secondairyDelim string) []string{

	pairs := strings.Split(text, primaryDelim)
	var split []string
	for _, v := range pairs {
		fmt.Printf("%s\n", v)
		split = append(split, strings.SplitN(v, secondairyDelim , 2)...)
	}
	return split
}


func newJob(input []string){

}

func runJob(input []string){
	var fido fidomsg;
	json.Unmarshal([]byte(input[1]), fido);
	


}

func removeJob(input []string){

}

func listJobs(input []string){

}