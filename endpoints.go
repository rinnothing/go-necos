package necos

import (
	"context"
	"fmt"
	"net/url"
)

// endpoint urls for wrapper to call
const (
	DefaultDomain   = "https://api.nekosapi.com/v3"
	Images          = "/images"
	RandomImages    = Images + "/random"
	ReportImage     = Images + "/report"
	Tags            = Images + "/tags"
	TagByID         = Tags + "/%d"
	TagImages       = TagByID + "/images"
	ImageByID       = Images + "/%d"
	ImageArtist     = ImageByID + "/artist"
	ImageCharacters = ImageByID + "/characters"
	ImageTags       = ImageByID + "/tags"

	Artists      = "/artists"
	ArtistByID   = Artists + "/%d"
	ArtistImages = ArtistByID + "/images"

	Characters      = "/characters"
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
	PolicyRepost bool `json:"policy_repost"`
	PolicyCredit bool `json:"policy_credit"`
	PolicyAI     bool `json:"policy_ai"`
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
//
// Request for GetImages supports those parameters:
//   - rating (array of strings)
//   - is_original (boolean)
//   - is_screenshot (boolean)
//   - is flagged (boolean) - whether the image is flagged by mods
//   - is animated (boolean)
//   - artist (integer) - the arist's ID
//   - character (array of integers) - the character's ID
//   - tag (array of integers) - the tag's ID
//   - limit (integer) - [1..100], default = 100
//   - offset (integer) - >= 0, default = 0
func (c *Client) GetImages(req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	err := c.Get(Images, req, &ret)
	return ret, err
}

// GetImagesWithContext is a wrapper for Images endpoint
//
// For more info on Request parameters see GetImages
func (c *Client) GetImagesWithContext(ctx context.Context, req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	err := c.GetWithContext(ctx, Images, req, &ret)
	return ret, err
}

// GetRandomImages is a wrapper for RandomImages endpoint
//
// Request for GetRandomImages supports those parameters:
//   - rating (array of strings)
//   - is_original (boolean)
//   - is_screenshot (boolean)
//   - is flagged (boolean) - whether the image is flagged by mods
//   - is animated (boolean)
//   - artist (integer) - the arist's ID
//   - character (array of integers) - the character's ID
//   - tag (array of integers) - the tag's ID
//   - limit (integer) - [1..100], default = 100
func (c *Client) GetRandomImages(req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	err := c.Get(RandomImages, req, &ret)
	return ret, err
}

// GetRandomImagesWithContext is a wrapper for RandomImages endpoint
//
// For more info on Request parameters see GetRandomImages
func (c *Client) GetRandomImagesWithContext(ctx context.Context, req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	err := c.GetWithContext(ctx, RandomImages, req, &ret)
	return ret, err
}

// PostReport is a wrapper for ReportImage endpoint
//
// Report for PostReport supports those parameters:
//   - id (integer) - probably the id of Image
//   - url (string) - probably the url of Image
func (c *Client) PostReport(req Report) error {
	return c.Post(ReportImage, req, nil)
}

// PostReportWithContext is a wrapper for ReportImage endpoint
//
// For more info on Report parameters see PostReport
func (c *Client) PostReportWithContext(ctx context.Context, req Report) error {
	return c.PostWithContext(ctx, ReportImage, req, nil)
}

// GetTags is a wrapper for Tags endpoint
//
// Request for GetTags supports those parameters:
//   - search (string) - search for a tag by name or description
//   - is_nsfw (boolean)
//   - limit (integer) - [1..100], default = 100
//   - offset (integer) - >= 0, default = 0
func (c *Client) GetTags(req Request) (MultipleContainer[Tag], error) {
	var ret MultipleContainer[Tag]
	err := c.Get(Tags, req, &ret)
	return ret, err
}

// GetTagsWithContext is a wrapper for Tags endpoint
//
// For more info on Request parameters see GetTags
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
//
// Request for GetTagImages supports those parameters:
//   - limit (integer) - [1..100], default = 100
//   - offset (integer) - >= 0, default = 0
func (c *Client) GetTagImages(tagID int, req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(TagImages, tagID)
	err := c.Get(path, req, &ret)
	return ret, err
}

// GetTagImagesWithContext is a wrapper for TagImages endpoint
func (c *Client) GetTagImagesWithContext(ctx context.Context, tagID int, req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(TagImages, tagID)
	err := c.GetWithContext(ctx, path, req, &ret)
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
//
// Request for GetImageCharacters supports those parameters:
//   - limit (integer) - [1..100], default = 100
//   - offset (integer) - >= 0, default = 0
//
// This method isn't recommended to use since most of the time the server only returns 500 (internal server error)
func (c *Client) GetImageCharacters(id int, req Request) (MultipleContainer[Character], error) {
	var ret MultipleContainer[Character]
	path := fmt.Sprintf(ImageCharacters, id)
	err := c.Get(path, req, &ret)
	return ret, err
}

// GetImageCharactersWithContext is a wrapper for ImageCharacters endpoint
//
// This method isn't recommended to use since most of the time the server only returns 500 (internal server error)
func (c *Client) GetImageCharactersWithContext(ctx context.Context, id int, req Request) (MultipleContainer[Character], error) {
	var ret MultipleContainer[Character]
	path := fmt.Sprintf(ImageCharacters, id)
	err := c.GetWithContext(ctx, path, req, &ret)
	return ret, err
}

// GetImageTags is a wrapper for ImageTags endpoint
//
// Request for GetImageTags supports those parameters:
//   - limit (integer) - [1..100], default = 100
//   - offset (integer) - >= 0, default = 0
//
// This method isn't recommended to use since most of the time the server only returns 500 (internal server error)
func (c *Client) GetImageTags(id int, req Request) (MultipleContainer[Tag], error) {
	var ret MultipleContainer[Tag]
	path := fmt.Sprintf(ImageTags, id)
	err := c.Get(path, req, &ret)
	return ret, err
}

// GetImageTagsWithContext is a wrapper for ImageTags endpoint
//
// This method isn't recommended to use since most of the time the server only returns 500 (internal server error)
func (c *Client) GetImageTagsWithContext(ctx context.Context, id int, req Request) (MultipleContainer[Tag], error) {
	var ret MultipleContainer[Tag]
	path := fmt.Sprintf(ImageTags, id)
	err := c.GetWithContext(ctx, path, req, &ret)
	return ret, err
}

// GetArtists is a wrapper for Artists endpoint
//
// Request for GetArtists supports those parameters:
//   - search (string) - Search term. Will return all tags with this term(s) in their name or description
//   - policy_repost (boolean) - Does this artist allow you to repost their art in other places?
//   - policy_credit (boolean) - Are you required to credit the artist when using their art?
//   - policy_ai (boolean) - Does the artist allow you to use their art for AI projects (AI training)?
//   - limit (integer) - [1..100], default = 100
//   - offset (integer) - >= 0, default = 0
func (c *Client) GetArtists(req Request) (MultipleContainer[Artist], error) {
	var ret MultipleContainer[Artist]
	err := c.Get(Artists, req, &ret)
	return ret, err
}

// GetArtistsWithContext is a wrapper for Artists endpoint
//
// For more info on Request parameters see GetArtists
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
//
// Request for GetArtistImages supports those parameters:
//   - limit (integer) - [1..100], default = 100
//   - offset (integer) - >= 0, default = 0
func (c *Client) GetArtistImages(id int, req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(ArtistImages, id)
	err := c.Get(path, req, &ret)
	return ret, err
}

// GetArtistImagesWithContext is a wrapper for ArtistImages endpoint
func (c *Client) GetArtistImagesWithContext(ctx context.Context, id int, req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(ArtistImages, id)
	err := c.GetWithContext(ctx, path, req, &ret)
	return ret, err
}

// GetCharacters is a wrapper for Characters endpoint
//
// Request for GetCharacters supports those parameters:
//   - search (string) - probably searches for the term in name and description
//   - age (array of integers) - One or more of the character's (official) ages.
//   - gender (string)
//   - species (string)
//   - nationality (string)
//   - occupation (array of strings) - Occupations the character officially has/has officially had.
//   - limit (integer) - [1..100], default = 100
//   - offset (integer) - >= 0, default = 0
func (c *Client) GetCharacters(req Request) (MultipleContainer[Character], error) {
	var ret MultipleContainer[Character]
	err := c.Get(Characters, req, &ret)
	return ret, err
}

// GetCharactersWithContext is a wrapper for Characters endpoint
//
// For more info on Request parameters see GetCharacters
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
//
// Request for GetCharacterImages supports those parameters:
//   - limit (integer) - [1..100], default = 100
//   - offset (integer) - >= 0, default = 0
func (c *Client) GetCharacterImages(id int, req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(CharacterImages, id)
	err := c.Get(path, req, &ret)
	return ret, err
}

// GetCharacterImagesWithContext is a wrapper for CharacterImages endpoint
//
// For more info on Request parameters see GetCharacterImages
func (c *Client) GetCharacterImagesWithContext(ctx context.Context, id int, req Request) (MultipleContainer[Image], error) {
	var ret MultipleContainer[Image]
	path := fmt.Sprintf(CharacterImages, id)
	err := c.GetWithContext(ctx, path, req, &ret)
	return ret, err
}
