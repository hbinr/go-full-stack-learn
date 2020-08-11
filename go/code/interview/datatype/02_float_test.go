package datatype

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

// TestFloat 验证浮点数有效位数
// float32 单精度浮点型，小数点有效位数为7位
// float64 双精度浮点型，小数点有效位数为16位
func TestFloat(t *testing.T) {
	var tmp32 float32
	tmp32 = 1.123456789                                // 小数点9位
	fmt.Printf("tmp32 的数据类型：%T，值为：%v\n", tmp32, tmp32) // tmp32 的数据类型：float32，值为：1.123457

	var tmp64 float64
	tmp64 = 1.12345678912345678                        // 小数点16位
	fmt.Printf("tmp64 的数据类型：%T，值为：%v\n", tmp64, tmp64) // tmp64 的数据类型：float64，值为：1.123457

}

//TestFloatConvert float64精度丢失问题
func TestFloatConvert(t *testing.T) {
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 19.90), 64)
	fmt.Println("num转为 float64后，貌似精度未丢失", num)
	fmt.Println("再乘以100，精度丢失", num*100)

	// 引入第三方库 decimal 来解决精度丢失为题
	num2, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", 19.90), 64)
	fmt.Println(num2)

	decimalValue := decimal.NewFromFloat(num2)
	decimalValue = decimalValue.Mul(decimal.NewFromInt(100))

	res, _ := decimalValue.Float64()
	fmt.Println("使用decimal库，精度未丢失", res)
}
