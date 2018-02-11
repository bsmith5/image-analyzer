package main

import (
	"log"
	"os"
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
	ip "image-analyzer/imageprocessing"
)

func main() {

	// Setup the google client
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ip.DetectLabelsURI(ctx, *client, os.Stdout, "https://i.ytimg.com/vi/C5Sxt0k5AG0/maxresdefault.jpg")
	ip.DetectLabelsFile(ctx, *client, os.Stdout, "testdata/Elmer_Keith.jpg")
}