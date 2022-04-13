package main

import (
	"log"

	"github.com/GeorgeHN666/understanding-mongogridfs/db"
	"github.com/GeorgeHN666/understanding-mongogridfs/routes"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("There was an error trying to connect wit database")
	}

	routes.Routes()

}

/*
how to read file

var bucket *gridfs.Bucket //creates a bucket
dbConnection, err := db.GetDBCollection() //connect db with your your
if err != nil {
    log.Fatal(err)
}
bucket, err = gridfs.NewBucket(dbConnection)
if err != nil {
    log.Fatal(err)
}
name := "br100_update.txt"
downloadStream, err := bucket.OpenDownloadStreamByName(name)
if err != nil {
    log.Printf("Failed to open %s: %v", name, err)
    http.Error(w, "something went wrong", http.StatusInternalServerError)
    return
}
defer func() {
    if err := downloadStream.Close(); err != nil {
        log.Fatal(err)
    }
}()

// Use SetReadDeadline to force a timeout if the download does not succeed in
// 2 seconds.
 if err = downloadStream.SetReadDeadline(time.Now().Add(2 * time.Second)); err
  != nil {
  log.Fatal(err)
 }

//  This read the file
fileBuffer := bytes.NewBuffer(nil)
 if _, err := io.Copy(fileBuffer, downloadStream); err != nil {
  log.Fatal(err)

*/
