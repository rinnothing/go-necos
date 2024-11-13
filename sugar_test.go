package necos

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/image/webp"
	"math/rand/v2"
	"os"
	"path/filepath"
	"testing"
)

// Basically the same as TestSaveToSlice
func TestDownloadImage(t *testing.T) {
	t.Parallel()
	c := NewClient()

	images, err := c.GetImages(OneValue())
	require.NoError(t, err)
	require.GreaterOrEqual(t, images.Count, 1)

	image := images.Items[0]
	imageSlice := make([]byte, 0)

	err = c.DownloadImage(&image, SaveToSlice(&imageSlice))
	require.NoError(t, err)

	_, err = webp.Decode(bytes.NewReader(imageSlice))
	require.NoError(t, err)
}

func TestDownloadSample(t *testing.T) {
	t.Parallel()
	c := NewClient()

	images, err := c.GetImages(OneValue())
	require.NoError(t, err)
	require.GreaterOrEqual(t, images.Count, 1)

	image := images.Items[0]
	imageSlice := make([]byte, 0)

	err = c.DownloadSample(&image, SaveToSlice(&imageSlice))
	require.NoError(t, err)

	_, err = webp.Decode(bytes.NewReader(imageSlice))
	require.NoError(t, err)
}

func TestSave(t *testing.T) {
	t.Parallel()
	c := NewClient()

	images, err := c.GetImages(OneValue())
	require.NoError(t, err)
	require.GreaterOrEqual(t, images.Count, 1)

	image := images.Items[0]

	name := filepath.Join(os.TempDir(), fmt.Sprint(rand.Uint()))
	writer, err := Save(name)
	require.NoError(t, err)
	defer os.Remove(name)

	err = c.DownloadSample(&image, writer)
	require.NoError(t, err)

	f, err := os.Open(name)
	require.NoError(t, err)
	defer f.Close()

	_, err = webp.Decode(f)
	require.NoError(t, err)
}

func TestSaveTemp(t *testing.T) {
	t.Parallel()
	c := NewClient()

	images, err := c.GetImages(OneValue())
	require.NoError(t, err)
	require.GreaterOrEqual(t, images.Count, 1)

	image := images.Items[0]

	writer, name, err := SaveTemp(image.GetPattern())
	require.NoError(t, err)
	defer os.Remove(name)

	err = c.DownloadSample(&image, writer)
	require.NoError(t, err)

	f, err := os.Open(name)
	require.NoError(t, err)
	defer f.Close()

	_, err = webp.Decode(f)
	require.NoError(t, err)
}

func TestOneValue(t *testing.T) {
	t.Parallel()
	c := NewClient()

	images, err := c.GetImages(OneValue())
	require.NoError(t, err)

	require.Equal(t, 1, len(images.Items))
}

func TestSafeRequest(t *testing.T) {
	t.Parallel()
	c := NewClient()

	images, err := c.GetImages(SafeRequest())
	require.NoError(t, err)

	for _, im := range images.Items {
		assert.Equal(t, "safe", im.Rating)
	}
}

func TestAddFields(t *testing.T) {
	goal := Request{
		"clean": {"shirt"},
		"new":   {"shoes"},
		"and":   {"I", "don't", "know"},
	}

	req := AddFields(nil,
		"clean", "shirt",
		"new", "shoes",
		"and", "I",
		"and", "don't",
		"and", "know")

	require.Equal(t, goal, req)

	AddFields(req, "where", 1)
	goal.Add("where", "1")

	require.Equal(t, goal, req)
}

func TestSetFields(t *testing.T) {
	goal := Request{
		"fly":         {"me"},
		"to":          {"the"},
		"moon":        {"and"},
		"let me play": {"among the stars"},
	}

	req := SetFields(nil,
		"fly", "me",
		"to", "the",
		"moon", "and",
		"let me play", "amogus",
		"let me play", "among the stars")

	require.Equal(t, goal, req)

	SetFields(req,
		"moon", 1,
		"billy", "jean")
	goal.Set("moon", "1")
	goal.Add("billy", "jean")

	require.Equal(t, goal, req)
}
