package entities

import (
	"context"
	"time"
)

type Feed struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
  CreatedAt time.Time `json:"created_at"`
}

type FeedModel interface {
  Insert(feed Feed) error
  GetAll(ctx context.Context) ([]*Feed, error)
}
