package main

import "fmt"

func main() {

	// 数组是同一种数据类型元素的集合。数组的长度（大小）是不可变化的

	// 1、数据定义或声明
	// var 数组变量名 [元素数量]T
	//长度是必须的。
	var a [3]int
	var b [4]int
	fmt.Println(a, b)
	// a = b //不可以这样做，因为此时a和b是不同的类型
	fmt.Println("-----------------------")

	// 2、数组初始化
	//方式1：初始化数组时可以使用初始化列表来设置数组元素的值。
	var testArray1 [3]int                        //数组会初始化为int类型的零值
	var numArray1 = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray1 = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray1)                      //[0 0 0]
	fmt.Println(numArray1)                       //[1 2 0]
	fmt.Println(cityArray1)                      //[北京 上海 深圳]
	// 方式2：编译器根据初始值的个数自行推断数组的长度
	var numArray2 = [...]int{1, 2}
	var cityArray2 = [...]string{"北京", "上海", "深圳"}
	fmt.Println(numArray2)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray2)   //type of numArray2:[2]int
	fmt.Println(cityArray2)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray2) //type of cityArray2:[3]string
	// 方式3：指定索引值的方式来初始化数组
	c := [...]int{1: 1, 3: 5}
	fmt.Println(c)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", c) //type of c:[4]int
	fmt.Println("-----------------------")

	// 3、遍历数组
	var citys = [...]string{"北京", "上海", "深圳"}
	// 方式1
	for i := 0; i < len(citys); i++ {
		fmt.Printf("索引：%v,值：%s\n", i, citys[i])
	}

	// 方式2
	for index, value := range citys {
		fmt.Printf("索引：%v,值：%s\n", index, value)

	}

}
