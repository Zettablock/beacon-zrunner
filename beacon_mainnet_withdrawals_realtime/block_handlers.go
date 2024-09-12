package main

import (
	"sort"

	"github.com/Zettablock/beacon-zrunner/dao"
	"github.com/Zettablock/zsource/dao/beacon"
	"github.com/Zettablock/zsource/utils"
)

func HandleWithdrawalsRealtimeByBlock(blockNumber int64, deps *utils.Deps) (bool, error) {
	var withdrawals []beacon.Withdrawal

	err := deps.SourceDB.Table(beacon.TableNameWithdrawal).Where("slot_number = ?", blockNumber).Find(&withdrawals).Limit(100).Error
	if err != nil {
		deps.Logger.Error("failed to get withdrawals", "block number", blockNumber, "error", err)
		return false, err

	}

	sortWithdrawalsByIndex(withdrawals)

	withdrawalsRealtime := make([]dao.WithdrawalsEnhancedRealtime, len(withdrawals))
	for i, withdrawal := range withdrawals {
		withdrawalsRealtime[i] = dao.WithdrawalsEnhancedRealtime{
			SlotNumber:     withdrawal.SlotNumber,
			Index:          withdrawal.Index,
			ValidatorIndex: withdrawal.ValidatorIndex,
			IndexPosition:  int32(i),
			Address:        withdrawal.Address,
			Amount:         withdrawal.Amount,
			BlockTime:      withdrawal.BlockTime,
			BlockNumber:    withdrawal.BlockNumber,
			BlockHash:      withdrawal.BlockHash,
			BlockDate:      withdrawal.BlockDate,
		}
	}

	if len(withdrawalsRealtime) == 0 {
		return false, nil
	}

	err = deps.DestinationDB.Table(dao.TableNameWithdrawalsEnhancedRealtime).Save(&withdrawalsRealtime).Error
	if err != nil {
		deps.Logger.Error("failed to save withdrawals enhanced realtime", "block number", blockNumber, "error", err)
		return false, err
	}

	return false, nil
}

func sortWithdrawalsByIndex(withdrawals []beacon.Withdrawal) {
	// sort by index
	sort.Slice(withdrawals, func(i, j int) bool {
		return withdrawals[i].Index < withdrawals[j].Index
	})
}
