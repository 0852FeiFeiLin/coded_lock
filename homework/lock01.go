package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		对称加密解密：(a~z)
			1.定义16进制密钥，然后转为int类型，
			2.设置key范围，通过判断转为10进制的密钥值有没有超过26，(26最大值，再大就进行转换就乱码了)
			3.明文赋值到新切片
			4.然后移位 加密+=key
			5.然后按string转换为字符串存入空字符串

			6.然后将string[] 切片的每一位继续转化为int类型的ASCII码
			5.然后将ASCII移位  解密-=
			6.由于对照的是ASCII码值，所以会有一些符号掺杂，所以进行处理
			7.先把value > 123 或者 value < 97 的赋值为0,说明他们不是a~z
			8.然后创建新切片，接收不为0的元素 append方法
			9.然后切片里的数据就是解密后的数据了，然后通过空字符串去+=接收他，
			10.输出，加密解密成功
				注意：移位只能操作于int类型的ASCII码，就是先将abc...这样的先转换为能进行移位的ASCII码，string(h) == int(104)
					加密和解密用的密钥需要一样，就是加密后移了10位，解密就要前移10位
					41 ~ 90 == A~Z ； 97 ~ 122 == a~z
	*/
	coded_Lock02()

}

//加密解密函数
func coded_Lock02() {
	//密钥
	key := "0x1a" //1A==26，最大key

	//明文
	str := "hello"

	//过渡数组，存放ASCII
	var Testslice = make([]int64, len(str))
	//空字符串,用来存放加密后的明文
	var copystr string = ""

	//密钥转换为int类型()
	h2, _ := strconv.ParseInt(key, 0, 0) //该方法看见0x自动转换为16进制，看见0 自动转换为8进制,其他情况，转为int
	// 参数3:0 --> int 、8-->int8 、16 --->int16
	fmt.Println("key是：", h2)

	if !(h2 > 0 && h2 <= 26) { //把范围控制在1~26字母之间
		fmt.Println("密钥长度有问题！")
		return
	}
	//加密
	for i := 0; i < len(str); i++ {
		//将明文数据的ASCII码值转型放入testslice里面
		Testslice[i] = int64(str[i])
		fmt.Print("原数据：", Testslice[i], "\t")
		//testslice 的每一位进行移位
		Testslice[i] += h2 //这就是加密算法
		fmt.Printf("加密后：%d\n", Testslice[i])

		//然后将testslice的每一位放进空字符串，就是加密后的明文存入空字符串
		copystr += string(Testslice[i]) //注意：这个就是密文，int类型的ASCII码值转化为对应的字符串
	}
	fmt.Println(Testslice)

	fmt.Print("加密后的数据是：", copystr, "\n")

	//解密
	var Test2slice = make([]int64, len(copystr)) //解密过渡切片
	var endstr = ""                              //存放解密后的密文
	for i := 0; i < len(copystr); i++ {
		//将密文数据放进过渡数组
		Test2slice[i] = int64(copystr[i])
		//将密文移位得到最终数据，解密
		Test2slice[i] -= h2 //解密算法
		//最终的数据放入空字符串里面
		endstr += string(Test2slice[i])
	}
	fmt.Println("解密后的明文是：", endstr)
	fmt.Println(Test2slice)

	fmt.Println("进行处理。。。。。。。。。。")

	//把数组元素>123 的赋值为0
	for i := 0; i < len(Test2slice)-1; i++ {
		if Test2slice[i] > 122 || Test2slice[i] < 97 { //97~122
			Test2slice[i] = 0
		}
	}
	fmt.Println(Test2slice)

	//定义新切片接收不为0 的值。
	endAlice := []int64{}
	for i := 0; i < len(Test2slice); i++ {
		if Test2slice[i] != 0 {
			endAlice = append(endAlice, Test2slice[i])
		}
	}
	fmt.Println(endAlice)
	fmt.Printf("解密后数据为：%c\n", endAlice) //转为字符方法
	var overStr string = ""
	for i := 0; i < len(endAlice); i++ {
		overStr += string(endAlice[i]) //字符串添加切片元素，直接拼接成字符串
	}
	fmt.Printf("解密后数据为：%s\n", overStr)

}
