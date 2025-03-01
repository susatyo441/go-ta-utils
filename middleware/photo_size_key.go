package middleware

// for using save multi image
type PhotoSizeKey string

const (
	Low    = PhotoSizeKey("low")
	Medium = PhotoSizeKey("medium")
	High   = PhotoSizeKey("high")
)
