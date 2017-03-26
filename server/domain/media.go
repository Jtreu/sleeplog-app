package domain

type MediaRepository interface {
	Create(media Media) (Media, error)
}

type Media struct {
	UID   string `json:"uid" bson:"uid"`
	Type  string `json:"type" bson:"uid"` // image, video
	Owner struct {
		UID  string `json:"uid" bson:"uid"`
		Type string `json:"type" bson:"type"` // user, game
	} `json:"owner" bson:"owner"`
	Source string      `json:"source" bson:"source"`
	Meta   interface{} `json:"meta,omitempty" bson:"meta"`
}

const (
	MediaOwnerUser = "user"
	MediaOwnerGame = "game"
)
