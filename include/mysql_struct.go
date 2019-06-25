package conf

/*

	使用前请将工作目录 ln -s 做一个链接至 go的根目录/src 下
	本项目是  /home/liushuojia/go/src/admin

	结构体处理数据， 结构体的好处： 结构体初始化时候数据格式已经完全处理好了

	由于设计习惯， 数据库的表字段都是小写， 所有字段操作这边都会过滤大小写
	
*/

import (
	"reflect"
	"fmt"
	"errors"
	"strings"

	"github.com/astaxie/beego"

	"github.com/gohouse/gorose"
	_ "github.com/gohouse/gorose/driver/mysql"
)

var connection *gorose.Connection


type SqlDB struct {
	Session *gorose.Session
	TableName string
	PrimaryKey string
	FieldString string
	FieldList map[string]interface{}
	ShowSqlFlag bool
}


func ( sqlDB *SqlDB ) FieldInit() () {
	FieldString := ""
	for k,_ := range sqlDB.FieldList {
		FieldString += "," + k
	}
	sqlDB.FieldString = FieldString[ 1: ]
	sqlDB.ShowSqlFlag = false
	return
}

func ( sqlDB *SqlDB ) Connect() ( err error) {
	if connection == nil {

		var DbConfig = &gorose.DbConfigSingle {
			Driver:          "mysql", // 驱动: mysql/sqlite/oracle/mssql/postgres
			EnableQueryLog:  true,   // 是否开启sql日志
			SetMaxOpenConns: 0,    // (连接池)最大打开的连接数，默认值为0表示不限制
			SetMaxIdleConns: 0,    // (连接池)闲置的连接数
			Prefix:          "", // 表前缀
			Dsn:             beego.AppConfig.String("mysql::host"), // 数据库链接
		}
		connection, err = gorose.Open(DbConfig)

		if err != nil {
			return
		}
	}

	if sqlDB.Session == nil {
		sqlDB.Session = connection.NewSession()
	}

	return
}

func ( sqlDB *SqlDB ) Close() {
	if connection != nil {
		//如果是持久性连接, 可以注释掉下面
		connection.Close()
		connection = nil
	}
	return
}



func ( sqlDB *SqlDB ) checkStatus() ( err error ) {
	if connection == nil || sqlDB.Session == nil {
		err = sqlDB.Connect()
		if err != nil {
			return
		}        
	}
	return
}

func (sqlDB SqlDB ) mapKey2Lower( data map[string]interface{} ) ( returnMap map[string]interface{} ) {
	returnMap = make( map[string]interface{} )
	for k,v := range data {
		returnMap[ strings.ToLower( k ) ] = v
	}
	return
}

func (sqlDB SqlDB ) Data2map( data interface{} ) ( returnMap map[string]interface{}, err error ) {
	returnMap = make( map[string]interface{} )
	elem := reflect.ValueOf(data).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		returnMap[ strings.ToLower(relType.Field(i).Name)] = elem.Field(i).Interface()
	}
	return
}

