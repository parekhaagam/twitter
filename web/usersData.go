package web

import "fmt"

type UsersRecord struct{
	users map[string]string
}

type User struct{
	Name string
	Isfollowed bool
}

type UserList struct{
	List []User
	nextPage bool
}

// key: email id , value: password
func InitUsers() *UsersRecord{
		return &UsersRecord{
			users: map[string]string{},
		}
}

func (usersRecord *UsersRecord)UserExist(userName string, password string)bool{

	pass, exists := usersRecord.users[userName]
	if exists {
		if pass != password{
			return false
		}
		return true
	}else {
		return false
	}

}

func (usersRecord *UsersRecord) InsertUser(newUserName string, password string)bool{
	if ! usersRecord.UserExist(newUserName, password) {
		usersRecord.users[newUserName] = password
		return true
	}else {
		return false
	}
}

var usersDataCache []User = nil

func (usersRecord *UsersRecord) GetUsers(pageNumber int, limit int)UserList{

	if (usersDataCache == nil ){
		usersDataCache = []User{}
		var tmp User
		for k := range usersRecord.users {
			tmp = User{
				Name:k,
				Isfollowed:false,
			}
			usersDataCache = append(usersDataCache, tmp)
		}
	}

	startUser := pageNumber * limit
	endUser := startUser + limit

	var nextPageStatus = true
	if (startUser < len(usersDataCache) && endUser > len(usersDataCache)) {
		fmt.Printf("Display for page number %d for users starting:%i\t till  ", pageNumber, startUser, endUser)
		endUser = len(usersDataCache)
		nextPageStatus = false
	}

	return UserList{
		List:usersDataCache[startUser:endUser],
		nextPage:nextPageStatus,
	}
}
