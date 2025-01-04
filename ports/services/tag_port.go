package services

import "mhakheancode/api/model"

type TagsPort interface {
	GetAllTags() (*[]model.Tag, error)
}
