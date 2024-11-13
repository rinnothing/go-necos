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
	r := bufio.NewReader(os.Stdin)

	c := necos.NewClient()

	c.DefaultQuery = necos.AddFields(nil,
		"limit", 1,
		"rating", "safe")

	var tag necos.Tag
	tagReq := necos.Request{}
	for {
		fmt.Println("What kind of tags do you want to have?")
		search, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Can't read string:", err)
			return
		}

		search = strings.TrimSpace(search)
		necos.SetFields(tagReq, "search", search)

		tags, err := c.GetTags(tagReq)
		if err != nil {
			fmt.Println("Cant get tags:", err)
			return
		}

		if tags.Count != 0 {
			tag = tags.Items[0]
			break
		}
		fmt.Println("There's no such tags")
	}
	fmt.Printf("Tag found:\nname: %s\ndescription: %s\n", tag.Name, tag.Description)

	images, err := c.GetRandomImages(necos.AddFields(nil, "tag", tag.ID))

	if err != nil {
		fmt.Println("Can't get images:", err)
		return
	}

	if len(images.Items) == 0 {
		fmt.Println("No images")
		return
	}
	image := images.Items[0]

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Can't get current working directory:", err)
		return
	}

	path := filepath.Join(wd, image.GetName())
	w, err := necos.Save(path)
	if err != nil {
		fmt.Println("Can't create file to save image:", err)
		return
	}

	err = c.DownloadImage(&image, w)
	if err != nil {
		fmt.Println("Can't download image:", err)
		return
	}

	fmt.Printf("Download finished successfully! path: %s\n", path)
}
