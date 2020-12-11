package server

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Koneksi()(*sql.DB,error)  {
	db,err := sql.Open("mysql","root:@tcp(localhost:3306)/db_pizza")
	if err != nil {
		return nil,err
		
	}
	return db,nil
	
}