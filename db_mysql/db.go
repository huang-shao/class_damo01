package db_mysql

import (
	"beego_damo02/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbuser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	//dbPort:=config.String("db_port")
	dbName := config.String("db_name")
	connUrl := dbuser + ":" + dbPassword + "@tcp(" + dbIp +")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据库理解错误，请检查配置")
	}
	Db = db

	//fmt.Println(port)
}
func InserUser(user models.User)(int64,error){
	//将用户密码进行hash脱敏，使用md5计算密码hash值，并存储hash值
	hashMd5:=md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes:=hashMd5.Sum(nil)
	user.Password=hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名：")
	result,err:=Db.Exec("insert into user (nick,password,birthday,hobby) value (?,?,?,?)",user.Nick,user.Password,user.Birthday,user.Hobby)
	if err!=nil{
		return -1,err
	}
	id,err:=result.RowsAffected()
	if err!=nil {
		return -1,err
	}
	return id,nil
}

//查询用户
func QueryUser(){
	Db.QueryRow("select *from user ")
}