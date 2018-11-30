package utils

import (
	"fmt"
	"os"
)

/*
  用于环境变量读取工具
 */

func ReadEnvByName(name string) string {
	fmt.Println("getting env %s", name)

	// get env value
	var env string
	if env := os.Getenv(name); env == ""{
		fmt.Println("no %s env can be get",name)
		return env
	}
	return env
}
