package auth

import (
	"github.com/gin-gonic/gin"
)

var loginPage = `
<html>
  <head>
    <style>
      #credentials-box {
        width: 300px;
        position: absolute;
        bottom: 50%;
        right: 50%;
        transform: translate(50%, 50%);
      }
    </style>
  </head>
  <body>
    <form method="POST">
      <div id="credentials-box">
        <input id="username" type="text" name="username" />
        <input type="password" name="password" />
        <br />
        <input type="submit" value="Login" />
      </div>
    </form>
  </body>
</html>
`

func PasswordAuthentication(credentials map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := c.MustGet("session").(*Session)

		if i_username, isset := session.Get("username"); isset {
			session.Delete("request_url")
			c.Set("username", i_username.(string))
			c.Set("authorized", true)
			return

		} else if i_request_url, isset := session.Get("request_url"); isset {
			username := c.PostForm("username")
			password := c.PostForm("password")
			if username != "" {
				if control_value, isset := credentials[username]; isset && password == control_value {
					session.Set("username", username)
					c.Redirect(303, i_request_url.(string))
					return
				}
			}

		} else {
			session.Set("request_url", c.Request.URL.String())

		}

		c.Writer.Header().Add("Content-Type", "text/html")
		c.String(401, loginPage)
		c.Abort()
	}
}
