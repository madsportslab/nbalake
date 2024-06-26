package nbalake

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

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
		Creds:  credentials.NewStaticV4(viper.GetString(
		NBALAKE_AUTH_KEY), viper.GetString(NBALAKE_AUTH_SECRET),
		""), Secure: false,
	})

	if err != nil {
		log.Fatal(err)
	}

	blobs = c

	ctx = context.Background()

} // ConnectionNew


func BucketName(y string, id string) string {
	return fmt.Sprintf("%s.%s", y, id)
} // BucketName


func Exists(b string, k string) bool {
  
	_, err := blobs.StatObject(ctx, b, k,
		minio.GetObjectOptions{})

	if err != nil {
		return false
	} else {
		return true
	}

} // Exists


func List(b string) <-chan minio.ObjectInfo {
	
	return blobs.ListObjects(ctx, b,
		minio.ListObjectsOptions{})

} // List


func Put(b string, k string, r []byte) {

	buf := bytes.NewReader(r)

	_, err := blobs.PutObject(ctx, b, k, buf, int64(buf.Len()),
	  minio.PutObjectOptions{ContentType: CONTENT_TYPE_JSON}) 

	if err != nil {
		log.Println(err)
	}

} // Put


func PutFile(b string, k string) {

	_, err := blobs.FPutObject(ctx, b, k, k, minio.PutObjectOptions{
		ContentType: CONTENT_TYPE_JSON})

	if err != nil {
		log.Println(err)
	}

} // PutFile


func Get(b string, f string) []byte {

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

} // Get


func ParseDate(f string) string {

	if len(f) == 0 {
		return STR_EMPTY
	}

	tokens := strings.Split(f, NBALAKE_DELIMITER)

	if len(tokens) == 0 {
		return STR_EMPTY
	} else {
		return tokens[0]
	}

} // ParseDate


func GetLatestLeaders(b string) string {

	var key, latest string

	blobs := List(b)

	for blob := range blobs {

		log.Println(blob.Key)

		if strings.Contains(blob.Key, PREFIX_LEADERS) {

			tokens := strings.Split(blob.Key, STR_PERIOD)

			d := tokens[0]

			if len(latest) == 0 {

				key 			= blob.Key
				latest 		= d

			} else {
		
				if latest < d {

					key 		= blob.Key
					latest 	= d

				}				

			}
	
		}

	}

	return key

} // GetLatestLeaders
