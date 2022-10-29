package main

import (
	"fmt"
	"io"
	"os"
)

func WriteFile(filename string, string2 string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("打开文件出错")
	}
	file.Write([]byte(string2))
	file.Close()
}

func ReadFile(filename string) []byte {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	//Read方法不是很好用，它是文件内部字符串的指针，我们想获取文件内容，只能通过遍历指针获得
	if err != nil {
		fmt.Println("打开文件出错")
	}
	//给data分配内存，不然无法赋值成功
	data := make([]byte, 1145)

	for {
		n, err := file.Read(data)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
		}
		//说明读取结束
		if n == 0 {
			break
		}

	}
	file.Close()
	return data
}
func CreateFile(filename string) {
	os.Create(filename)
}

func main() {
	filename := "./HomeWork/plan.txt"
	CreateFile(filename)
	str := "I’m not afraid of difficulties and insist on learning programming"
	WriteFile(filename, str)
	fmt.Println(string(ReadFile(filename)))
}
