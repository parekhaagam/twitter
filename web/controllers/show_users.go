package controllers

type user struct {
	Name string
	IsFollowed bool
}


type UsersList struct {
	List []user
}

func Get_all_users() (ul UsersList){

	users := []user{
		{
			"manish.n",
			true,
		},
		{
			"ysd",
			false,
		},
		{
			"agamp",
			false,
		},
		{
			"srk",
			true,
		},
		{
			"dhoni007",
			false,
		},
	}

	return UsersList{users}
}
