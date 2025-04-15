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

func TestListCardGrades(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-grades", r.URL.Path)

		// Check query parameters
		query := r.URL.Query()
		assert.Equal(t, "1", query.Get("cardId"))

		// Create mock response
		response := ListResponse[CardGrade]{
			Items: []CardGrade{
				{
					ID:             1,
					CardID:         1,
					GradeCompanyID: 1,
					GradeValue:     "PSA 10",
					CertificateID:  "12345678",
					GradedAt:       time.Now(),
					Notes:          "Gem Mint",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				},
			},
			ItemCount:      1,
			TotalItemCount: 1,
			Page:           1,
			PageCount:      1,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Create request parameters
	cardID := 1
	params := &ListCardGradesParams{
		CardID: &cardID,
	}

	// Make request
	result, err := client.ListCardGrades(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Items))
	assert.Equal(t, "PSA 10", result.Items[0].GradeValue)
}

func TestGetCardGrade(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-grades/1", r.URL.Path)

		// Create mock response
		response := CardGrade{
			ID:             1,
			CardID:         1,
			GradeCompanyID: 1,
			GradeValue:     "PSA 10",
			CertificateID:  "12345678",
			GradedAt:       time.Now(),
			Notes:          "Gem Mint",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Make request
	result, err := client.GetCardGrade(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "PSA 10", result.GradeValue)
}

func TestCreateCardGrade(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/card-grades", r.URL.Path)

		// Create mock response
		response := CardGrade{
			ID:             1,
			CardID:         1,
			GradeCompanyID: 1,
			GradeValue:     "PSA 10",
			CertificateID:  "12345678",
			GradedAt:       time.Now(),
			Notes:          "Gem Mint",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Create request
	grade := &CardGrade{
		CardID:         1,
		GradeCompanyID: 1,
		GradeValue:     "PSA 10",
		CertificateID:  "12345678",
		GradedAt:       time.Now(),
		Notes:          "Gem Mint",
	}

	// Make request
	result, err := client.CreateCardGrade(context.Background(), grade)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "PSA 10", result.GradeValue)
}

func TestUpdateCardGrade(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/api/card-grades/1", r.URL.Path)

		// Create mock response
		response := CardGrade{
			ID:             1,
			CardID:         1,
			GradeCompanyID: 1,
			GradeValue:     "PSA 9",
			CertificateID:  "12345678",
			GradedAt:       time.Now(),
			Notes:          "Mint",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Create request
	grade := &CardGrade{
		CardID:         1,
		GradeCompanyID: 1,
		GradeValue:     "PSA 9",
		CertificateID:  "12345678",
		GradedAt:       time.Now(),
		Notes:          "Mint",
	}

	// Make request
	result, err := client.UpdateCardGrade(context.Background(), 1, grade)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "PSA 9", result.GradeValue)
}

func TestDeleteCardGrade(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/api/card-grades/1", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	// Create client with test server URL
	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	// Make request
	err := client.DeleteCardGrade(context.Background(), 1)
	assert.NoError(t, err)
}
