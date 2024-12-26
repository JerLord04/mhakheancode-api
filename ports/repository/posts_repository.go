package repository

import "mhakheancode/api/model"

type PostRepository interface {
	Save(model *model.Post) error
}
