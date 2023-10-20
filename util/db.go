package util

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

type DBConn struct {
	db *gorm.DB
}

var dbInstance *DBConn
var dbOnce sync.Once

// GetDBConnection get DB connection
func NewDBConnection(db DB) *gorm.DB {

	dbOnce.Do(func() {
		log.Infoln("Initialize database connection...")

		connStr := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s"
		psqlconn := fmt.Sprintf(connStr, db.Host, db.Port, db.Username,
			db.Password, db.DBName, db.DBSchemaName)

		db, err := gorm.Open("postgres", psqlconn)
		if err != nil {
			log.Println(err)
		}

		/**
		* NOTES: this will set connection lifetime in connection pool to 1 minute.
		* 		  If the connection in the pool is idle > 1 min, Golang will close it
		* 		  and will create new connection if #connections in the pool < pool max num
		* 		  of connection. This to avoid invalid connection issue
		 */

		db.DB().SetConnMaxLifetime(time.Second * 60)

		if err != nil {
			panic(err)
		}

		dbInstance = &DBConn{
			db: db,
		}

		log.Infoln("Database has been initialized")
	})

	return dbInstance.db
}
