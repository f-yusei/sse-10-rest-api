package usecase

import (
	"gcp_go_cloud_run/app/domain/repository"
	"gcp_go_cloud_run/app/dto"
	"gcp_go_cloud_run/app/mapper"
)

// BellService はBellに関連するビジネスロジックを提供します。
type BellService struct {
	BellRepository repository.IBellRepository
}

// NewBellService はBellServiceの新しいインスタンスを作成します。
func NewBellService(bellRepo repository.IBellRepository) *BellService {
	return &BellService{
		BellRepository: bellRepo,
	}
}

// GetActiveBells は現在呼び出し状態のベルを取得し、それをDTOに変換して返します。
func (s *BellService) GetActiveBells() ([]*dto.BellDTO, error) {
	// 現在呼び出し状態のベルを取得
	activeBells, err := s.BellRepository.GetActiveBells()
	if err != nil {
		return nil, err
	}

	// エンティティのリストをDTOに変換
	var activeBellsDTO []*dto.BellDTO
	for _, bell := range activeBells {
		activeBellsDTO = append(activeBellsDTO, mapper.BellToDTO(bell))
	}

	return activeBellsDTO, nil
}

// CreateBell は新しいベルを作成します。
func (s *BellService) CreateBell(bellDTO *dto.BellDTO) (*dto.BellDTO, error) {
	// DTOをエンティティに変換
	bellEntity := mapper.DTOToBell(bellDTO)

	// ベルをリポジトリに保存
	err := s.BellRepository.Create(bellEntity)
	if err != nil {
		return nil, err
	}

	// 作成したエンティティをDTOに変換して返す
	return mapper.BellToDTO(bellEntity), nil
}

// UpdateBell はベルの情報を更新します。
func (s *BellService) UpdateBell(bellDTO *dto.BellDTO) (*dto.BellDTO, error) {
	// DTOをエンティティに変換
	bellEntity := mapper.DTOToBell(bellDTO)

	// ベルをリポジトリで更新
	err := s.BellRepository.Update(bellEntity)
	if err != nil {
		return nil, err
	}

	// 更新したエンティティをDTOに変換して返す
	return mapper.BellToDTO(bellEntity), nil
}

// DeleteBell は指定されたIDのベルを削除します。
func (s *BellService) DeleteBell(id int) error {
	// ベルをリポジトリで削除
	return s.BellRepository.Delete(id)
}
