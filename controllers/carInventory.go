package controllers

import . "ExamCode/conf"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetCarHandler(c *gin.Context) {
	CarLock.Lock()
	if CarInventoryModel.List.Len() == 0 {
		CarInventoryModel.SellCount++
		CarInventoryModel.TotalCount++
		CarLock.Unlock()
		c.JSON(200, gin.H{
			"car": time.Now().UnixNano(),
		})
	} else {
		carId := CarInventoryModel.List.Front()
		CarInventoryModel.SellCount++
		CarInventoryModel.List.Remove(CarInventoryModel.List.Front())
		CarLock.Unlock()
		c.JSON(200, gin.H{
			"car": carId.Value,
		})
	}
}

func GetRateHandler(c *gin.Context) {
	CarLock.RLock()
	var r = CarInventoryModel.Rate
	CarLock.RUnlock()
	c.JSON(200, gin.H{
		"rate": strconv.FormatFloat(r, 'f', 2, 32),
	})
}

func GetBufferHandler(c *gin.Context) {
	CarLock.RLock()
	var count = CarInventoryModel.Rate * CarInventoryModel.X
	fmt.Println(count)
	CarLock.RUnlock()
	c.JSON(200, gin.H{
		"buffer": int(count),
	})
}
