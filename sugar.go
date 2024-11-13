package necos

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var (
	SafeRequest = Request{"rating": []string{"safe"}}
	OneValue    = Request{"limit": {"1"}}
)

// GetName returns Image name
func (im *Image) GetName() string {
	return filepath.Base(im.ImageURL)
}

// GetSampleName returns ImageSample name
func (im *Image) GetSampleName() string {
	return filepath.Base(im.SampleURL)
}

// GetPattern makes a pattern to use in SaveTemp
//
// it assumes that ImageSample and Image have the same extension
func (im *Image) GetPattern() string {
	return "*" + filepath.Ext(im.ImageURL)
}

type fileWriter struct {
	*bufio.Writer
	f *os.File
}

func newFileWriter(f *os.File) *fileWriter {
	return &fileWriter{
		Writer: bufio.NewWriter(f),
		f:      f,
	}
}

func (fw *fileWriter) Close() error {
	if err := fw.Flush(); err != nil {
		return err
	}
	return fw.f.Close()
}

// Save writes file under given name
func Save(name string) (io.WriteCloser, error) {
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return newFileWriter(file), nil
}

// SaveTemp writes file in a temporary directory and returns it's name
//
// it's the callers responsibility to delete file after use
func SaveTemp(pattern string) (io.WriteCloser, string, error) {
	file, err := os.CreateTemp("", pattern)
	if err != nil {
		return nil, "", err
	}

	return newFileWriter(file), file.Name(), nil
}

type sliceWriter struct {
	slice *[]byte
}

func (sw sliceWriter) Write(p []byte) (int, error) {
	*sw.slice = append(*sw.slice, p...)
	return len(p), nil
}

func (sw sliceWriter) Close() error {
	*sw.slice = (*sw.slice)[0:len(*sw.slice):len(*sw.slice)]
	return nil
}

// SaveToSlice return writer that saves it's content to RAM
func SaveToSlice(dst *[]byte) io.WriteCloser {
	return sliceWriter{slice: dst}
}

// DownloadAppend is the method used to do append downloaded to given writer
//
// it makes a GET request to given url and writes received content to dst
func (c *Client) DownloadAppend(ctx context.Context, url string, dst io.Writer) error {
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
	return err
}

// Download is the method used to download Images
//
// Closes the file after finished reading
func (c *Client) Download(ctx context.Context, url string, dst io.WriteCloser) error {
	if err := c.DownloadAppend(ctx, url, dst); err != nil {
		return err
	}
	return dst.Close()
}

// DownloadImage downloads the Image with default context
//
// doesn't close the Writer
func (c *Client) DownloadImage(im *Image, dst io.WriteCloser) error {
	return c.DownloadImageWithContext(context.Background(), im, dst)
}

// DownloadImageWithContext downloads the Image with given context
//
// doesn't close the Writer
func (c *Client) DownloadImageWithContext(ctx context.Context, im *Image, dst io.WriteCloser) error {
	return c.Download(ctx, im.ImageURL, dst)
}

// DownloadSample downloads the sample of Image with default context
//
// doesn't close the Writer
func (c *Client) DownloadSample(im *Image, dst io.WriteCloser) error {
	return c.DownloadSampleWithContext(context.Background(), im, dst)
}

// DownloadSampleWithContext downloads the sample of Image with given context
//
// doesn't close the Writer
func (c *Client) DownloadSampleWithContext(ctx context.Context, im *Image, dst io.WriteCloser) error {
	return c.Download(ctx, im.SampleURL, dst)
}
