//Package profilehandler keeps all logic behind http requests and responses
package profilehandler
import(
	"net/http"
	"sync"
	"github.com/tahasevim/webserver/templates"
)

//ProfileHandler is a handler type that handles all incoming requests and responds them
type ProfileHandler struct {
	Entries map[string]func(http.ResponseWriter, *http.Request)//register all functions to a url
	Users	[]*Person //holds all user of application
	Mu sync.Mutex // protects Entries
}

//Person is a basic type that is created when a signup operation is performed
type Person struct {
	Username string //name of person
	Password string //password of person
	Email    string //email os person
}

//NewProfileHandler simply create a instance of type NewProfileHandler
//Before it returns,it registers all urls to corresponding functions
func NewProfileHandler()*ProfileHandler{

	m := &ProfileHandler{
		Entries:make(map[string]func(http.ResponseWriter,*http.Request)),
	}
	m.Entries["/signup/"] = m.SignupHandler
	m.Entries["/"] = m.IndexHandler
	m.Entries["/save/"] = m.SaveHandler
	m.Entries["/login/"] = m.LoginHandler
	m.Entries["/profiles/"] = m.ProfilesHandler
	return m
}

//FindUserFromUserName finds a user according to its argument that is username of target user
//If there is no such a user,it returns nil.
func (m *ProfileHandler) FindUserFromUserName(username string) *Person {
	for _, v := range m.Users {
		if v.Username == username {
			return v
		}
	}
	return nil
}

//SignupHandler is a simple handler of "/signup/"
func (m *ProfileHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	templates.TmplSignup.ExecuteTemplate(w, "signup.html",nil)
}

//IndexHandler is a simple handler of "/"
func (m *ProfileHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	templates.TmplIndex.ExecuteTemplate(w, "index.html", nil)
}

//LoginHandler is a handler of login operations
//It will be executed when a user attempt to login
func (m *ProfileHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	for _, u := range m.Users {
		if u.Username == username && u.Password == password {
			http.Redirect(w, r, "/profiles/"+u.Username, http.StatusFound)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

//SaveHandler will handle creation of a user
//When a user signs up,it will be executed
func (m *ProfileHandler) SaveHandler(w http.ResponseWriter, r *http.Request) {
	username, password, email := r.FormValue("username"), r.FormValue("password"), r.FormValue("email")
	user := &Person{username, password, email}
	m.Users = append(m.Users, user)
	m.Mu.Lock()
	defer m.Mu.Unlock()
	m.Entries["/profiles/"+user.Username] = m.ProfilesHandler
	http.Redirect(w, r, "/", http.StatusFound)
	return
}

//ProfilesHandler handle two types url
//If request url ends with only "/profiles/" it show all links of signed up Users
//If request url includes a username such as "/profiles/username", it will show that user's page
func (m *ProfileHandler) ProfilesHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/profiles/"):]
	if url := r.URL.Path; len(url) > len("/profiles/") {
		if user := m.FindUserFromUserName(username); user != nil {
			templates.TmplProfile.ExecuteTemplate(w, "profiles.html", user)
			return
		}
	} else {
		for _, s := range m.Users {
			templates.TmplProfileList.ExecuteTemplate(w, "profileList.html", s)
		}
	}
}

//ServeHTTP function main handler of all this webserver.It will call functions according to typed url
//This function implementation will allow to ProfileHandler type to be a handler
func (m *ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handle, ok := m.Entries[r.URL.String()]; ok {
		handle(w, r)
	}
}
