package api

import (
	"fmt"
	"net/http"
	"os"

	blob "github.com/rpdg/vercel_blob"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	err := r.ParseMultipartForm(0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if r.Form.Has("api_key") {
		key := r.Form.Get("api_key")
		// TODO: check key against DB
		keyToMatch := os.Getenv("IMG_UPLOAD_API_KEY")

		if key == keyToMatch {
			fmt.Println("Key handshake successful")
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

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(file.Pathname))
		}
	} else {
		fmt.Println("No key provided in request")
		for k, v := range r.Form {
			fmt.Printf("Key: %s   Values: %#v\n", k, v)
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}
