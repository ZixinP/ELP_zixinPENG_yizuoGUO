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

func image_to_base64strings() string {    // read image and convert it to base64 strings
	
	// create a Scanner object
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter path of image: ")

	var inputpath string
	// use the Scan() method to read input
	if scanner.Scan() {
		// 获取输入的字符串
		inputpath = scanner.Text()
		fmt.Println("You entered:", inputpath)
	} else {
		err := scanner.Err()
		fmt.Println("Error reading input:", err)
		return ""
	}
    
	// convert the image to image.Image type
	img, err := fct.LoadImage(inputpath)
	if err != nil {
		panic(err)
	}
    
	// convert the image to base64 strings
	base64String, err := fct.Encode_image(img)
	if err != nil {
		panic(err)
	}
	return base64String
}


func read_parameter() int {       // read the parameter from the user
	// create a Scanner object
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter intParameter: ")

	// use the Scan() method to read input
	if scanner.Scan() {
		// obtain the input string
		input_str := scanner.Text()
		fmt.Println("You entered:", input_str)

		// convert the input string to an integer
		input_int, err := strconv.Atoi(input_str)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			return 0
		}
		return input_int
	} else {
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
    
	base64String:=image_to_base64strings()

	parameter:=read_parameter()

	// create a Request object
	request := Request{
		Image:        base64String,
		IntParameter:  parameter,
	}

	// encode the Request object as JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// send the JSON data to the server
	_, err = conn.Write(jsonData)
	if err != nil {
		fmt.Println("Error sending JSON:", err)
		return
	}

	// read the server's response
	buffer := make([]byte, 1024)
	image_back, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// decode the base64 strings to image data
	imageData, err := base64.StdEncoding.DecodeString(string(buffer[:image_back]))
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return
	}
	
	// save the image data to a file
	err = os.WriteFile("image_back.jpg", imageData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}