package main

import (
	"github.com/fragrance/src/fra_images/src"
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

/*
  入口函数
 */
func main() {

	readConfig()
	scanProject()
	scanVersion()
	scanHostComputer()

	//utils.CheckDockerExit()
	utils.ReadEnvByName("go")
	images.GetImagesList()
}

/*
  读取配置文件
 */
func readConfig(){

}

/*
  扫描项目文件
 */
func scanProject(){

}

/*
  检测项目版本
 */
func scanVersion(){

}

/*
  检测宿主机环境
 */
func scanHostComputer(){

}