package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func GetWorkPath() string {
	if file, err := exec.LookPath(os.Args[0]); err == nil {
		return filepath.Dir(file) + "/"
	}
	return "./"
}
func StringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("agent 类型转换失败, 请检查配置文件中 agentid 配置是否为纯数字(%v)", err)
		return 0
	}
	return n
}
func StringToInt64(s string) int64 {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("agent 类型转换失败, 请检查配置文件中 agentid 配置是否为纯数字(%v)", err)
		return 0
	}
	return int64(n)
}
func StringToFloat64(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("agent 类型转换失败, 请检查配置文件中 agentid 配置是否为纯数字(%v)", err)
		return 0
	}
	return n
}
func FloatTo2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
