package storage

import (
	"github.com/developersismedika/sqlx"

	"github.com/ardihikaru/go-chi-example-part-2/pkg/config"
	"github.com/ardihikaru/go-chi-example-part-2/pkg/logger"
	mySqlx "github.com/ardihikaru/go-chi-example-part-2/pkg/mysqldb"
)

// DbConnect opens MySQL database connection
func DbConnect(log *logger.Logger, dbCfg config.DbMySQL) (*sqlx.DB, error) {
	return mySqlx.DbConnect(log, dbCfg)
}