//查询条件初始化 根据 sqlDB.FieldList 扩展查询
//
//searchKey 除了 order_by map[string] string   query_page map[string]int
//          其他的key都是 string => string/int64
//
func ( sqlDB SqlDB ) searchKeyInit( Session *gorose.Session, searchKey map[string]interface{}  ) ( err error ) {

	var kLeft string
	var kRight string
	var kMiddle string

	haveSearchKeyFlag := false
	for k,v := range searchKey {
		k = strings.ToLower(k)
		switch( k ){
		case "order_by":
			//排序
			orderByString := ""
			orderBy:=v.(map[string]string)
			for ko,vo := range orderBy {
				ko = strings.ToLower(ko)
				vo = strings.ToLower(vo)
				if _, ok := sqlDB.FieldList[ko]; ok {
					if vo=="" || vo=="desc" || vo=="asc" {
						orderByString += "," + ko  + " " + vo
					}
				}
			}
			if orderByString!="" {
				Session.Order( orderByString[1:] )
			}
		case "query_page":
			queryPage:=v.(map[string]int)
			//获取数量
			if num, ok := queryPage["num"]; ok {
				Session.Limit(num)
				if start, ok := queryPage["start"]; ok {
					Session.Offset(start)
				}
			}else{
				//不带limit限制下获取数量，最多10000
				//Session.Limit( 10000 )
			}

		default:
			// have key
			if slv, ok := sqlDB.FieldList[k]; ok {
				//存在
				Session.Where( k, "=", SetValueToType( v, reflect.TypeOf(slv).Name() ) )
				haveSearchKeyFlag = true
				continue
			}

			// min or max
			kLeft = k[ :(len(k)-4) ]
			kRight = k[ (len(k)-4): ]
			kMiddle = ""
			if kRight=="_min" {
				kMiddle = ">="
			}
			if kRight=="_max" {
				kMiddle = "<="
			}
			if kMiddle!="" {
				if slv, ok := sqlDB.FieldList[kLeft]; ok {
					Session.Where( kLeft, kMiddle, SetValueToType( v, reflect.TypeOf(slv).Name() ) )
					haveSearchKeyFlag = true       
					continue
				}
			}

			// like
			kLeft = k[ :(len(k)-5) ]
			kRight = k[ (len(k)-5): ]
			kMiddle = "like"
			if kRight=="_like" {
				if _, ok := sqlDB.FieldList[kLeft]; ok {
					Session.Where(kLeft,"like", "%" + SetValueToType( v, "string" ).(string) + "%" ) 
					haveSearchKeyFlag = true
					continue                    
				}
			}
		}
	}

	if _, ok := searchKey["query_page"]; !ok {
		//暂时放开限制
		//Session.Limit( 1 )  
	}
	if !haveSearchKeyFlag {
		err = errors.New("searchKey 查询条件为空");
		return
	}
	return 
}

//  data 传递必须是指针  Create( &admin )  admin为结构体
func ( sqlDB SqlDB ) Create( data interface{} ) ( err error ) {
	err=sqlDB.checkStatus();
	if err!=nil {
		return
	}

	dataEdit,err := sqlDB.Data2map( data )
	if err!=nil {
		return
	}

	//去掉主键 所有数据库都给个自动编码
	delete(dataEdit, sqlDB.PrimaryKey )

	id, err := sqlDB.Session.Table( data  ).Data( dataEdit ).InsertGetId()

	if sqlDB.ShowSqlFlag {
		fmt.Println(sqlDB.Session.LastSql )
	}

	if err != nil {
		return
	}
	if id <=0 {
		err = errors.New("创建数据失败");
		return
	}
	
	dataEdit[ sqlDB.PrimaryKey ] = id
	Map2Struct(data,dataEdit)
	return
}

//  data 传递必须是指针  Update( &admin )  admin为结构体
func ( sqlDB SqlDB ) Update( data interface{} ) ( err error ) {
	err=sqlDB.checkStatus();
	if err!=nil {
		return
	}

	dataEdit,err := sqlDB.Data2map( data )
	if err!=nil {
		return
	}

	//获取主键，并清理主键
	primaryKey := dataEdit[ sqlDB.PrimaryKey ]
	delete(dataEdit, sqlDB.PrimaryKey )

	num, err := sqlDB.Session.Table( data ).
		Data(dataEdit).
		Where(sqlDB.PrimaryKey, primaryKey).
		Update()

	if sqlDB.ShowSqlFlag {
		fmt.Println(sqlDB.Session.LastSql )
	}

	if err != nil {
		return
	}
	if num <=0 {
		err = errors.New("更新数据失败");
		return
	}

	return
}

/*
获取一个或多个数据 多个数据传进来的data必须是 []结构
searchKey 查询条件， 必须过滤 order_by make(map[string]string)  query_page := make(map[string]int)
*/
func ( sqlDB SqlDB ) QueryData( data interface{}, searchKey map[string]interface{}  ) ( err error ) {
	err=sqlDB.checkStatus();
	if err!=nil {
		return
	}
	dataDb := sqlDB.Session.Table( data ).Fields( sqlDB.FieldString )
	err = sqlDB.searchKeyInit( dataDb, searchKey )
	if err!=nil {
		return
	}
	err = dataDb.Select()
	if sqlDB.ShowSqlFlag {
		fmt.Println(sqlDB.Session.LastSql )
	}
	
	return
}

