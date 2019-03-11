package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/MySQL"
	"io"
)

type Userinfo struct {
	Id       int64
	Username string `orm:"size(128)"`
	Age      int
}

func init() {
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/user?charset=utf8")
	orm.RegisterModel(new(Userinfo))

	orm.RunSyncdb("default",false,true)
}

func Insertuser(u *Userinfo) (id int64,err error){
	o :=orm.NewOrm()
	num,err:=o.Insert(u)
	return num,err
}
func Selectuser(u *Userinfo,index int ,setoff int ,conds map[string]string,append string) (data interface{},err error){
	//var Defaultlimit =10
	o :=orm.NewOrm()
	userdata :=make([]*Userinfo,0)
	qs :=o.QueryTable(u)
	if conds!=nil {

		for k, v := range conds {
			//fmt.Println(k, v)
			qs =qs.Filter(k, v)
		}
	}
	if setoff!=0{
		qs =qs.Limit(index,setoff)
	}
	var w io.Writer

	// 设置为你的 io.Writer

	orm.DebugLog = orm.NewLog(w)
	qs.All(&userdata)
	return userdata,err
}
func Countuser(u *Userinfo)(cnt int ,err error){
	o:=orm.NewOrm()
	num,err := o.QueryTable(u).Count() // SELECT COUNT(*) FROM USER
	return int(num) ,err
}
func Updae(u *Userinfo)(n int ,e error){
	o :=orm.NewOrm()
	uid :=o.Read(u)
	num,err :=o.Update(uid)
	if err!=nil{
		return
	}
	return int(num), nil
}
func Del(u *Userinfo)(cnt int ,err error){
	o := orm.NewOrm()
	var num int
	if num, err := o.Delete(u); err == nil {
		fmt.Println(num)
		return int(num),err
	}
	return num ,err
}
