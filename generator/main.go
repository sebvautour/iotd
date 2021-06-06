package main

// Go util to generate an array of JSON images

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/colors"
	extractor "github.com/marekm4/color-extractor"
)

type Data struct {
	Title string
	URLs  []ImageMeta
}

type ImageMeta struct {
	URL               string
	URLHash           string
	PrimaryColorHex   string
	SecondaryColorHex string
}

func main() {
	// set log output to stderr to make piping the results easier
	log.SetOutput(os.Stderr)

	if len(os.Args) != 2 {
		log.Fatalln("Exactly one argument of the URL list must be provided")
	}

	// open URL list
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to open %s: %v", os.Args[1], err)
	}
	defer f.Close()

	var images []ImageMeta
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		imageURL := sc.Text()
		meta, err := NewImageMeta(imageURL)
		if err != nil {
			log.Printf("Failed to get image meta from %s\nError: %v\n Skipping...", imageURL, err)
			continue
		}
		images = append(images, meta)
	}

	j := json.NewEncoder(os.Stdout)
	j.SetIndent("", " ")
	err = j.Encode(Data{Title: "Memes", URLs: images})
	if err != nil {
		log.Fatalf("Failed write JSON output: %v", err)
	}
}

func NewImageMeta(imageURL string) (meta ImageMeta, err error) {
	resp, err := http.DefaultClient.Get(imageURL)
	if err != nil {
		return meta, err
	}

	if resp.StatusCode != 200 {
		return meta, errors.New("invalid HTTP status code: " + resp.Status)
	}

	defer resp.Body.Close()
	image, _, err := image.Decode(resp.Body)
	if err != nil {
		return meta, err
	}

	// Fill URL
	meta.URL = imageURL

	// Fill hash
	h := md5.New()
	h.Write([]byte(imageURL))
	meta.URLHash = hex.EncodeToString(h.Sum(nil))

	// Fill primary and secondary colors
	// defaults
	meta.PrimaryColorHex = "#9999FF"
	meta.SecondaryColorHex = "#FFFFFF"
	imgColors := extractor.ExtractColors(image)

	if len(imgColors) != 0 {
		meta.PrimaryColorHex = colors.FromStdColor(imgColors[0]).ToHEX().String()
	}
	if len(imgColors) >= 2 {
		meta.SecondaryColorHex = colors.FromStdColor(imgColors[1]).ToHEX().String()
	}

	return meta, nil
}
