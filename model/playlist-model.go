package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Playlist struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string            `json:"title,omitempty" bson:"title,omitempty"`
	Author string            `json:"author,omitempty" bson:"author,omitempty"`
	Songs  []string          `json:"songs,omitempty" bson:"songs,omitempty"`
}

func (p *Playlist) IsEmpty() bool {
	return p.Title == "" && len(p.Songs) == 0
}