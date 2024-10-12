package go_necos

import (
	"context"
	"fmt"
	"net/url"
)

// endpoint urls for wrapper to call
const (
	Domain = "https://api.nekosapi.com/v3"

	Images          = Domain + "/images"
	RandomImages    = Images + "/random"
	ReportImage     = Images + "/report"
	Tags            = Images + "/tags"
	TagByID         = Tags + "/%d"
	TagImages       = TagByID + "/images"
	ImageByID       = Images + "/%d"
	ImageArtist     = ImageByID + "/artist"
	ImageCharacters = ImageByID + "/characters"
	ImageTags       = ImageByID + "/tags"

	Artists      = Domain + "/artists"
	ArtistByID   = Artists + "/%d"
	ArtistImages = ArtistByID + "/images"

	Characters      = Domain + "/characters"
	CharacterByID   = Characters + "/%d"
	CharacterImages = CharacterByID + "/images"
)

// MultipleContainer is struct that is returned when there can be
// more than one answer to the request
//
// returned by: /images, /images/random, /images/tags, /images/tags/{id}/images, /images/{id}/characters,
// /images/{id}/tags, /artists, /artist/{id}/images, /characters, /characters/{id}/images
type MultipleContainer[T any] struct {
	Items []T
	Count int
}

// Image is struct representing the image data returned by API
//
// returned by: /images/{id}
type Image struct {
	ID             int
	IDv2           string `json:"id_v2"`
	ImageURL       string `json:"image_url"`
	SampleURL      string `json:"sample_url"`
	ImageSize      int    `json:"image_size"`
	ImageWidth     int    `json:"image_width"`
	ImageHeight    int    `json:"image_height"`
	SampleSize     int    `json:"sample_size"`
	SampleWidth    int    `json:"sample_width"`
	SampleHeight   int    `json:"sample_height"`
	Source         string
	SourceID       int     `json:"source_id"`
	Rating         string  `json:"rating"`
	Verification   string  `json:"verification"`
	HashMD5        string  `json:"hash_md5"`
	HashPerceptual string  `json:"hash_perceptual"`
	ColorDominant  Color   `json:"color_dominant"`
	ColorPalette   []Color `json:"color_palette"`
	Duration       int     `json:"duration"`
	IsOriginal     bool    `json:"is_original"`
	IsScreenshot   bool    `json:"is_screenshot"`
	IsFlagged      bool    `json:"is_flagged"`
	IsAnimated     bool    `json:"is_animated"`
	Artist         Artist
	Characters     []Character
	Tags           []Tag
	CreatedAt      float64 `json:"created_at"`
	UpdatedAt      float64 `json:"updated_at"`
}

// Artist is data type that represents artist data returned by API
//
// returned by: /artists/{id}
type Artist struct {
	ID           int
	IDv2         string `json:"id_v2"`
	Name         string
	Aliases      []string
	ImageURL     string `json:"image_url"`
	Links        []string
	PolicyRepost string `json:"policy_repost"`
	PolicyCredit string `json:"policy_credit"`
	PolicyAI     string `json:"policy_ai"`
}

// Character is data type that represents character data returned by API
//
// returned by: /characters/{id}
type Character struct {
	ID          int
	IDv2        string `json:"id_v2"`
	Name        string
	Aliases     []string
	Description string
	Ages        []int
	Height      int
	Weight      int
	Gender      string
	Species     string
	Birthday    string
	Nationality string
	Occupations []string
}

// Tag is data type that represents tag data returned by API
//
// returned by: /images/tags/{id}
type Tag struct {
	ID          int
	IDv2        string `json:"id_v2"`
	Name        string
	Description string
	Sub         string
	IsNSFW      bool `json:"is_nsfw"`
}

// Report contains data needed to make POST request to report an image.
// Should contain id (integer) or url (string)
type Report = url.Values

// Color is custom data used in parsing of colors
type Color [3]int

// Request is data needed to make GET request to any of endpoints using URL query
//
// since all fields are optional and nothing breaks in the API when providing extra fields
// I decided it would be easier to make it from url.Values and don't make extra struct
// (url.Values is also easy to encode to query syntax)
//
// list of possible fields: search (string), id (integer), rating (array of strings), is_original (boolean),
// is_screenshot (boolean), is_flagged (boolean), is_animated (boolean), is_nsfw (boolean), policy_repost (boolean),
// policy_credit (boolean), policy_ai (boolean), artist (integer), character (array of integers),
// age (array of integers), gender (string), species (string), nationality (string), occupation (array of strings),
// tag (array of integers), limit (integer) [1...100, 100 by default], offset (integer) [>=0, 0 by default]
type Request = url.Values

