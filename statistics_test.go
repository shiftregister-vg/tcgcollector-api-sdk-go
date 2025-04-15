package tcgcollector

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStatistics(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/statistics", r.URL.Path)
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))

		response := UserStatistics{
			UserCount:              1000,
			MonthlyActiveUserCount: 500,
			Premium: PremiumStatistics{
				UserCount:                          200,
				UserWithoutSubscriptionCount:       50,
				UserWithSubscriptionCount:          150,
				ActiveSubscriptionCount:            120,
				ActiveNonExpiringSubscriptionCount: 80,
				ActiveExpiringSubscriptionCount:    40,
				SuspendedSubscriptionCount:         10,
				ExpiredSubscriptionCount:           30,
				CanceledSubscriptionCount:          20,
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	stats, err := client.GetStatistics(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, 1000, stats.UserCount)
	assert.Equal(t, 500, stats.MonthlyActiveUserCount)
	assert.Equal(t, 200, stats.Premium.UserCount)
	assert.Equal(t, 50, stats.Premium.UserWithoutSubscriptionCount)
	assert.Equal(t, 150, stats.Premium.UserWithSubscriptionCount)
	assert.Equal(t, 120, stats.Premium.ActiveSubscriptionCount)
	assert.Equal(t, 80, stats.Premium.ActiveNonExpiringSubscriptionCount)
	assert.Equal(t, 40, stats.Premium.ActiveExpiringSubscriptionCount)
	assert.Equal(t, 10, stats.Premium.SuspendedSubscriptionCount)
	assert.Equal(t, 30, stats.Premium.ExpiredSubscriptionCount)
	assert.Equal(t, 20, stats.Premium.CanceledSubscriptionCount)
}
