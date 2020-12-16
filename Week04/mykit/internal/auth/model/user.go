package model

type User struct {
	Id         int64  `json:"id" xorm:"not null pk autoincr INT(11) 'id'"`
	Gid        uint32 `json:"gid" xorm:"not null comment('展示给客户端的ID') INT(11) 'gid'"`
	UserName   string `json:"user_name" xorm:"default 'NULL' comment('用户名') VARCHAR(60) 'user_name'"`
	Password   string `json:"password" xorm:"default 'NULL' comment('用户密码') VARCHAR(150) 'password'"`
	Channel    uint32 `json:"channel" xorm:"default NULL comment('渠道') INT(11) 'channel'"`
	Imei       string `json:"imei" xorm:"default 'NULL' comment('设备唯一标识') VARCHAR(60) 'imei'"`
	DeviceType uint32 `json:"device_type" xorm:"default NULL comment('设备类型（1 pc,2 web,3 android,4 ios）') INT(11) 'device_type'"`
	Status     int32  `json:"status" xorm:"default NULL comment('用户状态 （1 正常,2 冻结）') INT(11) 'status'"`
	CreateTime int64  `json:"create_time" xorm:"default NULL BIGINT(20) 'create_time'"`
	UpdateTime int64  `json:"update_time" xorm:"default NULL BIGINT(20) 'update_time'"`
}

func (u *User) TableName() string {
	return "im_user"
}
