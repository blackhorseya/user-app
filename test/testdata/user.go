package testdata

import (
	"github.com/blackhorseya/user-app/internal/pkg/entity/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// User1 user 1
	User1 = &user.Profile{
		ID:           primitive.ObjectID{},
		OpenIds:      nil,
		Name:         "",
		Nickname:     "",
		Email:        "",
		Token:        "",
		AccessToken:  "",
		RefreshToken: "",
		PictureURL:   "",
		CreatedAt:    0,
		UpdatedAt:    0,
	}
)
