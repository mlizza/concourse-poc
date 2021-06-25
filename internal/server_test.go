package internal

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_Init(t *testing.T) {
	type fields struct {
		Port string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "index",
			fields: fields{
				Port: "8080",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				Port: tt.fields.Port,
			}
			router := s.Init()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			response := httptest.NewRecorder()
			router.ServeHTTP(response, req)
			assert.Equal(t, 200, response.Code, "OK response is expected")
		})
	}
}
