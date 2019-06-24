package conf

/*
	使用前请将工作目录 ln -s 做一个链接至 go的根目录/src 下

	本项目是  /home/liushuojia/admin
*/
import (
	"reflect"
	"encoding/json"
	"strings"
	"time"
	"math"
	"math/rand"
	"crypto/md5"
	
	"encoding/hex"
	"fmt"
	"strconv"
	"os"
	"sort"
	"regexp"
)

func Encryption( data map[string]interface{} ) ( md5String string ) {
   // To store the keys in slice in sorted order
   // 与php的 ksort($array,SORT_STRING); 结果一致
    var sslice []string
    for key, _ := range data {
        sslice = append(sslice, key)
    }

    sort.Strings(sslice)
    //在将key输出
    md5String = ""
    for _, v := range sslice {
    	md5String += "&" + v + "=" + SetValueToType(data[v],"string").(string)
    }
  	md5String =  MD5(md5String[1:])
	return
}

// 生成32位MD5
func MD5(text string) string {
	/*
	data := []byte("hello world")
	s := fmt.Sprintf("%x", md5.Sum(data))
	fmt.Println(s)

	// 也可以用这种方式
	h := md5.New()
	h.Write(data)
	s = hex.EncodeToString(h.Sum(nil))
	fmt.Println(s)
	*/
	ctx := md5.New()
	ctx.Write([]byte(text))
	return strings.ToUpper( hex.EncodeToString(ctx.Sum(nil)) )
}

//判断是否手机号码
func CheckMobile( mobile string) (flag bool) {
   reg := `^1\d{10}$`
   rgx := regexp.MustCompile(reg)
   flag = rgx.MatchString(mobile)
   return
}

// 用的是 obj => json => map
func Obj2Map(obj interface{}) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})
	j, err := json.Marshal(obj)
	if err != nil {
		return
	}
	err = json.Unmarshal(j, &data)
	return
}

// 保留原来Obj的子项数据类型
func Obj2Map2(obj interface{}) (data map[string]interface{}, err error) {
	data = make( map[string]interface{} )
	elem := reflect.ValueOf(obj).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		data[ relType.Field(i).Name ] = elem.Field(i).Interface()
	}
	return
}

//注意这里的 objType 可以说map/或者struct.   必须带 &才会更改数据成功
func Map2Struct( ObjType interface{}, data map[string]interface{} ) (err error) {
	j, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = json.Unmarshal(j, ObjType)
	if err != nil {
		return
	}
	return
}

func Json2Map( jsonString string ) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})
	err = json.Unmarshal( []byte(jsonString), &data)
	return
}

//字符串过滤
func FilteString(str string) (string) {

	str = strings.Replace(str, "--", "", -1)
	str = strings.Replace(str, "/*", "", -1)
	str = strings.Replace(str, "*/", "", -1)
	str = strings.Replace(str, "\"", "", -1)
	str = strings.Replace(str, "'", "", -1)
	return str
}

func PrintObj( view interface{} ) {
	fmt.Println( view )
}

func DoShowMap( data map[string] interface{} ) {
	fmt.Println( "显示map[string]interface{}" )
	fmt.Println( "============================================================================" )
	for k,v := range data {

		t := reflect.TypeOf(v)

		fmt.Println( k, "=>" , "("+t.Name()+")", v  )
	}
	fmt.Println( "============================================================================\n" )
}

