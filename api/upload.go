package api

import (
	"net/http"
	"os"

	blob "github.com/rpdg/vercel_blob"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if r.PostForm.Has("api_key") {
		key := r.PostForm.Get("api_key")
		// TODO: check key against DB

		keyToMatch := os.Getenv("IMG_UPLOAD_API_KEY")
		if keyToMatch == "" {
			keyToMatch = "hazz1233"
		}
		if key == keyToMatch {
			client := blob.NewVercelBlobClient()
			formFile, header, err := r.FormFile("uploadFile")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}

			defer formFile.Close()

			file, err := client.Put("images/"+header.Filename, formFile, blob.PutCommandOptions{AddRandomSuffix: true})
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}

			w.Write([]byte(file.URL))
		}
	}
}