// GetImages is a wrapper for Images endpoint
func (c *Client) GetImages(req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	err := c.Get(Images, req, &ret)
	return ret, err
}

// GetImagesWithContext is a wrapper for Images endpoint
func (c *Client) GetImagesWithContext(ctx context.Context, req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	err := c.GetWithContext(ctx, Images, req, &ret)
	return ret, err
}

// GetRandomImages is a wrapper for RandomImages endpoint
func (c *Client) GetRandomImages(req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	err := c.Get(RandomImages, req, &ret)
	return ret, err
}

// GetRandomImagesWithContext is a wrapper for RandomImages endpoint
func (c *Client) GetRandomImagesWithContext(ctx context.Context, req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	err := c.GetWithContext(ctx, RandomImages, req, &ret)
	return ret, err
}

// Report is a wrapper for Report endpoint
func (c *Client) Report(req Report) error {
	return c.Post(ReportImage, req, nil)
}

// ReportWithContext is a wrapper for Report endpoint
func (c *Client) ReportWithContext(ctx context.Context, req Report) error {
	return c.PostWithContext(ctx, ReportImage, req, nil)
}

// GetTags is a wrapper for Tags endpoint
func (c *Client) GetTags(req Request) (MultipleContainer[Tag], error) {
	var ret MultipleContainer[Tag]
	err := c.Get(Tags, req, &ret)
	return ret, err
}

// GetTagsWithContext is a wrapper for Tags endpoint
func (c *Client) GetTagsWithContext(ctx context.Context, req Request) (MultipleContainer[Tag], error) {
	var ret MultipleContainer[Tag]
	err := c.GetWithContext(ctx, Tags, req, &ret)
	return ret, err
}

// GetTagByID is a wrapper for TagByID endpoint
func (c *Client) GetTagByID(id int) (Tag, error) {
	var ret Tag
	path := fmt.Sprintf(TagByID, id)
	err := c.Get(path, nil, &ret)
	return ret, err
}

// GetTagByIDWithContext is a wrapper for TagByID endpoint
func (c *Client) GetTagByIDWithContext(ctx context.Context, id int) (Tag, error) {
	var ret Tag
	path := fmt.Sprintf(TagByID, id)
	err := c.GetWithContext(ctx, path, nil, &ret)
	return ret, err
}

// GetTagImages is a wrapper for TagImages endpoint
func (c *Client) GetTagImages(tagID int) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(TagImages, tagID)
	err := c.Get(path, nil, &ret)
	return ret, err
}

// GetTagImagesWithContext is a wrapper for TagImages endpoint
func (c *Client) GetTagImagesWithContext(ctx context.Context, tagID int) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(TagImages, tagID)
	err := c.GetWithContext(ctx, path, nil, &ret)
	return ret, err
}

// GetImageByID is a wrapper for ImageByID endpoint
func (c *Client) GetImageByID(id int) (Image, error) {
	var ret Image
	path := fmt.Sprintf(ImageByID, id)
	err := c.Get(path, nil, &ret)
	return ret, err
}

// GetImageByIDWithContext is a wrapper for ImageByID endpoint
func (c *Client) GetImageByIDWithContext(ctx context.Context, id int) (Image, error) {
	var ret Image
	path := fmt.Sprintf(ImageByID, id)
	err := c.GetWithContext(ctx, path, nil, &ret)
	return ret, err
}

// GetImageArtist is a wrapper for ImageArtist endpoint
func (c *Client) GetImageArtist(id int) (Artist, error) {
	var ret Artist
	path := fmt.Sprintf(ImageArtist, id)
	err := c.Get(path, nil, &ret)
	return ret, err
}

// GetImageArtistWithContext is a wrapper for ImageArtist endpoint
func (c *Client) GetImageArtistWithContext(ctx context.Context, id int) (Artist, error) {
	var ret Artist
	path := fmt.Sprintf(ImageArtist, id)
	err := c.GetWithContext(ctx, path, nil, &ret)
	return ret, err
}

// GetImageCharacters is a wrapper for ImageCharacters endpoint
func (c *Client) GetImageCharacters(id int) (MultipleContainer[Character], error) {
	var ret MultipleContainer[Character]
	path := fmt.Sprintf(ImageCharacters, id)
	err := c.Get(path, nil, &ret)
	return ret, err
}

