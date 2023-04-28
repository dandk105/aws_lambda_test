package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	// local package
	"dandk105/aws_lambda_test/controller"
	"dandk105/aws_lambda_test/schema"
)

// ginのengine設定
// TODO: ここで与えられた引数によって、loggerやrecoveryなどのmiddlewareが設定されるようにしたい
func SetupRouter() *gin.Engine {
	var router = gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/test", test)
		v1.GET("/echoq", echoq)
		v1.GET("/clock", testClock)
		v1.PATCH("/clock", testClock)
	}
	return router
}

// ginのmode設定
func SetupGinmode() {
	if os.Getenv("ENV") == "dev" {
		gin.SetMode(gin.DebugMode)
		log.Printf("INFO: Set a server mode as development mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
		log.Printf("INFO: Set a server mode as release mode")
	}
}

// 固定文字を返却
func test(c *gin.Context) {
	c.JSON(http.StatusOK, "Test")
}

// 入力されたクエリをそのままの形で返却
func echoq(c *gin.Context) {
	// リクエストのクエリが定義された構造体
	var theme schema.Testtype
	if c.ShouldBind(&theme) == nil {
		s := theme.Themename + "\n"
		c.JSON(http.StatusOK, gin.H{"echo": s})
	} else {
		log.Printf("Happend error where parse test query")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "happend a server error"})
	}
}

// リクエストクエリのうちnameとfamily_nameをdbに登録するための関数
func subscriveUserinformation(c *gin.Context) {
	// リクエストのクエリが定義された構造体
	var user schema.User
	if c.ShouldBind(&user) == nil {
		// リクエストのクエリをdbに登録
		//s3.AddUser(user)
		//c.Request.Response(http.StatusOK)
	} else {
		log.Printf("Happend error when subscrive user information")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "happend a server error"})
	}
}

// リクエスト時の日付を返却。標準ライブラリのtime packageをそのまま使用
func testClock(c *gin.Context) {
	var t = time.Now()
	c.JSON(http.StatusOK, t)
}

func main() {
	// 環境変数の読み込み
	controller.LoadEnvfile()

	SetupGinmode()
	router := SetupRouter()
	port := os.Getenv("PORT")
	if err := router.Run(":" + port); err != nil {
		log.Printf("server didnt start")
	}

}
