package protocol

import "time"

// 设备信息
type Device struct {
	Id         int64     //主键
	Uuid       string    `orm:"index"` //手机UUID
	Serial     string    `orm:"index"` //手机序号
	Usb        string    `orm:"index"` //Devpath
	Model      string    //手机型号
	SdkVersion string    //SDK版本
	Version    string    //版本
	Abi        string    //二进制接口
	Height     int       //高度
	Width      int       //宽度
	Imei       string    //IMEI
	Online     int       //设备是否在线 0 不在线 1 在线
	GroupId    int64     `orm:"index"` //组别
	NickName   string    //手机备注
	Order      int       //排序
	UpdateTime time.Time //更新时间
	CreateTime time.Time `orm:"type(datetime);auto_now_add"` //时间
	Remark     string    `orm:"type(text)"`                  // 备注
	Uid        int64     //所属用户ID
	Mode       int       `orm:"default(0)"` //是否群控 [0: 未群控 1：群控 ]

	OrgId       int64  `orm:"-"` //组织Id
	GroupName   string `orm:"-"`
	OnlineName  string `orm:"-"`
	Fans        int64  `orm:"-"` // 每个手机的粉丝数量
	Wechat      string `orm:"-"` // 我的微信号
	DayMsgTotal int64  `orm:"-"` // 消息总量
}
