package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ms-xy/logtools"
)

type Session struct {
	SessionID string
	Start     time.Time
	Data      map[string]interface{}
}

func (this *Session) Set(key string, value interface{}) (interface{}, bool) {
	prevData, existed := this.Data[key]
	this.Data[key] = value
	return prevData, existed
}

func (this *Session) Get(key string) (interface{}, bool) {
	data, exists := this.Data[key]
	return data, exists
}

func (this *Session) Delete(key string) (interface{}, bool) {
	data, existed := this.Data[key]
	delete(this.Data, key)
	return data, existed
}

func Sessions() gin.HandlerFunc {
	sessions := make(map[string]*Session)

	return func(c *gin.Context) {
		var session *Session

		if sid, err := c.Cookie("sid"); err == nil {
			session, _ = sessions[sid]
		}

		if session == nil {
			session = &Session{
				SessionID: uuid.Must(uuid.NewUUID()).String(),
				Start:     time.Now(),
				Data:      make(map[string]interface{}),
			}
			sessions[session.SessionID] = session
			// name: sid
			// value: uuid4-string
			// maxAge: none
			// path: /
			// domain: current
			// secure: false
			// httpOnly: true
			c.SetCookie("sid", session.SessionID, 0, "/", "", false, true)
		}

		logtools.Debugf("session: %+v", session)

		c.Set("session", session)

		c.Next()
	}
}
