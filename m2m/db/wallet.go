package db

import (
	"time"

	pg "github.com/mxc-foundation/mxprotocol-server/m2m/db/postgres_db"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type walletDBInterface interface {
	CreateWalletTable() error
	CreateWalletFunctions() error
	InsertWallet(orgId int64, walletType types.WalletType) (insertIndex int64, err error)
	InsertSuperNodeIncomeWallet() (insertIndex int64, err error)
	InsertStakeStorageWallet() (insertIndex int64, err error)
	GetWalletIdFromOrgId(orgIdLora int64) (int64, error)
	GetWalletBalance(walletId int64) (float64, error)
	SyncTmpBalance(walletId int64) (balance float64, err error)
	GetWalletIdofActiveAcnt(acntAdr string, externalCur string) (walletId int64, err error)
	GetWalletIdSuperNode() (walletId int64, err error)
	GetWalletIdSuperNodeIncome() (walletId int64, err error)
	GetWalletIdStakeStorage() (walletId int64, err error)
	TmpBalanceUpdatePktTx(dvWalletId, gwWalletId int64, amount float64) error
	GetMaxWalletId() (maxWalletId int64, err error)
	GetSupernodeIncomeAmount(since time.Time, until time.Time) (val float64, err error)
}

var Wallet = walletDBInterface(&pg.PgWallet)
