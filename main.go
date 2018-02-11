package main

import (
	"log"
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
	ip "image-analyzer/imageprocessing"
	"fmt"
)

func main() {

	// Setup the google client
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uri_labels, err := ip.DetectLabelsURI(ctx, *client, "https://i.ytimg.com/vi/C5Sxt0k5AG0/maxresdefault.jpg")
	if err != nil {
		log.Fatalf("Failed to get labels: %v", err)
	}

	file_labels, err := ip.DetectLabelsFile(ctx, *client, "testdata/Elmer_Keith.jpg")
	if err != nil {
		log.Fatalf("Failed to get labels: %v", err)
	}

	fmt.Println("uri labels:", uri_labels)
	fmt.Println("file_labels:", file_labels)

	/* example for how to print
	w = os.Stdout
	fmt.Fprintln(w, "Labels from", URI, ":")
	for _, annotation := range annotations {
		fmt.Fprintln(w, annotation.Description)
	}

	fmt.Println("Labels from", filename, ":")
	for _, label := range labels {
		fmt.Println(label.Description)
	}
	*/
}
