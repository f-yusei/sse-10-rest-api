package mapper

import (
	"gcp_go_cloud_run/app/dto"
	"gcp_go_cloud_run/app/infrastructure/mysql/entity"
)

func StoreToDTO(store *entity.Store) *dto.StoreDTO {
	if store == nil {
		return nil
	}

	return &dto.StoreDTO{
		ID:             store.ID,
		Name:           store.Name,
		DisplayMessage: store.DisplayMessage,
	}
}

func DTOToStore(storeDTO *dto.StoreDTO) *entity.Store {
	if storeDTO == nil {
		return nil
	}

	return &entity.Store{
		ID:             storeDTO.ID,
		Name:           storeDTO.Name,
		DisplayMessage: storeDTO.DisplayMessage,
	}
}
