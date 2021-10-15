package service

import (
	"errors"
	"golang-hexagonal-template/repository"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type rewardService struct {
	rewardRepo repository.RewardRepository
}

func NewRewardService(rewardRepo repository.RewardRepository) rewardService {
	return rewardService{rewardRepo: rewardRepo}
}

func (service rewardService) GetRewardList() ([]RewardResponse, error) {

	rewardList, err := service.rewardRepo.GetAll()

	if err != nil {

		log.Println(err)

		return nil, err
	}

	rewardListResponse := []RewardResponse{}

	for _, reward := range rewardList {

		rewardResponse := RewardResponse{
			RewardID: reward.RewardID,
			Name:     reward.Name,
		}

		rewardListResponse = append(rewardListResponse, rewardResponse)
	}

	return rewardListResponse, nil
}

func (service rewardService) GetReward(id string) (*RewardResponse, error) {

	reward, err := service.rewardRepo.GetByID(id)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, errors.New("reward not found")
		}

		log.Println(err)
		return nil, err
	}

	rewardResponse := RewardResponse{
		RewardID: reward.RewardID,
		Name:     reward.Name,
	}

	return &rewardResponse, nil
}
