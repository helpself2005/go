package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string, destPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth+PthSep+fi.Name(), destPth)
		} else {
			ok := strings.HasSuffix(fi.Name(), ".mp4")
			if ok {
				fileName := fi.Name()
				srcFile := dirPth + PthSep + fileName
				nameRune := []rune(fileName)
				fileNameRune := string(nameRune[4:])
				destFile := destPth + PthSep + fileNameRune
				if !checkFileIsExist(destFile) {
					CopyFile(destFile, srcFile)
				}
				files = append(files, srcFile)
			}
		}
	}

	return files, dirs, nil
}

//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".mp4")
			if ok {
				files = append(files, fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

//判断文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CopyFile(des, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

func main() {
	/*
		xfiles, dirs, _ := GetFilesAndDirs("D:\\S2-备课资料-20180605", "D:\\S2MP4")

		fmt.Printf("获取的文件个数为[%d]\n", len(xfiles))

		for _, file := range xfiles {
			fmt.Printf("获取的文件为[%s]\n", file)
		}

		fmt.Printf("获取的文件夹个数为[%d]\n", len(dirs))
	*/

	xfiles, _ := GetAllFiles("D:\\儿歌-需要上传11.10")

	fmt.Printf("获取的文件个数为[%d]\n", len(xfiles))
	for _, file := range xfiles {
		fmt.Println(file)
	}

}
