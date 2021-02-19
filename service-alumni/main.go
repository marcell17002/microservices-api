package main

import (
	"log"
	"service-alumni/handler"
	"service-alumni/informations"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/alumni_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	informationRepository := informations.NewRepository(db)
	informationService := informations.NewService(informationRepository)
	informationHandler := handler.NewInformationsHandler(informationService)

	router := gin.Default()

	router.POST("/alumni-information", informationHandler.CreateInformations)
	router.GET("/alumni-informations", informationHandler.GetInformations)
	router.GET("/alumni-information/:id", informationHandler.GetInformationsDetail)
	router.PUT("/alumni-information/:id", informationHandler.UpdateInformations)
	router.DELETE("/alumni-information/:id", informationHandler.DeleteInformation)

	router.Run(":9000")
}
