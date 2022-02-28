package user

import (
	"github.com/blackhorseya/user-app/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Profile declare a user information
type Profile struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	OpenIds      map[string]string  `json:"open_ids" bson:"open_ids"`
	Name         string             `json:"name" bson:"name"`
	Nickname     string             `json:"nickname" bson:"nickname"`
	Email        string             `json:"email" bson:"email"`
	Token        string             `json:"token" bson:"token"`
	AccessToken  string             `json:"access_token" bson:"access_token"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token"`
	PictureURL   string             `json:"picture_url" bson:"picture_url"`
	CreatedAt    primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt    primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

// NewProfileResponse new a *pb.Profile form profile entity
func NewProfileResponse(info *Profile) *pb.Profile {
	return &pb.Profile{
		Id:      info.ID.Hex(),
		Name:    info.Name,
		Email:   info.Email,
		Picture: info.PictureURL,
		Token:   info.Token,
	}
}
