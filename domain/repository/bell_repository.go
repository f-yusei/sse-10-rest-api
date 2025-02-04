package repository

import "gcp_go_cloud_run/app/infrastructure/mysql/entity"

type IBellRepository interface {
	FindAll() ([]*entity.Bell, error)
	FindByID(id int) (*entity.Bell, error)
	Create(bell *entity.Bell) error
	Update(bell *entity.Bell) error
	UpdateStatus(id int, status string) error
	Delete(id int) error
	GetActiveBells() ([]*entity.Bell, error)
	GetStoreIDByBellID(bell_id int) (store_id int, err error)
}
