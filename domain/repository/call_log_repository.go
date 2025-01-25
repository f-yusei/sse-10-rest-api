package repository

import "gcp_go_cloud_run/app/infrastructure/mysql/entity"

type ICallLogRepository interface {
	FindAll() ([]*entity.CallLog, error)
	FindByID(id int) (*entity.CallLog, error)
	Create(callLog *entity.CallLog) error
	Update(callLog *entity.CallLog) error
}
