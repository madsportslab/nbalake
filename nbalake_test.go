package nbalake

import (
	"testing"
)


func TestConnectionNew(t *testing.T) {
	
  ConnectionNew()

	if blobs == nil {
		t.Error("blobs nil")
	}

} // TestConnectionNew
