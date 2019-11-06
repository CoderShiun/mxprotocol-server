package db

import (
	pg "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db/postgres_db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type stakeRevenueDBInterface interface {
	CreateStakeRevenueTable() error
	InsertStakeRevenue(stakeId int64, stakeReveneuPeriodId int64, revenueAmount float64) (insertIndex int64, err error)
	GetStakeReveneuHistory(walletId int64, offset int64, limit int64) (stakeRevenueHists []types.StakeRevenueHist, err error)

	GetStakeReveneuHistoryCnt(walletId int64) (recCnt int64, err error)
}

var StakeRevenue = stakeRevenueDBInterface(&pg.PgStakeRevenue)