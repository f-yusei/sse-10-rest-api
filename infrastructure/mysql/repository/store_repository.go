package mysql

import (
	"gcp_go_cloud_run/app/domain/repository"
	"gcp_go_cloud_run/app/infrastructure/mysql/entity"

	"gorm.io/gorm"
)

type StoreRepository struct {
	*gorm.DB
}

func NewStoreRepository(db *gorm.DB) repository.IStoreRepository {
	return &StoreRepository{db}
}

func (r *StoreRepository) FindAll() ([]*entity.Store, error) {
	var stores []*entity.Store
	if err := r.DB.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func (r *StoreRepository) FindByID(id int) (*entity.Store, error) {
	var store entity.Store
	if err := r.DB.First(&store, id).Error; err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *StoreRepository) Create(store *entity.Store) error {
	if err := r.DB.Create(store).Error; err != nil {
		return err
	}
	return nil
}

func (r *StoreRepository) Update(store *entity.Store) error {
	if err := r.DB.Save(store).Error; err != nil {
		return err
	}
	return nil
}

func (r *StoreRepository) Delete(id int) error {
	if err := r.DB.Delete(&entity.Store{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *StoreRepository) UpdateDisplayMessage(id int, displayMessage string) error {
	if err := r.DB.Model(&entity.Store{}).Where("id = ?", id).Update("display_message", displayMessage).Error; err != nil {
		return err
	}
	return nil
}
