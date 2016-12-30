// 查看苹果服务器ip到ping值
// author: 张良志
// emial:  521401@qq.com
// time:   2016-11-24

package main

import (
	"fmt"
	"os"
)

func chaeckApplePing() {
	// 写入txt文件
	userFile := "aip.txt"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}


	for index := 1; index <= 2000; index++ {
		stri := fmt.Sprintf("%d", index) 
		str:="ping a" + stri + ".phobos.apple.com -c 4\n"
		fmt.Println(str)
		fout.WriteString(str)

	}

}

func main() {
	chaeckApplePing()

}
