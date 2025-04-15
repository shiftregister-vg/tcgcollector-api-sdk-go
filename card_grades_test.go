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
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/card-grades", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

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

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	grade := &CardGrade{
		CardID:         1,
		GradeCompanyID: 1,
		GradeValue:     "PSA 10",
		CertificateID:  "12345678",
		GradedAt:       time.Now(),
		Notes:          "Gem Mint",
	}

	result, err := client.CreateCardGrade(context.Background(), grade)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "PSA 10", result.GradeValue)
	assert.Equal(t, "Gem Mint", result.Notes)
}

func TestUpdateCardGrade(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/api/card-grades/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

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
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	grade := &CardGrade{
		CardID:         1,
		GradeCompanyID: 1,
		GradeValue:     "PSA 9",
		CertificateID:  "12345678",
		GradedAt:       time.Now(),
		Notes:          "Mint",
	}

	result, err := client.UpdateCardGrade(context.Background(), 1, grade)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "PSA 9", result.GradeValue)
	assert.Equal(t, "Mint", result.Notes)
}

func TestDeleteCardGrade(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/api/card-grades/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.DeleteCardGrade(context.Background(), 1)
	assert.NoError(t, err)
}

func TestCreateCardGradeError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Invalid grade data",
			Code:    "INVALID_DATA",
		})
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	grade := &CardGrade{
		GradeValue: "", // Invalid grade value
	}
	result, err := client.CreateCardGrade(context.Background(), grade)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Invalid grade data")
}

func TestUpdateCardGradeError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Grade not found",
			Code:    "NOT_FOUND",
		})
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	grade := &CardGrade{
		GradeValue: "PSA 10",
	}
	result, err := client.UpdateCardGrade(context.Background(), 999, grade)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Grade not found")
}

func TestDeleteCardGradeError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Grade not found",
			Code:    "NOT_FOUND",
		})
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	err := client.DeleteCardGrade(context.Background(), 999)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Grade not found")
}

func TestListCardGradesWithAllParams(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-grades", r.URL.Path)

		query := r.URL.Query()
		assert.Equal(t, "1", query.Get("cardId"))
		assert.Equal(t, "2", query.Get("gradeCompanyId"))
		assert.Equal(t, "PSA 10", query.Get("gradeValue"))
		assert.Equal(t, "1", query.Get("page"))
		assert.Equal(t, "10", query.Get("pageSize"))

		response := ListResponse[CardGrade]{
			Items: []CardGrade{
				{
					ID:             1,
					CardID:         1,
					GradeCompanyID: 2,
					GradeValue:     "PSA 10",
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

	client := NewClient("test-api-key", WithBaseURL(ts.URL))

	cardID := 1
	gradeCompanyID := 2
	gradeValue := "PSA 10"
	page := 1
	pageSize := 10

	params := &ListCardGradesParams{
		CardID:         &cardID,
		GradeCompanyID: &gradeCompanyID,
		GradeValue:     &gradeValue,
		Page:           &page,
		PageSize:       &pageSize,
	}

	result, err := client.ListCardGrades(context.Background(), params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Items))
}

func TestListCardGradesWithNilParams(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-grades", r.URL.Path)
		assert.Empty(t, r.URL.Query())

		response := ListResponse[CardGrade]{
			Items:          []CardGrade{},
			ItemCount:      0,
			TotalItemCount: 0,
			Page:           1,
			PageCount:      1,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	result, err := client.ListCardGrades(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Empty(t, result.Items)
}

func TestListCardGradesError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Internal server error",
			Code:    "INTERNAL_ERROR",
		})
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	result, err := client.ListCardGrades(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Internal server error")
}

func TestListCardGradesInvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	result, err := client.ListCardGrades(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid character")
}

func TestGetCardGradeNotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "Card grade not found",
			Code:    "NOT_FOUND",
		})
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	result, err := client.GetCardGrade(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Card grade not found")
}

func TestGetCardGradeInvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	result, err := client.GetCardGrade(context.Background(), 1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid character")
}

func TestCreateCardGradeInvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	grade := &CardGrade{
		CardID:         1,
		GradeCompanyID: 1,
		GradeValue:     "PSA 10",
	}
	result, err := client.CreateCardGrade(context.Background(), grade)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid character")
}

func TestUpdateCardGradeInvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	grade := &CardGrade{
		CardID:         1,
		GradeCompanyID: 1,
		GradeValue:     "PSA 10",
	}
	result, err := client.UpdateCardGrade(context.Background(), 1, grade)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid character")
}
