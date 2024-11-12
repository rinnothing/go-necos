package go_necos

import (
	"bufio"
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

type imageFilePrivate interface {
	io.ReadWriteCloser
	Remove() error
	Save() error
}

type ImageFile interface {
	io.ReadCloser
	Remove() error
	getPrivate() imageFilePrivate
}

type DiskFile struct {
	bufio.ReadWriter
	f     *os.File
	saved bool
}

func NewDiskFile(f *os.File) ImageFile {
	return &DiskFile{
		ReadWriter: bufio.ReadWriter{
			Reader: bufio.NewReader(f),
			Writer: bufio.NewWriter(f),
		},
		f:     f,
		saved: false,
	}
}

func (d *DiskFile) Close() error {
	if err := d.Flush(); err != nil {
		return err
	}
	return d.f.Close()
}

func (d *DiskFile) Remove() error {
	if err := d.Close(); err != nil {
		return err
	}
	return os.Remove(d.f.Name())
}

func (d *DiskFile) Save() error {
	if d.saved {
		return nil
	}

	if err := d.Flush(); err != nil {
		return err
	}
	d.saved = true
	return nil
}

func (d *DiskFile) getPrivate() imageFilePrivate {
	return d
}

// Save returns file under given name
func Save(name string) (ImageFile, error) {
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return NewDiskFile(file), nil
}

// SaveTemp returns file created in temporary directory
//
// it's the callers responsibility to delete file after use
func SaveTemp() (ImageFile, error) {
	file, err := os.CreateTemp("", "go-necos")
	if err != nil {
		return nil, err
	}

	return NewDiskFile(file), nil
}

type RamFile struct {
	io.ReadWriter
	saved bool
}

func NewRamFile() ImageFile {
	return &RamFile{
		ReadWriter: &bytes.Buffer{},
		saved:      false,
	}
}

func (r *RamFile) Close() error {
	return nil
}

func (r *RamFile) Remove() error {
	r.ReadWriter = nil
	return nil
}

func (r *RamFile) Save() error {
	if r.saved {
		return nil
	}

	r.saved = true
	return nil
}

func (r *RamFile) getPrivate() imageFilePrivate {
	return r
}

// SaveRAM return writer that saves it's content to RAM
func SaveRAM() (*bytes.Buffer, error) {
	return new(bytes.Buffer), nil
}

// download is the method used to do all downloads
//
// it makes a GET request to given url and writes received content to dst, saves the content when finished
func (c *Client) download(ctx context.Context, url string, dstCmn ImageFile) error {
	dst := dstCmn.getPrivate()

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

	if _, err = io.Copy(dst, response.Body); err != nil {
		return err
	}
	if err = dst.Save(); err != nil {
		return err
	}
	return err
}

// DownloadImage downloads the Image with default context
//
// doesn't close the Writer
func (c *Client) DownloadImage(im *Image, dst ImageFile) error {
	return c.DownloadImageWithContext(context.Background(), im, dst)
}

// DownloadImageWithContext downloads the Image with given context
//
// doesn't close the Writer
func (c *Client) DownloadImageWithContext(ctx context.Context, im *Image, dst ImageFile) error {
	return c.download(ctx, im.ImageURL, dst)
}

// DownloadSample downloads the sample of Image with default context
//
// doesn't close the Writer
func (c *Client) DownloadSample(im *Image, dst ImageFile) error {
	return c.DownloadSampleWithContext(context.Background(), im, dst)
}

// DownloadSampleWithContext downloads the sample of Image with given context
//
// doesn't close the Writer
func (c *Client) DownloadSampleWithContext(ctx context.Context, im *Image, dst ImageFile) error {
	return c.download(ctx, im.SampleURL, dst)
}
