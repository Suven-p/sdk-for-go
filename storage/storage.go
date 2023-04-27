package storage

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/appwritemodels"
	"strings"
	"github.com/appwrite/sdk-for-go/appwrite"
)

// Storage service
type Storage struct {
	client appwrite.Client
}

func NewStorage(clt appwrite.Client) *Storage {
	return &Storage{
		client: clt,
	}
}

// ListBuckets get a list of all the storage buckets. You can use the query
// params to filter your results.
type ListBucketsOptions struct {
	Queries []interface{}
	Search string
}

type ListBucketsOption func(*ListBucketsOptions)

func WithListBucketsQueries(v []interface{}) ListBucketsOption {
	return func(o *ListBucketsOptions) {
		o.Queries = v
	}
}
func WithListBucketsSearch(v string) ListBucketsOption {
	return func(o *ListBucketsOptions) {
		o.Search = v
	}
}

	
func (srv *Storage) ListBuckets(optionalSetters ...ListBucketsOption)  (*appwriteModel.BucketList, error) {
	path := "/storage/buckets"

	options := &ListBucketsOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["queries"] = options.Queries
	params["search"] = options.Search
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.BucketList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.BucketList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateBucket create a new storage bucket.
type CreateBucketOptions struct {
	Permissions []interface{}
	FileSecurity bool
	Enabled bool
	MaximumFileSize int
	AllowedFileExtensions []interface{}
	Compression string
	Encryption bool
	Antivirus bool
}

type CreateBucketOption func(*CreateBucketOptions)

func WithCreateBucketPermissions(v []interface{}) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.Permissions = v
	}
}
func WithCreateBucketFileSecurity(v bool) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.FileSecurity = v
	}
}
func WithCreateBucketEnabled(v bool) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.Enabled = v
	}
}
func WithCreateBucketMaximumFileSize(v int) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.MaximumFileSize = v
	}
}
func WithCreateBucketAllowedFileExtensions(v []interface{}) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.AllowedFileExtensions = v
	}
}
func WithCreateBucketCompression(v string) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.Compression = v
	}
}
func WithCreateBucketEncryption(v bool) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.Encryption = v
	}
}
func WithCreateBucketAntivirus(v bool) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.Antivirus = v
	}
}

					
func (srv *Storage) CreateBucket(BucketId string, Name string, optionalSetters ...CreateBucketOption)  (*appwriteModel.Bucket, error) {
	path := "/storage/buckets"

	options := &CreateBucketOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["name"] = Name
	params["permissions"] = options.Permissions
	params["fileSecurity"] = options.FileSecurity
	params["enabled"] = options.Enabled
	params["maximumFileSize"] = options.MaximumFileSize
	params["allowedFileExtensions"] = options.AllowedFileExtensions
	params["compression"] = options.Compression
	params["encryption"] = options.Encryption
	params["antivirus"] = options.Antivirus
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Bucket
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Bucket)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetBucket get a storage bucket by its unique ID. This endpoint response
// returns a JSON object with the storage bucket metadata.
type GetBucketOptions struct {
}

type GetBucketOption func(*GetBucketOptions)


	
func (srv *Storage) GetBucket(BucketId string)  (*appwriteModel.Bucket, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}")

	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Bucket
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Bucket)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateBucket update a storage bucket by its unique ID.
type UpdateBucketOptions struct {
	Permissions []interface{}
	FileSecurity bool
	Enabled bool
	MaximumFileSize int
	AllowedFileExtensions []interface{}
	Compression string
	Encryption bool
	Antivirus bool
}

type UpdateBucketOption func(*UpdateBucketOptions)