//随机字符串
func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func  GetRandomNum(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 通过接口来获取任意参数，然后一一揭晓
func DoFiledAndMethod(input interface{}) {

	fmt.Println( "显示结构体内容" )
	fmt.Println( "============================================================================" )

	getType := reflect.TypeOf(input)
	if(getType.Name()=="") {
		fmt.Println( "只能处理结构，不能处理指针型变量" )
		return
	}

	fmt.Println("get Type is :", getType.Name())
	fmt.Println(getType)

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	/*
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)(err error)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
	*/
	fmt.Println( "============================================================================\n" )

}

/*
	处理类型 string bool int float
*/
func SetValueToType( input interface{}, returnType string ) (returnValue interface{}) {
	switch str := input.(type) {
		case string:
			switch returnType {
				case "string":
					returnValue = str
				case "int","int8","int16","int32","int64":
					//int 后面的整数
					if(returnType=="int"){
						returnValue,_ = strconv.Atoi( str )
					}else{
						s := string([]rune(returnType)[len("int"):])
						intType,_ := strconv.Atoi( s )
						returnValue,_ = strconv.ParseInt(str, 10, intType)
					}
				case "float32", "float64":
					//float 后面的整数
					s := string([]rune(returnType)[len("float"):])
					floatType,_ := strconv.Atoi( s )
					returnValue,_ = strconv.ParseFloat(str, floatType)

					//保留两位小数
					returnValue,_ = strconv.ParseFloat( fmt.Sprintf("%.2f", returnValue), 64 )
					if( floatType==32 ){
						returnValue = float32(returnValue.(float64))
					}
				case "bool":
					switch strings.ToLower(str) {
						case "true","1":
							returnValue = true
						case "false","0":
							returnValue = false
						default:
							returnValue = false
					}
			}
		case bool:
			switch returnType {
				case "string":
					if str {
						returnValue = "true"
					}else{
						returnValue = "false"
					}
				case "int","int8","int16","int32","int64":
					//int 后面的整数
					s := string([]rune(returnType)[len("int"):])
					intType,_ := strconv.Atoi( s )
					tmp := "0"
					if str {
						tmp = "1"
					}
					returnValue,_ = strconv.ParseInt(tmp, 10, intType)
				case "float32", "float64":
					//float 后面的整数
					s := string([]rune(returnType)[len("float"):])
					floatType,_ := strconv.Atoi( s )
					tmp := "0"
					if str {
						tmp = "1"
					}
					returnValue,_ = strconv.ParseFloat(tmp, floatType)

					if( floatType==32 ){
						returnValue = float32(returnValue.(float64))
					}
				case "bool":
					returnValue = str
			}
		case int:
			strInt := strconv.FormatInt(int64(str),10)
			returnValue = SetValueToType(strInt,returnType)
		case int8:
			strInt := strconv.FormatInt(int64(str),10)
			returnValue = SetValueToType(strInt,returnType)
		case int16:
			strInt := strconv.FormatInt(int64(str),10)
			returnValue = SetValueToType(strInt,returnType)
		case int32:
			strInt := strconv.FormatInt(int64(str),10)
			returnValue = SetValueToType(strInt,returnType)
		case int64:
			strInt := strconv.FormatInt(str,10)
			returnValue = SetValueToType(strInt,returnType)
		case float32:
			strfloat := strconv.FormatFloat(float64(str), 'f', -1, 32)
			returnValue = SetValueToType(strfloat,returnType)
		case float64:
			strfloat := strconv.FormatFloat(str, 'f', -1, 64)
			returnValue = SetValueToType(strfloat,returnType)
	}
	return
}

//      初始化结构里面的数据根据map
//      adminRead 为结构非指针
//      conf.SetStructFieldByJsonName( &adminRead, editMap )    必须带 & 
func SetStructFieldByJsonName( input interface{}, editFields map[string]interface{} ) (err error) {

	getType := reflect.TypeOf(input).Elem()
	getValue := reflect.ValueOf(input).Elem()

	for i := 0; i < getType.NumField(); i++ {

		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		if _,ok := editFields[ field.Name ]; ok {
			switch value.(type) {// 最终数据格式
			case string:
				setValue := SetValueToType( editFields[ field.Name ], "string" )
				getValue.FieldByName(field.Name).SetString( setValue.(string) )
			case int,int8,int16,int32,int64:
				setValue := SetValueToType( editFields[ field.Name ], "int64" )
				getValue.FieldByName(field.Name).SetInt( setValue.(int64) )
			case float32,float64:
				setValue := SetValueToType( editFields[ field.Name ], "float64" )
				getValue.FieldByName(field.Name).SetFloat( setValue.(float64) )
			}
		}
	}
	return
}

func Struct2String( obj interface{} ) () {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	fmt.Println( "" )
	fmt.Println( "======================================" )
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		fieldValue := v.Field(i)

		tag := field.Tag.Get("column")
		fmt.Printf("%d. %v(%v), tag:'%v'\n", i+1, field.Name,  field.Type.Name(), tag)
		fmt.Println( field.Type.Name() , " =>", fieldValue )

	}
	fmt.Println( "======================================" )
	fmt.Println( "" )

	return
}

func ShowStruct( obj interface{} ){
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		fmt.Printf( "%d. %v(%v), tag:'%v'\n", i+1, field.Name, field.Type.Name() )

		fieldValue := v.Field(i)

		switch fieldValue.Kind() {
		case reflect.String:
			fmt.Println( fieldValue.String())
		case reflect.Int:
			fmt.Println( fieldValue.Int())
		}

	}
	return
}


func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}


func HexDec(h string) (n int64) {
	s := strings.Split(strings.ToUpper(h), "")
	l := len(s)
	i := 0
	d := float64(0)
	hex := map[string]string{"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15"}
	for i = 0; i < l; i++ {
	  	c := s[i]
	  	if v, ok := hex[c]; ok {
		 	c = v
	  	}
	  	f, err := strconv.ParseFloat(c, 10)
	  	if err != nil {
		 	return -1
	  	}
	  	d += f * math.Pow(16, float64(l-i-1))
	}
	return int64(d)
}
