package main

import (
	"fmt"
	"bakapp/data"
	"os"

	"github.com/astaxie/beego"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "backup" {
			data.GetData()
		} else if os.Args[1] == "update" {
			data.SaveData()
		}
	} else {
		fmt.Println("Usage:")
		fmt.Println("备份数据(backup): bakapp backup")
		fmt.Println("更新数据(update): bakapp update")
	}
	beego.Run()
}
