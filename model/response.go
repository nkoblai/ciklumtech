package model

// Response represents response from external API calls
type Response struct {
	HTTPStatus int
	Response   struct {
		Items []ContentMarketing
	}
}
