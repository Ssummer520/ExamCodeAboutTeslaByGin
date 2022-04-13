package main

import (
	. "ExamCode/conf"
	. "ExamCode/controllers"
	_ "ExamCode/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
)

// @title ExamCodeAboutTeslaByGin
// @version 1.0
// @description ExamCodeAboutTeslaByGin
// @termsOfService http://swagger.io/terms/
// @contact.name 1362
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @host 124.220.12.138:8888
// @BasePath

func main() {
	//envPort:= os.Getenv("ASPNETCORE_PORT")
	CarInventoryModel.X = 2
	job()
	r := gin.New()
	r.GET("/age", GetAgeHandler)
	r.GET("/car", GetCarHandler)
	r.POST("/age", AddAgeHandler)
	r.GET("/rate", GetRateHandler)
	r.GET("/buffer", GetBufferHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(cors.Default())
	//r.Run(":"+envPort)
	r.Run()
}

func job() {

	go func() {
		for {
			checkCarInventoryData()
			//每秒钟执行一次 检查库存
			time.Sleep(time.Second)
			//fmt.Println("checkCarModel", CarInventoryModel.X, CarInventoryModel.List, CarInventoryModel.Rate, CarInventoryModel.SellCount, CarInventoryModel.TotalCount)
		}
	}()
	//启动协程来单独根据要求更新X 和R
	go func() {
		for {
			//fmt.Println("updateRate", CarInventoryModel)
			time.Sleep(time.Second * 60)
			//每分钟更新一次 R 窗口滑动
			CarLock.Lock()
			CarInventoryModel.X = (CarInventoryModel.X*60 + 1) / 60
			CarInventoryModel.Rate = float64(CarInventoryModel.SellCount) / CarInventoryModel.X
			CarLock.Unlock()
		}
	}()
}

func checkCarInventoryData() {
	CarLock.Lock()
	var inventoryQuantity = CarInventoryModel.Rate * CarInventoryModel.X
	var currentCount = CarInventoryModel.TotalCount - CarInventoryModel.SellCount
	if (inventoryQuantity == 0 && currentCount == 0) || CarInventoryModel.TotalCount-CarInventoryModel.SellCount < int(inventoryQuantity) {
		CarInventoryModel.TotalCount++
		CarInventoryModel.List.PushBack(time.Now().UnixNano())
	}
	CarLock.Unlock()
}
