package challenges

import "fmt"

type User struct {
	Name    string
	Email   string
	Friends []User
}

func (u *User) AddFriends(friend User) {
	u.Friends = append(u.Friends, friend)
}

func (u *User) RemoveFriends(friend User) {
	var index int
	for i := 0; i < len(u.Friends); i++ {
		if u.Friends[i].Name == friend.Name {
			index = i
		}
	}

	u.Friends = append(u.Friends[:index], u.Friends[index+1:]...)
}

func (u User) Display() {
	fmt.Println("Ismi: ", u.Name)
	fmt.Println("Emaili", u.Email)
	fmt.Println("Do'stlari: ")
	for _, v := range u.Friends {
		fmt.Println("\t", v.Name)
	}
}
