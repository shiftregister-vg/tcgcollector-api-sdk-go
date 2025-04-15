package tcgcollector

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/auth/login", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := LoginResponse{
			Token:     "test-token",
			ExpiresAt: time.Now().Add(24 * time.Hour),
			User: User{
				ID:                                  1,
				DisplayName:                         "testuser",
				EmailAddress:                        "test@example.com",
				IsEmailAddressVerified:              true,
				HasApiAccessToken:                   true,
				IsAdmin:                             false,
				CanWriteApiExpansions:               false,
				CanReadApiCards:                     true,
				CanReadApiCardsMinimal:              true,
				CanWriteApiCards:                    false,
				CanReadApiCardVariants:              true,
				CanReadApiCardVariantsMinimal:       true,
				CanWriteApiCardVariants:             false,
				CanWriteApiCardVariantTypes:         false,
				CanWriteApiCardIllustrators:         false,
				CanWriteApiCardLists:                false,
				CanReadApiStatistics:                true,
				CanWriteApiTcgPrices:                false,
				CanReadApiUsers:                     true,
				IsPremiumEnabled:                    false,
				IsPremiumWithoutSubscriptionEnabled: false,
				IsPremiumWithSubscriptionEnabled:    false,
				LastVisitDateTime:                   time.Now(),
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	request := &LoginRequest{
		Username: "testuser",
		Password: "password123",
	}
	result, err := client.Login(context.Background(), request)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test-token", result.Token)
	assert.Equal(t, "testuser", result.User.DisplayName)
	assert.Equal(t, "test@example.com", result.User.EmailAddress)

	// Update the client's apiKey with the new token
	client.apiKey = result.Token
}

func TestRegister(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/auth/register", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := RegisterResponse{
			User: User{
				ID:                                  1,
				DisplayName:                         "testuser",
				EmailAddress:                        "test@example.com",
				IsEmailAddressVerified:              false,
				HasApiAccessToken:                   false,
				IsAdmin:                             false,
				CanWriteApiExpansions:               false,
				CanReadApiCards:                     true,
				CanReadApiCardsMinimal:              true,
				CanWriteApiCards:                    false,
				CanReadApiCardVariants:              true,
				CanReadApiCardVariantsMinimal:       true,
				CanWriteApiCardVariants:             false,
				CanWriteApiCardVariantTypes:         false,
				CanWriteApiCardIllustrators:         false,
				CanWriteApiCardLists:                false,
				CanReadApiStatistics:                true,
				CanWriteApiTcgPrices:                false,
				CanReadApiUsers:                     true,
				IsPremiumEnabled:                    false,
				IsPremiumWithoutSubscriptionEnabled: false,
				IsPremiumWithSubscriptionEnabled:    false,
				LastVisitDateTime:                   time.Now(),
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	request := &RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	result, err := client.Register(context.Background(), request)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "testuser", result.User.DisplayName)
	assert.Equal(t, "test@example.com", result.User.EmailAddress)
}

func TestLogout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/auth/logout", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewClient("test-token", WithBaseURL(server.URL))
	err := client.Logout(context.Background())
	assert.NoError(t, err)
}

func TestRefreshToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/auth/refresh", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := LoginResponse{
			Token:     "new-test-token",
			ExpiresAt: time.Now().Add(24 * time.Hour),
			User: User{
				ID:                                  1,
				DisplayName:                         "testuser",
				EmailAddress:                        "test@example.com",
				IsEmailAddressVerified:              true,
				HasApiAccessToken:                   true,
				IsAdmin:                             false,
				CanWriteApiExpansions:               false,
				CanReadApiCards:                     true,
				CanReadApiCardsMinimal:              true,
				CanWriteApiCards:                    false,
				CanReadApiCardVariants:              true,
				CanReadApiCardVariantsMinimal:       true,
				CanWriteApiCardVariants:             false,
				CanWriteApiCardVariantTypes:         false,
				CanWriteApiCardIllustrators:         false,
				CanWriteApiCardLists:                false,
				CanReadApiStatistics:                true,
				CanWriteApiTcgPrices:                false,
				CanReadApiUsers:                     true,
				IsPremiumEnabled:                    false,
				IsPremiumWithoutSubscriptionEnabled: false,
				IsPremiumWithSubscriptionEnabled:    false,
				LastVisitDateTime:                   time.Now(),
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-token", WithBaseURL(server.URL))
	result, err := client.RefreshToken(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "new-test-token", result.Token)
	assert.Equal(t, "testuser", result.User.DisplayName)
	assert.Equal(t, "test@example.com", result.User.EmailAddress)

	// Update the client's apiKey with the new token
	client.apiKey = result.Token
}
