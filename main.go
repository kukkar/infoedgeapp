package main

import (
	"fmt"
	"infoedge/journalapp/lib"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		// take commands from file
		fileName := argsWithoutProg[0]
		lib.ReadAndProcessFromFile(fileName)
	} else {
		fmt.Println("\n INTERACTIVE COMMANDS :- \n LOGIN <User> <password> \n SIGNUP <Id> <email> <password> \n LISTJOURNAL \n CREATEJOURNAL <message>")
		//We need to make it interactive session now
		lib.ReadAndProcessStdIn()
	}
}
