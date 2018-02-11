package google

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	vision "cloud.google.com/go/vision/apiv1"
	"log"
)


func TestDetectLabelsFile(t *testing.T) {
	// Setup the google client
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	file_labels, err := DetectLabelsFile(ctx, *client, "../testdata/Elmer_Keith.jpg")
	//assert.Contains(t, file_labels, "cowboy")
	assert.NotEmptyf(t, file_labels, "file_labels should not be empty")
	assert.Nil(t, err)
}


func TestDetectLabelsURI(t *testing.T) {
	// Setup the google client
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uri_labels, err := DetectLabelsURI(ctx, *client, "https://upload.wikimedia.org/wikipedia/en/thumb/d/dc/Elmer_Keith.jpg/220px-Elmer_Keith.jpg")
	assert.NotEmptyf(t, uri_labels, "uri_labels should not be empty")
	assert.Nil(t, err)
}
