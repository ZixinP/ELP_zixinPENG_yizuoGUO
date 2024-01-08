package Gray

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"sync"
)

func main() {
	// 打开 JPEG 图片文件
	file, err := os.Open("F:/ELP_zixinPENG_yizuoGUO/Gray/input.jpg")
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

	// 创建等待组
	var wg sync.WaitGroup

	// 遍历每个像素，将彩色图像转换为灰度图像
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			// 增加等待组计数
			wg.Add(1)

			// 使用 goroutine 并发处理每个像素
			go func(x, y int) {
				defer wg.Done()

				// 获取原始颜色
				originalColor := img.At(x, y)
				// 计算灰度值
				grayValue := color.GrayModel.Convert(originalColor).(color.Gray).Y
				// 设置灰度图像像素
				grayImg.SetGray(x, y, color.Gray{grayValue})
			}(x, y)
		}
	}

	// 等待所有 goroutine 完成
	wg.Wait()

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
