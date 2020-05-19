package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/farnetani/exemplo-rotas-simples/utils"
)

func Connection() *sql.DB {
	var db, err = sql.Open("mysql", "root:farsoft01@/go_course?charset=utf8")
	utils.CheckErr(err)
	return db
}