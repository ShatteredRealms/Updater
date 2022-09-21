package srv

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func NewServer() *gin.Engine {
    r := gin.Default()

    v1 := r.Group("/v1")
    v1.GET("/health", health)

    return r
}

func health(c *gin.Context) {
    c.JSON(http.StatusOK, "ok")
}
