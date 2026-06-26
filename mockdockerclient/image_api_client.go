package mockdockerclient

import (
	"context"
	"io"
	"iter"

	"github.com/moby/moby/api/types/jsonstream"
	"github.com/moby/moby/client"
)

var _ client.ImageAPIClient = (*ImageAPIClient)(nil)

// ImageAPIClient is a mock implementation of the Docker's client.ImageAPIClient interface
type ImageAPIClient struct {
	PullResp *MockImagePullResponse
}

// MockImagePullResponse is a mock implementation of the client.ImagePullResponse interface
type MockImagePullResponse struct {
	io.ReadCloser
}

// Wait implements client.ImagePullResponse
func (ip *MockImagePullResponse) Wait(context.Context) error {
	return nil
}

// JSONMessages implements client.ImagePullResponse
func (ip *MockImagePullResponse) JSONMessages(context.Context) iter.Seq2[jsonstream.Message, error] {
	return nil
}

// ImageHistory is a mock implementation of Docker's client.ImageAPIClient.ImageHistory()
//
// TODO: properly implement
func (c *ImageAPIClient) ImageHistory(context.Context, string, ...client.ImageHistoryOption) (client.ImageHistoryResult, error) {
	return client.ImageHistoryResult{}, nil
}

// ImageImport is a mock implementation of Docker's client.ImageAPIClient.ImageImport()
//
// TODO: properly implement
func (c *ImageAPIClient) ImageImport(context.Context, client.ImageImportSource, string,
	client.ImageImportOptions) (client.ImageImportResult, error) {
	return nil, nil
}

// ImageInspect is a mock implementation of Docker's client.ImageAPIClient.ImageInspect()
//
// TODO: properly implement
func (c *ImageAPIClient) ImageInspect(context.Context, string, ...client.ImageInspectOption) (client.ImageInspectResult, error) {
	return client.ImageInspectResult{}, nil
}

// ImageList is a mock implementation of Docker's client.ImageAPIClient.ImageList()
//
// TODO: properly implement
func (c *ImageAPIClient) ImageList(context.Context, client.ImageListOptions) (client.ImageListResult, error) {
	return client.ImageListResult{}, nil
}

// ImageLoad is a mock implementation of Docker's client.ImageAPIClient.ImageLoad()
//
// TODO: properly implement
func (c *ImageAPIClient) ImageLoad(context.Context, io.Reader, ...client.ImageLoadOption) (client.ImageLoadResult, error) {
	return nil, nil
}

// ImagePull is a mock implementation of Docker's client.ImageAPIClient.ImagePull()
func (c *ImageAPIClient) ImagePull(context.Context, string, client.ImagePullOptions) (client.ImagePullResponse, error) {
	if c.PullResp == nil {
		return nil, Err
	}
	return c.PullResp, nil
}

// ImagePush is a mock implementation of Docker's client.ImageAPIClient.ImagePush()
//
// TODO: properly implement
func (c *ImageAPIClient) ImagePush(context.Context, string, client.ImagePushOptions) (client.ImagePushResponse, error) {
	return nil, nil
}

// ImageRemove is a mock implementation of Docker's client.ImageAPIClient.ImageRemove()
//
// TODO: properly implement
func (c *ImageAPIClient) ImageRemove(context.Context, string,
	client.ImageRemoveOptions) (client.ImageRemoveResult, error) {
	return client.ImageRemoveResult{}, nil
}

// ImageSearch is a mock implementation of Docker's client.ImageAPIClient.ImageSearch() (via RegistrySearchClient)
//
// TODO: properly implement
func (c *ImageAPIClient) ImageSearch(context.Context, string,
	client.ImageSearchOptions) (client.ImageSearchResult, error) {
	return client.ImageSearchResult{}, nil
}

// ImageSave is a mock implementation of Docker's client.ImageAPIClient.ImageSave()
//
// TODO: properly implement
func (c *ImageAPIClient) ImageSave(context.Context, []string, ...client.ImageSaveOption) (client.ImageSaveResult, error) {
	return nil, nil
}

// ImageTag is a mock implementation of Docker's client.ImageAPIClient.ImageTag()
//
// TODO: properly implement
func (c *ImageAPIClient) ImageTag(context.Context, client.ImageTagOptions) (client.ImageTagResult, error) {
	return client.ImageTagResult{}, nil
}

// ImagePrune is a mock implementation of Docker's client.ImageAPIClient.ImagePrune()
//
// TODO: properly implement
func (c *ImageAPIClient) ImagePrune(context.Context, client.ImagePruneOptions) (client.ImagePruneResult, error) {
	return client.ImagePruneResult{}, nil
}

// ImageAttestations is a mock implementation of Docker's client.ImageAPIClient.ImageAttestations()
//
// TODO: properly implement
func (c *ImageAPIClient) ImageAttestations(context.Context, string, ...client.ImageAttestationsOption) (client.ImageAttestationsResult, error) {
	return client.ImageAttestationsResult{}, nil
}
