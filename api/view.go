package api

import (
	"net/http"
	"strings"

	blob "github.com/rpdg/vercel_blob"
)

func RetrieveImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	client := blob.NewVercelBlobClient()
	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) < 3 {
		http.NotFound(w, r)
		return
	}

	blobResult, err := client.Head(parts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte(blobResult.URL))
}
