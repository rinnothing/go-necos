package go_necos

import (
	"net/url"
)

type MultipleContainer[T any] struct {
	Items []T
	Count int
}

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

type Tag struct {
	ID          int
	IDv2        string `json:"id_v2"`
	Name        string
	Description string
	Sub         string
	IsNSFW      bool `json:"is_nsfw"`
}

type Color [3]int

type Request url.Values
