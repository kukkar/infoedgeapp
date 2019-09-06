package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const PATH = "/tmp/"

func GetFileName(st string) string {

	s := strings.Replace(st, "@", "_", -1)
	s = strings.Replace(s, ".", "_", -1)
	return s
}

func Write2File(data []byte, fileName string) error {
	// open file using READ & WRITE permission

	var file, err = os.OpenFile(PATH+fileName, os.O_RDWR, 0644)
	if isError(err) {
		return err
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.Write(data)
	if isError(err) {
		return err
	}

	// save changes
	err = file.Sync()
	if isError(err) {
		return err
	}
	return nil
}

func GetFileData(fileName string) ([]byte, error) {

	file, err := ioutil.ReadFile(PATH + fileName)
	if err != nil {
		return nil, err
	}
	return file, nil

}

func CreateFile(fileName string) error {
	// detect if file exists
	var _, err = os.Stat(PATH + fileName)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(PATH + fileName)
		if isError(err) {
			return err
		}
		defer file.Close()
	} else {
		return fmt.Errorf("User Already Exists with this email")
	}
	return nil
}

func CheckFileExists(fileName string) error {
	var _, err = os.Stat(PATH + fileName)
	if os.IsNotExist(err) {
		return fmt.Errorf("SignUp First No User exists")
	}
	return nil
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func DeleteFile(fileName string)error {
	// delete file
	var err = os.Remove(PATH + fileName)
	if isError(err) { 
		return err 
	}
	return nil
}
