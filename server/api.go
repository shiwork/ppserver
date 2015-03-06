package server

import (
	"encoding/json"
	"net/http"

	"github.com/zenazn/goji/web"
	"github.com/shiwork/ppserver/ppserver"
)

func SampleApi(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Scheme == "" {
		r.URL.Scheme = "http"
	}
	movies := &[]ppserver.Movie{
		ppserver.Movie{
			Url: r.URL.Scheme + "://" + r.Host + "/apps/ppserver/demo/assets/sample_movie.mp4",
			Title: "PPServer Sample",
		},
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(movies)
}
