// 删除文件夹中含output的jpg文件
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func deleteFile(dir string) {
	// Replace with the actual folder path
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), "output") && strings.HasSuffix(file.Name(), ".jpg") {
			err := os.Remove(filepath.Join(dir, file.Name()))
			if err != nil {
				fmt.Println("Error deleting file:", err)
			} else {
				fmt.Println("Deleted file:", file.Name())
			}
		}
	}
}

func main() {
	deleteFile("F:/ELP_zixinPENG_yizuoGUO/")
}
