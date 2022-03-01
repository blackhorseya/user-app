package testdata

import (
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// User1 user 1
	User1 = &user.Profile{
		ID:           primitive.ObjectID{},
		OpenIds:      map[string]string{"line": "line"},
		Name:         "",
		Nickname:     "",
		Email:        "",
		Token:        "token",
		AccessToken:  "",
		RefreshToken: "",
		PictureURL:   "",
		CreatedAt:    0,
		UpdatedAt:    0,
	}
)
