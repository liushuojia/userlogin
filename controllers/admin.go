package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"fmt"
	"time"
	"github.com/astaxie/beego"

	"user-api/models"
	"user-api/include"
)

// AdminController operations for Admin
type AdminController struct {
	beego.Controller
}

var Admin models.Admin
//登录的用户admin
var AdminToken models.Admin

// Post ...
// @Title Post
// @Description create Admin
// @Param	body		body 	models.Admin	true		"body for Admin content"
// @Success 201 {int} models.Admin
// @Failure 403 body is empty
// @router / [post]
func (c *AdminController) Post() {
	if err := c.checkToken(); err!= nil {
		c.error(err)
		return
	}
	if AdminToken.Admin_role != 1 {
		c.Ctx.Output.SetStatus(401)
		c.error(errors.New("无权限"))
		return
	}
 
	AdminDB := Admin.InitDB()
	defer AdminDB.Close()

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &Admin); err == nil {

		if err := AdminDB.Create(&Admin); err == nil {
			//c.Ctx.Output.SetStatus(201)
			returnMap := make (map[string] interface{})
			returnMap["flag"] = 0
			returnMap["msg"] = "获取成功"
			returnMap["data"] = Admin

			c.Data["json"] = returnMap
			c.ServeJSON()
			return
		} else {
			//c.Ctx.Output.SetStatus(400)
			c.error(err)
			return
		}
	} else {
		c.error(err)
		return
	}
}

// GetOne ...
// @Title Get One
// @Description get Admin by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Admin
// @Failure 410 Gone
// @router /:id [get]
func (c *AdminController) GetOne() {
	if err := c.checkToken(); err!= nil {
		c.error(err)
		return
	}
	if AdminToken.Admin_role != 1 {
		c.Ctx.Output.SetStatus(401)
		c.error(errors.New("无权限"))
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	if( id<=0 ){
		c.error(errors.New("参数传递错误"))
		return
	}

	AdminDB := Admin.InitDB()
	defer AdminDB.Close()

	err := AdminDB.GetOne( &Admin, int64(id) )

	if err != nil {
		//c.Ctx.Output.SetStatus(410)
		c.error(errors.New("无数据"))
		return
	}
	//c.Ctx.Output.SetStatus(200)
	returnMap := make (map[string] interface{})
	returnMap["flag"] = 0
	returnMap["msg"] = "获取成功"
	returnMap["data"] = Admin

	c.Data["json"] = returnMap	
	c.ServeJSON()
	return
}

// GetAll ...
// @Title Get All
// @Description get Admin
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Admin
// @Failure 403
// @router / [get]
func (c *AdminController) GetAll() {
	if err := c.checkToken(); err!= nil {
		c.error(err)
		return
	}
	if AdminToken.Admin_role != 1 {
		c.Ctx.Output.SetStatus(401)
		c.error(errors.New("无权限"))
		return
	}

	var data []models.Admin 		
	AdminDB := Admin.InitDB()
	AdminDB.ShowSqlFlag = true
	defer AdminDB.Close()

	if editMap,err := conf.Json2Map( c.GetString("query") ); err != nil {
		c.error(err)
		return
	}else{
		if err := AdminDB.QueryData( &data, editMap ); err!=nil {
			c.error(err)
			return
		}
	}

	returnMap := make (map[string] interface{})
	returnMap["flag"] = 0
	returnMap["msg"] = "获取数据成功"
	returnMap["data"] = data

	c.Data["json"] = returnMap	
	c.ServeJSON()
	return
}

// Put ...
// @Title Put
// @Description update the Admin
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Admin	true		"body for Admin content"
// @Success 200 {object} models.Admin
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AdminController) Put() {
	if err := c.checkToken(); err!= nil {
		c.error(err)
		return
	}
	if AdminToken.Admin_role != 1 {
		c.Ctx.Output.SetStatus(401)
		c.error(errors.New("无权限"))
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	if( id<=0 ){
		c.error(errors.New("参数传递错误"))
		return
	}


	editMap := make( map[string]interface{} )
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &editMap); err != nil {
		c.error(err)
		return
	}

	AdminDB := Admin.InitDB()
	defer AdminDB.Close()
	//AdminDB.ShowSqlFlag = true

	searchKey := make( map[string]interface{} )
	searchKey["Admin_id"] = id

	num, err := AdminDB.UpdateMore(editMap,searchKey)
	if( err!=nil ){
		c.error(err)
		return
	}

	returnMap := make (map[string] interface{})	
	returnMap["flag"] = 0
	tmp := conf.SetValueToType(num,"string")
	returnMap["msg"] = "更新成功, 影响 " + tmp.(string) + " 记录"
	c.Data["json"] = returnMap	
	c.ServeJSON()
	return
}


