package Gray

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	// 打开 JPEG 图片文件
	file, err := os.Open("input.jpg")
	if err != nil {
		fmt.Println("Error opening image:", err)
		return
	}
	defer file.Close()

	// 解码 JPEG 图片
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// 创建一个灰度图
	grayImg := image.NewGray(img.Bounds())

	// 遍历每个像素，将彩色图像转换为灰度图像
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			// 获取原始颜色
			originalColor := img.At(x, y)
			// 计算灰度值
			grayValue := color.GrayModel.Convert(originalColor).(color.Gray).Y
			// 设置灰度图像像素
			grayImg.SetGray(x, y, color.Gray{grayValue})
		}
	}

	// 创建输出文件
	outputFile, err := os.Create("output_gray.jpg")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// 编码并保存灰度图像为 JPEG 格式
	jpeg.Encode(outputFile, grayImg, nil)

	fmt.Println("Conversion complete. Output saved as output_gray.jpg.")
}
