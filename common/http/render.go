package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	json "github.com/json-iterator/go"
)

var environment = os.Getenv("TKPENV")

// Render render http response with cache (in minutes) and callback param if exist
func Render(w http.ResponseWriter, response interface{}, cache int, callback string) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {

		//log.WithField("error", err).Error("Failed to marshal response")
	}

	if callback != "" {
		w.Header().Add("Content-Type", "text/javascript")
	} else {
		w.Header().Add("Content-Type", "application/json")
		if environment != "production" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}

	if cache > 0 {
		now := time.Now()
		t := now.Add(time.Duration(cache) * time.Minute)

		w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(cache*60))
		w.Header().Set("Expires", t.Format(time.RFC1123))
		w.Header().Set("Date", now.Format(time.RFC1123))
	}

	if callback != "" {
		fmt.Fprintf(w, "%s(%s)", callback, jsonResponse)
	} else {
		w.Write(jsonResponse)
	}
}

//RenderHTTPJSON middleware to write json to http
func RenderHTTPJSON(w http.ResponseWriter, data interface{}, code int, callback string) {
	b, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "StatusInternalServerError", http.StatusInternalServerError)
		return
	}

	if callback != "" {
		w.Header().Set("Content-Type", "text/javascript")
		fmt.Fprintf(w, "%s(%s)", callback, b)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}

//RenderHTTPCode middle to write http header code to header
func RenderHTTPCode(w http.ResponseWriter, message string, code int) {
	type response struct {
		M string `json:"message"`
	}

	b, err := json.Marshal(response{M: message})
	if err != nil {
		http.Error(w, "StatusInternalServerError", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}

//RenderHTTPText middleware to write plain text to http
func RenderHTTPText(w http.ResponseWriter, message string, code int) {
	b := []byte(message)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write(b)
}

//RenderHTTPTextWithHeader middleware to write plain text to http
func RenderHTTPTextWithHeader(w http.ResponseWriter, message string, code int, additionalHeader map[string]string) {
	b := []byte(message)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	for k := range additionalHeader {
		w.Header().Set(k, additionalHeader[k])
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write(b)
}

//GetJSON parse http response from the given url to targeted struct
//Returned response must be in json format
func GetJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
