package tcgcollector

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListUsers(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/api/users" {
			t.Errorf("Expected path /api/users, got %s", r.URL.Path)
		}

		// Verify query parameters
		query := r.URL.Query()
		if query.Get("page") != "1" {
			t.Errorf("Expected page=1, got %s", query.Get("page"))
		}
		if query.Get("pageSize") != "10" {
			t.Errorf("Expected pageSize=10, got %s", query.Get("pageSize"))
		}
		if query.Get("search") != "test" {
			t.Errorf("Expected search=test, got %s", query.Get("search"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"items": [
				{
					"id": 1,
					"username": "testuser",
					"email": "test@example.com",
					"createdAt": "2024-01-01T00:00:00Z",
					"updatedAt": "2024-01-01T00:00:00Z"
				}
			],
			"itemCount": 1,
			"totalItemCount": 1,
			"page": 1,
			"pageCount": 1
		}`))
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Create parameters
	page := 1
	pageSize := 10
	search := "test"

	params := &ListUsersParams{
		Page:     &page,
		PageSize: &pageSize,
		Search:   &search,
	}

	// Make request
	result, err := client.ListUsers(context.Background(), params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result.Items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(result.Items))
	}

	if result.Items[0].ID != 1 {
		t.Errorf("Expected item ID to be 1, got %d", result.Items[0].ID)
	}
}

func TestGetUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/users/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := User{
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
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.GetUser(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "testuser", result.DisplayName)
	assert.Equal(t, "test@example.com", result.EmailAddress)
}

func TestCreateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/users", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := User{
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
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	params := &CreateUserParams{
		DisplayName:  "testuser",
		EmailAddress: "test@example.com",
		Password:     "password123",
	}
	result, err := client.CreateUser(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "testuser", result.DisplayName)
	assert.Equal(t, "test@example.com", result.EmailAddress)
}

func TestUpdateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/api/users/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := User{
			ID:                                  1,
			DisplayName:                         "updateduser",
			EmailAddress:                        "updated@example.com",
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
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	displayName := "updateduser"
	emailAddress := "updated@example.com"
	params := &UpdateUserParams{
		DisplayName:  &displayName,
		EmailAddress: &emailAddress,
	}
	result, err := client.UpdateUser(context.Background(), 1, params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "updateduser", result.DisplayName)
	assert.Equal(t, "updated@example.com", result.EmailAddress)
}

func TestDeleteUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/api/users/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	err := client.DeleteUser(context.Background(), 1)
	assert.NoError(t, err)
}

func TestGetCurrentUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/users/me", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := User{
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
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	result, err := client.GetCurrentUser(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "testuser", result.DisplayName)
	assert.Equal(t, "test@example.com", result.EmailAddress)
}

func TestUpdateCurrentUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/api/users/me", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := User{
			ID:                                  1,
			DisplayName:                         "updateduser",
			EmailAddress:                        "updated@example.com",
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
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	displayName := "updateduser"
	emailAddress := "updated@example.com"
	params := &UpdateUserParams{
		DisplayName:  &displayName,
		EmailAddress: &emailAddress,
	}
	result, err := client.UpdateCurrentUser(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "updateduser", result.DisplayName)
	assert.Equal(t, "updated@example.com", result.EmailAddress)
}

func TestDeleteCurrentUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/api/users/me", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	err := client.DeleteCurrentUser(context.Background())
	assert.NoError(t, err)
}

func TestGetUserCount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/users/count", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"count": 42}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	count, err := client.GetUserCount(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 42, count)
}

func TestPruneActivityLogs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/users/prune-activity-logs", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.PruneActivityLogs(context.Background())
	assert.NoError(t, err)
}

func TestDisableUserPremium(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/users/1/disable-premium", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.DisableUserPremium(context.Background(), 1)
	assert.NoError(t, err)
}

func TestEnableUserPremiumWithoutSubscription(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/users/1/enable-premium-without-subscription", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.EnableUserPremiumWithoutSubscription(context.Background(), 1)
	assert.NoError(t, err)
}

func TestGenerateAPIAccessToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/users/1/generate-api-access-token", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"token": "test-token"}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	token, err := client.GenerateAPIAccessToken(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "test-token", token)
}

func TestGetUserPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/users/1/permissions", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"permissions": ["read", "write"]}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	permissions, err := client.GetUserPermissions(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, []string{"read", "write"}, permissions)
}

func TestRevokeAPIAccessToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/users/1/revoke-api-access-token", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.RevokeAPIAccessToken(context.Background(), 1)
	assert.NoError(t, err)
}

func TestGetUserPreferences(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/users/1/preferences", r.URL.Path)

		// Create mock response
		response := UserPreferences{
			ID:              1,
			UserID:          1,
			DefaultCurrency: "USD",
			Language:        "en",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	result, err := client.GetUserPreferences(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "USD", result.DefaultCurrency)
	assert.Equal(t, "en", result.Language)
}

func TestUpdateUserPreferences(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/api/users/1/preferences", r.URL.Path)

		// Create mock response
		response := UserPreferences{
			ID:              1,
			UserID:          1,
			DefaultCurrency: "EUR",
			Language:        "fr",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	preferences := &UserPreferences{
		DefaultCurrency: "EUR",
		Language:        "fr",
	}

	result, err := client.UpdateUserPreferences(context.Background(), 1, preferences)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "EUR", result.DefaultCurrency)
	assert.Equal(t, "fr", result.Language)
}
