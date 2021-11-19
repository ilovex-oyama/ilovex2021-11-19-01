package main

import(
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")
  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "/",
      "h1": "/",
    })
  })
  router.GET("/hoge", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "hoge",
      "h1": "/hoge",
    })
  })
  router.GET("/fuga", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "fuga",
      "h1": "/",
    })
  })
  router.Run()
}
