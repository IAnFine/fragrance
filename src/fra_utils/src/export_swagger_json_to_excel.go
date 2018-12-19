package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
	"strconv"
	"time"
)

// 遍历json数据
// copy模板文件 根据json中的tag来进行文件目录分类
// 根据模板规则写入数据
// 返回结果
func ExportSwaggerJsonToExcel(json SwaggerJson, src string) {
	//for v, item := range json.Paths {
	//	fmt.Print("url:  ")
	//	fmt.Println(v)
	//
	//	if item.Get.OperationId != "" {
	//		fmt.Println("this get request")
	//		fmt.Print("Consumes:  ")
	//		fmt.Println(item.Get.Consumes)
	//		fmt.Print("Description:  ")
	//		fmt.Println(item.Get.Description)
	//		fmt.Print("OperationId:  ")
	//		fmt.Println(item.Get.OperationId)
	//		fmt.Println("Parameters:  ")
	//		for _, param := range item.Get.Parameters {
	//			fmt.Print("    Description:  ")
	//			fmt.Println(param.Description)
	//			fmt.Print("    Format:  ")
	//			fmt.Println(param.Format)
	//			fmt.Print("    In:  ")
	//			fmt.Println(param.In)
	//			fmt.Print("    Required:  ")
	//			fmt.Println(param.Required)
	//			fmt.Print("    Type:  ")
	//			fmt.Println(param.Type)
	//		}
	//
	//	}
	//	if item.Post.OperationId != "" {
	//		fmt.Println("this post request")
	//
	//		fmt.Print("Consumes:  ")
	//		fmt.Println(item.Post.Consumes)
	//		fmt.Print("Description:  ")
	//		fmt.Println(item.Post.Description)
	//		fmt.Print("OperationId:  ")
	//		fmt.Println(item.Post.OperationId)
	//		fmt.Println("Parameters:  ")
	//		for _, param := range item.Post.Parameters {
	//			fmt.Print("    Description:  ")
	//			fmt.Println(param.Description)
	//			fmt.Print("    Format:  ")
	//			fmt.Println(param.Format)
	//			fmt.Print("    In:  ")
	//			fmt.Println(param.In)
	//			fmt.Print("    Required:  ")
	//			fmt.Println(param.Required)
	//			fmt.Print("    Type:  ")
	//			fmt.Println(param.Type)
	//		}
	//	}
	//}
	for v, item := range json.Paths {
		// v 为接口地址
		// item 为对应key 的值 分为get post 或者还有put delete
		fmt.Println(v)
		// 复制模板到指定接口文档目录下
		var dst string

		if len(item.Get.Tags) != 0 {

			// 组合api文档地址
			dst = "./api/" + item.Get.Tags[0] + "/"

			// copy 模板到指定地址
			_, err := copy(src, dst, item.Get.OperationId+".xlsx")

			if err != nil {
				fmt.Println(err)
			}

			err = writeDataToExcel(item.Get, dst+item.Get.OperationId+".xlsx", "GET", v)
			if err != nil {
				fmt.Println(err)
			}
		}

		if len(item.Post.Tags) != 0 {
			// 组合api文档地址
			dst = "./api/" + item.Post.Tags[0] + "/"

			// copy 模板到指定地址
			_, err := copy(src, dst, item.Post.OperationId+".xlsx")

			if err != nil {
				fmt.Println(err)
			}

			err = writeDataToExcel(item.Post, dst+item.Post.OperationId+".xlsx", "POST", v)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// 将swagger接口数据写入excel中
func writeDataToExcel(info SwaggerPathInfo, src, requestType, url string) error {
	xlsx, err := excelize.OpenFile(src)

	if err != nil {
		fmt.Println(err)
		return err
	}

	for v, sheet := range xlsx.Sheet {
		fmt.Println(v)
		fmt.Println(sheet)
	}

	// 写入更新记录
	xlsx.SetCellValue("更新记录", "B3", "1.0.0")
	xlsx.SetCellValue("更新记录", "C3", time.Now().Format("2006-01-02"))
	xlsx.SetCellValue("更新记录", "D3", "自动化")
	xlsx.SetCellValue("更新记录", "E3", "swagger读取工具创建")

	// 写入API定义sheet
	// 写入第一块

	// 写入模块数据
	xlsx.SetCellValue("API定义", "F2", info.Tags[0])

	// 写入模块-接口数据
	xlsx.SetCellValue("API定义", "Y2", info.Description+info.Summary)

	// 写入模块-接口id数据
	xlsx.SetCellValue("API定义", "Y3", info.OperationId)

	// 写入模块数据
	xlsx.SetCellValue("API定义", "AH2", "")

	// 写入模块数据
	xlsx.SetCellValue("API定义", "F3", info.Tags[0])

	// 写入模块数据
	xlsx.SetCellValue("API定义", "AH3", time.Now().Format("2006-01-02"))

	// 写入第二块
	xlsx.SetCellValue("API定义", "G6", info.Description+info.Summary)
	xlsx.SetCellValue("API定义", "G7", requestType)
	xlsx.SetCellValue("API定义", "L7", url)

	// 写入第三块
	var start_line = 11
	for index, item := range info.Parameters {
		var col = start_line + index
		if index > 3 {
			xlsx.InsertRow("API定义", col)
			//xlsx.GetCellStyle()
		}

		xlsx.SetCellValue("API定义", "B"+strconv.Itoa(col), item.Name)
		xlsx.SetCellValue("API定义", "G"+strconv.Itoa(col), item.Description)
		if item.Name == "id" {
			xlsx.SetCellValue("API定义", "G"+strconv.Itoa(col), item.Name)
		}
		xlsx.SetCellValue("API定义", "M"+strconv.Itoa(col), item.Type)
		if item.Required {
			xlsx.SetCellValue("API定义", "R"+strconv.Itoa(col), "●")
		}

	}

	errSave := xlsx.Save()
	if errSave != nil {
		fmt.Println(errSave)
	}

	return nil
}

// 复制目标文件到目的地
func copy(src, dst, fileName string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	// 判断目的地址存不存在
	_, errDst := os.Stat(dst)
	if os.IsNotExist(errDst) {
		mkErr := os.Mkdir(dst, os.ModePerm)

		if mkErr != nil {
			fmt.Printf("mkdir failed![%v]\n", mkErr)
		}
	}

	destination, err := os.Create(dst + fileName)

	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
