package admin

import (
	"context"
	"os"
	"pool-event/dao"
	"testing"

	"github.com/Viva-con-Agua/vcago"
)

var (
	tester *vcago.Test
)

func TestMain(m *testing.M) {
	e := vcago.NewServer()
	tester = vcago.NewTest(e)
	ctx := context.Background()
	dao.InitialTestDatabase()
	dao.InitialCollections()
	ret := m.Run()
	dao.Database.Database.Drop(ctx)
	os.Exit(ret)
}
