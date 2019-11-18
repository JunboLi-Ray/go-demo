package object

/*
   存储及返回结构
*/
type User struct {
	Id   int64    `json:"id"`
	Name string   `json:"name"`
	Type UserType `json:"type"`
}

type UserRela struct {
	Id      int64        `json:"-"`
	User_id int64        `json:"user_id"`
	State   string       `json:"state"`
	Type    UserRelaType `json:"type"`
}

//错误返回结果
type Error struct {
	Err string `json:"error"`
}

/*
	用户表，type字段
*/
type UserType string

var NormalUser UserType = "user"

/*
	关系表，type字段
*/
type UserRelaType string

var LikeRela UserRelaType = "relationship"

/*
	关系表喜欢关系
*/
type RelaState string

var (
	Liked RelaState = "liked"

	DisLiked RelaState = "disliked"

	Matched RelaState = "matched"
)
