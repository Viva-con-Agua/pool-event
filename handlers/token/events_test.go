package token

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pool-event/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type EventResponse struct {
	Type    string       `json:"type"`
	Message string       `json:"message"`
	Model   string       `json:"model"`
	Payload models.Event `json:"payload"`
}

var (
	event1             = &models.Event{ID: "20000000-0000-0000-0000-000000000001"}
	event2             = &models.Event{ID: "20000000-0000-0000-0000-000000000002"}
	event3             = &models.Event{ID: "20000000-0000-0000-0000-000000000003"}
	eventCreate        = &models.EventCreate{}
	eventCreateInvalid = &models.EventCreate{}
	eventUpdate        = &models.EventUpdate{}
	eventID            = "20000000-0000-0000-0000-000000000001"
)

func TestEventCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	create, _ := json.Marshal(eventCreate)
	token.Claims = adminToken
	c := tester.POSTContext(string(create), rec, token)
	if assert.NoError(t, Event.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
	rec = httptest.NewRecorder()
	c = tester.POSTContext(string(create), rec, token)
	if assert.NoError(t, Event.Create(c)) {
		assert.Equal(t, http.StatusConflict, rec.Code)
	}
	rec = httptest.NewRecorder()
	create, _ = json.Marshal(eventCreateInvalid)
	c = tester.POSTContext(string(create), rec, token)
	if assert.NoError(t, Event.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

/*
func TestEventGetByID(t *testing.T) {
	rec := httptest.NewRecorder()
	c := tester.GETByIDContext(eventID, rec, nil)
	if assert.NoError(t, Event.GetByID(c))
}*/
