package handler

import (
	"encoding/json"
	"fmt"
	"golang-hexagonal-template/service"
	"net/http"

	"github.com/gorilla/mux"
)

type rewardHandler struct {
	rewardService service.RewardService
}

func NewRewardHandler(rewardService service.RewardService) rewardHandler {
	return rewardHandler{rewardService: rewardService}
}

func (handle rewardHandler) GetRewardList(write http.ResponseWriter, req *http.Request) {

	rewardList, err := handle.rewardService.GetRewardList()

	if err != nil {

		write.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(write, err)
		return
	}

	write.Header().Set("content-type", "application/json")

	json.NewEncoder(write).Encode(rewardList)
}

func (handle rewardHandler) GetReward(write http.ResponseWriter, req *http.Request) {

	rewardID := mux.Vars(req)["rewardID"]

	reward, err := handle.rewardService.GetReward(rewardID)

	if err != nil {

		write.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(write, err)
		return
	}

	write.Header().Set("content-type", "application/json")

	json.NewEncoder(write).Encode(reward)
}
