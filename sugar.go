package go_necos

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	SafeRequest = Request{"rating": []string{"safe"}}
)

// todo: make a general interface that could allow for files to be read more than once and to Close them (for SaveRam to nil it's slice and for SaveTemp to delete)

// Save returns file under given name
func Save(name string) (*os.File, error) {
	return os.Create(name)
}

// SaveTemp returns file created in temporary directory
//
// it's the callers responsibility to delete file after use
func SaveTemp() (*os.File, error) {
	return os.CreateTemp("", "go-necos")
}

// SaveRAM return writer that saves it's content to RAM
func SaveRAM() (*bytes.Buffer, error) {
	return new(bytes.Buffer), nil
}

// download is the method used to do all downloads
//
// it makes a GET request to given url and writes received content to dst, doesn't close the Writer
func (c *Client) download(ctx context.Context, url string, dst io.Writer) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}

	response, err := c.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %s", BadStatusError, response.Status)
	}

	_, err = io.Copy(dst, response.Body)
	if err != nil {
		return err
	}
	return err
}

// DownloadImage downloads the Image with default context
//
// doesn't close the Writer
func (c *Client) DownloadImage(im *Image, dst io.Writer) error {
	return c.DownloadImageWithContext(context.Background(), im, dst)
}

// DownloadImageWithContext downloads the Image with given context
//
// doesn't close the Writer
func (c *Client) DownloadImageWithContext(ctx context.Context, im *Image, dst io.Writer) error {
	return c.download(ctx, im.ImageURL, dst)
}

// DownloadSample downloads the sample of Image with default context
//
// doesn't close the Writer
func (c *Client) DownloadSample(im *Image, dst io.Writer) error {
	return c.DownloadSampleWithContext(context.Background(), im, dst)
}

// DownloadSampleWithContext downloads the sample of Image with given context
//
// doesn't close the Writer
func (c *Client) DownloadSampleWithContext(ctx context.Context, im *Image, dst io.Writer) error {
	return c.download(ctx, im.SampleURL, dst)
}
