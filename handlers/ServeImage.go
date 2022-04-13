package handlers

import (
	"net/http"

	"github.com/GeorgeHN666/understanding-mongogridfs/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func ServeImage(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	finalID, _ := primitive.ObjectIDFromHex(id)

	db := db.DBConnect

	// Create new bucket
	bucket, _ := gridfs.NewBucket(db.Database("test"))

	// Here we download the file and serve it to the response writer
	_, _ = bucket.DownloadToStream(finalID, w)

}
