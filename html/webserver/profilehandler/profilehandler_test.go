package profilehandler

import(
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestIndexHandler(t *testing.T){
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET","/",nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := NewProfileHandler()
	handler.IndexHandler(w, r)
	if w.Code != http.StatusOK{
		t.Fatal("Index handler doesnt work")
	}
}
func TestProfilesHandler(t *testing.T){
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET","/profiles/",nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := NewProfileHandler()
	handler.ProfilesHandler(w, r)
	if w.Code != http.StatusOK{
		t.Fatal("Profiles handler doesnt work")
	}
}
func TestSignupHandler(t *testing.T){
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET","/signup/",nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := NewProfileHandler()
	handler.SignupHandler(w, r)
	if w.Code != http.StatusOK{
		t.Fatal("Signup handler doesnt work")
	}
}