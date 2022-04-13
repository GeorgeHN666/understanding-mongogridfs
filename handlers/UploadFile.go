package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/GeorgeHN666/understanding-mongogridfs/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var FILEID primitive.ObjectID

func UploadFileEndPoint(w http.ResponseWriter, r *http.Request) {
	// Here we take the file as multipart
	_, handler, _ := r.FormFile("image")

	uploadFile(handler, handler.Filename, w)

}

func uploadFile(file *multipart.FileHeader, filename string, w http.ResponseWriter) {
	// open the file
	filecontent, _ := file.Open()

	// Converted to []bytes
	dataFile, _ := ioutil.ReadAll(filecontent)

	// db.DBConnect connects with your database
	db := db.DBConnect

	// We set the metadata
	uploadopt := options.GridFSUpload().SetMetadata(bson.D{{"metadata tag", "tag"}})

	// we create a new bucket and set the database to create the collections fs.files , fs.chunks
	bucket, _ := gridfs.NewBucket(
		db.Database("test"),
	)

	// Here we upload the file passing the filename(you can change the name)
	// Also we pass the datafile that is a slice of bytes and create a new buffer
	// And finally we pass the metadata
	FILEID, _ = bucket.UploadFromStream(filename, bytes.NewBuffer(dataFile), uploadopt)

	log.Println("New File Created, The file size::", file.Size)
	// we print the id of the file
	// with this id we're gonna be able to download the file
	log.Println("This is the ID::", FILEID)

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	// In this function we pass the response writer and the FileID
	json.NewEncoder(w).Encode(FILEID.Hex())

}

/*
func UploadFileEndPoint(w http.ResponseWriter, r *http.Request) {

	_, filename, _ := r.FormFile("image")

	da,_ := os.Create(filename.Filename)



	uploadFile(filename, filename.Filename, w)

}

func uploadFile(f *multipart.FileHeader, filename string, w http.ResponseWriter) {

	filecontent, _ := f.Open()

	data, _ := ioutil.ReadAll(filecontent)

	db := db.DBConnect

	// Here we initialize the Bucket
	bucket, err := gridfs.NewBucket(
		db.Database("test"),
	)
	if err != nil {
		log.Println("Error in ln 55  error::", err.Error())
		return
	}

	write, err := bucket.OpenUploadStream(filename)
	if err != nil {
		log.Println("Error ln::57 error::", err.Error())
		return
	}

	filesize, err := write.Write(data)
	if err != nil {
		log.Println("Error in ln 63  error::", err.Error())
	}

	w.WriteHeader(http.StatusCreated)

	log.Printf("Write File succesfully to database, filesize::%d", filesize)

}

*/
