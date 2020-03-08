package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	commandInit   = "init"
	commandStatus = "status"
	commandInput  = "input"
	commandLeave  = "leave"
	commandFind   = "find"
	commandExit   = "exit"
)

type identity struct {
	identityType   string
	identityNumber int
}

func main() {
	lockerNumber, err := initLocker()
	for err != nil {
		fmt.Println("failed to init: ", err, "Please try again")
		lockerNumber, err = initLocker()
	}
	fmt.Println("locker number is successfully set to ", lockerNumber)

	data := make(map[int]identity, lockerNumber)

	isExit := false
	for !isExit {
		isExit, err = mainFunction(&data, lockerNumber)
		if err != nil {
			fmt.Println("failed to execute command: ", err)
		}
	}

	fmt.Println("program closed by user request")
}

func initLocker() (int, error) {
	var command string
	_, err := fmt.Scan(&command)
	if err != nil {
		return 0, err
	}
	if strings.ToLower(command) != commandInit {
		return 0, errors.New("first command must be \"init\"")
	}

	var lockerNumber int
	_, err = fmt.Scanln(&lockerNumber)
	return lockerNumber, err
}

func mainFunction(data *map[int]identity, lockerNumber int) (bool, error) {
	var command string
	_, err := fmt.Scan(&command)
	if err != nil {
		return false, err
	}

	switch command {
	case strings.ToLower(commandExit):
		return true, nil
	case strings.ToLower(commandInput):
		return false, inputLocker(data, lockerNumber)
	case strings.ToLower(commandLeave):
		return false, leaveLocker(data, lockerNumber)
	case strings.ToLower(commandFind):
		return false, findLocker(data, lockerNumber)
	case strings.ToLower(commandStatus):
		return false, statusLocker(data, lockerNumber)
	default:
		return false, errors.New("invalid command")
	}
}

func inputLocker(data *map[int]identity, lockerNumber int) error {
	var identityType string
	var identityNumber int

	_, err := fmt.Scan(&identityType)
	if err != nil {
		return err
	}

	_, err = fmt.Scanln(&identityNumber)
	if err != nil {
		return err
	}

	for i := 0; i <= lockerNumber; i++ {
		_, exist := (*data)[i]
		if exist {
			if i < lockerNumber-1 {
				continue
			}
			return errors.New("locker already full")
		}

		(*data)[i] = identity{
			identityType:   identityType,
			identityNumber: identityNumber,
		}
		fmt.Println("locker input to number ", i+1)
		return nil
	}

	return nil
}

func leaveLocker(data *map[int]identity, lockerNumber int) error {
	var inputNumber int

	_, err := fmt.Scan(&inputNumber)
	if err != nil {
		return err
	}

	if inputNumber-1 == lockerNumber {
		return errors.New("input exceed locker number")
	}
	if inputNumber <= 0 {
		return errors.New("input must be positive")
	}

	delete(*data, inputNumber-1)
	fmt.Println("locker number ", inputNumber, " deleted")
	return nil
}

func findLocker(data *map[int]identity, lockerNumber int) error {
	var inputNumber int

	_, err := fmt.Scan(&inputNumber)
	if err != nil {
		return err
	}

	for i := 0; i < lockerNumber; i++ {
		d, exist := (*data)[i]
		if exist {
			if d.identityNumber == inputNumber {
				fmt.Println("identity found in locker number ", i+1)
				return nil
			}
		}
	}
	fmt.Println("locker not found")
	return nil

}

func statusLocker(data *map[int]identity, lockerNumber int) error {
	fmt.Printf("no\ttype\tidentity no\n")

	for i := 0; i < lockerNumber; i++ {
		d, exist := (*data)[i]
		if exist {
			fmt.Printf("%d\t%s\t%d\n", i+1, d.identityType, d.identityNumber)
		}
	}
	return nil
}