func WithUpdateBucketPermissions(v []interface{}) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.Permissions = v
	}
}
func WithUpdateBucketFileSecurity(v bool) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.FileSecurity = v
	}
}
func WithUpdateBucketEnabled(v bool) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.Enabled = v
	}
}
func WithUpdateBucketMaximumFileSize(v int) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.MaximumFileSize = v
	}
}
func WithUpdateBucketAllowedFileExtensions(v []interface{}) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.AllowedFileExtensions = v
	}
}
func WithUpdateBucketCompression(v string) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.Compression = v
	}
}
func WithUpdateBucketEncryption(v bool) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.Encryption = v
	}
}
func WithUpdateBucketAntivirus(v bool) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.Antivirus = v
	}
}

					
func (srv *Storage) UpdateBucket(BucketId string, Name string, optionalSetters ...UpdateBucketOption)  (*appwriteModel.Bucket, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}")

	options := &UpdateBucketOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["name"] = Name
	params["permissions"] = options.Permissions
	params["fileSecurity"] = options.FileSecurity
	params["enabled"] = options.Enabled
	params["maximumFileSize"] = options.MaximumFileSize
	params["allowedFileExtensions"] = options.AllowedFileExtensions
	params["compression"] = options.Compression
	params["encryption"] = options.Encryption
	params["antivirus"] = options.Antivirus
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.Bucket
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.Bucket)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// DeleteBucket delete a storage bucket by its unique ID.
type DeleteBucketOptions struct {
}

type DeleteBucketOption func(*DeleteBucketOptions)


	
func (srv *Storage) DeleteBucket(BucketId string)  (*interface{}, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}")

	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// ListFiles get a list of all the user files. You can use the query params to
// filter your results.
type ListFilesOptions struct {
	Queries []interface{}
	Search string
}

type ListFilesOption func(*ListFilesOptions)

func WithListFilesQueries(v []interface{}) ListFilesOption {
	return func(o *ListFilesOptions) {
		o.Queries = v
	}
}
func WithListFilesSearch(v string) ListFilesOption {
	return func(o *ListFilesOptions) {
		o.Search = v
	}
}

			
func (srv *Storage) ListFiles(BucketId string, optionalSetters ...ListFilesOption)  (*appwriteModel.FileList, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}/files")

	options := &ListFilesOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["queries"] = options.Queries
	params["search"] = options.Search
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.FileList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.FileList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// CreateFile create a new file. Before using this route, you should create a
// new bucket resource using either a [server
// integration](/docs/server/storage#storageCreateBucket) API or directly from
// your Appwrite console.
// 
// Larger files should be uploaded using multiple requests with the
// [content-range](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Range)
// header to send a partial request with a maximum supported chunk of `5MB`.
// The `content-range` header values should always be in bytes.
// 
// When the first request is sent, the server will return the **File** object,
// and the subsequent part request must include the file's **id** in
// `x-appwrite-id` header to allow the server to know that the partial upload
// is for the existing file and not for a new one.
// 
// If you're creating a new file using one of the Appwrite SDKs, all the
// chunking logic will be managed by the SDK internally.
// 
type CreateFileOptions struct {
	Permissions []interface{}
}

type CreateFileOption func(*CreateFileOptions)

func WithCreateFilePermissions(v []interface{}) CreateFileOption {
	return func(o *CreateFileOptions) {
		o.Permissions = v
	}
}

							
func (srv *Storage) CreateFile(BucketId string, FileId string, File appwrite.InputFile, optionalSetters ...CreateFileOption)  (*appwriteModel.File, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}/files")

	options := &CreateFileOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	params["file"] = File
	params["permissions"] = options.Permissions
	headers := map[string]interface{}{
		"content-type": "multipart/form-data",
	}

    paramName := "file"


    uploadId := ""
    uploadId = FileId

    resp, err := srv.client.FileUpload(path, headers, params, paramName, uploadId)
    if err != nil {
		return nil, err
	}
	parsed, ok := resp.Result.(appwriteModel.File)
	if !ok {
		return nil, errors.New("Unexpected response type")
	}
	return &parsed, nil

}
// GetFile get a file by its unique ID. This endpoint response returns a JSON
// object with the file metadata.
type GetFileOptions struct {
}

