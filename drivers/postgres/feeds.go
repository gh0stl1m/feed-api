package postgres

import (
	"context"
	"database/sql"

	"github.com/gh0stl1m/feed-api/domains/entities"
  logger "github.com/sirupsen/logrus"
)

type FeedRepository struct {
  db *sql.DB
}

func NewFeedClient(db *sql.DB) entities.FeedModel {

  return &FeedRepository{ db }
}

func (fr *FeedRepository) Insert(feed entities.Feed) error {

  query := `INSERT INTO feeds(title, description, created_at) VALUES ($1, $2, $3)`

  _, err := fr.db.Exec(query, feed.Title, feed.Description, feed.CreatedAt)

  return err
}

func (fr *FeedRepository) GetAll(ctx context.Context) ([]*entities.Feed, error) {

  query := `SELECT title, description, created_at FROM feeds`
  rows, err := fr.db.Query(query)

  if err != nil {

    logger.WithFields(logger.Fields{
      "error": err,
    }).Error("Something went wrong reading all feeds")

    return nil, err
  }

  defer rows.Close()

  var feeds []*entities.Feed

  for rows.Next() {

    feed := &entities.Feed{}

    err := rows.Scan(&feed.Title, &feed.Description, &feed.CreatedAt)

    if err != nil {

      logger.WithFields(logger.Fields{
        "error": err,
      }).Error("Something went wrong reading row")

      return nil, err
    }

    feeds = append(feeds, feed)

  }

  if err := rows.Err(); err != nil {
  
    logger.WithFields(logger.Fields{
      "error": err,
    }).Error("Error reading row")
     
    return nil, err
  }

  return feeds, nil
}
