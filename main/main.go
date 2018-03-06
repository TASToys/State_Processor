package main

import (
	"sync"
	"fmt"
	"math/rand"
	//"os"
	
	"State_processor/netcode"
	"strings"
)

const(
	differentInfoDelimiter string = "¡"
	nameTextDelimiter string = ":"
)


func main() {

	controller()
	//go processor("1.1.1.1:80")
	//for {}
	
}

func controller(){
	var l sync.Mutex
	var sum = 0
	var num = rand.Int()%rand.Int()%rand.Int()%100000
	var threads = 0

	fmt.Printf("mux:%d\nstart:%d", num,(num*num ))

	for i := 0; i < num; i++{
		l.Lock()
		threads = threads+1
		l.Unlock()
		go add(&l,&sum,num, &threads)

	}
	
	for threads > 0{
	}
	fmt.Printf("total:%d", sum)

}

func add(l *sync.Mutex, sum *int, num int, threads *int){
	var s = 0
	for i := 0; i < num; i++{
		s++
	}
	l.Lock()
		*sum = *sum + s

	*threads = *threads-1
	l.Unlock()

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


