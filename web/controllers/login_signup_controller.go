package controllers

import (
	"fmt"
	"../../globals"
)

func UserExist(userName string, password string)bool{

	pass, exists := globals.UsersRecord[userName]
	if exists {
		if pass != password{
			return false
		}
		return true
	}else {
		return false
	}

}

func InsertUser(newUserName string, password string)bool{
	if ! UserExist(newUserName, password) {
		globals.UsersRecord[newUserName] = password
		return true
	}else {
		return false
	}
}

var usersDataCache []globals.UserFollowed = nil

func GetUsers(pageNumber int, limit int) globals.UserList{

	if (usersDataCache == nil ){
		usersDataCache = []globals.UserFollowed{}
		var tmp globals.UserFollowed
		for k := range globals.UsersRecord {
			tmp = globals.UserFollowed{
				UserName:k,
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

	return globals.UserList{
		List:usersDataCache[startUser:endUser],
		NextPage:nextPageStatus,
	}
}
