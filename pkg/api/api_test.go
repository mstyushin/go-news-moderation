package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mstyushin/go-news-moderation/pkg/config"

	"github.com/stretchr/testify/assert"
)

func TestAPI_moderation(t *testing.T) {
	cfg := config.DefaultConfig()
	api := New(cfg)
	api.initMux()

	tests := []struct {
		sample   ModerationRequest
		expected int
	}{
		{
			sample: ModerationRequest{
				Author: "Alice",
				Text:   "Totally legal comment",
			},
			expected: 200,
		},
		{
			sample: ModerationRequest{
				Author: "Bob qwerty",
				Text:   "Not so legal nickname",
			},
			expected: 400,
		},
		{
			sample: ModerationRequest{
				Author: "Jane",
				Text:   "Absolutely illegal asdfg comment",
			},
			expected: 400,
		},
	}

	var buf bytes.Buffer

	for idx, test := range tests {
		t.Run(fmt.Sprintf("sample-%d", idx), func(t *testing.T) {
			json.NewEncoder(&buf).Encode(test.sample)

			req := httptest.NewRequest(http.MethodPost, "/moderation", &buf)
			rr := httptest.NewRecorder()

			api.mux.ServeHTTP(rr, req)

			assert.True(t, rr.Code == test.expected)
		})
	}
}
