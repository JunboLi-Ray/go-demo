package service

import (
	"github.com/JunboLi-Ray/go-demo/app/store"
	"github.com/JunboLi-Ray/go-demo/app/object"
	"strconv"
)

func CloseDb() {
	store.CloseDb()
}

//获取所有普通用户
func AllNormalUser() (*[]object.User, error) {
	return store.GetUsersByType(object.NormalUser)
}

//添加用户
func AddNormalUser(name string) (*object.User, error) {
	id, err := store.AddUser(name, object.NormalUser)
	if err != nil {
		return nil, err
	}
	changeId, _ := strconv.ParseInt(id, 10, 64)
	return &object.User{
		Id:   changeId,
		Name: name,
		Type: object.NormalUser,
	}, nil
}

//获取某个用户所有关系
func UserAllLikeRela(userId int64) (*[]object.UserRela, error) {
	return store.GetUserAllRelasByType(userId, object.LikeRela)
}

//更新用户的某一条关系
func UpdateLikeRela(userId int64, otherUserId int64, state string) (*object.UserRela, error) {
	newState, err := store.UpdateOneRelaState(state, userId, otherUserId, object.LikeRela)
	if err != nil {
		return nil, err
	}
	return &object.UserRela{
		User_id: otherUserId,
		State:   newState,
		Type:    object.LikeRela,
	}, err
}
