package dao

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pool-event/models"
	"testing"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/stretchr/testify/assert"
)

var (
	user1 = models.User{
		ID:          "10000000-0000-0000-0000-000000000001",
		Email:       "user1@test.org",
		FirstName:   "test",
		LastName:    "user1",
		FullName:    "test user1",
		DisplayName: "test user1",
		Profile: models.Profile{
			ID:        "10000000-0000-0000-0000-000000000001",
			Gender:    "diverse",
			Phone:     "0177 77777777",
			Birthdate: 1659521082,
			UserID:    "10000000-0000-0000-0000-000000000001",
			Modified: vmod.Modified{
				Updated: 1659521082,
				Created: 1659521082,
			},
		},
	}
	user2 = models.User{
		ID:          "10000000-0000-0000-0000-000000000002",
		Email:       "user2@test.org",
		FirstName:   "test",
		LastName:    "user2",
		FullName:    "test user2",
		DisplayName: "test user2",
		Profile: models.Profile{
			ID:        "10000000-0000-0000-0000-000000000002",
			Gender:    "diverse",
			Phone:     "0177 77777777",
			Birthdate: 1659521082,
			UserID:    "10000000-0000-0000-0000-000000000002",
			Modified: vmod.Modified{
				Updated: 1659521082,
				Created: 1659521082,
			},
		},
	}
	user3 = models.User{
		ID:          "10000000-0000-0000-0000-000000000003",
		Email:       "user3@test.org",
		FirstName:   "test",
		LastName:    "user3",
		FullName:    "test user3",
		DisplayName: "test user3",
		Profile: models.Profile{
			ID:        "10000000-0000-0000-0000-000000000003",
			Gender:    "diverse",
			Phone:     "0177 77777777",
			Birthdate: 1659521082,
			UserID:    "10000000-0000-0000-0000-000000000003",
			Modified: vmod.Modified{
				Updated: 1659521082,
				Created: 1659521082,
			},
		},
	}
	userList = []models.User{user1, user2, user3}

	userResponse = vcago.NewCreated("user", userList)
)

func TestUserGetRequest(t *testing.T) {
	t.Run("TestOne", func(t *testing.T) {
		response, _ := json.Marshal(userResponse)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, string(response))
		}))
		AdminRequest.URL = server.URL
		rsp, err := UserGetRequest()
		if assert.NoError(t, err) {
			assert.Equal(t, rsp, userList)
		}

	})
}
