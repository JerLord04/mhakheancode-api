package repository

import "mhakheancode/api/model"

type TagsRepository interface {
	GetAll() (*[]model.Tag, error)
}
