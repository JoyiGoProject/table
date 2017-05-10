package utils

import (
	"fmt"
	"os"
)

//将读取的数据写入到文件中(文件不存在则自动创建)
func WriteLog(data string) {
	f, err := os.OpenFile("D:/log.log", os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// data := strings.Join([]string{data}, "")
		f.Write([]byte(data + "\r\n" + "日志记录成功" + "\r\n"))
		// f.WriteString(data)
	}
}
