package types

import "time"

type DlCategory string

const (
	JOIN_ANS    DlCategory = "JOIN_ANS"
	MAC_COMMAND DlCategory = "MAC_COMMAND"
	PAYLOAD     DlCategory = "PAYLOAD"
	UNKNOWN     DlCategory = "UNKNOWN"
)

type DlPkt struct {
	Id        int64      `db:"id"`
	FkDevice  int64      `db:"dev_eui"` // fk in App server
	FkGateway int64      `db:"fk_gateway"`
	Nonce     int64      `db:"nonce"`
	SentAt    time.Time  `db:"sent_at"`
	Size      float64    `db:"size"`
	Category  DlCategory `db:"category"`
}