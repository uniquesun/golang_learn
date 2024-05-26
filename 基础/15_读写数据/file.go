package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// 目录 curd
func createDir(dirPath string) error {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	}
	fmt.Printf("Directory %s created successfully\n", dirPath)
	return nil
}

func updateDir(oldDirPath, newDirPath string) error {
	err := os.Rename(oldDirPath, newDirPath)
	if err != nil {
		return err
	}
	fmt.Printf("Directory %s renamed to %s successfully\n", oldDirPath, newDirPath)
	return nil
}

func deleteDir(dirPath string) error {
	err := os.RemoveAll(dirPath)
	if err != nil {
		return err
	}
	fmt.Printf("Directory %s deleted successfully\n", dirPath)
	return nil
}

func readDir(dirPath string) ([]string, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	files, err := dir.ReadDir(0)
	if err != nil {
		return nil, err
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, filepath.Join(dirPath, file.Name()))
	}
	return fileNames, nil
}

func OSDir() {
	// 创建目录
	var err error
	//err = createDir("./olaf/1")
	//if err != nil {
	//	fmt.Println("Error creating directory:", err)
	//}
	//// 改名
	//err = updateDir("./olaf/1", "./olaf/2")
	//if err != nil {
	//	fmt.Println("Error renaming directory:", err)
	//}
	// 删除目录
	//err = deleteDir("./olaf/1")
	//if err != nil {
	//	fmt.Println("Error delete directory:", err)
	//}

	// 读取目录
	var dir []string
	dir, err = readDir("./olaf/1")
	if err != nil {
		return
	}
	fmt.Println(dir)

}

// 文件 curd

func OSFile() {
	// 创建文件
	var data []byte
	err := os.WriteFile("./olaf/1/olaf1.txt", data, 0755)
	if err != nil {
		return
	}

	// 改名
	//os.Rename("./olaf/1/olaf1.txt", "./olaf/1/olaf2.txt")

	// 修改权限
	//os.Chmod("./olaf/1/olaf1.txt", 0777)

	// 删除文件
	//os.Remove("./olaf/1/olaf2.txt")
}

// 文件内容 curd

func OSFileContentWrite() error {
	// 覆盖写
	data := "Hello, World!"
	err := os.WriteFile("./olaf/1/olaf1.txt", []byte(data), 0777)
	if err != nil {
		return err
	}

	// 追加写
	file, err1 := os.OpenFile("./olaf/1/olaf1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err1 != nil {
		return err1
	}
	defer file.Close()

	_, err = file.WriteString("21313123")
	if err != nil {
		return err
	}
	fmt.Printf("Content appended to file %s\n", "./olaf/1/olaf1.txt")
	return nil

}

func OSFileContentRead() {
	// 一次性读取全部
	//file, _ := os.ReadFile("./olaf/1/olaf1.txt")
	//fmt.Println(string(file))

	// 一次性读取全部
	//file, err := os.Open("./olaf/1/olaf1.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//
	//data, err1 := io.ReadAll(file)
	//if err1 != nil {
	//	log.Fatal(err1)
	//}
	//fmt.Println(string(data))

	// 按行读取内容
	//file, _ := os.Open("./olaf/1/olaf1.txt")
	//defer file.Close()
	//scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}

	// buffer读
	file, _ := os.Open("./olaf/1/olaf1.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Print(string(buffer[:n]))
	}
}

func main() {
	OSFileContentRead()
}
