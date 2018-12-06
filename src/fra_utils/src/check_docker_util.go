package utils

import (
	"fmt"
	"os/exec"
)

func CheckDockerExit() bool{
	fmt.Println("checking docker exists.......")
	cmd := exec.Command("docker","version") ///查看当前目录下文件
	out, err := cmd.Output()
	if err != nil {
		panic(err)
		return false
	}
	fmt.Println(string(out))

	fmt.Println("docker exist")
	return true
}


