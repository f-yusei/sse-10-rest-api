package mysql

import (
	"gcp_go_cloud_run/app/domain/repository"
	"gcp_go_cloud_run/app/infrastructure/mysql/entity"

	"gorm.io/gorm"
)

type BellRepository struct {
	*gorm.DB
}

func NewBellRepository(db *gorm.DB) repository.IBellRepository {
	return &BellRepository{db}
}

func (r *BellRepository) FindAll() ([]*entity.Bell, error) {
	var bells []*entity.Bell
	if err := r.DB.Find(&bells).Error; err != nil {
		return nil, err
	}
	return bells, nil
}

func (r *BellRepository) FindByID(id int) (*entity.Bell, error) {
	var bell entity.Bell
	if err := r.DB.First(&bell, id).Error; err != nil {
		return nil, err
	}
	return &bell, nil
}

func (r *BellRepository) Create(bell *entity.Bell) error {
	if err := r.DB.Create(bell).Error; err != nil {
		return err
	}
	return nil
}

func (r *BellRepository) Update(bell *entity.Bell) error {
	if err := r.DB.Save(bell).Error; err != nil {
		return err
	}
	return nil
}

func (r *BellRepository) Delete(id int) error {
	if err := r.DB.Delete(&entity.Bell{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *BellRepository) GetActiveBells() ([]*entity.Bell, error) {
	var bells []*entity.Bell
	if err := r.DB.Where("status = ?", "calling").Find(&bells).Error; err != nil {
		return nil, err
	}
	return bells, nil
}

func (r *BellRepository) UpdateStatus(id int, status string) error {
	if err := r.DB.Model(&entity.Bell{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *BellRepository) GetStoreIDByBellID(bell_id int) (int, error) {
	var bell entity.Bell
	if err := r.DB.First(&bell, bell_id).Error; err != nil {
		return 0, err
	}
	return bell.StoreID, nil
}
