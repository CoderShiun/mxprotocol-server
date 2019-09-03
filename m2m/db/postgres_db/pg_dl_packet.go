package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type DlPkt struct {
	Id        int64     `db:"id"`
	FkDevice  int64     `db:"dev_eui"` // fk in App server
	FkGateway int64     `db:"fk_gateway"`
	Nonce     int64     `db:"nonce"`
	SentAt    time.Time `db:"sent_at"`
	Size      float64   `db:"size"`
	Category  string    `db:"category"`
}

func (pgDbp *PGHandler) CreateDlPktTable() error {
	_, err := pgDbp.DB.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS 
				(SELECT 1 FROM pg_type WHERE typname = 'dl_category') 
			THEN
				CREATE TYPE dl_category AS ENUM (
					'JOIN_ANS',
					 'MAC_COMMAND',
					 'PAYLOAD',
					 'UNKNOWN'
		);
		END IF;
			CREATE TABLE IF NOT EXISTS dl_pkt (
			id SERIAL PRIMARY KEY,
			fk_device INT REFERENCES device (id) NOT NULL,
			fk_gateway INT REFERENCES gateway (id) NOT NULL,
			nonce INT,
			sent_at     TIMESTAMP,
			size FLOAT ,
			category  dl_category
		);

		END$$;
		
	`)
	return errors.Wrap(err, "db/pg_dl_pkt/CreateDlPktTable")
}

func (pgDbp *PGHandler) InsertDlPkt(dlPkt DlPkt) (insertIndex int64, err error) {
	err = pgDbp.DB.QueryRow(`
		INSERT INTO dl_pkt (
			fk_device,
			fk_gateway,
			nonce ,
			sent_at,
			size ,
			category
			)
		VALUES
			($1,$2,$3,$4,$5,$6)
		RETURNING id ;
	`,
		dlPkt.FkDevice,
		dlPkt.FkGateway,
		dlPkt.Nonce,
		dlPkt.SentAt,
		dlPkt.Size,
		dlPkt.Category,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_dl_pkt/InsertDlPkt")
}
