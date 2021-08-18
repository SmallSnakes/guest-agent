package main

import (
	"log"
	"troila-guest-agent/utils"

	"github.com/gin-gonic/gin"
)

func execOrder(c *gin.Context) {
	action := c.Query("action")
	code := utils.ExecOrder(action)

	if code == 200 {
		c.JSON(200, gin.H{"message": "The cli have been executed"})
	} else if code == 400 {
		c.JSON(400, gin.H{"warning": "The cli exec fail"})
	} else if code == 403 {
		c.JSON(403, gin.H{"error": "The command is not allowed"})
	}
}

//获取状态信息
func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//获取cpu信息
func getCPUInfo(c *gin.Context) {
	info, err := utils.GetCPUInfo()
	if err != nil {
		log.Println("get cpu info error ", err)
	}
	c.JSON(200, gin.H{
		"cpu": info,
	})
}

//获取memory
func getMemoryInfo(c *gin.Context) {
	info, err := utils.GetMemory()
	if err != nil {
		log.Println("get memory info error ", err)
	}
	c.JSON(200, gin.H{
		"memory": info,
	})
}

//获取disk
func getDiskInfo(c *gin.Context) {
	info, err := utils.GetDiskInfo()
	if err != nil {
		log.Println("get disk info error ", err)
	}
	c.JSON(200, gin.H{
		"disk": info,
	})
}

//get network
func getNetInfo(c *gin.Context) {
	info, err := utils.GetNetInfo()
	if err != nil {
		log.Println("get network info error ", err)
	}
	c.JSON(200, gin.H{
		"interface": info,
	})
}


func main() {
	router := gin.Default()

	router.GET("/ping", ping)
	router.POST("/power", execOrder)
	router.GET("/CPUInfo", getCPUInfo)
	router.GET("/MemoryInfo", getMemoryInfo)
	router.GET("/DiskInfo", getDiskInfo)
	router.GET("/NetInfo", getNetInfo)

	router.Run(":1234")
}
