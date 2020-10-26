package method

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

// SignUpParam 用户注册参数
type SignUpParam struct {
	Username   string `json:"username" form:"username" v:"username@required|length:6,30#请输入用户名|用户名长度应当在:min到:max之间"`  // 用户名
	Password   string `json:"password" form:"password" v:"password@required|length:6,16#请输入密码|密码长度应当在:min到:max之间"`    // 密码
	RePassword string `json:"rePassword" form:"rePassword" v:"rePassword@required|same:password#请输入密码|两次密码不一致，请重新输入"` // 重复密码
	NickName   string `json:"nickName" form:"nickName" v:"nickName@required#请输入中文名"`                                  // 中文名
	Email      string `json:"email" form:"email" v:"email@required|email#请输入邮箱|邮箱不合法"`                                // 中文名
}

func BindAndValid(c *gin.Context, params interface{}) error {
	_ = c.ShouldBind(params) // 展示校验库，就先不多写err判断了
	if err := gvalid.CheckStruct(params, nil); err != nil {
		return err
	}
	return nil
}

// TestMethod 方法变量和方法表达式 study
func TestMethod(t *testing.T) {
	r := gin.Default()
	var uParam SignUpParam
	r.POST("/signup", func(c *gin.Context) {
		if err := BindAndValid(c, &uParam); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"data": nil,
				"msg":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": &uParam,
			"msg":  "注册成功",
		})

	})

}
