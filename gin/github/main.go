package main

import (
	"os"

	"github.com/captncraig/ghauth"
	"github.com/captncraig/temple"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
)

var templateManager temple.TemplateStore

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	// first create the auth handler
	conf := &ghauth.Conf{
		ClientId:     os.Getenv("ac6f19e183a949b77b50"),
		ClientSecret: os.Getenv("866d87172df23008f0ea166126a10578511bbfbd"),
		Scopes:       []string{"user", "read:public_key", "repo"},
		CookieName:   "ghuser",
		CookieSecret: "any random string can go here, but make sure it is truly random and secret to secure your cookies",
	}
	auth := ghauth.New(conf)

	// register oauth routes
	auth.RegisterRoutes("/login", "/authorize", "/logout", r)
	r.Use(auth.AuthCheck())
	// all unauthorized routes can go here. Will still have user populated if logged in
	r.GET("/", home)

	// require authorization for these routes. Will redirect to login if not logged-in
	authRequired := r.Group("/", auth.RequireAuth())
	authRequired.GET("/repo/:owner/:repo", repo)

	r.Run(":80") // listen and serve on 0.0.0.0:8080
}

func home(ctx *gin.Context) {
	user := ghauth.User(ctx)
	var repos = []github.Repository{}
	if user != nil {
		var err error
		opts := &github.RepositoryListOptions{}
		opts.PerPage = 100
		opts.Sort = "pushed"
		opts.Direction = "desc"
		repos, _, err = user.Client().Repositories.List("", opts) //may need to loop through pagination, but whatever.
		if err != nil {
			ctx.Error(err)
			return
		}
	}
	ctx.HTML(200, "main.html", gin.H{
		"User":  user,
		"Repos": repos,
	})
}

func repo(ctx *gin.Context) {
	user := ghauth.User(ctx)
	owner, repo := ctx.Param("owner"), ctx.Param("repo")
	opts := &github.PullRequestListOptions{}

	pulls, _, err := user.Client().PullRequests.List(owner, repo, opts)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.HTML(200, "repo.html", gin.H{
		"User":     user,
		"Pulls":    pulls,
		"FullName": owner + "/" + repo,
	})
}
