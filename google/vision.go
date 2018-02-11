package google

import (
	"io"
	"context"
	"fmt"
)


/*
detectLabels gets labels from the Vision API for an image at the given file path
this is from https://cloud.google.com/vision/docs/detecting-labels#vision-label-detection-go
*/
func DetectLabelsURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
	return err
	}

	image := vision.NewImageFromURI(file)
	annotations, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
	return err
	}

	if len(annotations) == 0 {
	fmt.Fprintln(w, "No labels found.")
	} else {
	fmt.Fprintln(w, "Labels:")
	for _, annotation := range annotations {
	fmt.Fprintln(w, annotation.Description)
	}
	}

	return nil
}
