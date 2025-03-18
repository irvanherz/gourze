package dto

import "time"

type BunnyCaption struct {
	Srclang string `json:"srclang"`
	Label   string `json:"label"`
}

type BunnyChapter struct {
	Title string `json:"title"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type BunnyMoment struct {
	Label     string `json:"label"`
	Timestamp int    `json:"timestamp"`
}

type BunnyMetaTag struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

type BunnyTranscodingMessage struct {
	TimeStamp time.Time `json:"timeStamp"`
	Level     int       `json:"level"`
	IssueCode int       `json:"issueCode"`
	Message   string    `json:"message"`
	Value     string    `json:"value"`
}

type BunnyCreateVideoResponse struct {
	VideoLibraryId       int                       `json:"videoLibraryId"`
	Guid                 string                    `json:"guid"`
	Title                string                    `json:"title"`
	DateUploaded         time.Time                 `json:"dateUploaded"`
	Views                int                       `json:"views"`
	IsPublic             bool                      `json:"isPublic"`
	Length               int                       `json:"length"`
	Status               int                       `json:"status"`
	Framerate            int                       `json:"framerate"`
	Rotation             int                       `json:"rotation"`
	Width                int                       `json:"width"`
	Height               int                       `json:"height"`
	AvailableResolutions string                    `json:"availableResolutions"`
	OutputCodecs         string                    `json:"outputCodecs"`
	ThumbnailCount       int                       `json:"thumbnailCount"`
	EncodeProgress       int                       `json:"encodeProgress"`
	StorageSize          int                       `json:"storageSize"`
	Captions             []BunnyCaption            `json:"captions"`
	HasMP4Fallback       bool                      `json:"hasMP4Fallback"`
	CollectionId         string                    `json:"collectionId"`
	ThumbnailFileName    string                    `json:"thumbnailFileName"`
	AverageWatchTime     int                       `json:"averageWatchTime"`
	TotalWatchTime       int                       `json:"totalWatchTime"`
	Category             string                    `json:"category"`
	Chapters             []BunnyChapter            `json:"chapters"`
	Moments              []BunnyMoment             `json:"moments"`
	MetaTags             []BunnyMetaTag            `json:"metaTags"`
	TranscodingMessages  []BunnyTranscodingMessage `json:"transcodingMessages"`
}
