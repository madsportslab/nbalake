package nbalake

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
  "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var blobs *minio.Client

var ctx context.Context


func InitBuckets(buckets []string) {

	for _, b := range buckets {

		ok, err := blobs.BucketExists(ctx, b)

		if err != nil {
			log.Println(err)
		}
	
		if !ok {
			blobs.MakeBucket(ctx, b, minio.MakeBucketOptions{})
		}
	
	}

} // InitBuckets


func ConnectionNew() {

	parseConfig()

	c, err := minio.New(GetServerAddr(), &minio.Options{
		Creds:  credentials.NewStaticV4(viper.GetString(BLOB_AUTH_KEY),
		viper.GetString(BLOB_AUTH_SECRET), ""),
		Secure: false,
	})

	if err != nil {
		log.Fatal(err)
	}

	blobs = c

	ctx = context.Background()

} // ConnectionNew


func BucketName(y string, id string) string {
	return fmt.Sprintf("%s.nba.%s", y, id)
} // BucketName


func BlobList(b string) <-chan minio.ObjectInfo {
	
	return blobs.ListObjects(ctx, b,
		minio.ListObjectsOptions{})

} // BlobList


func BlobPut(b string, k string, r []byte) {

	buf := bytes.NewReader(r)

	_, err := blobs.PutObject(ctx, b, k, buf, int64(buf.Len()),
	  minio.PutObjectOptions{ContentType: CONTENT_TYPE_JSON}) 

	if err != nil {
		log.Println(err)
	}

} // BlobPut


func BlobPutFile(b string, k string) {

	_, err := blobs.FPutObject(ctx, b, k, k, minio.PutObjectOptions{
		ContentType: CONTENT_TYPE_JSON})

	if err != nil {
		log.Println(err)
	}

} // BlobPutFile


func BlobGet(b string, f string) []byte {

	o, err := blobs.GetObject(ctx, b, f, minio.GetObjectOptions{})

	if err != nil {
		log.Println(err)
	} else {

		info, err := o.Stat()

		if err != nil {
			log.Println(err)
		} else {

			buf := make([]byte, info.Size)
	
			o.Read(buf)
	
			return buf
	
		}

	}

	return nil

} // BlobGet