package user

import (
	"encoding/json"
	"fmt"
	"infoedge/journalapp/helpers"
	"infoedge/journalapp/journal"
)

type User struct {
	Id       string
	Password string
}

type UserSignUP struct {
	Name     string
	Password string
	Email    string
}

func Login(u User) (*string, error) {
	fileName := helpers.GetFileName(u.Id)
	fmt.Println(fileName)
	err := helpers.CheckFileExists(fileName)
	if err != nil {
		return nil, err
	}

	data, err := helpers.GetFileData(fileName)
	if err != nil {
		return nil, err
	}
	var fData journal.File
	err = json.Unmarshal(data, &fData)
	if err != nil {
		fmt.Println("error")
		return nil, err
	}

	if fData.Pass != u.Password {
		return nil, fmt.Errorf("Kindly enter valid password")
	}
	return &fileName, nil
}

func SignUp(su UserSignUP) error {

	fileName := helpers.GetFileName(su.Email)
	d := journal.File{
		Pass: su.Password,
		Data: nil,
	}

	err := helpers.CreateFile(fileName)
	if err != nil {
		return err
	}
	dBytes, err := json.Marshal(d)
	if err != nil {
		return err
	}
	err = helpers.Write2File(dBytes, fileName)
	if err != nil {
		return err
	}

	return nil
}

func RemoveUser(email string) error {

	fn := helpers.GetFileName(email)
	err := helpers.DeleteFile(fn)
	if err != nil {
		return err
	}
	return nil
}
