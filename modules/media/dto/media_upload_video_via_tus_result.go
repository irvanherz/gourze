package dto

type MediaUploadVideoViaTusResult struct {
	Headers  MediaUploadVideoViaTusResultHeaders  `json:"headers"`
	Metadata MediaUploadVideoViaTusResultMetadata `json:"metadata"`
}

type MediaUploadVideoViaTusResultHeaders struct {
	AuthorizationSignature string `json:"authorizationSignature"`
	AuthorizationExpire    uint64 `json:"authorizationExpire"`
	VideoId                string `json:"videoId"`
	LibraryId              uint64 `json:"libraryId"`
}

type MediaUploadVideoViaTusResultMetadata struct {
	Filetype   string  `json:"filetype"`
	Title      string  `json:"title"`
	Collection *string `json:"collection"`
	UserID     uint    `json:"userId"`
}
