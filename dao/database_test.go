package dao

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/Viva-con-Agua/vcago"
)

var out io.Writer = os.Stdout

func TestMain(m *testing.M) {
	ctx := context.Background()
	InitialTestDatabase()
	vcago.Settings.Load()
	ret := m.Run()
	Database.Database.Drop(ctx)
	os.Exit(ret)
}

func TestInitialCollections(t *testing.T) {
	InitialCollections()
}
