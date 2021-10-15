package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type rewardRepositoryDB struct {
	db mongo.Database
}

func NewRewardRepositoryDB(db mongo.Database) rewardRepositoryDB {
	return rewardRepositoryDB{db: db}
}

func (repo rewardRepositoryDB) GetAll() ([]Reward, error) {

	type project struct {
		name   int `bson:"name"`
		status int `bson:"status"`
	}

	var rewardList []Reward

	query := bson.M{}

	projection := project{
		name:   1,
		status: 1,
	}

	options := options.Find()

	options.SetProjection(projection)

	options.SetLimit(2)

	cursor, err := repo.db.Collection("pantip_point_reward").Find(context.Background(), query, options)

	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &rewardList)

	if err != nil {
		return nil, err
	}

	return rewardList, nil
}

func (repo rewardRepositoryDB) GetByID(id string) (*Reward, error) {

	type project struct {
		name   int `bson:"name"`
		status int `bson:"status"`
	}

	var reward *Reward

	rewardID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	query := bson.M{"_id": rewardID}

	options := options.FindOne()

	projection := project{
		name:   1,
		status: 1,
	}

	options.SetProjection(projection)

	err = repo.db.Collection("pantip_point_reward").FindOne(context.Background(), query, options).Decode(&reward)

	if err != nil {
		return nil, err
	}

	return reward, nil
}
