package dao

import (
	"context"
	"pool-event/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestSubscribeUserCreate(t *testing.T) {
	SubscribeUserCreate(&user1)
	u := new(models.User)
	ctx := context.Background()
	err := UserCollection.FindOne(ctx, bson.D{{Key: "_id", Value: user1.ID}}, u)
	if assert.NoError(t, err) {
		assert.Equal(t, user1, *u)
	}

}

func TestSubscribeUserUpdate(t *testing.T) {
	updateUser := user2
	updateUser.ID = user1.ID
	SubscribeUserUpdate(&updateUser)
	u := new(models.User)
	ctx := context.Background()
	err := UserCollection.FindOne(ctx, bson.D{{Key: "_id", Value: user1.ID}}, u)
	if assert.NoError(t, err) {
		assert.Equal(t, updateUser, *u)
	}
}