func ( sqlDB SqlDB ) QueryNum( searchKey map[string]interface{} ) (num int64, err error) {
	err=sqlDB.checkStatus();
	if err!=nil {
		return
	}
	dataDb := sqlDB.Session.Table( sqlDB.TableName )
	delete( searchKey, "order_by" )
	delete( searchKey, "query_page" )

	err = sqlDB.searchKeyInit( dataDb, searchKey )
	if err!=nil {
		return
	}

	num, err = dataDb.Count( sqlDB.PrimaryKey )
	if sqlDB.ShowSqlFlag {
		fmt.Println(sqlDB.Session.LastSql )
	}

	if err!=nil {
		return
	}

	return
}

//获取一条数据
func ( sqlDB SqlDB ) GetOne( data interface{}, id int64) ( err error ) {
	searchKey := make(map[string]interface{})
	query_page := map[string]int { "num":1 }
	searchKey["query_page"] = query_page
	searchKey[ sqlDB.PrimaryKey ] = id
	err = sqlDB.QueryData(data,searchKey)
	return
}

func ( sqlDB SqlDB ) GetOneDataByMap( data interface{}, searchKey map[string]interface{} ) ( err error ) {
	query_page := map[string]int { "num":1 }
	searchKey["query_page"] = query_page
	err = sqlDB.QueryData(data,searchKey)
	return
}

//更新部分数据
func ( sqlDB SqlDB ) UpdateMore( editObj map[string]interface{}, searchKey map[string]interface{} )( num int64,err error ) {
	err=sqlDB.checkStatus();
	if err!=nil {
		return
	}

	editAction := sqlDB.mapKey2Lower(editObj)
	//去掉主键 所有数据库都给个自动编码
	delete(editAction, sqlDB.PrimaryKey )

	fmt.Println(editAction);
	haveEditFlag := false
	for k,v := range editAction {
		if slv, ok := sqlDB.FieldList[k]; ok {
			//存在
			editAction[k] = SetValueToType( v, reflect.TypeOf(slv).Name() )
			haveEditFlag = true
		}else{
			delete(editAction,k)
		}
	}
	if( !haveEditFlag ){
		err = errors.New("更新内容为空，请您检查 editObj 的项目");
		return
	}
	dataDb := sqlDB.Session.Table( sqlDB.TableName )
	dataDb.Data( editAction )

	err = sqlDB.searchKeyInit( dataDb, searchKey )
	if err!=nil {
		return
	}

	num, err = dataDb.Update()

	if sqlDB.ShowSqlFlag {
		fmt.Println(sqlDB.Session.LastSql )
	}

	if err!=nil {
		return
	}
	/*
	if num <=0 {
		err = errors.New("更新数据失败");
		return
	}
	*/
	return
}

// 软删 设置 is_delete
func ( sqlDB SqlDB ) Delete(searchKey map[string]interface{} )( num int64,err error ) {
	err=sqlDB.checkStatus();
	if err!=nil {
		return
	}

	editObj := make(map[string]interface{})
	editObj["is_delete"] = 1
	num, err = sqlDB.UpdateMore( editObj, searchKey ) 
	if err!=nil {
		err = errors.New("删除数据失败");		
		return
	}
	/*
	if num <=0 {
		err = errors.New("删除数据失败");
		return
	}
	*/
	return
}

//删除数据库记录 不保险
func ( sqlDB SqlDB ) DeleteData(searchKey map[string]interface{} )( num int64,err error ) {
	err=sqlDB.checkStatus();
	if err!=nil {
		return
	}

	dataDb := sqlDB.Session.Table( sqlDB.TableName )
	err = sqlDB.searchKeyInit( dataDb, searchKey )
	if err!=nil {
		return
	}

	num, err = dataDb.Delete()
	if sqlDB.ShowSqlFlag {
		fmt.Println(sqlDB.Session.LastSql )
	}
	if err!=nil {
		return
	}
	if num <=0 {
		err = errors.New("删除数据失败");
		return
	}
	return
}
