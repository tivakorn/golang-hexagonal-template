package repository

import "errors"

type rewardRepositoryMock struct {
	reward []Reward
}

func NewRewardRepositoryMock() rewardRepositoryMock {

	reward := []Reward{
		{
			RewardID: "5fe1bf3bf8f82045a5b7add8",
			Name:     "ของรางวัล 0001 (normal)",
			Status:   true,
		},
		{
			RewardID: "5fcda6eb30e183001956f7f2",
			Name:     "ของรางวัล 0002 (normal)",
			Status:   false,
		},
	}

	return rewardRepositoryMock{reward: reward}
}

func (repo rewardRepositoryMock) GetAll() ([]Reward, error) {
	return repo.reward, nil
}

func (repo rewardRepositoryMock) GetByID(id string) (*Reward, error) {

	for _, reward := range repo.reward {
		if reward.RewardID == id {
			return &reward, nil
		}
	}

	return nil, errors.New("reward not found")
}
