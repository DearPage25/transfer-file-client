package main

import (
	"log"
	"os"
)

func main(){
	cmd := os.Args[1];
	switch cmd {
		case "receive":
			receiveMode()
		case "send":
			sendMode()
		default:
			log.Fatal("Accions unknowns | Options: receive or send.")
		
	}
}

