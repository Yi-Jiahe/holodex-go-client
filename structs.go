package holodex

import (
	"time"
)

type VideoFull struct {
	ID             string             `json:"id"`
	Title          string             `json:"title"`
	Type           string             `json:"type"`
	TopicID        string             `json:"topic_id"`
	PublishedAt    time.Time          `json:"published_at"`
	AvailableAt    time.Time          `json:"available_at"`
	Duration       int                `json:"duration"`
	Status         string             `json:"status"`
	StartScheduled time.Time          `json:"start_scheduled"`
	StartActual    time.Time          `json:"start_actual"`
	EndActual      time.Time          `json:"end_actual"`
	LiveViewers    int                `json:"live_viewers"`
	Description    string             `json:"description"`
	Songcount      int                `json:"songcount"`
	ChannelID      string             `json:"channel_id"`
	Clips          []VideoWithChannel `json:"clips"`
	Sources        []VideoWithChannel `json:"sources"`
	Refers         []VideoWithChannel `json:"refers"`
	Simulcasts     []VideoWithChannel `json:"simulcasts"`
	Mentions       []Mention          `json:"mentions"`
	Songs          int                `json:"songs"`
}

type ChannelVideosResponse struct {
	Total int         `json:"total"`
	Items []VideoFull `json:"items"`
}

type Channel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
	Type        string `json:"type"`
	Photo       string `json:"photo"`
}

type VideoWithChannel struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Type           string    `json:"type"`
	TopicID        string    `json:"topic_id"`
	PublishedAt    time.Time `json:"published_at"`
	AvailableAt    time.Time `json:"available_at"`
	Duration       int       `json:"duration"`
	Status         string    `json:"status"`
	StartScheduled time.Time `json:"start_scheduled"`
	StartActual    time.Time `json:"start_actual"`
	EndActual      time.Time `json:"end_actual"`
	LiveViewers    int       `json:"live_viewers"`
	Description    string    `json:"description"`
	Songcount      int       `json:"songcount"`
	ChannelID      string    `json:"channel_id"`
	Channel        Channel   `json:"channel"`
}

type Mention struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
	Type        string `json:"type"`
	Photo       string `json:"photo"`
	Org         string `json:"org"`
}
