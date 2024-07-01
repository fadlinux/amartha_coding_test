package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConnection(driver, host string) (conn *sql.DB) {
	conn, err := sql.Open(driver, host)

	if err != nil {
		log.Fatalln(`can't open db:`, err)
		log.Panic("[error NewDBConnection] Failed to init mysql : ", err)
	}

	if err != nil {
		panic(err.Error())
	}

	// err = conn.Ping()
	// if err != nil {
	// 	log.Fatalln(`can't ping db:`, err)
	// }

	return
}
