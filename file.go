// 打印目录下面所有文件，并去除.svn .git等隐藏文件
// author: 张良志
// emial:  521401@qq.com
// time:   2016-11-09

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

// 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if filename[0] == '.' {
			// fmt.Println("...")
			return nil
		}
		if strings.Contains(filename, "/.") {
			// fmt.Println("...fff")
			return nil
		}
		if strings.Contains(filename, "file") {
			return nil
		}

		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

func WriteJson() {
	files, _ := WalkDir("./", "")

	// 写入文件
	userFile := "file.json"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	// json格式
	type FileName struct {
		Name string
	}
	type FileNamesSlice struct {
		FileNames []FileName
	}
	var s FileNamesSlice
	for index := 0; index < len(files); index++ {
		s.FileNames = append(s.FileNames, FileName{Name: files[index]})
	}
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}

	fout.WriteString(string(b))
}

func WriteTxt() {
	files, _ := WalkDir("./", "")

	// 写入txt文件
	userFile := "file.txt"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	fout.WriteString("[\n")
	for index := 0; index < len(files); index++ {
		fmt.Println(files[index])

		tmpname := files[index]
		if index+1 == len(files) {
			// 最后一行
			fout.WriteString("\"" + tmpname + "\"\n")
		} else {
			fout.WriteString("\"" + tmpname + "\",\n")
		}

	}
	fout.WriteString("]")
}

func main() {
	WriteTxt()

	// or 写入json文件格式里
	// WriteJson()
}
