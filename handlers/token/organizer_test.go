package token

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"pool-event/models"
	"testing"

	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/stretchr/testify/assert"
)

type OrganizerResponse struct {
	Type    string           `json:"type"`
	Message string           `json:"message"`
	Model   string           `json:"model"`
	Payload models.Organizer `json:"payload"`
}

var (
	organizer1 = &models.Organizer{
		ID:   "13a29971-3356-4325-8665-2b7a9360a79b",
		Name: "test_dummy_1",
		Modified: vmod.Modified{
			Updated: 1659521082,
			Created: 1659521082,
		},
	}
	organizer2 = &models.Organizer{
		ID:   "13a29971-3356-4325-8665-2b7a9360a79c",
		Name: "test_dummy_2",
		Modified: vmod.Modified{
			Updated: 1659521082,
			Created: 1659521082,
		},
	}
	organizer3 = &models.Organizer{
		ID:   "13a29971-3356-4325-8665-2b7a9360a79d",
		Name: "test_dummy_3",
		Modified: vmod.Modified{
			Updated: 1659521082,
			Created: 1659521082,
		},
	}
	organizerCreate         = &models.OrganizerCreate{Name: "test_artist"}
	organizerCreateNotValid = &models.OrganizerCreate{Name: ""}
	organizerUpdate         = &models.OrganizerUpdate{ID: "13a29971-3356-4325-8665-2b7a9360a79b", Name: "test_dummy_4"}
	organizerID             = "13a29971-3356-4325-8665-2b7a9360a79b"
	organizerDuplicateError = errors.New(`write exception: write errors: [E11000 duplicate key error collection: pool-event-test.organizers index: name_1 dup key: { name: "test_artist" }]`)
)

func TestOrganizerCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	dataValid, _ := json.Marshal(organizerCreate)
	token.Claims = adminToken
	t.Log(token.Claims)
	c := tester.POSTContext(string(dataValid), rec, token)
	if assert.NoError(t, Organizer.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
	rec = httptest.NewRecorder()
	c = tester.POSTContext(string(dataValid), rec, token)
	if assert.NoError(t, Organizer.Create(c)) {
		assert.Equal(t, http.StatusConflict, rec.Code)
	}
	rec = httptest.NewRecorder()
	dataInvalid, _ := json.Marshal(artistCreateNotValid)
	c = tester.POSTContext(string(dataInvalid), rec, token)
	if assert.NoError(t, Organizer.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
func TestOrganizerGetByID(t *testing.T) {
	rec := httptest.NewRecorder()
	c := tester.GETByIDContext(organizerID, rec, nil)
	if assert.NoError(t, Organizer.GetByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		rsp := new(OrganizerResponse)
		json.Unmarshal(rec.Body.Bytes(), rsp)
		assert.Equal(t, *organizer1, rsp.Payload)
	}
	rec = httptest.NewRecorder()
	c = tester.GETByIDContext("dummyID", rec, nil)
	if assert.NoError(t, Organizer.GetByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestOrganizerUpdate(t *testing.T) {
	rec := httptest.NewRecorder()
	au, _ := json.Marshal(organizerUpdate)
	c := tester.PUTContext(string(au), rec, nil)
	if assert.NoError(t, Organizer.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}

func TestOrganizerGet(t *testing.T) {
	rec := httptest.NewRecorder()
	c := tester.GETContext("", rec, nil)
	if assert.NoError(t, Organizer.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestOrganizerDelete(t *testing.T) {
	rec := httptest.NewRecorder()
	c := tester.DELETEContext(organizerID, rec, nil)
	if assert.NoError(t, Organizer.Delete(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
