package mapper

import (
	"gcp_go_cloud_run/app/dto"
	"gcp_go_cloud_run/app/infrastructure/mysql/entity"
)

func BellToDTO(bell *entity.Bell) *dto.BellDTO {
	if bell == nil {
		return nil
	}

	return &dto.BellDTO{
		ID:           bell.ID,
		StoreID:      bell.StoreID,
		DeviceID:     bell.DeviceID,
		Status:       bell.Status,
		LastCalledAt: bell.LastCalledAt,
	}
}

func DTOToBell(bellDTO *dto.BellDTO) *entity.Bell {
	if bellDTO == nil {
		return nil
	}

	return &entity.Bell{
		ID:           bellDTO.ID,
		StoreID:      bellDTO.StoreID,
		DeviceID:     bellDTO.DeviceID,
		Status:       bellDTO.Status,
		LastCalledAt: bellDTO.LastCalledAt,
	}
}
