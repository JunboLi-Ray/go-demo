package store

import (
	"github.com/go-pg/pg"
	"github.com/JunboLi-Ray/go-demo/app/object"
	"github.com/JunboLi-Ray/go-demo/app/config"
)

var db = pg.Connect(&pg.Options{
	Addr:     config.SysConfs.SqlAddr,
	User:     config.SysConfs.SqlUser,
	Database: config.SysConfs.SqlDatabase,
})

func CloseDb() {
	db.Close()
}

//获取user表所有信息
func GetUsersByType(userType object.UserType) (*[]object.User, error) {
	var users []object.User
	err := db.Model(&users).Where(" type=? ", userType).Select()
	if len(users) == 0 {
		return nil, err
	}
	return &users, err
}

//添加用户
func AddUser(name string, userType object.UserType) (string, error) {
	var id string
	err := db.RunInTransaction(func(tx *pg.Tx) error {
		_, err := tx.QueryOne(pg.Scan(&id), `INSERT INTO users (name, type) VALUES (?, ?) returning id`, name, userType)
		return err
	})
	return id, err
}

//获取某个用户所有关系
func GetUserAllRelasByType(id int64, relaType object.UserRelaType) (*[]object.UserRela, error) {
	var userRelas []object.UserRela
	err := db.Model(&userRelas).Where(" id=? AND type=?", id, relaType).Select()
	if len(userRelas) == 0 {
		return nil, err
	}
	return &userRelas, err
}

func UpdateOneRelaState(state string, userId int64, otherUserId int64, relaType object.UserRelaType) (string, error) {
	err := db.RunInTransaction(func(tx *pg.Tx) error {
		var bToaState object.RelaState
		_, err := tx.QueryOne(pg.Scan(&bToaState), `SELECT state FROM user_relas WHERE id=? AND user_id=? AND type=?`, otherUserId, userId, relaType)
		if err != nil && err.Error() != "pg: no rows in result set" {
			return err
		}
		if bToaState == object.Liked {
			if state == string(object.Liked) {
				state = string(object.Matched)
				_, err = tx.Exec(`UPDATE user_relas SET state = ? WHERE id=? AND user_id=? AND type=?`, object.Matched, otherUserId, userId, relaType)
				if err != nil {
					return err
				}
			}
		} else if bToaState == object.Matched {
			if state == string(object.DisLiked) {
				_, err = tx.Exec(`UPDATE user_relas SET state = ? WHERE id=? AND user_id=? AND type=?`, object.Liked, otherUserId, userId, relaType)
				if err != nil {
					return err
				}
			} else {
				state = string(object.Matched)
			}
		}
		var aTobState string
		_, err = tx.QueryOne(pg.Scan(&aTobState), `SELECT state FROM user_relas WHERE id=? AND user_id=? AND type=?`, userId, otherUserId, relaType)
		if err != nil && err.Error() != "pg: no rows in result set" {
			return err
		}
		err = nil
		if aTobState == "" {
			_, err = tx.Exec(`INSERT INTO user_relas (id, user_id, state, type) VALUES (?, ?, ?, ?)`, userId, otherUserId, state, relaType)
		} else {
			if aTobState != state {
				_, err = tx.Exec(`UPDATE user_relas SET state = ? WHERE id=? AND user_id=? AND type=?`, state, userId, otherUserId, relaType)
			}
		}
		return err
	})
	return state, err
}
