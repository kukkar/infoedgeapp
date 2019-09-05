package lib

import (
	"errors"
	"fmt"
	"infoedge/journalapp/journal"
	"infoedge/journalapp/user"
	"strings"
)

var sessionId string
var jr *journal.Journal

//map of allowed commands along with the arguments to read
var allowedCommands = map[string]int{
	"login":         2,
	"signup":        3,
	"listjournal":   0,
	"createjournal": 1,
}

var argumentsErrors = map[string]error{
	"login":         fmt.Errorf("login with user name and password"),
	"signup":        fmt.Errorf("atlease id password and email required"),
	"listjournal":   nil,
	"createjournal": fmt.Errorf("Make an entry with something please"),
}

const (
	UNSUPPORTED_COMMAND           = "Unsupported Command"
	UNSUPPORTED_COMMAND_ARGUMENTS = "Unsupported Command Arguments"
)

// Process the command taken in from file/stdin
// Separate the command and arguments for command
// Validate the command and then do the necessary action
func processCommand(command string) error {
	commandDelimited := strings.Split(command, " ")
	lengthOfCommand := len(commandDelimited)
	
	var err error
	arguments := []string{}
	if lengthOfCommand < 1 {
		err := errors.New(UNSUPPORTED_COMMAND)
		fmt.Println(err.Error())
		return err
	} else if lengthOfCommand == 1 {
		command = commandDelimited[0]
	} else {
		command = commandDelimited[0]
		arguments = commandDelimited[1:]
	}

	// check if command is one of the allowed commands
	if numberOfArguments, exists := allowedCommands[command]; exists {

		if len(arguments) != numberOfArguments {
			fmt.Println(argumentsErrors[command].Error())
			return argumentsErrors[command]
		}

		// after validation of number of arguments per command, perform the necessary command
		switch command {
		case "login":
			var err error
			sessionId, err = user.Login(user.User{})
			if err != nil {
				return err
			}
			fmt.Println("Login Successfully")
			return nil
		case "signup":
			user.SignUp(user.UserSignUP{})
			return nil

		case "listjournal":
			if jr == nil {
				if sessionId != "" {
					jr, err = journal.GetInstace(sessionId)
				} else {
					return fmt.Errorf("Login First")
				}
			}
			err = jr.ListEntries()
			if err != nil {
				return err
			}
			return nil
		case "createjournal":
			if jr == nil {
				fmt.Println(sessionId)
				if sessionId != "" {
					jr, err = journal.GetInstace(sessionId)
				} else {
					return fmt.Errorf("Login First")
				}
			}
			err = jr.InputEntry(string(arguments[0]))
			if err != nil {
				return err
			}
			return nil
		}
	} else {
		err := errors.New(UNSUPPORTED_COMMAND)
		fmt.Println(err.Error())

		return err
	}
	return errors.New("Not Reachable Code")
}
