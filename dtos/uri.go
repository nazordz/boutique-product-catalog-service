package dtos

type CatalogUri struct {
	ID string `uri:"id" binding:"required,uuid"`
}
