package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// MysqlConfig 结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     uint16 `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig 结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Password string `ini:"password"`
	Database uint8  `ini:"database"`
}

// Config 结构体
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func validParam(fileName string, data interface{}) (err error) {
	if len(fileName) == 0 {
		err = errors.New("fileName param is empty")
		return
	}

	if data == nil {
		err = errors.New("data param is nil")
		return
	}

	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return
	}

	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer")
		return
	}

	return
}

// loadIni 加载配置
func loadIni(fileName string, data interface{}) (err error) {
	// 1. 参数校验
	err = validParam(fileName, data)
	if err != nil {
		return
	}

	// 2. 加载文件到配置中
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("ioutil.ReadFile(fileName) failed, err=%v\n", err)
		return
	}

	dataT := reflect.TypeOf(data).Elem()
	dataV := reflect.ValueOf(data).Elem()
	var dataItemT *reflect.StructField
	var dataItemV reflect.Value
	lines := strings.Split(string(b), "\r\n")
	for idx, line := range lines {
		// fmt.Printf("%v\n", line)
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			//忽略空行
			continue
		}

		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			//忽略注释
			continue
		}

		if strings.HasPrefix(line, "[") {
			//处理section
			if !strings.HasSuffix(line, "]") {
				err = fmt.Errorf("line:%d, syntax error, missing ]", idx)
				return
			}

			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			for i := 0; i < dataT.NumField(); i++ {
				field := dataT.Field(i)
				if field.Tag.Get("ini") == sectionName {
					if field.Type.Kind() != reflect.Struct {
						fmt.Printf("field[%s] is not a struct", field.Name)
						continue
					}

					dataItemT = &field
					dataItemV = dataV.Field(i)
					break
				}
			}
		} else {
			if dataItemT == nil {
				fmt.Printf("dataItemT is nil, ignore line:%s\n", line)
				continue
			}

			//处理key value
			kvs := strings.Split(line, "=")
			if len(kvs) != 2 {
				err = fmt.Errorf("line:%d, syntax error, only permit key=value", idx)
				return
			}

			key := kvs[0]
			value := kvs[1]
			for i := 0; i < dataItemT.Type.NumField(); i++ {
				fieldT := dataItemT.Type.Field(i)
				if fieldT.Tag.Get("ini") == key {
					fieldV := dataItemV.Field(i)
					switch fieldT.Type.Kind() {
					case reflect.String:
						fieldV.SetString(value)
					case reflect.Bool:
						if val, err := strconv.ParseBool(value); err == nil {
							fieldV.SetBool(val)
						}
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						if val, err := strconv.ParseInt(value, 10, 64); err == nil {
							fieldV.SetInt(val)
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						if val, err := strconv.ParseUint(value, 10, 64); err == nil {
							fieldV.SetUint(val)
						}
					default:
						fmt.Printf("unspoort type:%s", fieldT.Type.Kind().String())
					}
					break
				}
			}
		}

	}

	return
}

func main() {
	config := new(Config)
	err := loadIni("./conf.ini", config)
	if err != nil {
		fmt.Printf("loadIni failed, err=%v\n", err)
		return
	}
	fmt.Printf("%+v", config)
}
