package google

import (
	"io"
	"fmt"
	"golang.org/x/net/context"
	vision "cloud.google.com/go/vision/apiv1"
	"os"
	"log"
	//pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)


/*
detectLabels gets labels from the Vision API for an image at the given file path
this is from https://cloud.google.com/vision/docs/detecting-labels#vision-label-detection-go
*/
func DetectLabelsURI(ctx context.Context, client vision.ImageAnnotatorClient, w io.Writer, URI string) error {

	image := vision.NewImageFromURI(URI)
	annotations, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No labels found.")
	} else {
		fmt.Fprintln(w, "Labels from", URI, ":")
		for _, annotation := range annotations {
			fmt.Fprintln(w, annotation.Description)
		}
	}

	return nil
}


// detects labesl from file
func DetectLabelsFile(ctx context.Context, client vision.ImageAnnotatorClient, w io.Writer, filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	labels, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	fmt.Println("Labels from", filename, ":")
	for _, label := range labels {
		fmt.Println(label.Description)
	}

	return nil
}