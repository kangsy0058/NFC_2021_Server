package firebaseauth

import (
	"context"
	"log"
	"net/http"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// firebase 세팅용함수 retun firebase Cilent객체
func SetupFirebase() *auth.Client {
	serviceAccountKeyFilePath, err :=
		filepath.Abs("./firebaseauth/nfc-2021-firebase-adminsdk-25wsf-b7931547ae.json")
	if err != nil {
		log.Fatal("Unable to load dserviceAccountkey json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	config := &firebase.Config{ProjectID: "nfc-2021"}
	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("err initializing app: %v \n", err)
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	return auth
}

// API 요청시 중간에서 firebase ID token 을 확인하는 Middleware
func FirebaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)

		// firebase 토큰 확인
		restoken, err := firebaseAuth.VerifyIDToken(context.Background(), token)
		// 토큰값이 이상하다면
		if err != nil {
			// Forbiddem status return 후 종료
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		// 토큰 확인 완료시 진행
		// token := c.MustGet("token").(*auth.Token)
		c.Set("token", restoken)
		c.Next()
	}
}

// admin API 요청시 중간에서 firebase Id token과 admin권한을 확인하는 middleware
