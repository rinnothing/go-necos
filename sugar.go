package go_necos

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

var (
	SafeRequest = Request{"rating": []string{"safe"}}
)

// download is the method used to do all downloads
//
// it makes a GET request to given url and writes received content to dst, if all operations finished successfully closes it
func (c *Client) download(ctx context.Context, url string, dst io.WriteCloser) error {
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
	err = dst.Close()
	return err
}

// DownloadImage downloads the Image with default context
func (c *Client) DownloadImage(im *Image, dst io.WriteCloser) error {
	return c.DownloadImageWithContext(context.Background(), im, dst)
}

// DownloadImageWithContext downloads the Image with given context
func (c *Client) DownloadImageWithContext(ctx context.Context, im *Image, dst io.WriteCloser) error {
	return c.download(ctx, im.ImageURL, dst)
}

// DownloadSample downloads the sample of Image with default context
func (c *Client) DownloadSample(im *Image, dst io.WriteCloser) error {
	return c.DownloadSampleWithContext(context.Background(), im, dst)
}

// DownloadSampleWithContext downloads the sample of Image with given context
func (c *Client) DownloadSampleWithContext(ctx context.Context, im *Image, dst io.WriteCloser) error {
	return c.download(ctx, im.SampleURL, dst)
}
