package imageupload

type FileIdentifier struct {
	BucketName string `validate:"required,min=1"`
	ObjectKey  string `validate:"required,min=3"`
}
