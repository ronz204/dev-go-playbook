package demos

type User struct {
	Name     string
	Email    string
	Password string
}

// Receiver by value, as we are not modifying the User struct
func (u User) DisplayInfo() {
	println("User Name:", u.Name)
	println("User Email:", u.Email)
	println("User Password:", u.Password)
}

// Receiver by pointer, as we are modifying the User struct
func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

func StructuringDemo() {
	var user User = User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "secure123",
	}

	user.DisplayInfo()
	user.UpdateEmail("john.doe.updated@example.com")
	user.DisplayInfo()
}
