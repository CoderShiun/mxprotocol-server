package postgres_db

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
)

type aggWalletUsageInterface struct{}

var PgAggWalletUsage aggWalletUsageInterface

func (*aggWalletUsageInterface) CreateAggWltUsgTable() error {
	_, err := PgDB.Exec(`
	
		CREATE TABLE IF NOT EXISTS agg_wallet_usage (
			id SERIAL PRIMARY KEY,
			fk_agg_period INT REFERENCES agg_period (id) NOT NULL,
			fk_wallet INT REFERENCES wallet (id) NOT NULL,
			dl_cnt_dv INT    DEFAULT 0 ,
			dl_cnt_dv_free INT    DEFAULT 0 ,
			ul_cnt_dv     INT DEFAULT 0,
			ul_cnt_dv_free INT DEFAULT 0,
			dl_cnt_gw    INT DEFAULT 0,
			dl_cnt_gw_free INT  DEFAULT 0,
			ul_cnt_gw INT  DEFAULT 0,
			ul_cnt_gw_free INT  DEFAULT 0,
			spend  NUMERIC(28,18) DEFAULT 0,
			income  NUMERIC(28,18) DEFAULT 0,
			balance_increase  NUMERIC(28,18) DEFAULT 0,
			updated_balance  NUMERIC(28,18) DEFAULT 0

		);		
	`)
	return errors.Wrap(err, "db/pg_agg_wallet_usage/CreateAggWltUsgTable")

}

func (*aggWalletUsageInterface) InsertAggWltUsg(awu types.AggWltUsg) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
		INSERT INTO agg_wallet_usage (
			fk_wallet,
			fk_agg_period,
			dl_cnt_dv,
			dl_cnt_dv_free,
			ul_cnt_dv,
			ul_cnt_dv_free,
			dl_cnt_gw,
			dl_cnt_gw_free,
			ul_cnt_gw,
			ul_cnt_gw_free,
			spend,
			income,
			balance_increase,
			updated_balance
			) 
		VALUES 
			($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
		RETURNING id ;
	`,
		awu.FkWallet,
		awu.FkAggPeriod,
		awu.DlCntDv,
		awu.DlCntDvFree,
		awu.UlCntDv,
		awu.UlCntDvFree,
		awu.DlCntGw,
		awu.DlCntGwFree,
		awu.UlCntGw,
		awu.UlCntGwFree,
		awu.Spend,
		awu.Income,
		awu.BalanceIncrease,
		awu.UpdatedBalance,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_agg_wallet_usage/InsertAggWltUsg")
}

func (*aggWalletUsageInterface) GetWalletUsageHist(walletId int64, offset int64, limit int64) (awuList []types.AggWltUsg, err error) {

	rows, err := PgDB.Query(
		`SELECT
			awu.id,
			awu.fk_agg_period,
			awu.fk_wallet,
			awu.dl_cnt_dv,
			awu.dl_cnt_dv_free , 
			awu.ul_cnt_dv,  
			awu.ul_cnt_dv_free,
			awu.dl_cnt_gw,
			awu.dl_cnt_gw_free,
			awu.ul_cnt_gw, 
			awu.ul_cnt_gw_free,
			awu.spend,  
			awu.income, 
			awu.balance_increase,
			awu.updated_balance,  
			ap.start_at,
			ap.duration_minutes

		FROM
			agg_wallet_usage awu,
			agg_period ap
		WHERE
			awu.fk_agg_period = ap.id
			AND
			awu. fk_wallet = $1 
		ORDER BY id DESC
		LIMIT $2 
		OFFSET $3
	;`, walletId, limit, offset)

	if err != nil {
		return awuList, errors.Wrap(err, "db/pg_agg_wallet_usage/GetWalletUsageHist")
	}

	defer rows.Close()

	awu := types.AggWltUsg{}

	for rows.Next() {
		rows.Scan(
			&awu.Id,
			&awu.FkAggPeriod,
			&awu.FkWallet,
			&awu.DlCntDv,
			&awu.DlCntDvFree,
			&awu.UlCntDv,
			&awu.UlCntDvFree,
			&awu.DlCntGw,
			&awu.DlCntGwFree,
			&awu.UlCntGw,
			&awu.UlCntGwFree,
			&awu.Spend,
			&awu.Income,
			&awu.BalanceIncrease,
			&awu.UpdatedBalance,
			&awu.StartAt,
			&awu.DurationMinutes,
		)

		awuList = append(awuList, awu)
	}
	return awuList, errors.Wrap(err, "db/pg_agg_wallet_usage/GetWalletUsageHist")

}

func (*aggWalletUsageInterface) GetWalletUsageHistCnt(walletId int64) (recCnt int64, err error) {
	err = PgDB.QueryRow(`
		SELECT
			COUNT(*)
		FROM
			agg_wallet_usage 
		WHERE
			fk_wallet = $1 
	`, walletId).Scan(&recCnt)

	return recCnt, errors.Wrap(err, "db/pg_agg_wallet_usage/GetWalletUsageHistCnt")

}
