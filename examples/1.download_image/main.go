package main

import (
	"bufio"
	"fmt"
	"github.com/rinnothing/go-necos"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// opening stdin to read input by lines
	r := bufio.NewReader(os.Stdin)

	// creating new client with default domain
	c := necos.NewClient()

	// we want to get only one value and need it to be safe rating
	// providing default query we avoid need to set it with every Request
	c.DefaultQuery = necos.AddFields(nil,
		"limit", 1,
		"rating", "safe")

	// value we will put tag in
	var tag necos.Tag
	// request we will set 'search' value to
	tagReq := necos.Request{}
	for {
		// getting tag value to search for
		fmt.Println("What kind of tags do you want to have?")
		search, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Can't read string:", err)
			return
		}

		// trimming string to ease search
		search = strings.TrimSpace(search)
		// set 'search' value to Request
		necos.SetFields(tagReq, "search", search)

		// getting tags from API
		tags, err := c.GetTags(tagReq)
		if err != nil {
			fmt.Println("Cant get tags:", err)
			return
		}

		// if tag is found - assign and leave
		if tags.Count != 0 {
			tag = tags.Items[0]
			break
		}
		fmt.Println("There's no such tags")
	}
	fmt.Printf("Tag found:\nname: %s\ndescription: %s\n", tag.Name, tag.Description)

	// getting image by tag ID (also creating request in place (attention, it should not be used in cycles, due to constant creation on new maps))
	images, err := c.GetRandomImages(necos.AddFields(nil, "tag", tag.ID))

	if err != nil {
		fmt.Println("Can't get images:", err)
		return
	}

	if len(images.Items) == 0 {
		fmt.Println("No images")
		return
	}
	// taking image
	image := images.Items[0]

	// getting working directory to construct path for image
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Can't get current working directory:", err)
		return
	}

	// constructing path for image
	path := filepath.Join(wd, image.GetName())
	// creating file to save image by path
	w, err := necos.Save(path)
	// when using temp files better add
	//	defer os.Remove(path)
	if err != nil {
		fmt.Println("Can't create file to save image:", err)
		return
	}

	// downloading image to given writer
	err = c.DownloadImage(&image, w)
	if err != nil {
		fmt.Println("Can't download image:", err)
		return
	}

	fmt.Printf("Download finished successfully! path: %s\n", path)
}