// Delete ...
// @Title Delete
// @Description delete the Admin
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AdminController) Delete() {
	if err := c.checkToken(); err!= nil {
		c.error(err)
		return
	}
	if AdminToken.Admin_role != 1 {
		c.Ctx.Output.SetStatus(401)
		c.error(errors.New("无权限"))
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	if( id<=0 ){
		c.error(errors.New("参数传递错误"))
		return
	}

	AdminDB := Admin.InitDB()
	defer AdminDB.Close()

	searchKey := make( map[string]interface{} )
	searchKey["Admin_id"] = id

	if num,err := AdminDB.Delete(searchKey); err == nil {
		//c.Ctx.Output.SetStatus(204)
		returnMap := make (map[string] interface{})
		returnMap["flag"] = 0
		tmp := conf.SetValueToType(num,"string")
		returnMap["msg"] = "删除成功, 删除 " + tmp.(string) + " 条记录"
		c.Data["json"] = returnMap
		c.ServeJSON()
		return
	} else {
		//c.Ctx.Output.SetStatus(500)
		c.error(err)
		return
	}

}


// @Title Login
// @Description Logs user into the system
// @Param	mobile		query 	string	true		"The mobile for login"
// @Param	smscode		query 	string	true		"Mobile verification code"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (c *AdminController) Login() {

	if err := c.checkToken(); err!= nil {
		c.error(err)
		return
	}

	mobile := c.GetString("mobile")
	smscode := c.GetString("code")

	checkKey := beego.AppConfig.String("sms::checkKey") + "_" + mobile

	var redisObj conf.RedisConn
	defer redisObj.Close()

	value, err := redisObj.Get( checkKey )
	if err!=nil {
		c.error(errors.New("验证码已经过期,请您重新获取"))
		return
	}

	if value!=smscode {
		c.error(errors.New("验证码输入错误,请重新输入"))
		return
	}

	AdminDB := Admin.InitDB()
	defer AdminDB.Close()

	searchKey := make( map[string]interface{} )
	searchKey["admin_mobile"] = mobile

	if err := AdminDB.GetOneDataByMap(&Admin,searchKey); err != nil {
		c.error(err)
		return
	}

	if( Admin.Admin_status==0 ){
		c.error(errors.New("账号已停用"))
		return
	}
	if( Admin.Is_delete==1 ){
		c.error(errors.New("账号已删"))
		return
	}
	
	userAgent := c.Ctx.Request.Header.Get("User-Agent")
	token := buildUid( Admin, userAgent )

	adminJson, _ := json.Marshal(Admin)
	redisKey := "admin_id_" + conf.SetValueToType(Admin.Admin_id,"string").(string)
	if err := redisObj.Set( redisKey, string(adminJson), 24*60*60); err != nil {
		c.error(err)
		return
	}

	returnMap := make (map[string] interface{})
	returnMap["flag"] = 0
	returnMap["msg"] = "获取成功"
	returnMap["data"] = token
	c.Data["json"] = returnMap
	c.ServeJSON()
	return
	fmt.Println(mobile)
	return
}


