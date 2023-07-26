package postgres

import (
  "database/sql"
  logger "github.com/sirupsen/logrus"
)

func newConnection(uri string) *sql.DB {

  db, err := sql.Open("postgres", uri)

  if err != nil {

    logger.WithFields(logger.Fields{
      "error": err,
    }).Fatal("Connection failed")
  }

  defer db.Close()

  return db
}
