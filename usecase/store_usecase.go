package usecase

import (
	"gcp_go_cloud_run/app/domain/repository"
	"gcp_go_cloud_run/app/dto"
	"gcp_go_cloud_run/app/mapper"
)

type StoreService struct {
	StoreRepository repository.IStoreRepository
}

func NewStoreService(storeRepo repository.IStoreRepository) *StoreService {
	return &StoreService{
		StoreRepository: storeRepo,
	}
}

func (s *StoreService) GetAllStores() ([]*dto.StoreDTO, error) {
	stores, err := s.StoreRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var storesDTO []*dto.StoreDTO
	for _, store := range stores {
		storesDTO = append(storesDTO, mapper.StoreToDTO(store))
	}

	return storesDTO, nil
}

func (s *StoreService) GetStoreByID(id int) (*dto.StoreDTO, error) {
	store, err := s.StoreRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return mapper.StoreToDTO(store), nil
}

func (s *StoreService) CreateStore(storeDTO *dto.StoreDTO) (*dto.StoreDTO, error) {
	storeEntity := mapper.DTOToStore(storeDTO)

	err := s.StoreRepository.Create(storeEntity)
	if err != nil {
		return nil, err
	}

	return mapper.StoreToDTO(storeEntity), nil
}

func (s *StoreService) UpdateStore(storeDTO *dto.StoreDTO) (*dto.StoreDTO, error) {
	storeEntity := mapper.DTOToStore(storeDTO)

	err := s.StoreRepository.Update(storeEntity)
	if err != nil {
		return nil, err
	}

	return mapper.StoreToDTO(storeEntity), nil
}

func (s *StoreService) DeleteStore(id int) error {
	return s.StoreRepository.Delete(id)
}

func (s *StoreService) UpdateDisplayMessage(id int, displayMessage string) error {
	return s.StoreRepository.UpdateDisplayMessage(id, displayMessage)
}
