package dao

import (
	"bytes"
	"context"
	"io"
	"os"
	"pool-event/models"
	"testing"
	"time"

	"github.com/Viva-con-Agua/vcago"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

var out io.Writer = os.Stdout

func TestMain(m *testing.M) {
	ctx := context.Background()
	InitialTestDatabase()
	vcago.Nats.Connect()
	SubscribeUserCreate()
	SubscribeUserUpdate()
	ret := m.Run()
	Database.Database.Drop(ctx)
	os.Exit(ret)
}

func TestInitialCollections(t *testing.T) {
	assert.NoError(t, InitialCollections())

}

func TestSubscribeUserCreate(t *testing.T) {
	buf := &bytes.Buffer{}
	out = buf
	vcago.Nats.Publish("pool-user.user.created", user1)
	time.Sleep(500 * time.Millisecond)
	u := new(models.User)
	ctx := context.Background()
	err := UserCollection.FindOne(ctx, bson.D{{Key: "_id", Value: user1.ID}}, u)
	if assert.NoError(t, err) {
		assert.Equal(t, user1, *u)
	}

}

func TestSubscribeUserUpdate(t *testing.T) {
	buf := &bytes.Buffer{}
	out = buf
	updateUser := user2
	updateUser.ID = user1.ID
	vcago.Nats.Publish("pool-user.user.updated", updateUser)
	time.Sleep(500 * time.Millisecond)
	u := new(models.User)
	ctx := context.Background()
	err := UserCollection.FindOne(ctx, bson.D{{Key: "_id", Value: user1.ID}}, u)
	if assert.NoError(t, err) {
		assert.Equal(t, updateUser, *u)
	}
}
