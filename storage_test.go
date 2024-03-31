package nbalake

import (
  "testing"
)


func init() {

	ConnectionNew()

} // init

func TestGetLastDate(t *testing.T) {

	t.Log(GetLastDate(BucketName("2023", BUCKET_ANALYTICS)))

} // TestGetLastDate
