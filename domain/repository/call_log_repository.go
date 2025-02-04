package repository

import "gcp_go_cloud_run/app/infrastructure/mysql/entity"

type ICallLogRepository interface {
	FindAll() ([]*entity.CallLog, error)
	FindByID(id int) (*entity.CallLog, error)
	Create(bell_id int, store_id int) error
	Update(callLog *entity.CallLog) error
	UpdateStatus(bell_id int, status string) error
}
