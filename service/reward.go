package service

type RewardResponse struct {
	RewardID string `json:"_id"`
	Name     string `json:"name"`
}

type RewardService interface {
	GetRewardList() ([]RewardResponse, error)
	GetReward(string) (*RewardResponse, error)
}
