package repository

import "gcp_go_cloud_run/app/infrastructure/mysql/entity"

type IStoreRepository interface {
	FindAll() ([]*entity.Store, error)
	FindByID(id int) (*entity.Store, error)
	Create(store *entity.Store) error
	Update(store *entity.Store) error
	Delete(id int) error
	UpdateDisplayMessage(id int, displayMessage string) error
}
