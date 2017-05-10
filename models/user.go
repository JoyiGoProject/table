package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	ID   int64
	Uid  string
	Name string
	Age  int64
}

var userid string

//表名
func (r *User) TableName() string {
	return "user"
}

//模型注册
func init() {
	orm.RegisterModel(new(User))
}

//获取用户id
func GetUid(userid string) string {
	return userid
}

//通过传递过来的userid生成数据表
func SetTableByUid(userid string) string {
	o := orm.NewOrm()
	res, err := o.Raw(" CREATE TABLE IF NOT EXISTS `" + "user_" + userid + "` (" +
		"`i_d` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY," +
		"`uid` varchar(255) NOT NULL DEFAULT '' ," +
		"`name` varchar(255) NOT NULL DEFAULT '' ," +
		"`age` bigint NOT NULL DEFAULT 0" +
		") ENGINE=InnoDB;").Exec()
	//res, err := o.Raw("UPDATE user_test SET name = ?", userid).Exec()
	//beego.Info("执行数据打印：", res.LastInsertId)
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return userid
}

//添加用户
func AddUser(user *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	return id, err
}

//根据uid查询表是否存在
func CheckUserIsExist(uid string) bool {
	o := orm.NewOrm()
	b := o.QueryTable("user_test").Filter("uid", uid).Exist()
	return b
}
