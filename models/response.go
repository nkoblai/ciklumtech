package models

type Response struct {
	HTTPStatus int
	Response   struct {
		Items []ContentMarketing
	}
}
