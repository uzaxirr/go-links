package main

import (
	"github.com/gin-gonic/gin"
)

func RootRedirect(c *gin.Context) {
	c.Redirect(302, "https://uzair.biz/")
}

func InstaRedirect(c *gin.Context) {
	c.Redirect(302, "https://www.instagram.com/uzaxirr/")
}

func ResumeRedirect(c *gin.Context) {
	c.Redirect(302, "https://drive.google.com/file/d/1sasOj2nxrrHZ1xad0EztJ8SzZQRGiYSP/view?usp=sharing")
}

func GithubRedirect(c *gin.Context) {
	c.Redirect(302, "https://github.com/uzaxirr/")
}

func LinkedinRedirect(c *gin.Context) {
	c.Redirect(302, "https://www.linkedin.com/in/uzaxirr/")
}

func TwitterRedirct(c *gin.Context) {
	c.Redirect(302, "https://twitter.com/uzaxirr/")
}

func DevRedirect(c *gin.Context) {
	c.Redirect(302, "https://dev.to/uzaxirr/")
}

func LeetcodeRedirect(c *gin.Context) {
	c.Redirect(302, "https://leetcode.com/uzairmq/")
}

func main() {
	server := gin.Default()

	server.GET("/", RootRedirect)
	server.GET("/instagram", InstaRedirect)
	server.GET("/resume", ResumeRedirect)
	server.GET("/cv", ResumeRedirect)
	server.GET("/github", GithubRedirect)
	server.GET("/Github", GithubRedirect)
	server.GET("/GitHub", GithubRedirect)
	server.GET("/lkdn", LinkedinRedirect)
	server.GET("/linkedin", LinkedinRedirect)
	server.GET("/twitter", TwitterRedirct)
	server.GET("/Twitter", TwitterRedirct)
	server.GET("/dev", DevRedirect)
	server.GET("/dev.to", DevRedirect)
	server.GET("/leetcode", LeetcodeRedirect)

	server.Run()
}
