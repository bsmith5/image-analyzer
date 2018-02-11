package google

import (
	"golang.org/x/net/context"
	vision "cloud.google.com/go/vision/apiv1"
	"os"
	"log"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

// this borrows heavily from https://cloud.google.com/vision/docs/detecting-labels#vision-label-detection-go


// detects labels from uri
func DetectLabelsURI(ctx context.Context, client vision.ImageAnnotatorClient, URI string) ([]*pb.EntityAnnotation, error) {

	image := vision.NewImageFromURI(URI)
	annotations, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		return nil, err
	}

	if len(annotations) == 0 {
		log.Fatal("No labels found.")
		return nil, err
	}

	return annotations, nil
}


// detects labels from file
func DetectLabelsFile(ctx context.Context, client vision.ImageAnnotatorClient, filename string) ([]*pb.EntityAnnotation, error) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
		return nil, err
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
		return nil, err
	}

	labels, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
		return nil, err
	}

	return labels, nil
}