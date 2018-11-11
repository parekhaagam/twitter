package web

type UsersRecord struct{
	users map[string]string
}

// key: email id , value: password
func initUsers() *UsersRecord{
		return &UsersRecord{
			users: map[string]string{},
		}
}

func (usersRecord *UsersRecord)userExist(userName string, password string)bool{

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

func (usersRecord *UsersRecord) insertUser(newUserName string, password string)bool{
	if ! usersRecord.userExist(newUserName, password) {
		usersRecord.users[newUserName] = password
		return true
	}else {
		return false
	}
}

func (usersRecord *UsersRecord) getUsers()[]string{

	keys := make([]string, 0, len(usersRecord.users))
	for k := range usersRecord.users {
		keys = append(keys, k)
	}

	return keys
}

