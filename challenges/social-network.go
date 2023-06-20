package challenges

// Social Network:
// Create a User struct with fields Name, Email, and Friends. Implement methods to add friends, remove friends, and display the user's friend list.

import "fmt"

type User struct {
	Name    string
	Email   string
	Friends []User
}

func (u *User) AddFriend(friend User) {
	u.Friends = append(u.Friends, friend)
	fmt.Println("Friend added.")
}

func (u *User) RemoveFriend(friendName string) {
	var updatedFriends []User
	for _, friend := range u.Friends {
		if friend.Name != friendName {
			updatedFriends = append(updatedFriends, friend)
		}
	}
	u.Friends = updatedFriends
	fmt.Println("Friend removed.")
}

func (u User) DisplayFriendList() {
	fmt.Printf("Friend list for %s:\n", u.Name)
	for _, friend := range u.Friends {
		fmt.Println(friend.Name)
	}
}

// func main() {
// 	user1 := User{Name: "John", Email: "john@example.com"}
// 	user2 := User{Name: "Jane", Email: "jane@example.com"}
// 	user3 := User{Name: "Alice", Email: "alice@example.com"}

// 	user1.AddFriend(user2)
// 	user1.AddFriend(user3)

// 	user1.DisplayFriendList()

// 	user1.RemoveFriend("Jane")

// 	user1.DisplayFriendList()
// }
// Output:

// vbnet
// Copy code
// Friend added.
// Friend added.
// Friend list for John:
// Jane
// Alice
// Friend removed.
// Friend list for John:
// Alice
