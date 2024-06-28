package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender(t *testing.T) {
	type args struct {
		response interface{}
		cache    int
		callback string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Complete Argument",
			args: args{
				response: `{"alive": true}`,
				callback: "callback",
				cache:    0,
			},
		},
		{
			name: "Empty Callback",
			args: args{
				response: `{"alive": true}`,
				callback: "",
				cache:    0,
			},
		},
		{
			name: "Set Cache",
			args: args{
				response: `{"alive": true}`,
				callback: "",
				cache:    60,
			},
		},
		{
			name: "Fail Json Marshal",
			args: args{
				response: `fail`,
				callback: "",
				cache:    60,
			},
		},
	}
	for _, tt := range tests {
		var ts *httptest.Server
		if tt.args.response == "fail" {
			ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Render(w, make(chan int), tt.args.cache, tt.args.callback)
			}))
		} else {
			ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Render(w, tt.args.response, tt.args.cache, tt.args.callback)
			}))
		}
		defer ts.Close()
		http.Get(ts.URL)
	}
}
