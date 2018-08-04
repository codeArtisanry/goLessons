package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/sunfmin/reflectutils"
)

func UnmarshalByNames(r *http.Request, v interface{}, names []string) (err error) {
	f := func(key string) (r string, skip bool) {
		skip = true
		for _, n := range names {
			if n == key {
				skip = false
				break
			}
		}
		r = key
		return
	}
	UnmarshalFunc(r, v, f)
	return
}

func UnmarshalByPrefix(r *http.Request, v interface{}, prefix string) (err error) {
	f := func(key string) (stripped string, skip bool) {
		if prefix == "" {
			stripped = key
			skip = false
			return
		}
		if strings.Index(key, prefix) != 0 {
			skip = true
			return
		}
		stripped = key[len(prefix):]
		skip = false
		return
	}
	UnmarshalFunc(r, v, f)
	return
}

func UnmarshalFunc(r *http.Request, v interface{}, f func(key string) (string, bool)) (err error) {
	if r.Form == nil && r.MultipartForm == nil {
		r.ParseMultipartForm(32 << 20)
	}

	var vals map[string][]string
	if r.MultipartForm != nil {
		vals = r.MultipartForm.Value
	} else if r.Form != nil {
		vals = map[string][]string(r.Form)
	}

	for fk, fv := range vals {
		key, skip := f(fk)
		if skip {
			continue
		}
		for _, velem := range fv {
			reflectutils.Set(v, key, velem)
		}
	}

	if r.MultipartForm != nil {

		for filek, filev := range r.MultipartForm.File {
			key, skip := f(filek)
			if skip {
				continue
			}
			for _, velem := range filev {
				reflectutils.Set(v, key, velem)
			}
		}
		return
	}

	return
}

type Person struct {
	Name        string
	GivenName   string
	Photo       *multipart.FileHeader
	Resume      *multipart.FileHeader
	Gender      int
	Company     *Company
	Departments []*Department
	Projects    []*Project
	Phones      map[string]string
}

type Company struct {
	Name string
}

type Department struct {
	Id   string
	Name string
}

type Project struct {
	Id      string
	Name    string
	Members []*Person
}

var html = `
<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8" />
	<title>Multipart Test</title>
</head>
<body>

<form accept-charset="UTF-8" action="/" enctype="multipart/form-data" id="edit_employee_91919" method="post">

<input name="Employee.FamilyName" size="50" type="text" value="秦" />
<input name="Employee.GivenName" size="50" type="text" value="俊滨" />
<input name="Employee.Photo" type="file" />
<input name="Employee.Resume" type="file" />
<input type="submit" />

</form>

</body>
</html>
`

func mpart(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet == r.Method {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(200)
		io.WriteString(w, html)
		return
	}
	var a *Person
	UnmarshalByPrefix(r, &a, "Person")
	f1, _ := a.Photo.Open()
	f2, _ := a.Resume.Open()

	fc1, _ := ioutil.ReadAll(f1)
	fc2, _ := ioutil.ReadAll(f2)
	fmt.Fprint(w, string(fc1))
	fmt.Fprint(w, string(fc2))
}
func main() {

	http.HandleFunc("/", mpart)
	fmt.Println("sever on port :12345")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
