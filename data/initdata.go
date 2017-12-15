package data

import (
	"bakapp/models"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/astaxie/beego"
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

func GetData() {
	device, _, err := models.Connect()
	if err != nil {
		beego.Error("get the data is error", err)
		return
	}
	var d [1000]Device
	// var g [1000]Group
	for j := 0; j < len(device); j++ {
		for i := 0; i < len(device); i++ {
			if i == j {
				d[i].Uuid = device[j]["uuid"].(string)
				gid := device[j]["group_id"].(string)
				id, err := strconv.ParseInt(gid, 10, 64)
				if err != nil {
					beego.Error("error", err)
					return
				}
				d[i].GroupId = id
				d[i].NickName = device[j]["nick_name"].(string)
			}
		}
	}
	beego.Info(d)
	b, err := json.Marshal(d)
	beego.Info("----", string(b))
	SaveFile(b, 1)

	/* 	for j := 0; j < len(group); j++ {
	   		for i := 0; i < len(group); i++ {
	   			if i == j {
	   				id := group[j]["id"].(string)
	   				gid, err := strconv.ParseInt(id, 10, 64)
	   				if err != nil {
	   					beego.Error("error", err)
	   					return
	   				}
	   				g[i].Id = gid
	   				g[i].Name = group[j]["name"].(string)
	   			}
	   		}
	   	}
	   	beego.Info(g)
	   	groupbyte, err := json.Marshal(g)
	   	beego.Info("序列化后的组数据：", string(groupbyte))
	   	SaveFile(groupbyte, 2) */
}

func SaveFile(b []byte, num int) {
	var filename string
	if num == 1 {
		filename = "./devicebak.json"
	} else {
		filename = "./groupbak.json"
	}
	var f *os.File
	var err error
	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666) //打开文件
		os.Remove(filename)
		f, err = os.Create(filename) //创建文件
		fmt.Println("文件已创建")
	} else {
		f, err = os.Create(filename) //创建文件
		fmt.Println("文件已创建")
	}
	n, err := io.WriteString(f, string(b)) //写入文件(字符串)
	if err != nil {
		beego.Error("write the data error ", err)
		return
	}
	fmt.Printf("备份完毕，本次共写入 %d 个字节\n", n)
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SaveData() {
	filename := "./devicebak.json"
	dat, err := ioutil.ReadFile(filename)
	var msgs []models.Device
	err = json.Unmarshal([]byte(dat), &msgs)
	if err != nil {
		fmt.Println("Can't decode json message", err)
	} else {
		fmt.Println("====================", msgs)
	}
	check(err)
	fmt.Print("解析之后的数据：", msgs)
	err = models.ConnectHcluster(msgs, nil, 1)
	if err != nil {
		beego.Error("update the device data error", err)
		return
	}
	/* 	filename = "./groupbak.json"
	   	gbyte, err := ioutil.ReadFile(filename)
	   	var groups []models.Group
	   	err = json.Unmarshal([]byte(gbyte), &groups)
	   	if err != nil {
	   		fmt.Println("Can't decode json message", err)
	   	} else {
	   		fmt.Println("====================", groups)
	   	}
	   	check(err)
	   	fmt.Print("解析之后的数据：", groups)
	   	err = models.ConnectHcluster(nil, groups, 2)
	   	derr := models.SaveData(nil, groups, 2)
	   	if derr != nil {
	   		beego.Error("update the device data error", err)
	   		return
	   	} */
	beego.Info("更新完毕")
}
