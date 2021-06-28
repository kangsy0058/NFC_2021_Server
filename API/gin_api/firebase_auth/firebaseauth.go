package firebaseauth

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func SetupFirebase() *auth.Client {
	serviceAccountKeyFilePath, err :=
		filepath.Abs("./hoseo-covid-nfc-firebase-adminsdk-9tm7x-fde7ac1723.json")
	if err != nil {
		panic("Unable to load dserviceAccountkey json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Firebse load error")
	}
	return auth
}
