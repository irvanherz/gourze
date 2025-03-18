package dto

type BunnyCreateVideoInput struct {
	LibraryID     uint64  `json:"-"`
	Title         string  `json:"title"`
	CollectionID  *string `json:"collectionId"`
	ThumbnailTime *uint32 `json:"thumbnailTime"`
}
