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

func (r *CallLogRepository) Create(bell_id int, store_id int) error {
	callLog := entity.CallLog{
		BellID:  bell_id,
		StoreID: store_id,
		Status:  "active", //デフォルトでactiveにする
	}
	if err := r.DB.Create(&callLog).Error; err != nil {
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

func (r *CallLogRepository) UpdateStatus(bell_id int, status string) error {
	//最新の呼び出しログを取得
	var callLog entity.CallLog
	if err := r.DB.Where("bell_id = ?", bell_id).Order("id desc").First(&callLog).Error; err != nil {
		return err
	}

	//ステータスを更新
	callLog.Status = status
	if err := r.DB.Save(&callLog).Error; err != nil {
		return err
	}

	return nil
}
