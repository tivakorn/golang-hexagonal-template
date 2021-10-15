package repository

type Reward struct {
	RewardID string `json:"_id"`
	Name     string `json:"name"`
	Status   bool   `json:"status"`
}

type RewardRepository interface {
	GetAll() ([]Reward, error)
	GetByID(string) (*Reward, error)
}
