package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/hardikchoksi151/go-api-mux/types"
)

type mockUserStore struct{}

// CreateUser implements types.UserStore.
func (m *mockUserStore) CreateUser(types.User) error {
	panic("unimplemented")
}

// GetUserByEmail implements types.UserStore.
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	panic("unimplemented")
}

// GetUserById implements types.UserStore.
func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	panic("unimplemented")
}

func TestUserServiceHandlers(t *testing.T) {
	userStore := mockUserStore{}
	handler := NewHandler(&userStore)

	payload := types.RegisterUserPayload{
		FirstName: "user",
		LastName:  "123",
		Email:     "",
		Password:  "asd",
	}
	payloadBytes, _ := json.Marshal(payload)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/register", strings.NewReader(string(payloadBytes)))
		if err != nil {
			log.Fatal(err)
		}
	})
}
