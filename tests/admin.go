package main

import (
	"fmt"
	//"reflect"
	"encoding/json"
	//"github.com/goinggo/mapstructure"
	//"time"
	//"strconv"

	"user-api/models"
	"user-api/include"
)

func main() {

//    var str string
//    str = "测试 ' 大\" /*   */  --  over"
//    fmt.Println(str)
//    str = conf.FilteString(str)
///    fmt.Println(str)
//
//    return


	
	//json转map
	var err error
	var reqJson string

	reqJson = `
	{
		"Admin_id" : 2,
		"Login_username": "hiloy",
		"Realname": "刘硕嘉",
		"Admin_email": "hiloy@landtu.com",
		"Mobile": "13725588389",
		"Admin_status": 1,
		"Create_time": 123333,
		"Admin_role": 1,
		"Department_code": "000200010001",
		"Department_name": "线上事业部 - 淘宝事业部 - 综合B组",
		"Company_code": "0001",
		"Company_name": "深圳市刘墨轩股份有限公司",
		"Companffy_name": "深圳市刘墨轩股份有限公司",
		"Test": "测试",
		"Is_delete": 0,
		"Admin_balance":101.01
	}
	`

	//json 2 struct
	var admin models.Admin
	AdminDB := admin.InitDB()

	err = json.Unmarshal([]byte(reqJson), &admin)

	if err!=nil {
		fmt.Println(err)
		return
	}
	conf.DoFiledAndMethod(admin)

	//这里必须传结构体
	fmt.Println("\n创建数据开始")
	fmt.Println("========================")
	errCreate := AdminDB.Create(&admin)
	if errCreate!=nil {
		fmt.Println(errCreate)
	}else{
		fmt.Println(admin)
	}
	fmt.Println("========================")
/*

	fmt.Println("\n更新数据开始")
	fmt.Println("========================")

	d := make( map[string]interface{} )
	d["Create_time"] = time.Now().Unix()
	d["Login_username"] = "hiloy" + strconv.FormatInt(time.Now().Unix(),10)
	d["Realname"] = "刘硕嘉" + strconv.FormatInt(time.Now().Unix(),10)
	d["Realname_nooo"] = d["Realname"].(string) + strconv.FormatInt(time.Now().Unix(),10)

	conf.SetStructFieldByJsonName( &admin, d ) 
	errUpdate := AdminDB.Update(&admin)
	if errUpdate!=nil {
		fmt.Println(errCreate)
	}else{
		fmt.Println(admin)    
	}
	fmt.Println("========================")

	fmt.Println("\n获取一条数据 根据map")
	fmt.Println("========================")


	var adminRead models.Admin
	searchReadKey := make( map[string]interface{} )
	searchReadKey["Mobile"] = "13725588389"
	searchReadKey["Create_time_max"] = "1556274510"
	searchReadKey["Realname_like"] = "刘硕嘉"
	errGetOneDataByMap := AdminDB.GetOneDataByMap(&adminRead,searchReadKey)
	if errGetOneDataByMap!=nil {
		fmt.Println(errGetOneDataByMap)
	}else{
		fmt.Println(adminRead)


	}

	//conf.DoFiledAndMethod(adminRead)

	fmt.Println("========================")


	fmt.Println("\n获取多条数据 根据map")
	fmt.Println("========================")


	var adminReadMore []models.Admin
	searchReadKeyMore := make( map[string]interface{} )
	searchReadKeyMore["Realname_like"] = "刘硕嘉"
	searchReadKeyMore["Mobile"] = "13725588389"
	searchReadKeyMore["Create_time_min"] = "1556273857"

	order_by := make(map[string]string)
	order_by["Create_time"] = "desc"
	order_by["admin_id"] = ""
	searchReadKeyMore["order_by"] = order_by

	query_page := make(map[string]int)
	query_page["start"] = 0
	query_page["num"] = 6
	searchReadKeyMore["query_page"] = query_page

	errGetOneDataByMapMore := AdminDB.QueryData(&adminReadMore,searchReadKeyMore)
	if errGetOneDataByMapMore!=nil {
		fmt.Println(errGetOneDataByMapMore)
	}else{
		fmt.Println(adminReadMore)
	}

	num,errorNum := AdminDB.QueryNum(searchReadKeyMore)
	if errorNum!=nil {
		fmt.Println(errorNum)
	}else{
		fmt.Println("总共", num , "条记录")
	}
	

	fmt.Println("========================")


	fmt.Println("\n更新部分数据开始")
	fmt.Println("========================")

	editMap := make( map[string]interface{} )
	editMap["Login_username"] = "h" + strconv.FormatInt(time.Now().Unix(),10)
	editMap["Realname"] = "刘" + strconv.FormatInt(time.Now().Unix(),10)
	editMap["admin_balance"] = "999.99"

	searchKey := make( map[string]interface{} )
	searchKey["Admin_email"] = "hiloy@landtu.com"
	searchKey["Realname"] = "刘1556433877"

	num, errUpdateMore := AdminDB.UpdateMore(editMap,searchKey)
	if errUpdateMore!=nil {
		fmt.Println(errUpdateMore)
	}else{
		 fmt.Println( "成功更新", num , "记录"  )   
	}


	fmt.Println("========================")



	searchDelete := make( map[string]interface{} )
	searchDelete["admin_id"] = 11

	num, errDelete := AdminDB.Delete(searchDelete)
	if errDelete!=nil {
		fmt.Println(errUpdateMore)
	}else{
		fmt.Println( "删除更新", num , "记录"  )	
	}





/*
	err = admin.Create()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(admin)
 /*
	//json 2 map
	d,err := conf.Json2Map(reqJson)
	fmt.Println(d)
	if( err!=nil ){    
		fmt.Println(err)
		return
	}

	d["CreateTime"] = time.Now().Unix()
	d["LoginUsername"] = d["LoginUsername"].(string) + strconv.FormatInt(time.Now().Unix(),10)
	d["Realname"] = d["Realname"].(string) + strconv.FormatInt(time.Now().Unix(),10)
	d["Realname_nooo"] = d["Realname"].(string) + strconv.FormatInt(time.Now().Unix(),10)

	fmt.Println(d)

	fmt.Println("\n\n开始创建数据")
	var adminCreate models.Admin
	j, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(j, &adminCreate)
	err = adminCreate.Create()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(adminCreate)

	return
	//数据发生变更的话明确项目可以直接用对象赋值，如果有多个变量， 可以将 结构变量转成map后赋值后重新转为结构


	//获取单个记录
	fmt.Println("\n\n开始获取单条记录")

	//adminRead := new(models.Admin)
	var adminRead models.Admin

	adminRead.AdminId = 10
	err = adminRead.GetAdmin()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(adminRead)


/*
	fmt.Println("\n\n开始更新")
	adminRead.LoginUsername = adminRead.LoginUsername + "更新"
	adminRead.Realname = adminRead.Realname
	adminRead.AdminEmail = adminRead.AdminEmail + "更新"
	adminRead.Mobile = "13725588888"
	adminRead.AdminStatus = 0
	adminRead.CreateTime = 123456789
	adminRead.AdminRole = 0
	adminRead.DepartmentCode = "1001"
	adminRead.DepartmentName = "1001更新"
	adminRead.CompanyCode = "1002"
	adminRead.CompanyName = "1002更新"
	adminRead.IsDelete = 1

	err = adminRead.Update()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(adminRead)


	fmt.Println("\n\n开始更新部分")
	reqJson = `
	{
		"AminId": "12",
		"LoginUsername": "hiloy1",
		"Realname": 123,
		"AdminEmail": "hiloy@landtu.com3",
		"Mobile": "13725588384",
		"CreateTime":987654321,
		"AdminBalance":102.23
	}
	`
	//json 2 struct
	//json 2 map
	editMap,err := conf.Json2Map(reqJson)
	if( err!=nil ){    
		fmt.Println(err)
		return
	}

	fmt.Println(editMap)
	conf.DoFiledAndMethod(adminRead)
	conf.SetStructFieldByJsonName( &adminRead, editMap )
	conf.DoFiledAndMethod(adminRead)

/*

	*/
	/*
	searchKey := make(map[string]string)

	var adminList []models.Admin
	num,err := models.QueryAdmin( &adminList, searchKey )
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
	fmt.Println(adminList)

	/*
	//select one an update
	var adminRead models.Admin
	adminRead.AdminId = 4
	err = models.GetAdmin(&adminRead)
	if err!=nil {
		fmt.Println(err)
		return
	}
	adminRead.LoginUsername = adminRead.LoginUsername + "更新"
	adminRead.Realname = adminRead.Realname
	adminRead.AdminEmail = adminRead.AdminEmail + "更新"
	adminRead.Mobile = "13725588888"
	adminRead.AdminStatus = 0
	adminRead.CreateTime = 123456789
	adminRead.AdminRole = 0
	adminRead.DepartmentCode = "1001"
	adminRead.DepartmentName = "1001更新"
	adminRead.CompanyCode = "1002"
	adminRead.CompanyName = "1002更新"
	adminRead.IsDelete = 1

	err = models.DeleteAdmin( 5 )
	if err!=nil {
		fmt.Println(err)
		return
	}

	fmt.Println(adminRead)
	*/

	//conf.ShowStruct(admin)
	return
}
