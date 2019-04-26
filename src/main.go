package main

import (
	"github.com/fragrance/src/fra_boot"
	"os/exec"
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
   |- fra_message  消息管理
   |- fra_services 服务管理
   |- fra_stacks  集群管理
   |- fra_users  用户管理
*/

// 声明全局变量
// 项目版本
var FRAGRANCE_PROJECT_VERSION string = "1.0.0"

var DEFAULT_CONFIG_PATH string = "~/config/config.conf"

/*
  入口函数
*/
func main() {
	// 初始化项目 加载配置 扫描代码 启动单测
	boot.Init(DEFAULT_CONFIG_PATH, FRAGRANCE_PROJECT_VERSION)

	exec.Command("go", "test")
	//url := flag.String("url", "", "swagger_doc 数据网络地址")
	//flag.Parse()
	//
	//var realUrl = DEFAULT_URL
	//if *url != "" {
	//	realUrl = *url
	//}
	//
	//result, _ := utils.GetSwaggerJson(realUrl)
	//
	//utils.ExportSwaggerJsonToExcel(*result, "api_demo_v2.0.xlsx")
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

	// 扫描项目是否正常启动，项目启动后首先单元测试
	//ScanEnvSafe(DEFAULT_CONFIG_PATH)

	//fmt.Println(a)
	//a = 3
	//fmt.Println(a)
	//
	//resp, err := http.Get("http://127.0.0.1:5700/send_private_msg?user_id=1057885002&message=xxxxxxxxxx")
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(string(body))

	//for i := 0; i < 10; i++ {
	//	go message.SendPrivateMessage("1057885002",strconv.Itoa(i))
	//}
	//
	//runtime.Gosched()

}
