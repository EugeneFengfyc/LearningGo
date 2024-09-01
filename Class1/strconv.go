package main

import (
	"fmt"
	"strconv"
)

func main() {

	//string 和 int 的转换
	//int to string
	fmt.Println("v1:" + strconv.Itoa(100))

	//string to int
	v2, _ := strconv.Atoi("200")
	fmt.Println(v2)

	/*由于 string 类型可能无法转换为 int 类型，所以这个函数有两个返回值：第一个返回值是转换成的 int 类型的值，第二个返回值是转换是否成功。*/
	v3, err := strconv.Atoi("a")
	fmt.Println(v3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------------------------")
	//Parse 类函数用于将字符串转换为指定的数值类型，其中包含 ParseBool()、ParseFloat()、ParseInt()、ParseUint() 等方法
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 64)
	fmt.Println(b, f, i, u)

	/*
		ParseInt() 和 ParseUint() 有 3 个参数：
		func ParseInt(s string, base int, bitSize int) (i int64, err error)
		func ParseUint(s string, base int, bitSize int) (uint64, error)

		其中：
		s 表示需要转换的字符串
		bitSize 表示转换为多少位的 int 或 uint，有效值为 0、8、16、32、64，当值为 0 的时候，表示转换为 int 或 uint 类型。例如 bitSize=8 表示转换后的值的类型为 int8 或 uint8。
		base 表示以什么进制的方式去解析给定的字符串，有效值为 0 和 2 ~ 36。当 base=0 的时候，表示根据 string 的前缀来判断以什么进制去解析：0x 开头表示以 16 进制的方式去解析，0 开头的以 8 进制方式去解析，其它的以 10 进制方式解析。
	*/

	fmt.Println("----------------------------")
	//将指定数值类型转换为 string 类型，其中包含 FormatBool()、FormatFloat()、FormatInt()、FormatUint() 等方法。
	s1 := strconv.FormatBool(true)
	fmt.Println(s1)

	/*
		func FormatFloat(f float64, fmt byte, prec, bitSize int) string
		bitSize 表示 f 的来源类型（32 指 float32、64 指 float64），会据此进行四舍五入。
		fmt 表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则就用'f'格式）、'G'（指数很大时用'E'格式，否则就用'f'格式）。
		prec 控制精度（排除指数部分）：对 'f'、'e'、'E'，它表示小数点后的数字个数，对'g'、'G'，它控制总的数字个数。如果 prec 为 -1，则代表使用最少但又必需的数字来表示 f。
	*/
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(s2)

	/*
		其中，第二个参数 base 指定将第一个参数转换为多少进制，有效值为 2 ~ 36 之间，当指定的进制位大于 10 的时候，超出 10 的数值以 a~z 字母表示
		，例如 16 进制时，10~15 的数字分别使用 a~f 表示，17 进制时，10-16 的数值分别使用 a~g 表示，例如 FormatInt(-42, 16) 表示将 -42 转换为 16 进制数，转换的结果为 -2a。
	*/
	s3 := strconv.FormatInt(-42, 16)
	fmt.Println(s3)
	s4 := strconv.FormatUint(42, 16)
	fmt.Println(s4)

	fmt.Println("----------------------------")
	//Append 类函数用于将指定类型转换成字符串后追加到一个切片中，其中包含 AppendBool()、AppendFloat()、AppendInt()、AppendUint() 等方法。
	// 声明一个slice
	b10 := []byte("a1")

	// 将转换为10进制的string，追加到slice中
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("a2")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))

}