type GetFileOption func(*GetFileOptions)


			
func (srv *Storage) GetFile(BucketId string, FileId string)  (*appwriteModel.File, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")

	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.File
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.File)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// UpdateFile update a file by its unique ID. Only users with write
// permissions have access to update this resource.
type UpdateFileOptions struct {
	Permissions []interface{}
}

type UpdateFileOption func(*UpdateFileOptions)

func WithUpdateFilePermissions(v []interface{}) UpdateFileOption {
	return func(o *UpdateFileOptions) {
		o.Permissions = v
	}
}

					
func (srv *Storage) UpdateFile(BucketId string, FileId string, optionalSetters ...UpdateFileOption)  (*appwriteModel.File, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")

	options := &UpdateFileOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	params["permissions"] = options.Permissions
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed appwriteModel.File
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(appwriteModel.File)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// DeleteFile delete a file by its unique ID. Only users with write
// permissions have access to delete this resource.
type DeleteFileOptions struct {
}

type DeleteFileOption func(*DeleteFileOptions)


			
func (srv *Storage) DeleteFile(BucketId string, FileId string)  (*interface{}, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")

	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetFileDownload get a file content by its unique ID. The endpoint response
// return with a 'Content-Disposition: attachment' header that tells the
// browser to start downloading the file to user downloads directory.
type GetFileDownloadOptions struct {
}

type GetFileDownloadOption func(*GetFileDownloadOptions)


			
func (srv *Storage) GetFileDownload(BucketId string, FileId string)  (*[]byte, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/download")

	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetFilePreview get a file preview image. Currently, this method supports
// preview for image files (jpg, png, and gif), other supported formats, like
// pdf, docs, slides, and spreadsheets, will return the file icon image. You
// can also pass query string arguments for cutting and resizing your preview
// image. Preview is supported only for image files smaller than 10MB.
type GetFilePreviewOptions struct {
	Width int
	Height int
	Gravity string
	Quality int
	BorderWidth int
	BorderColor string
	BorderRadius int
	Opacity float64
	Rotation int
	Background string
	Output string
}

type GetFilePreviewOption func(*GetFilePreviewOptions)

func WithGetFilePreviewWidth(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Width = v
	}
}
func WithGetFilePreviewHeight(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Height = v
	}
}
func WithGetFilePreviewGravity(v string) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Gravity = v
	}
}
func WithGetFilePreviewQuality(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Quality = v
	}
}
func WithGetFilePreviewBorderWidth(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.BorderWidth = v
	}
}
func WithGetFilePreviewBorderColor(v string) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.BorderColor = v
	}
}
func WithGetFilePreviewBorderRadius(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.BorderRadius = v
	}
}
func WithGetFilePreviewOpacity(v float64) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Opacity = v
	}
}
func WithGetFilePreviewRotation(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Rotation = v
	}
}
func WithGetFilePreviewBackground(v string) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Background = v
	}
}
func WithGetFilePreviewOutput(v string) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Output = v
	}
}

					
func (srv *Storage) GetFilePreview(BucketId string, FileId string, optionalSetters ...GetFilePreviewOption)  (*[]byte, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/preview")

	options := &GetFilePreviewOptions{}
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	params["width"] = options.Width
	params["height"] = options.Height
	params["gravity"] = options.Gravity
	params["quality"] = options.Quality
	params["borderWidth"] = options.BorderWidth
	params["borderColor"] = options.BorderColor
	params["borderRadius"] = options.BorderRadius
	params["opacity"] = options.Opacity
	params["rotation"] = options.Rotation
	params["background"] = options.Background
	params["output"] = options.Output
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
// GetFileView get a file content by its unique ID. This endpoint is similar
// to the download method but returns with no  'Content-Disposition:
// attachment' header.
type GetFileViewOptions struct {
}

type GetFileViewOption func(*GetFileViewOptions)


			
func (srv *Storage) GetFileView(BucketId string, FileId string)  (*[]byte, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/view")

	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
