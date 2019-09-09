package lib

import (
	"errors"
	"fmt"
	"infoedge/journalapp/journal"
	"infoedge/journalapp/user"
	"strings"
)

var sessionId *string
var jr *journal.Journal

//map of allowed commands along with the arguments to read
var allowedCommands = map[string]int{
	"login":         2,
	"signup":        3,
	"listjournal":   0,
	"createjournal": 99,
	"removeuser":    1,
}

var argumentsErrors = map[string]error{
	"login":         fmt.Errorf("login with user name and password"),
	"signup":        fmt.Errorf("atlease id password and email required"),
	"listjournal":   nil,
	"createjournal": fmt.Errorf("Make an entry with something please"),
	"loginfirst":    fmt.Errorf("Need to Login first"),
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

		if len(arguments) != numberOfArguments && numberOfArguments != 99 {
			fmt.Println(argumentsErrors[command].Error())
			return argumentsErrors[command]
		}

		w := &ErrWrapper{}


		// after validation of number of arguments per command, perform the necessary command
		switch command {
		case "login":
			w.do(func() error {
				var err error
				sessionId, err = user.Login(user.User{
					Id:       arguments[0],
					Password: arguments[1],
				})
				if err != nil {
					return err
				}
				fmt.Println("Login Successfully")
				return nil
			})
		case "signup":
			return w.do(func () error {
				err := user.SignUp(user.UserSignUP{
					Name:     arguments[0],
					Password: arguments[2],
					Email:    arguments[1],
				})
				if err != nil {
					return err
				}
				fmt.Println("Signup successfully")
				return nil
			})		
		case "listjournal":
			return w.do(func() error {
				if jr == nil {
					if sessionId != nil {
						jr, err = journal.GetInstace(*(sessionId))
					} else {
						return argumentsErrors["loginfirst"]
					}
				}
				enteries, err := jr.ListEntries()
				if err != nil {
					return err
				}
				for index, eachEntry := range enteries {
					fmt.Printf("%d %s %s \n", index, eachEntry.Time, eachEntry.ToRemeber)
				}
				return nil
			})
		case "createjournal":
			return w.do(func() error {
				if jr == nil {
					if sessionId != nil {
						jr, err = journal.GetInstace(*(sessionId))
					} else {
						return argumentsErrors["loginfirst"]
					}
				}
				var input string
				for _, eachData := range arguments {
					input = input + " " + eachData
				}
				err = jr.InputEntry(input)
				if err != nil {
					return err
				}
				return nil
			})

		case "removeuser":
			return w.do(func ()error{
				return user.RemoveUser(arguments[0])
			})
		}
	} else {
		err := errors.New(UNSUPPORTED_COMMAND)
		fmt.Println(err.Error())

		return err
	}
	return errors.New("Not Reachable Code")
}
