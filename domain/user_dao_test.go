package domain

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserNotFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "Not expecting a user with id 0")
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, fmt.Sprintf("User %v was not found", 0), err.Message)
}

func TestGetUser(t *testing.T) {
	user, err := GetUser(1)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "One", user.FirstName)
	assert.EqualValues(t, "Test", user.LastName)
	assert.EqualValues(t, "test_one@gmail.com", user.Email)
}
