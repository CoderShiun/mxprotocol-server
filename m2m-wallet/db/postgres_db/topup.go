package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Topup struct {
	Id              int       `db:"id"`
	FkExtAcntSender int       `db:"fk_ext_account_sender"`
	FkExtAcntRcvr   int       `db:"fk_ext_account_receiver"`
	FkExtCurr       int       `db:"fk_ext_currency"`
	Value           float64   `db:"value"`
	TxAprvdTime     time.Time `db:"tx_approved_time"`
	TxHash          string    `db:"tx_hash"`
}

func (pgDbp DbSpec) CreateTopupTable() error {
	_, err := pgDbp.Db.Exec(`
	CREATE TABLE IF NOT EXISTS topup (
		id SERIAL PRIMARY KEY,
		fk_ext_account_sender INT REFERENCES  ext_account(id) NOT NULL,
		fk_ext_account_receiver INT REFERENCES  ext_account(id) NOT NULL,
		fk_ext_currency INT REFERENCES ext_currency(id) NOT NULL,
		value NUMERIC(28,18) NOT NULL,
		tx_approved_time TIMESTAMP,
		tx_hash varchar (128) NOT NULL UNIQUE
		);
	`)
	return errors.Wrap(err, "db: PostgreSQL connection error")
}

func (pgDbp DbSpec) InsertTopup(tu Topup) (insertIndex int, err error) {
	err = pgDbp.Db.QueryRow(`
		INSERT INTO topup (
			fk_ext_account_sender,
			fk_ext_account_receiver,
			fk_ext_currency,
			value ,
			tx_approved_time,
			tx_hash )
		VALUES (
			$1,	$2,	$3,	$4, $5, $6
		)
		RETURNING id;
		
	`,
		tu.FkExtAcntSender,
		tu.FkExtAcntRcvr,
		tu.FkExtCurr,
		tu.Value,
		tu.TxAprvdTime,
		tu.TxHash).Scan(&insertIndex)

	return insertIndex, errors.Wrap(err, "db: query error InsertWallet()")
}