// @Title SendSms
// @Description send sms
// @Param	mobile		query 	string	true		"The mobile for sms"
// @Success 200 {string} success
// @router /sendsms [get]
func (c *AdminController) SendSms() {
	if err := c.checkToken(); err!= nil {
		c.error(err)
		return
	}

	mobile := c.GetString("mobile")
	if( !conf.CheckMobile( mobile ) ){
		c.error(errors.New("手机号码输入错误"))
		return
	}

	AdminDB := Admin.InitDB()
	defer AdminDB.Close()

	searchKey := make( map[string]interface{} )
	searchKey["admin_mobile"] = mobile

	returnMap := make (map[string] interface{})
	if err := AdminDB.GetOneDataByMap(&Admin,searchKey); err != nil {
		//c.Ctx.Output.SetStatus(410)
		returnMap["flag"] = 1
		returnMap["msg"] = "手机还未注册"
	} else {

		ranNum := conf.GetRandomNum(6)
		if err := SendSmsAction(mobile, ranNum); err!= nil {
			c.error(err)
			return
		}

		returnMap["flag"] = 0
		returnMap["msg"] = "发送短信成功"
		if( beego.AppConfig.String("runmode")=="dev" ){
			returnMap["msg"] = "发送短信成功 - " + ranNum
		}

	}

	c.Data["json"] = returnMap
	c.ServeJSON()
}


// @Title GetSmsToken
// @Description send sms
// @Param	mobile		query 	string	true		"mobile"
// @Param	verify		query 	string	true		"verify"
// @Success 200 {string} success
// @router /getToken [get]
func (c *AdminController) GetToken() {

	mobile := c.GetString("mobile")
	verify := c.GetString("verify")
	userAgent := c.Ctx.Request.Header.Get("User-Agent")

	returnMap := make (map[string] interface{})
	if( mobile=="" ){

		data := map[string] interface{} {
			"Key" : Md5Key,
			"userAgent" : userAgent,
			"verify" : conf.MD5(userAgent),
			"id" : 0,
			"timestamp" : time.Now().Unix(),
		}
		md5String := BuildToken( data )

		returnMap["flag"] = 0
		returnMap["msg"] = "获取成功"
		returnMap["data"] = md5String

		c.Data["json"] = returnMap
		c.ServeJSON()
	
	}else{

		if( !conf.CheckMobile(mobile) ){
			c.error(errors.New("手机或密钥传递错误"))
			return
		}

		AdminDB := Admin.InitDB()
		defer AdminDB.Close()

		searchKey := make( map[string]interface{} )
		searchKey["admin_mobile"] = mobile

		if err := AdminDB.GetOneDataByMap(&Admin,searchKey); err != nil {
			c.error(errors.New("手机或密钥传递错误"))
			return
		}
		
		if( verify=="" ){
			c.error(errors.New("手机或密钥传递错误"))
			return
		}
		
		if( Admin.Admin_verify!=verify ){
			c.error(errors.New("手机或密钥传递错误"))
			return
		}
		
		if( Admin.Admin_status==0 ){
			c.error(errors.New("账号已停用"))
			return
		}
		if( Admin.Is_delete==1 ){
			c.error(errors.New("账号已删"))
			return
		}

		/*
			token 放在redis里面, 方便设置有效期
		*/
		token := buildUid( Admin, userAgent )
		adminJson, _ := json.Marshal(Admin)

		var redisObj conf.RedisConn
		defer redisObj.Close()

		redisKey := "admin_id_" + conf.SetValueToType(Admin.Admin_id,"string").(string)

		if err := redisObj.Set(redisKey,string(adminJson), 24*60*60); err != nil {
			c.error(err)
			return
		}

		returnMap["flag"] = 0
		returnMap["msg"] = "获取成功"
		returnMap["data"] = token
		c.Data["json"] = returnMap
		c.ServeJSON()
	}
}

// @Title Status
// @Description check status
// @Success 200 {string} success
// @router /status [get]
func (c *AdminController) Status() {
	if err := c.checkToken(); err != nil {
		c.error(err)
		return
	}

	returnMap := make (map[string] interface{})
	returnMap["flag"] = 0
	returnMap["msg"] = "token 有效"
	c.Data["json"] = returnMap
	c.ServeJSON()
}

/*
	操作function

*/


