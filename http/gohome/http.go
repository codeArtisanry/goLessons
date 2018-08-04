package main

import(
	"net/http"
    "errors"
    "mime/multipart"
)

//assumes you have already used parseMultipartForm on your request
func FormFiles(r *http.Request, key string) ([]*multipart.FileHeader, error) {

    if r.MultipartForm != nil && r.MultipartForm.File != nil {       
        return r.MultipartForm.File[key], nil
    }
    return nil, errors.New("missing file!")
}