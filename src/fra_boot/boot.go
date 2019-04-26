package boot

import (
	"fmt"
)

/*
	初始化检测和配置
*/
func Init(path string, version string) {
	scanEnv(path)
	readConfig(path)
	scanProject()
	scanVersion(version)
	scanHostComputer()
	scanCoolq()
}

func scanEnv(conf string) {
	fmt.Println("开始扫描配置文件....")
	//var configMap = readConfig(conf)
	//var coolurl = configMap["coolq_url"]
	//fmt.Println(coolurl)
	fmt.Println("配置文件扫描完毕")
}

/*
  读取配置文件
*/
func readConfig(path string) map[string]string {

	fmt.Println("开始读取配置文件....")
	//var src = path
	//
	//configMap := utils.InitConfig(src)
	//fmt.Println(configMap)

	fmt.Println("读取配置文件完毕")
	return nil
}

/*
  扫描项目文件
*/
func scanProject() {

	fmt.Println("开始扫描项目文件....")
	fmt.Println("扫描项目文件完毕")
}

/*
  检测项目版本
*/
func scanVersion(version string) {
	fmt.Println("开始检测项目版本....")
	fmt.Println("当前项目版本为: " + version)
	fmt.Println("检测项目版本完毕")
}

/*
  检测宿主机环境
*/
func scanHostComputer() {
	fmt.Println("开始检测宿主机环境....")
	fmt.Println("检测宿主机环境完毕")
}

/*
  检测coolq是否正常启动
*/
func scanCoolq() {
	fmt.Println("开始检测coolq是否正常启动....")
	fmt.Println("检测coolq是否正常启动完毕")
}
