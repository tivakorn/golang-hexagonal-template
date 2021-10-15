package main

import (
	"fmt"
	"golang-hexagonal-template/handler"
	"golang-hexagonal-template/lib"
	"golang-hexagonal-template/repository"
	"golang-hexagonal-template/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {

	lib.InitConfig()
	lib.InitTimeZone()

	// mongodb := lib.MongodbConnection()

	// rewardRepositoryDB := repository.NewRewardRepositoryDB(mongodb)

	rewardRepositoryMock := repository.NewRewardRepositoryMock()

	// _ = rewardRepositoryMock

	rewardService := service.NewRewardService(rewardRepositoryMock)
	rewardHandler := handler.NewRewardHandler(rewardService)

	router := mux.NewRouter()

	router.HandleFunc("/reward", rewardHandler.GetRewardList).Methods(http.MethodGet)
	router.HandleFunc("/reward/{rewardID}", rewardHandler.GetReward).Methods(http.MethodGet)

	log.Printf("Service started at port %v", viper.GetInt("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}
