package main

import (
	"flag"
	"fmt"
	"github.com/fragrance/src/fra_utils/src"
)

/*
  author : Mr.zheng.
  design_date: 2018-11-25.
  source_by: docker-api.

  src
   |- fra_containers 容器管理
   |- fra_db 数据管理
   |- fra_hardware 硬件管理
   |- fra_images  镜像管理
   |- fra_services 服务管理
   |- fra_stacks  集群管理
   |- fra_users  用户管理
*/

// 声明全局变量
// 项目版本
var FRAGRANCE_PROJECT_VERSION string = "1.0.0"

var DEFAULT_CONFIG_PATH string = "../config/config.conf"

var DEFAULT_URL string = "http://ec.singularitychina.com/cw/api/swagger_doc"

/*
  入口函数
*/
func main() {
	url := flag.String("url", "", "swagger_doc 数据网络地址")
	u := flag.String("u", "", "swagger_doc 数据网络地址")
	flag.Parse()

	var realUrl = DEFAULT_URL
	if *url != "" {
		realUrl = *url
	}

	if *u != "" {
		realUrl = *u
	}

	result, _ := utils.GetSwaggerJson(realUrl)

	//for v, item := range result.Paths {
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

	utils.ExportSwaggerJsonToExcel(*result, "api_demo_v2.1.xlsx")
	//config := flag.String("config","","weather need config")
	//version := flag.String("version","","weather need config")
	//
	//readConfig("")
	//flag.Parse()
	//
	//fmt.Println(*config)
	//fmt.Println(*version)
	//readConfig()
	//scanProject()
	//scanVersion()
	//scanHostComputer()
	//
	////utils.CheckDockerExit()
	//utils.ReadEnvByName("DOCKER_API_VERSION")
	////images.GetImagesList()
}

/*
  读取启动命令行参数
  命令行参数优先级大于配置文件优先级
*/
func readArg() {

}

/*
  读取配置文件
  默认配置文件地址为空字符串
*/
func readConfig(path string) map[string]string {
	var src = DEFAULT_CONFIG_PATH
	if path != "" {
		src = path
	}

	configMap := utils.InitConfig(src)
	fmt.Println(configMap)
	return configMap
}

/*
  扫描项目文件
*/
func scanProject() {

}

/*
  检测项目版本
*/
func scanVersion() {

}

/*
  检测宿主机环境
*/
func scanHostComputer() {

}
