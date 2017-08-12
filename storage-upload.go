package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	"google.golang.org/api/option"

	// Imports the Google Cloud Storage client package.
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

var verb bool

func main() {
	ctx := context.Background()

	// Get args and process them
	keypath := flag.String("k", "", "Service key path")
	dest := flag.String("d", "", "Upload path, of form 'gs://bucket/folder/file.txt'")
	help := flag.Bool("help", false, "Prints usage")
	flag.BoolVar(&verb, "verbose", false, "Outputs infos")
	projectID := flag.String("p", "", "Create the bucket in the corresponding given projectID if it doesn't exist")

	flag.Parse()

	switch {
	case *help:
		flag.PrintDefaults()
		os.Exit(0)
	case *keypath == "":
		log.Fatal("You must provide the service key path with option '-k'")
	case *dest == "":
		log.Fatal("You must provide the destination path with option '-d'")
	}

	// Get bucket and path
	slices := strings.SplitN(*dest, "/", 4)
	if len(slices) < 4 {
		log.Fatal("Wrong path, must be of kind 'gs://bucket/folder/file.txt'")
	}
	bucket := slices[2]
	path := slices[3]
	print(bucket)
	print(path)

	// Creates a client.
	client, err := storage.NewClient(ctx, option.WithServiceAccountFile(*keypath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Object to interact with our bucket
	bkt := client.Bucket(bucket)
	if *projectID != "" {
		if err := bkt.Create(ctx, *projectID, nil); err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}
	}

	// Read from stdin
	print("Reading data")
	w := bkt.Object(path).NewWriter(ctx)
	n, err := bufio.NewReader(os.Stdin).WriteTo(w)
	if err != nil {
		log.Fatalf("Uploading error: %v", err)
	}
	log.Printf("Sucessfully uploaded %d byte(s) of data", n)
	if err := w.Close(); err != nil {
		log.Fatalf("Close w2 error: %v", err)
	}

}

func print(text string) {
	if verb {
		log.Println(text)
	}
}
