package usecase

import (
	"fmt"
	"gcp_go_cloud_run/app/domain/repository"
	"gcp_go_cloud_run/app/dto"
	"gcp_go_cloud_run/app/mapper"
)

type BellService struct {
	BellRepository    repository.IBellRepository
	CallLogRepository repository.ICallLogRepository
}

/*
type IBellRepository interface {
    FindAll() ([]*entity.Bell, error)
    FindByID(id int) (*entity.Bell, error)
    Create(bell *entity.Bell) error
    Update(bell *entity.Bell) error
    Delete(id int) error
    GetActiveBells() ([]*entity.Bell, error)
}
*/

func NewBellService(bellRepo repository.IBellRepository) *BellService {
	return &BellService{
		BellRepository: bellRepo,
	}
}

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

func (s *BellService) CreateBell(bellDTO *dto.BellDTO) (*dto.BellDTO, error) {
	bellEntity := mapper.DTOToBell(bellDTO)

	err := s.BellRepository.Create(bellEntity)
	if err != nil {
		return nil, err
	}

	return mapper.BellToDTO(bellEntity), nil
}

func (s *BellService) UpdateBell(bellDTO *dto.BellDTO) (*dto.BellDTO, error) {
	bellEntity := mapper.DTOToBell(bellDTO)

	err := s.BellRepository.Update(bellEntity)
	if err != nil {
		return nil, err
	}

	return mapper.BellToDTO(bellEntity), nil
}

func (s *BellService) DeleteBell(id int) error {
	return s.BellRepository.Delete(id)
}

/*
 1. 呼び出し開始
    ○ お店側がAPIを叩いて、指定のベルを呼び出します。
    ○ call_log に新しいレコードが作成され、bell の status が calling に更新されます。
 2. 親機のポーリング
    ○ 親機がポーリングを通してAPIを叩くと、call_log テーブルで status = 'active' のレコードを取得し、対象の bell の device_id を確認します。
 3. 呼び出し終了

呼び出しが終了したら、APIを介して call_log の status を completed に更新し、bell の status を idle に戻します。
*/

func (s *BellService) CallBell(bellID int) error {
	fmt.Printf("bellId is %d", bellID)
	//bellIDからstoreIDを取得
	storeID, err := s.BellRepository.GetStoreIDByBellID(bellID)
	if err != nil {
		fmt.Printf("error:%s", err)
		return err
	}

	//call_logに新しいレコードを追加
	err = s.CallLogRepository.Create(storeID, bellID)
	if err != nil {
		fmt.Printf("error:%s", err)
		return err
	}

	//bellのstatusをcallingに更新
	err = s.BellRepository.UpdateStatus(bellID, "calling")
	if err != nil {
		fmt.Printf("error:%s", err)
		return err
	}

	return nil
}

func (s *BellService) CompleteCall(bellID int) error {
	//call_logのstatusをcompletedに更新
	err := s.CallLogRepository.UpdateStatus(bellID, "completed")
	if err != nil {
		return err
	}

	//bellのstatusをidleに更新
	err = s.BellRepository.UpdateStatus(bellID, "idle")
	if err != nil {
		return err
	}

	return nil
}
