package nbalake

import (
  "testing"
)


func init() {

	ConnectionNew()

} // init


func TestGetLatestLeaders(t *testing.T) {

	t.Log(GetLatestLeaders(BucketName("2023", BUCKET_ANALYTICS)))

} // TestGetLatestLeaders
