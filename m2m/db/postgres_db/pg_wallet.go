package postgres_db

import (
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type walletInterface struct{}

var PgWallet walletInterface

type wallet struct {
	Id int64 `db:"id"`

	// FkOrgLa: foreign_key of the organization in the LPWAN App Server DB
	FkOrgLa int64  `db:"fk_org_la"`
	TypeW   string `db:"type"`

	// balance is updated during the aggregations (always references with a row in the internal_tx table)
	Balance float64 `db:"balance"`

	// tmp_balance is updated per packet transmission (downlink) and is used to check the balance as the balance is not updated.
	TmpBalance float64 `db:"tmp_balance"`

	// During the aggregation, balance will be updated and tmp_balance will get sync with the value of balance.
	// When a topup, withdraw, stake triggers, both tmp_balance and balance will get updated!
}

func (*walletInterface) CreateWalletTable() error {
	_, err := PgDB.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS 
				(SELECT 1 FROM pg_type WHERE typname = 'wallet_type') 
			THEN
				CREATE TYPE WALLET_TYPE AS ENUM (
					'SUPER_ADMIN',
					'SUPER_NODE_INCOME',
					 'STAKE_STORAGE',
					 'USER'
		);
		END IF;
			CREATE TABLE IF NOT EXISTS wallet (
	
			id SERIAL PRIMARY KEY,
			fk_org_la INT UNIQUE NOT NULL, 
			type WALLET_TYPE NOT NULL,
			balance NUMERIC(28,18) NOT NULL DEFAULT 0,
			tmp_balance NUMERIC(28,18) NOT NULL DEFAULT 0
		);

		END$$;
		
	`)
	return errors.Wrap(err, "db/CreateWalletTable")
}

func (*walletInterface) CreateWalletFunctions() error {
	_, err := PgDB.Exec(`

	CREATE OR REPLACE FUNCTION tmp_balance_update_pkt_tx (dv_wallet_id INT, gw_wallet_id INT, amount NUMERIC(28,18)) RETURNS void
		LANGUAGE plpgsql
			AS $$
			BEGIN
			
			UPDATE
				wallet 
			SET
				tmp_balance =  tmp_balance - amount
			WHERE
				id = dv_wallet_id
			;
			UPDATE
				wallet
			SET 
				tmp_balance = tmp_balance + amount
			WHERE
				id = gw_wallet_id
			;
			END;
		$$;
		`)

	return errors.Wrap(err, "db/CreateWalletFunctions")

}

func insertWallet(w wallet) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
		INSERT INTO wallet (
			fk_org_la ,
			type,
			balance,
			tmp_balance ) 
		VALUES 
			($1,$2,$3,$4)
		RETURNING id ;
	`,
		w.FkOrgLa,
		w.TypeW,
		w.Balance,
		w.TmpBalance).Scan(&insertIndex)

	return insertIndex, errors.Wrap(err, "db/insertWallet")
}

func (*walletInterface) InsertWallet(orgId int64, walletType types.WalletType) (insertIndex int64, err error) {
	w := wallet{
		FkOrgLa:    orgId,
		TypeW:      string(walletType),
		Balance:    0.0,
		TmpBalance: 0.0,
	}

	return insertWallet(w)
}

func (*walletInterface) InsertSuperNodeIncomeWallet() (insertIndex int64, err error) {

	id, err := PgWallet.GetWalletIdSuperNodeIncome()
	if err == nil {
		return id, err
	} else if strings.HasSuffix(err.Error(), types.DbError.NoRowQueryRes.Error()) {
		//  there is no existing record for NodeIncome
		w := wallet{
			FkOrgLa:    -1,
			TypeW:      string(types.SUPER_NODE_INCOME),
			Balance:    0.0,
			TmpBalance: 0.0,
		}
		return insertWallet(w)

	} else {
		return 0, errors.Wrap(err, "db/InsertSuperNodeIncomeWallet")
	}

}

func (*walletInterface) InsertStakeStorageWallet() (insertIndex int64, err error) {

	id, err := PgWallet.GetWalletIdStakeStorage()
	if err == nil {
		return id, err
	} else if strings.HasSuffix(err.Error(), types.DbError.NoRowQueryRes.Error()) {
		//  there is no existing record for StakeStorage
		w := wallet{
			FkOrgLa:    -2,
			TypeW:      string(types.STAKE_STORAGE),
			Balance:    0.0,
			TmpBalance: 0.0,
		}
		return insertWallet(w)

	} else {
		return 0, errors.Wrap(err, "db/InsertStakeStorageWallet")
	}

}

func (*walletInterface) GetWalletIdFromOrgId(orgIdLora int64) (int64, error) {
	id := int64(0)
	err := PgDB.QueryRow(
		`SELECT id
		FROM wallet
		WHERE
			fk_org_la = $1;`,
		orgIdLora).Scan(&id)

	return id, errors.Wrap(err, "db/GetWalletIdFromOrgId")
}

func (*walletInterface) GetWalletBalance(walletId int64) (float64, error) {
	balance := float64(0)
	err := PgDB.QueryRow(
		`SELECT tmp_balance
		FROM wallet
		WHERE
			id = $1;`,
		walletId).Scan(&balance)

	return balance, errors.Wrap(err, "db/GetWalletBalance")
}

