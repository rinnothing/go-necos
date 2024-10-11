package go_necos

import (
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

	Characters       = Domain + "/characters"
	CharacterByID   = Characters + "/%d"
	CharactersImages = CharacterByID + "/images"
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
type Report url.Values

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
type Request url.Values
