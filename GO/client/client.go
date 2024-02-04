package client

import (
	"bufio"
	"fmt"
	fct "../fonctions"
	"net"
	"os"
	"encoding/json"
	"encoding/base64"
	"strconv"
)

type Request struct {
	Image        string `json:"image"`
	IntParameter int    `json:"intParameter"`
}

func image_to_base64strings() string {    // 读取图片文件并将其转为 base64 字符串
	
	// 创建一个 Scanner 对象
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter path of image: ")

	var inputpath string
	// 使用 Scan() 方法读取输入
	if scanner.Scan() {
		// 获取输入的字符串
		inputpath = scanner.Text()
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
	base64String, err := fct.Encode_image(img)
	if err != nil {
		panic(err)
	}
	return base64String
}

func read_parameter() int {       // 读取intParameter
	// 创建一个 Scanner 对象
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter intParameter: ")

	// 使用 Scan() 方法读取输入
	if scanner.Scan() {
		// 获取输入的字符串
		input_str := scanner.Text()
		fmt.Println("You entered:", input_str)

		// 将输入的字符串转为整数
		input_int, err := strconv.Atoi(input_str)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			return 0
		}
		return input_int
	} else {
		// 处理扫描错误
		err := scanner.Err()
		fmt.Println("Error reading input:", err)
		return 0
	}
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

	// 读取服务器返回的数据
	buffer := make([]byte, 1024)
	image_back, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// 解码 base64 字符串为图像数据
	imageData, err := base64.StdEncoding.DecodeString(string(buffer[:image_back]))
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return
	}
	
	// 将图像数据保存为文件
	err = os.WriteFile("image_back.jpg", imageData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}