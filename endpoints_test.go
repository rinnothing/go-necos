package go_necos

import (
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

// Since I'm lazy to try to mimic the way API should respond to my requests
// I will simply send requests to API itself

// I also don't want it to take an era to accomplish so limit is 1
var query = url.Values{"limit": {"1"}}

func TestGetImages(t *testing.T) {
	t.Parallel()
	c := newClient()

	_, err := c.GetImages(query)
	require.NoError(t, err)
}

func TestGetRandomImages(t *testing.T) {
	t.Parallel()
	c := newClient()

	_, err := c.GetRandomImages(query)
	require.NoError(t, err)
}

// I don't want to make any false alarms, so there'll be no test for PostReport

func TestGetTags(t *testing.T) {
	t.Parallel()
	c := newClient()

	_, err := c.GetTags(query)
	require.NoError(t, err)
}

func TestGetTagByID(t *testing.T) {
	t.Parallel()
	c := newClient()

	tags, err := c.GetTags(query)
	require.NoError(t, err)

	_, err = c.GetTagByID(tags.Items[0].ID)
	require.NoError(t, err)
}

func TestGetTagImages(t *testing.T) {
	t.Parallel()
	c := newClient()

	tags, err := c.GetTags(query)
	require.NoError(t, err)

	_, err = c.GetTagImages(tags.Items[0].ID)
	require.NoError(t, err)
}

func TestGetImageByID(t *testing.T) {
	t.Parallel()
	c := newClient()

	images, err := c.GetImages(query)
	require.NoError(t, err)

	_, err = c.GetImageByID(images.Items[0].ID)
	require.NoError(t, err)
}

func TestGetImageArtist(t *testing.T) {
	t.Parallel()
	c := newClient()

	images, err := c.GetImages(query)
	require.NoError(t, err)

	_, err = c.GetImageArtist(images.Items[0].ID)
	require.NoError(t, err)
}

func TestGetImageCharacters(t *testing.T) {
	t.Parallel()
	c := newClient()

	images, err := c.GetImages(query)
	require.NoError(t, err)

	_, err = c.GetImageCharacters(images.Items[0].ID)
	require.NoError(t, err)
}

func TestGetImageTags(t *testing.T) {
	t.Parallel()
	c := newClient()

	images, err := c.GetImages(query)
	require.NoError(t, err)

	_, err = c.GetImageTags(images.Items[0].ID)
	require.NoError(t, err)
}

func TestGetArtists(t *testing.T) {
	t.Parallel()
	c := newClient()

	_, err := c.GetArtists(query)
	require.NoError(t, err)
}

func TestGetArtistByID(t *testing.T) {
	t.Parallel()
	c := newClient()

	artists, err := c.GetArtists(query)
	require.NoError(t, err)

	_, err = c.GetArtistByID(artists.Items[0].ID)
	require.NoError(t, err)
}

func TestGetArtistImages(t *testing.T) {
	t.Parallel()
	c := newClient()

	artists, err := c.GetArtists(query)
	require.NoError(t, err)

	_, err = c.GetArtistImages(artists.Items[0].ID)
	require.NoError(t, err)
}

func TestGetCharacters(t *testing.T) {
	t.Parallel()
	c := newClient()

	_, err := c.GetCharacters(query)
	require.NoError(t, err)
}

func TestGetCharacterByID(t *testing.T) {
	t.Parallel()
	c := newClient()

	characters, err := c.GetCharacters(query)
	require.NoError(t, err)

	_, err = c.GetCharacterByID(characters.Items[0].ID)
	require.NoError(t, err)
}

func TestGetCharacterImages(t *testing.T) {
	t.Parallel()
	c := newClient()

	characters, err := c.GetCharacters(query)
	require.NoError(t, err)

	_, err = c.GetCharacterImages(characters.Items[0].ID, query)
	require.NoError(t, err)
}
