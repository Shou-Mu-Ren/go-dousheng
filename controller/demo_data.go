package controller

var DemoVideos = []Video{
	{
		Id:     1,
		Author: DemoUser,
		// PlayUrl: "https://www.w3schools.com/html/movie.mp4",
		PlayUrl: "http://192.168.1.103:8080/douyin/file/download?filename=bear.mp4",
		// CoverUrl: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		CoverUrl:      "http://192.168.1.103:8080/douyin/file/download?filename=666.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
