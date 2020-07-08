package model

type Response struct {
	HTTPStatus int
	Response   struct {
		Items []ContentMarketing
	}
}
