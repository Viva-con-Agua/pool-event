package token

import (
	"context"
	"os"
	"pool-event/dao"
	"testing"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	tester     *vcago.Test
	token      = new(jwt.Token)
	adminToken = &vcapool.AccessToken{
		ID:             "5d914df8-792d-4390-b35a-0a92f6f01b3a",
		FirstName:      "test",
		LastName:       "user",
		FullName:       "test user",
		Roles:          vmod.RoleListCookie{"admin", "member"},
		StandardClaims: jwt.StandardClaims{},
	}
)

func TestMain(m *testing.M) {
	e := vcago.NewServer()
	tester = vcago.NewTest(e)
	ctx := context.Background()
	dao.InitialTestDatabase()
	dao.InitialCollections()
	dao.ArtistCollection.InsertMany(ctx, bson.A{artist1, artist2, artist3})
	dao.OrganizerCollection.InsertMany(ctx, bson.A{organizer1, organizer2, organizer3})
	ret := m.Run()
	dao.Database.Database.Drop(ctx)
	os.Exit(ret)
}
