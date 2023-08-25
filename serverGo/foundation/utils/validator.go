package utils

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

// 定义一个全局翻译器T
var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

	}
	return
}

// 反射获取字段tag上的错误信息
func GetValigMsg(err error, obj interface{}) string {
	// obj 为及结构体指针
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			// 根据报错的字段名 来获取 是结构体中具体哪个字段
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				fmt.Println("f.tag = ", f.Tag.Get("msg"))
				return f.Tag.Get("msg") // 仅返回找到的第一个错误
			}
		}
	}
	return err.Error()
}
