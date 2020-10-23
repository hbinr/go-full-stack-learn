package interfaces

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

/*
	如何优雅的使用接口(Go-style)，记住一个原则:
		"Accept interfaces, return structs"
		即：接口作为参数，结构体作为返回值
*/
type CURD interface {
	insert()
}

type UserDao struct {
	Name string
}

func (u *UserDao) insert() {
	fmt.Println("UserDao")
}

type UserService struct {
}

func (s *UserService) insert() {
	fmt.Println("userservice")
}

func GetUserScoreHandler(user CURD) gin.HandlerFunc {
	return func(c *gin.Context) {
		user.insert()
		c.String(200, "ssssss")
	}
}

func TestInterface(t *testing.T) {

	//u := &UserService{}
	//var i interface{} = u
	//o, ok := i.(CURD) // 类型断言，判断实例u是否实现了接口CURD
	//if ok {
	//	fmt.Println(o, ok)
	//} else {
	//	fmt.Println("ss")
	//}
	//r := gin.Default()
	//r.GET("/get", GetUserScoreHandler(u))
	//r.Run()
}
