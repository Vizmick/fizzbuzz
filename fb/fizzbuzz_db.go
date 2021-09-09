package fb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type FizzbuzzRequest struct {
	id    int64
	Int1  int    `form:"int1" json:"int1"`
	Int2  int    `form:"int2" json:"int2"`
	Limit int    `form:"limit" json:"limit"`
	Str1  string `form:"str1" json:"str1"`
	Str2  string `form:"str2" json:"str2"`
}

type FizzbuzzResponse struct {
	Int1  int    `form:"int1" json:"int1"`
	Int2  int    `form:"int2" json:"int2"`
	Limit int    `form:"limit" json:"limit"`
	Str1  string `form:"str1" json:"str1"`
	Str2  string `form:"str2" json:"str2"`
	Count int    `json:"count"`
}

func mostFrequentRequest() (FizzbuzzResponse, error) {
	var fbResponse FizzbuzzResponse

	row := db.QueryRow("SELECT `int1`, `int2`, `limit`, `str1`, `str2`,count(*) FROM fb_request GROUP BY `int1`, `int2`, `limit`, `str1`, `str2` ORDER BY count(*) DESC LIMIT 1")
	if err := row.Scan(&fbResponse.Int1, &fbResponse.Int2, &fbResponse.Limit, &fbResponse.Str1, &fbResponse.Str2, &fbResponse.Count); err != nil {
		if err == sql.ErrNoRows {
			return fbResponse, fmt.Errorf("no requests")
		}
		return fbResponse, err
	}
	return fbResponse, nil
}

func addRequest(fbRequest FizzbuzzRequest) (int64, error) {
	result, err := db.Exec("INSERT INTO fb_request (`int1`, `int2`, `limit`, `str1`, `str2`) values (?,?,?,?,?)",
		fbRequest.Int1, fbRequest.Int2, fbRequest.Limit, fbRequest.Str1, fbRequest.Str2)
	if err != nil {
		return 0, fmt.Errorf("fbRequest :%v", fbRequest)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
func init() {
	cfg := mysql.Config{
		User:                 "victor@localhost",
		Passwd:               "miam_du_chocolat",
		AllowNativePasswords: true,
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "fizzbuzz",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}
