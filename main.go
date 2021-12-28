package main

import (
	"github.com/gin-gonic/gin"
)

func RootRedirect(c *gin.Context) {
	c.Redirect(302, "https://uzair.biz/")
}

func ResumeRedirect(c *gin.Context) {
	c.Redirect(302, "https://drive.google.com/file/d/1zYg0T3f8lLLpgewshXUvxU1Ei-KbqhWZ/view?usp=sharing")
}

func GithubRedirect(c *gin.Context) {
	c.Redirect(302, "https://github.com/uzair-ali10")
}

func LinkedinRedirect(c *gin.Context) {
	c.Redirect(302, "https://www.linkedin.com/in/uzair-ali10")
}

func TwitterRedirct(c *gin.Context) {
	c.Redirect(302, "https://twitter.com/UzairAli101")
}

func DevRedirect(c *gin.Context) {
	c.Redirect(302, "https://dev.to/uzairali10/")
}

func LeetcodeRedirect(c *gin.Context) {
	c.Redirect(302, "https://leetcode.com/uzair-ali10/")
}

func main() {
	server := gin.Default()

	server.GET("/", RootRedirect)
	server.GET("/resume", ResumeRedirect)
	server.GET("/cv", ResumeRedirect)
	server.GET("/github", GithubRedirect)
	server.GET("/lkdn", LinkedinRedirect)
	server.GET("/linkedin", LinkedinRedirect)
	server.GET("/twitter", TwitterRedirct)
	server.GET("/dev", DevRedirect)
	server.GET("/leetcode", LeetcodeRedirect)

	server.Run()
}
