package main

import (

	"fmt"

	"os"
	
	"State_processor/netcode"
	"strings"
)

const(
	differentInfoDelimiter string = "¡"
	nameTextDelimiter string = ":"
	controllerArgumentString = "controller"
	processorArgumentString = "processor"
)


func main() {



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
	inputChannel := make(chan string, 1024) 
	commandAddress := netcode.ArbitraryHost(inputChannel)
	commandAddressString := "newProcessor" + nameTextDelimiter + commandAddress
	netcode.SendMessage(commandAddressString, controllerAddress)
	var commandString string
	var commands []string
	for{
		commandString = <- inputChannel
		commands = textSplitter(commandString, differentInfoDelimiter, nameTextDelimiter)

		switch strings.ToLower(commands[0]) {
		case "newjob"   :
		case "runjob"   :
		case "removejob":
		case "listjobs" :
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


