package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Device struct {
	Uuid     string
	GroupId  int64
	NickName string
}

func Connect() (d []orm.Params, derr error) {
	dns, _ := getConfig(1)
	beego.Debug(dns)
	fmt.Printf("数据库is %s", dns)
	err := orm.RegisterDataBase("default", "mysql", dns)
	if err != nil {
		fmt.Println("\n========>  数据库连接failed  <==========")
	} else {
		fmt.Println("\n========>  数据库连接sucess  <==========")
		d, derr = GetHtransData()
		return d, derr
	}
	return nil, nil
}

func getConfig(flag int) (string, string) {
	var dns string
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	beego.Debug(db_host, db_name, db_pass, db_port, db_user)
	if flag == 1 {
		fmt.Println("========>     连接数据库     <==========")
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", db_user, db_pass, db_host, db_port, db_name)
	} else {
		fmt.Println("========>     创建数据库     <==========")
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", db_user, db_pass, db_host, db_port)
	}

	return dns, db_name
}

func GetHtransData() (device []orm.Params, err error) {
	o := orm.NewOrm()
	sql := "SELECT * FROM device"
	_, err = o.Raw(sql).Values(&device)
	return device, err
}

func ConnectHcluster(d []Device) (derr error) {
	dns, _ := getConfigHcluster(1)
	beego.Debug(dns)
	fmt.Printf("数据库is %s", dns)
	err := orm.RegisterDataBase("default", "mysql", dns)
	if err != nil {
		fmt.Println("\n========>  数据库连接failed  <==========")
	} else {
		fmt.Println("\n========>  数据库连接sucess  <==========")
		derr := SaveData(d)
		return derr
	}
	return nil
}

func getConfigHcluster(flag int) (string, string) {
	var dns string
	hcluster_host := beego.AppConfig.String("hcluster_host")
	hcluster_port := beego.AppConfig.String("hcluster_port")
	hcluster_user := beego.AppConfig.String("hcluster_user")
	hcluster_pass := beego.AppConfig.String("hcluster_pass")
	hcluster_name := beego.AppConfig.String("hcluster_name")
	beego.Debug(hcluster_host, hcluster_name, hcluster_pass, hcluster_port, hcluster_user)
	if flag == 1 {
		fmt.Println("========>     连接数据库     <==========")
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", hcluster_user, hcluster_pass, hcluster_host, hcluster_port, hcluster_name)
	} else {
		fmt.Println("========>     创建数据库     <==========")
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", hcluster_user, hcluster_pass, hcluster_host, hcluster_port)
	}

	return dns, hcluster_name
}

func SaveData(d []Device) (err error) {
	beego.Info(">>>>>>>>>>>>>>>>>>", d)
	o := orm.NewOrm()
	for i := 0; i < len(d); i++ {
		if d[i].Uuid == "" {
			break
		} else {
			sql := "UPDATE device SET group_id= " + fmt.Sprintf("%d", d[i].GroupId) + " , nick_name= \"" + d[i].NickName + "\" WHERE uuid=\"" + d[i].Uuid + "\";"
			beego.Info(">>>>>>>>>>>>>>>>", d[i].Uuid)
			beego.Info("+++++++++++++++", sql)
			o.Raw(sql).Exec()
			o.Commit()
		}
	}
	return nil
}
