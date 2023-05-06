package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

// ginのengine設定
// TODO: ここで与えられた引数によって、loggerやrecoveryなどのmiddlewareが設定されるようにしたい
func SetupRouter() *gin.Engine {
	var router = gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", ping)
		v1.GET("/echoq", echoq)
		v1.GET("/clock", testClock)
		v1.PATCH("/clock", testClock)
		v1.POST("/user", subscribeUserAccountInfo)
		v1.GET("/user", getUserAccountinfo)
	}
	return router
}

// ginのmode設定
// 環境変更のENVがdevの場合はginのmodeをdebugに、それ以外の場合はreleaseに設定する
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
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "response": "pong"})
}

// 入力されたクエリをそのままの形で返却
func echoq(c *gin.Context) {
	// リクエストのクエリが定義された構造体
	var echoschema Testtype
	if c.ShouldBind(&echoschema) == nil {
		s := echoschema.Themename + "\n"
		c.JSON(http.StatusOK, gin.H{"echo": s})
	} else {
		log.Printf("Happend error where parse test query")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "happend a server error"})
	}
}

// リクエストクエリのうちnameとfamily_nameをdbに登録するための関数
func subscribeUserAccountInfo(c *gin.Context) {
	// リクエストのクエリが定義された構造体
	var user User
	// TODO: User構造体のbirthdayやidの値がdefaultで設定されるので、
	// ShouldBind()だとnameなどの他の値が空欄でバインドされてしまう。
	// 上記の制限を回避する実装を行う
	if c.ShouldBind(&user) == nil {
		// リクエストが送信されたら、固定された値を返却する
		c.JSON(http.StatusOK, gin.H{"name": user.Name, "family_name": user.Family_name, "from": user.From, "address": user.Address, "birthday": user.Birthday, "id": user.Id})
	} else {
		log.Printf("Happend error when get user information")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "happend a server error"})
	}
}

// UserAccountinfoを取得するための関数
func getUserAccountinfo(c *gin.Context) {
	// リクエストのクエリが定義された構造体
	var user User
	if c.ShouldBind(&user) == nil {
		c.JSON(http.StatusOK, gin.H{"name": "test01", "family_name": "test", "from": "Japan", "address": "Japan/Tokyo", "birthday": "2023-05-02", "id": 1})
	} else {
		log.Printf("Happend error when get user information")
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
	godotenv.Load()
	log.Println("INFO: Load .env file")

	SetupGinmode()

	// dbの接続を確立するための処理
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	endpoint := os.Getenv("DB_ENDPOINT")
	database := os.Getenv("DB_DATABASE")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, endpoint, database))
	if err != nil && gin.Mode() == gin.ReleaseMode {
		log.Fatalf("ERROR: Failed to connect to database")
	}
	// dbの接続が完了した場合は,deferでdbを閉じる処理を遅らせる
	defer db.Close()

	router := SetupRouter()
	port := os.Getenv("PORT")
	log.Printf("INFO: Waite for request on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Printf("server didnt start")
	}
}
