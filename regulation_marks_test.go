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

func TestListRegulationMarks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/regulation-marks", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := ListRegulationMarksResponse{
			Items: []RegulationMark{
				{
					ID:          1,
					Name:        "D",
					Description: "Regulation Mark D",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
			ItemCount:      1,
			TotalItemCount: 1,
			Page:           1,
			PageCount:      1,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	marks, err := client.ListRegulationMarks(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, marks)
	assert.Len(t, marks.Items, 1)
	assert.Equal(t, 1, marks.Items[0].ID)
	assert.Equal(t, "D", marks.Items[0].Name)
	assert.Equal(t, "Regulation Mark D", marks.Items[0].Description)
}

func TestGetRegulationMark(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/regulation-marks/1", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := RegulationMark{
			ID:          1,
			Name:        "D",
			Description: "Regulation Mark D",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	mark, err := client.GetRegulationMark(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, mark)
	assert.Equal(t, 1, mark.ID)
	assert.Equal(t, "D", mark.Name)
	assert.Equal(t, "Regulation Mark D", mark.Description)
}