// GetImageCharactersWithContext is a wrapper for ImageCharacters endpoint
func (c *Client) GetImageCharactersWithContext(ctx context.Context, id int) (MultipleContainer[Character], error) {
	var ret MultipleContainer[Character]
	path := fmt.Sprintf(ImageCharacters, id)
	err := c.GetWithContext(ctx, path, nil, &ret)
	return ret, err
}

// GetImageTags is a wrapper for ImageTags endpoint
func (c *Client) GetImageTags(id int) (MultipleContainer[Tag], error) {
	var ret MultipleContainer[Tag]
	path := fmt.Sprintf(ImageTags, id)
	err := c.Get(path, nil, &ret)
	return ret, err
}

// GetImageTagsWithContext is a wrapper for ImageTags endpoint
func (c *Client) GetImageTagsWithContext(ctx context.Context, id int) (MultipleContainer[Tag], error) {
	var ret MultipleContainer[Tag]
	path := fmt.Sprintf(ImageTags, id)
	err := c.GetWithContext(ctx, path, nil, &ret)
	return ret, err
}

// GetArtists is a wrapper for Artists endpoint
func (c *Client) GetArtists(req Request) (MultipleContainer[Artist], error) {
	var ret MultipleContainer[Artist]
	err := c.Get(Artists, req, &ret)
	return ret, err
}

// GetArtistsWithContext is a wrapper for Artists endpoint
func (c *Client) GetArtistsWithContext(ctx context.Context, req Request) (MultipleContainer[Artist], error) {
	var ret MultipleContainer[Artist]
	err := c.GetWithContext(ctx, Artists, req, &ret)
	return ret, err
}

// GetArtistByID is a wrapper for ArtistByID endpoint
func (c *Client) GetArtistByID(id int) (Artist, error) {
	var ret Artist
	path := fmt.Sprintf(ArtistByID, id)
	err := c.Get(path, nil, &ret)
	return ret, err
}

// GetArtistByIDWithContext is a wrapper for ArtistByID endpoint
func (c *Client) GetArtistByIDWithContext(ctx context.Context, id int) (Artist, error) {
	var ret Artist
	path := fmt.Sprintf(ArtistByID, id)
	err := c.GetWithContext(ctx, path, nil, &ret)
	return ret, err
}

// GetArtistImages is a wrapper for ArtistImages endpoint
func (c *Client) GetArtistImages(id int) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(ArtistImages, id)
	err := c.Get(path, nil, &ret)
	return ret, err
}

// GetArtistImagesWithContext is a wrapper for ArtistImages endpoint
func (c *Client) GetArtistImagesWithContext(ctx context.Context, id int) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(ArtistImages, id)
	err := c.GetWithContext(ctx, path, nil, &ret)
	return ret, err
}

// GetCharacters is a wrapper for Characters endpoint
func (c *Client) GetCharacters(req Request) (MultipleContainer[Character], error) {
	var ret MultipleContainer[Character]
	err := c.Get(Characters, req, &ret)
	return ret, err
}

// GetCharactersWithContext is a wrapper for Characters endpoint
func (c *Client) GetCharactersWithContext(ctx context.Context, req Request) (MultipleContainer[Character], error) {
	var ret MultipleContainer[Character]
	err := c.GetWithContext(ctx, Characters, req, &ret)
	return ret, err
}

// GetCharacterByID is a wrapper for CharacterByID endpoint
func (c *Client) GetCharacterByID(id int) (Character, error) {
	var ret Character
	path := fmt.Sprintf(CharacterByID, id)
	err := c.Get(path, nil, &ret)
	return ret, err
}

// GetCharacterByIDWithContext is a wrapper for CharacterByID endpoint
func (c *Client) GetCharacterByIDWithContext(ctx context.Context, id int) (Character, error) {
	var ret Character
	path := fmt.Sprintf(CharacterByID, id)
	err := c.GetWithContext(ctx, path, nil, &ret)
	return ret, err
}

// GetCharacterImages is a wrapper for CharacterImages endpoint
func (c *Client) GetCharacterImages(id int) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(CharacterImages, id)
	err := c.Get(path, nil, &ret)
	return ret, err
}

// GetCharacterImagesWithContext is a wrapper for CharacterImages endpoint
func (c *Client) GetCharacterImagesWithContext(ctx context.Context, id int) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(CharacterImages, id)
	err := c.GetWithContext(ctx, path, nil, &ret)
	return ret, err
}
