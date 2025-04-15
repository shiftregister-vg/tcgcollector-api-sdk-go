package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCardGradeCompanies(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-grade-companies", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "PSA",
				"description": "Professional Sports Authenticator",
				"website": "https://www.psacard.com",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "BGS",
				"description": "Beckett Grading Services",
				"website": "https://www.beckett.com/grading",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	companies, err := client.ListCardGradeCompanies(context.Background())
	assert.NoError(t, err)
	assert.Len(t, companies, 2)
	assert.Equal(t, 1, companies[0].ID)
	assert.Equal(t, "PSA", companies[0].Name)
	assert.Equal(t, "Professional Sports Authenticator", companies[0].Description)
	assert.Equal(t, "https://www.psacard.com", companies[0].Website)
	assert.Equal(t, 2, companies[1].ID)
	assert.Equal(t, "BGS", companies[1].Name)
	assert.Equal(t, "Beckett Grading Services", companies[1].Description)
	assert.Equal(t, "https://www.beckett.com/grading", companies[1].Website)
}

func TestGetCardGradeCompany(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/card-grade-companies/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "PSA",
			"description": "Professional Sports Authenticator",
			"website": "https://www.psacard.com",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	company, err := client.GetCardGradeCompany(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, company.ID)
	assert.Equal(t, "PSA", company.Name)
	assert.Equal(t, "Professional Sports Authenticator", company.Description)
	assert.Equal(t, "https://www.psacard.com", company.Website)
}

func TestListCardGradeCompaniesError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{
			"message": "Internal server error",
			"code": "INTERNAL_ERROR"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	companies, err := client.ListCardGradeCompanies(context.Background())
	assert.Error(t, err)
	assert.Nil(t, companies)
	assert.Contains(t, err.Error(), "Internal server error")
}

func TestListCardGradeCompaniesInvalidResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	companies, err := client.ListCardGradeCompanies(context.Background())
	assert.Error(t, err)
	assert.Nil(t, companies)
	assert.Contains(t, err.Error(), "invalid character")
}

func TestGetCardGradeCompanyError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{
			"message": "Card grade company not found",
			"code": "NOT_FOUND"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	company, err := client.GetCardGradeCompany(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, company)
	assert.Contains(t, err.Error(), "Card grade company not found")
}

func TestGetCardGradeCompanyInvalidResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	company, err := client.GetCardGradeCompany(context.Background(), 1)
	assert.Error(t, err)
	assert.Nil(t, company)
	assert.Contains(t, err.Error(), "invalid character")
}
