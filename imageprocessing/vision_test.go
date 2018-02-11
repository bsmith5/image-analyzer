package google

import (
	"testing"
	"golang.org/x/net/context"
	vision "cloud.google.com/go/vision/apiv1"
	"log"
	"os"
)


func TestDetectLabelsFile(t *testing.T) {
	// Setup the google client
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// these should be the same image so should theoretically return identical labels, for now just check for no error
	DetectLabelsFile(ctx, *client, os.Stdout, "../testdata/Elmer_Keith.jpg")
}


func TestDetectLabelsURI(t *testing.T) {
	// Setup the google client
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	DetectLabelsURI(ctx, *client, os.Stdout, "https://upload.wikimedia.org/wikipedia/en/thumb/d/dc/Elmer_Keith.jpg/220px-Elmer_Keith.jpg")

}
