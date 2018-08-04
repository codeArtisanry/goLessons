package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "这是中间件设置的request")
		c.Next()
		fmt.Println("before middleware")
	}
}
func main() {
	router := gin.Default()

	router.GET("/auth/signin",
		func(c *gin.Context) {
			cookie := &http.Cookie{Name: "session_id", Value: "123", Path: "/", HttpOnly: true}
			http.SetCookie(c.Writer, cookie)
			c.String(http.StatusOK, "Login successful")
		})
	router.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note than you are using the copied context "c_cp", IMPORTANT
			message := "Done! in path " + cp.Request.URL.Path
			log.Println(message)
			c.String(http.StatusOK, message)
		}()
	})
	// 全局中间件

	router.Use(MiddleWare())
	{
		router.GET("/middleware", func(c *gin.Context) {
			request := c.MustGet("request").(string)
			req, _ := c.Get("request")
			c.JSON(http.StatusOK, gin.H{
				"middile_request": request,
				"request":         req,
			})
		})
	}
	s := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}
	s.ListenAndServe()
}
