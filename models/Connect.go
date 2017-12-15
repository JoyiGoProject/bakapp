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

type Group struct {
	Id   int64  //主键
	Name string //组别名称
}

func Connect() (d []orm.Params, g []orm.Params, derr error) {
	dns, _ := getConfig(1)
	beego.Debug(dns)
	fmt.Printf("数据库is %s", dns)
	err := orm.RegisterDataBase("default", "mysql", dns)
	if err != nil {
		fmt.Println("\n========>  数据库连接failed  <==========")
	} else {
		fmt.Println("\n========>  数据库连接sucess  <==========")
		d, g, derr = GetHtransData()
		return d, g, derr
	}
	return d, g, derr
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

func GetHtransData() (device []orm.Params, group []orm.Params, err error) {
	o := orm.NewOrm()
	sql := "SELECT * FROM device"
	_, err = o.Raw(sql).Values(&device)
	sql2 := "SELECT * FROM `GROUP`"
	o.Raw(sql2).Values(&group)
	return device, group, err
}

func ConnectHcluster(d []Device, g []Group, num int) (derr error) {
	dns, _ := getConfigHcluster(1)
	beego.Debug(dns)
	fmt.Printf("数据库is %s", dns)
	err := orm.RegisterDataBase("default", "mysql", dns)
	if err != nil {
		fmt.Println("\n========>  数据库连接failed  <==========", err)
	} else {
		fmt.Println("\n========>  数据库连接sucess  <==========")
		derr := SaveData(d, g, num)
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

func SaveData(d []Device, g []Group, num int) (err error) {
	o := orm.NewOrm()
	if num == 1 {
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
	} else {
		for i := 0; i < len(g); i++ {
			if g[i].Id == 0 {
				break
			} else {
				sql := "UPDATE `group` SET name= \"" + g[i].Name + "\" WHERE id= \"" + fmt.Sprintf("%d", g[i].Id) + "\";"
				beego.Info(">>>>>>>>>>>>>>>>", g[i].Id)
				beego.Info("+++++++++++++++", sql)
				o.Raw(sql).Exec()
				o.Commit()
			}
		}
	}
	return nil
}
