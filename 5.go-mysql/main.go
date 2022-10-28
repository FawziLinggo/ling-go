package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getConnectionMysql() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(server2:3306)/crawling_mongo")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func sentToMongo(url string, judul string, penulis string, waktu_publish string, isi_berita string) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://server2:33202/?connect=direct"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.PrimaryPreferred())
	if err != nil {
		panic(err)
	}
	collection := client.Database("crawling").Collection("kompas")
	// defer client.Disconnect(ctx)
	res, err := collection.InsertOne(ctx, bson.D{{Key: "url", Value: url}, {Key: "judul", Value: judul}, {Key: "penulis", Value: penulis}, {Key: "waktu_publish", Value: waktu_publish}, {Key: "isi_berita", Value: isi_berita}})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("succes to mongoDB wiht : %v\n", res.InsertedID)

	return res.InsertedID
}

func main() {

	db := getConnectionMysql()
	defer db.Close()

	ctx := context.Background()

	// INSERT
	// result, err := db.ExecContext(ctx, "INSERT INTO kompas(nomor,url,judul,penulis,waktu_publish,isi_berita) VALUES ('16','https://www.kompas.com/properti/read/2022/10/27/083000621/raja-juli-serahkan-sertifikat-tanah-milik-yayasan-gus-dur','Raja Juli Serahkan Sertifikat Tanah Milik Yayasan Gus Dur','Suhaiela Bahfein','2022-10-27 08:30:00','JAKARTA, KOMPAS.com');")
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Println("Success insert data to database ", result)
	nomor_ := 1
	for {
		script := "select nomor,url,judul,penulis,waktu_publish,isi_berita from kompas where nomor = ? limit 1"
		rows, err := db.QueryContext(ctx, script, nomor_)
		if err != nil {
			panic(err)
		}

		if rows.Next() {
			var (
				nomor         int
				url           string
				judul         string
				penulis       string
				waktu_publish string
				isi_berita    string
			)
			err := rows.Scan(&nomor, &url, &judul, &penulis, &waktu_publish, &isi_berita)
			if err != nil {
				panic(err)
			}
			sentToMongo(url, judul, penulis, waktu_publish, isi_berita)
			nomor_ += 1
			rows.Close()

			// fmt.Println("nomor :", nomor)
			// fmt.Println("url :", url)
			// fmt.Println("judul :", judul)
			// fmt.Println("penulis :", penulis)
			// fmt.Println("waktu_publish :", waktu_publish)
			// fmt.Println("isi_berita :", isi_berita)

		}
	}

}
