package connection

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// id := flag.Int64("id", 0, "person ID to find")
	dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")
	flag.Parse()

	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}
	// if *id == 0 {
	// 	log.Fatal("missing person ID")
	// }

	// Opening a driver typically will not attempt to connect to the database.
	DB, err = sql.Open("postgres", *dsn)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}

	DB.SetConnMaxLifetime(0)
	DB.SetMaxIdleConns(3)
	DB.SetMaxOpenConns(3)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := DB.PingContext(ctx); err != nil {
		log.Fatal("Database not reachable:", err)
	}

}

// Query the database for the information requested and prints the results.
// If the query fails exit the program with an error.
func QueryData(ctx context.Context, query string) ([]map[string]interface{}, error) {
	rows, err := DB.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("query error:", err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	log.Println("Columns:", cols)

	// Buat slice untuk scan data (tipe []interface{})
	values := make([]interface{}, len(cols))
	valuePtrs := make([]interface{}, len(cols))
	var result []map[string]interface{}

	for rows.Next() {
		for i := range cols {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			log.Fatal("scan error:", err)
		}

		// Cetak data per baris
		rowData := make(map[string]interface{})
		for i, col := range cols {
			val := values[i]

			// Coba convert []byte ke string jika bisa
			b, ok := val.([]byte)
			if ok {
				rowData[col] = string(b)
			} else {
				rowData[col] = val
			}
		}
		result = append(result, rowData)
	}

	return result, nil
}
