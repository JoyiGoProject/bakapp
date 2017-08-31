package data

import (
	"encoding/json"
	"fmt"
	"bakapp/models"
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

func GetData() {
	device, err := models.Connect()
	if err != nil {
		beego.Error("get the data is error", err)
		return
	}
	var d [1000]Device
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
	SaveFile(b)
}

func SaveFile(b []byte) {
	filename := "./htrans.json"
	var f *os.File
	var err error
	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		os.Remove(filename)
		fmt.Println("文件不存在")
		f, err = os.Create(filename) //创建文件
		fmt.Println("文件已创建")
	} else {
		f, err = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	n, err := io.WriteString(f, string(b)) //写入文件(字符串)
	if err != nil {
		beego.Error("write the data error ", err)
		return
	}
	fmt.Printf("写入 %d 个字节n", n)
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
	filename := "./htrans.json"
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
	err = models.ConnectHcluster(msgs)
	if err != nil {
		beego.Error("get the data error", err)
		return
	}
	beego.Info("执行完毕")
}
