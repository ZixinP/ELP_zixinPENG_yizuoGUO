package client

import (
	"bufio"
	"fmt"
	fct "fonctions"
	"net"
	"os"
	"encoding/json"
)

type Request struct {
	Image        string `json:"image"`
	IntParameter int    `json:"intParameter"`
}

func image_to_base64strings() string {
	
	// 创建一个 Scanner 对象
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter path of image: ")

	var inputpath string
	// 使用 Scan() 方法读取输入
	if scanner.Scan() {
		// 获取输入的字符串
		inputpath:= scanner.Text()
		fmt.Println("You entered:", inputpath)
	} else {
		// 处理扫描错误
		err := scanner.Err()
		fmt.Println("Error reading input:", err)
		return ""
	}
    
	// 读取图片文件并将其转为 image.Image
	img, err := fct.LoadImage(inputpath)
	if err != nil {
		panic(err)
	}
    
	// 将图片转为 base64 字符串
	base64String := fct.Encode_image(img)
	return base64String
}

func read_parameter() int {
	// 创建一个 Scanner 对象
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter intParameter: ")

	var inputparameter int
	// 使用 Scan() 方法读取输入
	if scanner.Scan() {
		// 获取输入的字符串
		inputparameter:= scanner.Text()
		fmt.Println("You entered:", inputparameter)
	} else {
		// 处理扫描错误
		err := scanner.Err()
		fmt.Println("Error reading input:", err)
		return 0
	}
	return inputparameter
}


func main(){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()
    
	// 将图片转为 base64 字符串
	base64String:=image_to_base64strings()

	// 读取intParameter
	parameter:=read_parameter()

	// 创建一个 Request 对象
	request := Request{
		Image:        base64String,
		IntParameter:  parameter,
	}

	// 将 Request 对象编码为 JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// 发送 JSON 数据
	_, err = conn.Write(jsonData)
	if err != nil {
		fmt.Println("Error sending JSON:", err)
		return
	}

}