//  smsKey = "SmsQueue"             //短信发送队列
//    smsCheckKey = "SmsCheck"        //短信检验数据
func SendSmsAction(mobile string, smscode string) ( err error) {
	
	queueKey := beego.AppConfig.String("sms::queueKey")
	checkKey := beego.AppConfig.String("sms::checkKey") + "_" + mobile

	var redisObj conf.RedisConn
	defer redisObj.Close()

	if_exit, err := redisObj.EXISTS( checkKey )
	if err != nil {
		return
	}

	if( if_exit==true ) {
		err = errors.New("短信已发送过了, 请60秒后重新获取")
		return
	}

	err = redisObj.Set( checkKey, smscode, 60 )
	if err != nil {
		return
	}

	tmpArray := make( map[string]string )
	tmpArray["mobile"] = mobile
	tmpArray["content"] = "您的验证码是：" + smscode + "。请不要把验证码泄露给其他人。"
	j, _ := json.Marshal(tmpArray)

	err = redisObj.RPush( queueKey, string(j))
	if err != nil {
		return
	}
	return
}


func buildUid( dataAdmin models.Admin, userAgent string ) ( uid string ) {
	data := map[string]interface{}  {
		"Key" : Md5Key,
		"userAgent" : userAgent,
		"verify" : dataAdmin.Admin_verify,
		"id" : dataAdmin.Admin_id,
		"timestamp" : conf.SetValueToType(time.Now().Unix(),"string").(string),
	}
	uid = BuildToken( data )
	return
}

func (c *AdminController) error( err error ) {
	returnMap := make (map[string] interface{})
	returnMap["flag"] = 1
	returnMap["msg"] = err.Error()
	c.Data["json"] = returnMap
	c.ServeJSON()
}

func (c *AdminController) checkToken() (err error) {
	token := c.Ctx.Request.Header.Get("token")
	userAgent := c.Ctx.Request.Header.Get("User-Agent")

	chrstr := strings.Split(token,"-");

	data := map[string]interface{} {
		"Key" : Md5Key,
		"userAgent" : userAgent,
		//"verify" : verify,
		//"id" : conf.SetValueToType(id,"string").(string),
		//"timestamp" : conf.SetValueToType(time.Now().Unix(),"string").(string),
	}
	switch( len(chrstr) ){
	case 2:
		data["id"] = 0
		data["verify"] = conf.MD5(userAgent)
		data["timestamp"] = conf.SetValueToType( AnyToDecimal( chrstr[0], IDNUN ) , "string").(string)

		BMD5String := conf.Encryption(data)
		if( BMD5String != chrstr[1] ) {
			err = errors.New("token 传递错误");
		}
		break;
	case 3:
		id := AnyToDecimal( chrstr[0], IDNUN )
		timestamp := AnyToDecimal( chrstr[1], IDNUN )

		if( id<=0 ){
			err = errors.New("token 传递错误");
			break;
		}

		if( (time.Now().Unix() - int64(timestamp))>24*60*60 ) {
			err = errors.New("token 超时");
			break;
		}


		/*
			写进redis 判断用cache数据做
		*/
		var redisObj conf.RedisConn
		defer redisObj.Close()

		redisKey := "admin_id_" + conf.SetValueToType(id,"string").(string)

		adminJson,err := redisObj.Get(redisKey)
		if err!=nil {
			err = errors.New("token 超时");
			break;
		}

		err = json.Unmarshal( []byte(adminJson), &AdminToken)
		if err!=nil {
			err = errors.New("token 超时");
			break;
		}

		/*
		AdminDB := Admin.InitDB()
		defer AdminDB.Close()

		if err := AdminDB.GetOne( &AdminToken, int64(id) ); err!=nil {
			err = errors.New("token 传递错误");
			break;
		}
		*/

		if( AdminToken.Admin_status==0 ){
			err = errors.New("账号已停用");
			break;
		}

		if( AdminToken.Is_delete==1 ){
			err = errors.New("账号已删");
			break;
		}

		data["verify"] = AdminToken.Admin_verify
		data["id"] = id
		data["timestamp"] = timestamp
		BMD5String := conf.Encryption(data)
		if( BMD5String != chrstr[2] ) {
			err = errors.New("token 传递错误");
		}

		break;
	default:
		err = errors.New("token 传递错误");
		break;
	}

	return
}
