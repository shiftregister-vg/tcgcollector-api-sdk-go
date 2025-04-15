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

		var requestBody LoginRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		assert.NoError(t, err)
		assert.Equal(t, "testuser", requestBody.Username)
		assert.Equal(t, "password123", requestBody.Password)

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

	client := NewClient("", WithBaseURL(server.URL))
	req := &LoginRequest{
		Username: "testuser",
		Password: "password123",
	}
	result, err := client.Login(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test-token", result.Token)
	assert.Equal(t, "testuser", result.User.DisplayName)
	assert.Equal(t, "test@example.com", result.User.EmailAddress)
}

func TestLoginError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Invalid credentials",
			Code:    "INVALID_CREDENTIALS",
		})
	}))
	defer server.Close()

	client := NewClient("", WithBaseURL(server.URL))
	req := &LoginRequest{
		Username: "testuser",
		Password: "wrong-password",
	}
	result, err := client.Login(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Invalid credentials")
}

func TestLoginInvalidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"invalid json`))
	}))
	defer server.Close()

	client := NewClient("", WithBaseURL(server.URL))
	req := &LoginRequest{
		Username: "testuser",
		Password: "password123",
	}
	result, err := client.Login(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to decode response")
}

func TestRegister(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/auth/register", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var requestBody RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		assert.NoError(t, err)
		assert.Equal(t, "testuser", requestBody.Username)
		assert.Equal(t, "test@example.com", requestBody.Email)
		assert.Equal(t, "password123", requestBody.Password)

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

	client := NewClient("", WithBaseURL(server.URL))
	req := &RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	result, err := client.Register(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "testuser", result.User.DisplayName)
	assert.Equal(t, "test@example.com", result.User.EmailAddress)
}

func TestRegisterError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Email already exists",
			Code:    "EMAIL_EXISTS",
		})
	}))
	defer server.Close()

	client := NewClient("", WithBaseURL(server.URL))
	req := &RegisterRequest{
		Username: "testuser",
		Email:    "existing@example.com",
		Password: "password123",
	}
	result, err := client.Register(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Email already exists")
}

func TestRegisterInvalidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"invalid json`))
	}))
	defer server.Close()

	client := NewClient("", WithBaseURL(server.URL))
	req := &RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	result, err := client.Register(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to decode response")
}

func TestLogout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/auth/logout", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewClient("test-token", WithBaseURL(server.URL))
	err := client.Logout(context.Background())
	assert.NoError(t, err)
}

func TestLogoutError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Invalid token",
			Code:    "INVALID_TOKEN",
		})
	}))
	defer server.Close()

	client := NewClient("invalid-token", WithBaseURL(server.URL))
	err := client.Logout(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid token")
}

func TestRefreshToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/auth/refresh", r.URL.Path)
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		response := LoginResponse{
			Token:     "new-token",
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
	assert.Equal(t, "new-token", result.Token)
	assert.Equal(t, "testuser", result.User.DisplayName)
	assert.Equal(t, "test@example.com", result.User.EmailAddress)
}

func TestRefreshTokenError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Invalid token",
			Code:    "INVALID_TOKEN",
		})
	}))
	defer server.Close()

	client := NewClient("invalid-token", WithBaseURL(server.URL))
	result, err := client.RefreshToken(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Invalid token")
}

func TestRefreshTokenInvalidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"invalid json`))
	}))
	defer server.Close()

	client := NewClient("test-token", WithBaseURL(server.URL))
	result, err := client.RefreshToken(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to decode response")
}
