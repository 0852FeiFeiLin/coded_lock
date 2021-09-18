package main

import "fmt"

func main() {
	/*
	对称加密解密：
		1.先赋值到新切片
		2.然后移位 加密+=
		3.然后按string转换为字符串存入空字符串

		4.然后将string[] 切片的每一位继续转化为int类型的ASCII码
		5.然后将ASCII移位  解密-=
		6.将解密后的数据string() 放入空字符串拼接起来，输出
			注意：移位只能操作于int类型的ASCII码，就是先将abc...这样的先转换为能进行移位的ASCII码，string(h) == int(104)
				加密和解密用的密钥需要一样，就是加密后移了10位，解密就要前移10位
	*/
	coded_Lock01()

}
//加密解密函数
func coded_Lock01(){
	//密钥
	key := 10
	//明文
	str := "hello"
	//过渡数组，存放ASCII
	var Testslice = make([]int ,len(str))
	//var Testslice  []int  //不能这样定义切片，因为这样slice没有具体下标，然后len() 就是为0
	//空字符串,用来存放加密后的明文
	var copystr string = ""

	//加密
	for i:= 0;i< len(str);i++{
		//将明文数据的ASCII码值放入testslice里面
		Testslice[i] = int(str[i])
		//testslice 的每一位进行移位
		Testslice[i] += key   //这就是加密算法
		//然后将testslice的每一位放进空字符串，就是加密后的明文存入空字符串
		copystr += string(Testslice[i])  //注意：这个就是密文，int类型的ASCII码值转化为对应的字符串
	}
	fmt.Print("加密后的明文是：",copystr,"\n")

	//解密
	var Test2slice = make([]int ,len(copystr)) //解密过渡切片
	var endstr = ""   //存放解密后的密文
	for  i := 0 ;i<len(copystr);i++{
		//将密文数据放进过渡数组
		Test2slice[i] = int(copystr[i])
		//将密文移位得到最终数据，解密
		Test2slice[i] -= key  //解密算法
		//最终的数据放入空字符串里面
		endstr += string(Test2slice[i])
	}
	fmt.Println("解密后的明文是：",endstr)
	fmt.Println("此加密解密应用的是简单的位移法")

}