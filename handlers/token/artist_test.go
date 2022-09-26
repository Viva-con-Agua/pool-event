package token

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"pool-event/models"
	"testing"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/stretchr/testify/assert"
)

type ArtistResponse struct {
	Type    string        `json:"type"`
	Message string        `json:"message"`
	Model   string        `json:"model"`
	Payload models.Artist `json:"payload"`
}

var (
	response = &vcago.Response{Type: "success", Message: "successfully_selected", Model: "artist"}
	artist1  = &models.Artist{
		ID:   "13a29971-3356-4325-8665-2b7a9360a79b",
		Name: "test_dummy_1",
		Modified: vmod.Modified{
			Updated: 1659521082,
			Created: 1659521082,
		},
	}
	artist2 = &models.Artist{
		ID:   "13a29971-3356-4325-8665-2b7a9360a79c",
		Name: "test_dummy_2",
		Modified: vmod.Modified{
			Updated: 1659521082,
			Created: 1659521082,
		},
	}
	artist3 = &models.Artist{
		ID:   "13a29971-3356-4325-8665-2b7a9360a79d",
		Name: "test_dummy_3",
		Modified: vmod.Modified{
			Updated: 1659521082,
			Created: 1659521082,
		},
	}
	artistCreate         = &models.ArtistCreate{Name: "test_artist"}
	artistCreateNotValid = &models.ArtistCreate{Name: ""}
	artistUpdate         = &models.ArtistUpdate{ID: "13a29971-3356-4325-8665-2b7a9360a79b", Name: "test_dummy_4"}
	artistID             = "13a29971-3356-4325-8665-2b7a9360a79b"
	artistDuplicateError = errors.New(`write exception: write errors: [E11000 duplicate key error collection: pool-event-test.artists index: name_1 dup key: { name: "test_artist" }]`)
)

func TestArtistCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	ac, _ := json.Marshal(artistCreate)
	token.Claims = adminToken
	c := tester.POSTContext(string(ac), rec, token)
	if assert.NoError(t, Artist.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
	rec = httptest.NewRecorder()
	c = tester.POSTContext(string(ac), rec, token)
	if assert.NoError(t, Artist.Create(c)) {
		assert.Equal(t, http.StatusConflict, rec.Code)
	}
	rec = httptest.NewRecorder()
	acn, _ := json.Marshal(artistCreateNotValid)
	c = tester.POSTContext(string(acn), rec, token)
	if assert.NoError(t, Artist.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestArtistGetByID(t *testing.T) {
	rec := httptest.NewRecorder()
	c := tester.GETByIDContext(artistID, rec, nil)
	if assert.NoError(t, Artist.GetByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		rsp := new(ArtistResponse)
		json.Unmarshal(rec.Body.Bytes(), rsp)
		assert.Equal(t, *artist1, rsp.Payload)
	}
	rec = httptest.NewRecorder()
	c = tester.GETByIDContext("dummyID", rec, nil)
	if assert.NoError(t, Artist.GetByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestArtistUpdate(t *testing.T) {
	rec := httptest.NewRecorder()
	au, _ := json.Marshal(artistUpdate)
	c := tester.PUTContext(string(au), rec, nil)
	if assert.NoError(t, Artist.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}

func TestArtistGet(t *testing.T) {
	rec := httptest.NewRecorder()
	c := tester.GETContext("", rec, nil)
	if assert.NoError(t, Artist.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestArtistDelete(t *testing.T) {
	rec := httptest.NewRecorder()
	c := tester.DELETEContext(artistID, rec, nil)
	if assert.NoError(t, Artist.Delete(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
