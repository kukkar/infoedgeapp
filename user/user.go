package user

type User struct {
	Id       string
	Password string
}

type UserSignUP struct {
	Name     string
	Password string
	Email    string
}

func Login(u User) (string,error) {
	return "sahil",nil
}

func SignUp(su UserSignUP) error {

	return nil
}
