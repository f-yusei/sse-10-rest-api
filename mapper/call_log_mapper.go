package mapper

import (
	"gcp_go_cloud_run/app/dto"
	"gcp_go_cloud_run/app/infrastructure/mysql/entity"
)

func CallLogToDTO(callLog *entity.CallLog) *dto.CallLogDTO {
	if callLog == nil {
		return nil
	}

	return &dto.CallLogDTO{
		ID:       callLog.ID,
		BellID:   callLog.BellID,
		StoreID:  callLog.StoreID,
		CalledAt: callLog.CalledAt,
		Status:   callLog.Status,
	}
}

func DTOtoCallLog(callLogDTO *dto.CallLogDTO) *entity.CallLog {
	if callLogDTO == nil {
		return nil
	}

	return &entity.CallLog{
		ID:       callLogDTO.ID,
		BellID:   callLogDTO.BellID,
		StoreID:  callLogDTO.StoreID,
		CalledAt: callLogDTO.CalledAt,
		Status:   callLogDTO.Status,
	}
}
