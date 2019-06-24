package models
import (
	"user-api/include"
)

type Admin struct {
	Admin_id			int64
	Name_en				string			//英文名
	Realname			string			//姓名
	Admin_email			string			//邮件
	Admin_email_flag	int				//邮件验证状态
	Admin_mobile		string			//手机
	Admin_mobile_flag	int				//手机验证状态
	Admin_status		int				//状态 1有效 0暂停
	Admin_verify		string 			//密钥
	Create_time			int64			//创建时间
	Admin_role			int				//管理员权限
	Is_delete			int				//是否已删 1 已删 0 默认
}

func (u *Admin) TableName() string {
	return "TUserDB.admin"
}

func (u *Admin) InitDB() ( sqldb conf.SqlDB ) {
	sqldb.PrimaryKey = "admin_id"
	sqldb.TableName = u.TableName()
	sqldb.FieldList,_ = sqldb.Data2map(u)
	sqldb.FieldInit()
	return
}