func (*walletInterface) SyncTmpBalance(walletId int64) (balance float64, err error) {

	err = PgDB.QueryRow(
		`UPDATE
			wallet
		 SET
			tmp_balance = balance
		 WHERE
		  	id = $1
		RETURNING balance;
		`,
		walletId).Scan(&balance)

	return balance, errors.Wrap(err, "db/SyncTmpBalance")
}

func (*walletInterface) getWalletNotTmpeBalance(walletId int64) (float64, error) {
	balance := float64(0)
	err := PgDB.QueryRow(
		`SELECT balance
		FROM wallet
		WHERE
			id = $1;`,
		walletId).Scan(&balance)

	return balance, errors.Wrap(err, "db/getWalletNotTmpBalance")
}

func (*walletInterface) GetWalletIdofActiveAcnt(acntAdr string, externalCur string) (walletId int64, err error) {

	err = PgDB.QueryRow(
		`select 
			w.id as wallet_id 
			from
			wallet w,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet AND
			w.type = 'USER' AND
			ea.fk_ext_currency = ec.id AND
			ea.status = 'ACTIVE' AND
			account_adr = $1 AND
			ec.abv = $2 
		ORDER BY ea.id DESC 
		LIMIT 1 
		;`, acntAdr, externalCur).Scan(&walletId)

	return walletId, errors.Wrap(err, "db/GetWalletIdofActiveAcnt")
}

func getWalletIdofActiveAcntSuperAdmin(acntAdr string, externalCur string) (walletId int64, err error) {

	err = PgDB.QueryRow(
		`select 
			w.id as wallet_id 
			from
			wallet w,ext_account ea,ext_currency ec
		WHERE
			w.id = ea.fk_wallet AND
			w.type = 'SUPER_ADMIN' AND
			ea.fk_ext_currency = ec.id AND
			ea.status = 'ACTIVE' AND
			account_adr = $1 AND
			ec.abv = $2 
		ORDER BY ea.id DESC 
		LIMIT 1 
		;`, acntAdr, externalCur).Scan(&walletId)

	return walletId, errors.Wrap(err, "db/getWalletIdofActiveAcntSuperAdmin")
}

func (*walletInterface) GetWalletIdSuperNode() (walletId int64, err error) {

	err = PgDB.QueryRow(
		`SELECT
			id
		FROM
			wallet 
		WHERE
			type = 'SUPER_ADMIN' 
			ORDER BY id DESC 
			LIMIT 1  
		;`).Scan(&walletId)

	return walletId, errors.Wrap(err, "db/GetWalletIdSuperNode")
}

func (*walletInterface) GetWalletIdSuperNodeIncome() (walletId int64, err error) {

	err = PgDB.QueryRow(
		`SELECT
			id
		FROM
			wallet 
		WHERE
			type = 'SUPER_NODE_INCOME' 
			ORDER BY id DESC 
			LIMIT 1 
		;`).Scan(&walletId)

	return walletId, errors.Wrap(err, "db/GetWalletIdSuperNodeIncome")
}

func (*walletInterface) GetWalletIdStakeStorage() (walletId int64, err error) {

	err = PgDB.QueryRow(
		`SELECT
			id
		FROM
			wallet 
		WHERE
			type = 'STAKE_STORAGE' 
			ORDER BY id DESC 
			LIMIT 1 
		;`).Scan(&walletId)

	return walletId, errors.Wrap(err, "db/GetWalletIdSuperNodeIncome")
}

func (*walletInterface) updateBalanceByWalletId(walletId int64, newBalance float64) error {
	_, err := PgDB.Exec(`
	UPDATE
		wallet 
	SET
		balance = $1
	WHERE
		id = $2
	;
	`, newBalance, walletId)

	return errors.Wrap(err, "db/UpdateBalanceByWalletId")
}

func (*walletInterface) TmpBalanceUpdatePktTx(dvWalletId, gwWalletId int64, amount float64) error {
	_, err := PgDB.Exec(`
	select tmp_balance_update_pkt_tx ($1, $2 , $3)
	`, dvWalletId, gwWalletId, amount)

	return errors.Wrap(err, "db/TmpBalanceUpdatePktTx")
}

func (*walletInterface) GetMaxWalletId() (maxWalletId int64, err error) {
	err = PgDB.QueryRow(
		`SELECT
			COALESCE(MAX (id),0)
		FROM
			wallet 
		;`).Scan(&maxWalletId)

	return maxWalletId, errors.Wrap(err, "db/GetMaxWalletId")
}
func (*walletInterface) GetSupernodeIncomeAmount(since time.Time, until time.Time) (val float64, err error) {

	snIncomeWlt, err := PgWallet.GetWalletIdSuperNodeIncome()
	if err != nil {
		return 0, errors.Wrap(err, "Impossible to get supernode walletId db/getSupernodeIncome/ ")
	}

	err = PgDB.QueryRow(
		`SELECT 
			COALESCE(SUM(value),0) as sum_income
		FROM
			internal_tx
		WHERE
			fk_wallet_receiver = $1
		AND
			time_tx 
		BETWEEN  
			$2
		AND
			$3
		;`,
		snIncomeWlt,
		since,
		until,
	).Scan(&val)

	return val, errors.Wrap(err, "db/getSupernodeIncome")

}
