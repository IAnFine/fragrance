package utils

import (
	"fmt"
	"os"
)

/*
  用于环境变量读取工具
 */

func ReadEnvByName(name string) string {
	fmt.Printf("getting env %s \n", name)

	// get env value
	var env string
	if env := os.Getenv(name); env != ""{
		fmt.Printf("the %s is %s \n",name,env)
		return env
	}
	return env
}
