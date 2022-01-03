package repositories

import (
	"api/src/utils/config"
	"context"

	adminFirestore "cloud.google.com/go/firestore"
	adminFirebase "firebase.google.com/go"
	adminAuth "firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	app *adminFirebase.App
	Ctx = context.Background()
)

type FirebaseApp interface {
	GetFirestore() (*adminFirestore.Client, error)
	GetAuth() (*adminAuth.Client, error)
}

type firebase struct{}

func NewFirebaseApp() FirebaseApp {
	return &firebase{}
}

func initialize() (*adminFirebase.App, error) {
	opt := option.WithCredentialsJSON(config.SERVICE_ACCOUNT)
	conf := &adminFirebase.Config{ProjectID: config.PROJECT_ID}
	firebase, err := adminFirebase.NewApp(Ctx, conf, opt)

	if err != nil {
		return nil, err
	}

	return firebase, nil
}

func GetApp() (*adminFirebase.App, error) {
	var err error

	if app == nil {
		app, err = initialize()

		if err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (firebase *firebase) GetFirestore() (*adminFirestore.Client, error) {
	getApp, err := GetApp()
	if err != nil {
		return nil, err
	}

	firestore, err := getApp.Firestore(Ctx)
	if err != nil {
		return nil, err
	}

	return firestore, nil
}

func (firebase *firebase) GetAuth() (*adminAuth.Client, error) {
	getApp, err := GetApp()
	if err != nil {
		return nil, err
	}

	auth, err := getApp.Auth(Ctx)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
