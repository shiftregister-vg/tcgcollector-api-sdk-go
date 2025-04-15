package tcgcollector

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListEnergyTypes(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/energy-types", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"name": "Fire",
				"description": "Fire energy type",
				"symbol": "🔥",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			},
			{
				"id": 2,
				"name": "Water",
				"description": "Water energy type",
				"symbol": "💧",
				"createdAt": "2024-01-01T00:00:00Z",
				"updatedAt": "2024-01-01T00:00:00Z"
			}
		]`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	energyTypes, err := client.ListEnergyTypes(context.Background())
	assert.NoError(t, err)
	assert.Len(t, energyTypes, 2)
	assert.Equal(t, 1, energyTypes[0].ID)
	assert.Equal(t, "Fire", energyTypes[0].Name)
	assert.Equal(t, "Fire energy type", energyTypes[0].Description)
	assert.Equal(t, "🔥", energyTypes[0].Symbol)
	assert.Equal(t, 2, energyTypes[1].ID)
	assert.Equal(t, "Water", energyTypes[1].Name)
	assert.Equal(t, "Water energy type", energyTypes[1].Description)
	assert.Equal(t, "💧", energyTypes[1].Symbol)
}

func TestGetEnergyType(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/energy-types/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"id": 1,
			"name": "Fire",
			"description": "Fire energy type",
			"symbol": "🔥",
			"createdAt": "2024-01-01T00:00:00Z",
			"updatedAt": "2024-01-01T00:00:00Z"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	energyType, err := client.GetEnergyType(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, energyType.ID)
	assert.Equal(t, "Fire", energyType.Name)
	assert.Equal(t, "Fire energy type", energyType.Description)
	assert.Equal(t, "🔥", energyType.Symbol)
}

func TestListEnergyTypesError(t *testing.T) {
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
	energyTypes, err := client.ListEnergyTypes(context.Background())
	assert.Error(t, err)
	assert.Nil(t, energyTypes)
	assert.Contains(t, err.Error(), "Internal server error")
}

func TestListEnergyTypesInvalidResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	energyTypes, err := client.ListEnergyTypes(context.Background())
	assert.Error(t, err)
	assert.Nil(t, energyTypes)
	assert.Contains(t, err.Error(), "invalid character")
}

func TestGetEnergyTypeError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{
			"message": "Energy type not found",
			"code": "NOT_FOUND"
		}`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	energyType, err := client.GetEnergyType(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, energyType)
	assert.Contains(t, err.Error(), "Energy type not found")
}

func TestGetEnergyTypeInvalidResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	client := NewClient("test-api-key", WithBaseURL(ts.URL))
	energyType, err := client.GetEnergyType(context.Background(), 1)
	assert.Error(t, err)
	assert.Nil(t, energyType)
	assert.Contains(t, err.Error(), "invalid character")
}
