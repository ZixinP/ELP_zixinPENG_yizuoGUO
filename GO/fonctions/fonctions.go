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


// decode the image and intParameter from the client
func Decode_image(conn net.Conn) (image.Image, int) {

	// create a buffer to store the data received from the client
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

	// decode the JSON data to a Request object
	var request Request
	err := json.Unmarshal(data, &request)  
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, 0
	}


	// decode the base64 strings to image data
	imageData, err := base64.StdEncoding.DecodeString(request.Image)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return nil, 0
	}

	fmt.Println("Image data received from client.")

	// create an image.Image from the image data
	img, _, err := image.Decode(strings.NewReader(string(imageData)))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil, 0
	}

	return img, request.IntParameter
}


func Encode_image(image_obj image.Image) (string,error){     // encode image.Image to base64 string
	var buf bytes.Buffer

	// encode the image to JPEG format
	err := jpeg.Encode(&buf, image_obj, nil)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return "", err
	}

	// convert the image from bytes to base64
	base64String := base64.StdEncoding.EncodeToString(buf.Bytes())

	return base64String, nil
}
	

func LoadImage(imagePath string) (image.Image, error) {   // read image from file and return image.Image
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