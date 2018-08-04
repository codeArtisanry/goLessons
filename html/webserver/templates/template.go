//Package templates contains all HTML templates of webserver
package templates

import (
	"html/template"
)

var (
  //TmplSignup is a HTML template which will be shown in /signup/ url
  TmplSignup = template.Must(template.New("signup.html").Parse(`<html>
  <head></head>
  <body>
      <form action ="/save/" method="POST">
      <div class ="container">
        Username:<input type ="text" name="username"><br>
        Password:<input type ="password" name="password"><br>
        Email:<input type ="text" name="email"><br>
        <input type="submit" value="Sign up!">
        <button type="submit" name="cancel" formaction="/" class="signup">Cancel</button>
      </div>
      </form>
  </body>
  </html>`))

  //TmplIndex is aHTML template which will be shown in / url
	TmplIndex = template.Must(template.New("index.html").Parse(`<html>
    <head></head>
      <body>
        <form action="/login/" method="POST">
         <div class="container">
          <label><b>Username</b></label>
          <input type="text" placeholder="Enter Username" name="username">

          <label><b>Password</b></label>
          <input type="password" placeholder="Enter Password" name="password">

          <button type="submit">Login</button>
  </div>

    <button type="submit" name="signup" formaction="/signup/" class="signup">Sign up!</button>
</form>
</body>
</html>`))

  //TmplProfile is a HTML template which will be shown in /profiles/username url
  TmplProfile = template.Must(template.New("profiles.html").Parse(`<html>
  <head></head>
    <body>
        <form action = "/" method ="POST">
        <div><b>Username:</b>{{.Username}}<br>
            <b>Password:</b>000000000<br>
            <b>Email:</b> {{.Email}}</div>
        <input type ="submit" value="Back" name ="back"> 
        </form>
    </body>
    </html>`))
  
  //TmplProfileList is a HTML template which will be shown in /profiles/ url
  TmplProfileList = template.Must(template.New("profileList.html").Parse(`
    <a href ="/profiles/{{.Username}}">{{.Username}}</a>
`))
)