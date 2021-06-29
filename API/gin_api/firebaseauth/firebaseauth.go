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

func TestFirebaseToken(c *gin.Context) {
	token := c.Request.Header.Get("token")
	log.Println(token)
	client := SetupFirebase()
	restoken, err := client.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Println("err veryfy id token")
		c.JSON(http.StatusForbidden, gin.H{"res": "bad token"})
		c.Abort()
		return
	}
	log.Println(token)
	log.Printf("veryfied ID toke:%v \n", restoken)
	c.JSON(http.StatusOK, gin.H{
		"res": "OK",
	})
}

func Testid(c *gin.Context) {
	client := SetupFirebase()
	user, err := client.GetUserByEmail(context.Background(), "poor@poor.com")
	if err != nil {
		log.Println("not available user")
	}
	c.JSON(http.StatusOK, gin.H{
		"res": user,
	})
}
