package mysql

import (
	"gcp_go_cloud_run/app/domain/repository"
	"gcp_go_cloud_run/app/infrastructure/mysql/entity"

	"gorm.io/gorm"
)

type CallLogRepository struct {
	*gorm.DB
}

func NewCallLogRepository(db *gorm.DB) repository.ICallLogRepository {
	return &CallLogRepository{db}
}

func (r *CallLogRepository) FindAll() ([]*entity.CallLog, error) {
	var callLogs []*entity.CallLog
	if err := r.DB.Find(&callLogs).Error; err != nil {
		return nil, err
	}
	return callLogs, nil
}

func (r *CallLogRepository) FindByID(id int) (*entity.CallLog, error) {
	var callLog entity.CallLog
	if err := r.DB.First(&callLog, id).Error; err != nil {
		return nil, err
	}
	return &callLog, nil
}

func (r *CallLogRepository) Create(callLog *entity.CallLog) error {
	if err := r.DB.Create(callLog).Error; err != nil {
		return err
	}
	return nil
}

func (r *CallLogRepository) Update(callLog *entity.CallLog) error {
	if err := r.DB.Save(callLog).Error; err != nil {
		return err
	}
	return nil
}
