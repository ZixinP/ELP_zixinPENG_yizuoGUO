package fonctions

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"net"
	"os"
	"strings"
)


type Request struct {
	Image        string `json:"image"`
	IntParameter int    `json:"intParameter"`
}


// 解码客户端发送的图像数据,返回jpg格式的图像数据和int参数
func Decode_image(conn net.Conn) (image.Image, int) {

	// 读取客户端发送的数据
	buffer := make([]byte, 1024)
	var data []byte
    
	for {
        n, err := conn.Read(buffer)
        if err != nil {
            fmt.Println("Error reading data:", err)
            return nil, 0
        }

        data = append(data, buffer[:n]...)

        if n < len(buffer) {
            break
        }
    }

	// 解码 JSON 数据
	var request Request
	err := json.Unmarshal(data, &request) //intParameter在这里以及保存在request中
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, 0
	}


	// 将 base64 字符串解码为图像数据
	imageData, err := base64.StdEncoding.DecodeString(request.Image)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return nil, 0
	}

	fmt.Println("Image data received from client.")

	// 创建一个 image.Image 对象
	img, _, err := image.Decode(strings.NewReader(string(imageData)))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil, 0
	}

	return img, request.IntParameter
}


func Encode_image(image_jpg image.Image) (string,error){     // 将图像数据编码为 base64 字符串
	var buf bytes.Buffer

	// 将图片编码为 JPG 格式
	err := jpeg.Encode(&buf, image_jpg, nil)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return "", err
	}

	// 将编码后的字节切片转为 base64 编码的字符串
	base64String := base64.StdEncoding.EncodeToString(buf.Bytes())

	return base64String, nil
}
	

func LoadImage(imagePath string) (image.Image, error) {   // 读取图片文件并将其转为 image.Image
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}