package action

import (
	"fmt"
	"github.com/JunboLi-Ray/go-demo/app/constant"
	"github.com/JunboLi-Ray/go-demo/app/object"
	"github.com/JunboLi-Ray/go-demo/app/service"
	"github.com/JunboLi-Ray/go-demo/app/util"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//获取所有用户
func AllNormalUserFunc(w http.ResponseWriter, r *http.Request) {
	data, err := service.AllNormalUser()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write(util.ObjectToJson(&object.Error{Err: err.Error()}))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(util.ObjectToJson(data))
}

//添加用户
func AddNormalUserFunc(w http.ResponseWriter, r *http.Request) {
	//参数校验
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(util.ObjectToJson(&object.Error{Err: err.Error()}))
		return
	}
	var reqParam object.User
	util.JsonToObject(body, &reqParam)
	if strings.TrimSpace(reqParam.Name) == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write(util.ObjectToJson(&object.Error{Err: constant.ParamError.Error()}))
		return
	}
	data, err := service.AddNormalUser(reqParam.Name)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write(util.ObjectToJson(&object.Error{Err: err.Error()}))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(util.ObjectToJson(data))
}

//获取某个用户所有关系
func UserAllLikeRelaFunc(w http.ResponseWriter, r *http.Request) {
	//校验参数
	userId, err := strconv.ParseInt(mux.Vars(r)["user_id"], 10, 64)
	if err != nil || userId < 1 {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write(util.ObjectToJson(&object.Error{Err: constant.ParamError.Error()}))
		return
	}
	data, err := service.UserAllLikeRela(userId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write(util.ObjectToJson(&object.Error{Err: err.Error()}))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(util.ObjectToJson(data))
}

//更新用户的某一条关系
func UpdateLikeRelaFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//校验参数
	if strings.TrimSpace(vars["user_id"]) == "" || strings.TrimSpace(vars["other_user_id"]) == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write(util.ObjectToJson(&object.Error{Err: constant.ParamError.Error()}))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write(util.ObjectToJson(&object.Error{Err: err.Error()}))
		return
	}
	var reqParam object.UserRela
	util.JsonToObject(body, &reqParam)
	if reqParam.State != string(object.Liked) && reqParam.State != string(object.DisLiked) {
		w.WriteHeader(http.StatusNotFound)
		w.Write(util.ObjectToJson(&object.Error{Err: constant.ParamError.Error()}))
		return
	}

	changeUserId, err := strconv.ParseInt(vars["user_id"], 10, 64)
	changeOtherUserId, err := strconv.ParseInt(vars["other_user_id"], 10, 64)
	if err != nil || changeUserId < 1 || changeOtherUserId < 1 || changeUserId == changeOtherUserId {
		w.WriteHeader(http.StatusNotFound)
		w.Write(util.ObjectToJson(&object.Error{Err: constant.ParamError.Error()}))
		return
	}
	data, err := service.UpdateLikeRela(changeUserId, changeOtherUserId, reqParam.State)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write(util.ObjectToJson(&object.Error{Err: err.Error()}))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(util.ObjectToJson(data))
}
