package postgres_db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/mxc-foundation/mxprotocol-server/m2m/types"
)

type deviceInterface struct{}

var PgDevice deviceInterface

func (*deviceInterface) CreateDeviceTable() error {
	_, err := PgDB.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS 
				(SELECT 1 FROM pg_type WHERE typname = 'device_mode') 
			THEN
				CREATE TYPE device_mode AS ENUM (
					'DV_INACTIVE',
					 'DV_FREE_GATEWAYS_LIMITED',
					 'DV_WHOLE_NETWORK',
					 'DV_DELETED'
		);
		END IF;
			CREATE TABLE IF NOT EXISTS device (
			id SERIAL PRIMARY KEY,
			dev_eui VARCHAR(64) NOT NULL,
			fk_wallet INT REFERENCES wallet (id) NOT NULL,
			mode DEVICE_MODE    NOT NULL,
			created_at     TIMESTAMP,
			last_seen_at    TIMESTAMP,
			application_id INT    ,
			name  varchar(128)    
		);

		END$$;
		
	`)
	return errors.Wrap(err, "db/pg_device/CreateDevice")
}

func (*deviceInterface) InsertDevice(dv types.Device) (insertIndex int64, err error) {
	err = PgDB.QueryRow(`
		INSERT INTO device (
			dev_eui ,
			fk_wallet,
			mode,
			created_at,
			last_seen_at,
			application_id,
			name
			) 
		VALUES 
			($1,$2,$3,$4,$5,$6,$7)
		RETURNING id ;
	`,
		dv.DevEui,
		dv.FkWallet,
		dv.Mode,
		dv.CreatedAt,
		dv.LastSeenAt,
		dv.ApplicationId,
		dv.Name,
	).Scan(&insertIndex)
	return insertIndex, errors.Wrap(err, "db/pg_device/InsertDevice")
}

func (*deviceInterface) GetDeviceMode(dvId int64) (dvMode types.DeviceMode, err error) {
	err = PgDB.QueryRow(
		`SELECT
			mode
		FROM
			device 
		WHERE
			id = $1 
		;`, dvId).Scan(&dvMode)
	return dvMode, errors.Wrap(err, "db/pg_device/GetDeviceMode")
}

func (*deviceInterface) SetDeviceMode(dvId int64, dvMode types.DeviceMode) (err error) {
	_, err = PgDB.Exec(`
		UPDATE
			device 
		SET
			mode = $1
		WHERE
			id = $2
		;
		`, dvMode, dvId)

	return errors.Wrap(err, "db/pg_device/SetDeviceMode")

}

func (*deviceInterface) DeleteDevice(dvId int64) (err error) {
	return PgDevice.SetDeviceMode(dvId, types.DV_DELETED)
}

func (*deviceInterface) GetDeviceIdByDevEui(devEui string) (devId int64, err error) {
	err = PgDB.QueryRow(
		`SELECT
			id
		FROM
			device 
		WHERE
			dev_eui = $1 
		AND
			mode <> 'DV_DELETED'
		ORDER BY id DESC 
		LIMIT 1  
		;`, devEui).Scan(&devId)
	return devId, errors.Wrap(err, "db/pg_device/GetDeviceIdByDevEui")
}

func (*deviceInterface) UpdateDeviceLastSeen(dvId int64, newTime time.Time) (err error) {
	_, err = PgDB.Exec(`
		UPDATE
			device 
		SET
			last_seen_at = $1
		WHERE
			id = $2
		;
		`, newTime, dvId)

	return errors.Wrap(err, "db/pg_device/UpdateDeviceLastSeen")
}

func (*deviceInterface) GetDeviceProfile(dvId int64) (dv types.Device, err error) {

	err = PgDB.QueryRow(
		`SELECT
			*
		FROM
			device 
		WHERE
			id = $1 
		;`, dvId).Scan(
		&dv.Id,
		&dv.DevEui,
		&dv.FkWallet,
		&dv.Mode,
		&dv.CreatedAt,
		&dv.LastSeenAt,
		&dv.ApplicationId,
		&dv.Name)
	return dv, errors.Wrap(err, "db/pg_device/GetDeviceProfile")
}

func (*deviceInterface) GetDeviceListOfWallet(walletId int64, offset int64, limit int64) (dvList []types.Device, err error) {

	rows, err := PgDB.Query(
		`SELECT
			*
		FROM
			device 
		WHERE
			fk_wallet = $1 
		AND
			mode <> 'DV_DELETED'
		ORDER BY id DESC
		LIMIT $2 
		OFFSET $3
	;`, walletId, limit, offset)

	defer rows.Close()
	if err != nil {
		return dvList, errors.Wrap(err, "db/pg_device/GetDeviceListOfWallet")
	}

	var dv types.Device

	for rows.Next() {
		rows.Scan(
			&dv.Id,
			&dv.DevEui,
			&dv.FkWallet,
			&dv.Mode,
			&dv.CreatedAt,
			&dv.LastSeenAt,
			&dv.ApplicationId,
			&dv.Name,
		)

		dvList = append(dvList, dv)
	}
	return dvList, errors.Wrap(err, "db/pg_device/GetDeviceListOfWallet")
}

func (*deviceInterface) GetAllDevices() (devList []types.Device, err error) {
	rows, err := PgDB.Query(
		`SELECT
			*
		FROM
			device 
		WHERE 
			mode <> 'DV_DELETED'
		ORDER BY created_at DESC;`)

	defer rows.Close()
	if err != nil {
		return devList, errors.Wrap(err, "db/pg_device/GetAllDevices")
	}

	var dev types.Device
	for rows.Next() {
		rows.Scan(
			&dev.Id,
			&dev.DevEui,
			&dev.FkWallet,
			&dev.Mode,
			&dev.CreatedAt,
			&dev.LastSeenAt,
			&dev.ApplicationId,
			&dev.Name,
		)

		devList = append(devList, dev)
	}
	return devList, errors.Wrap(err, "db/pg_device/GetAllDevices")
}

func (*deviceInterface) GetDeviceRecCnt(walletId int64) (recCnt int64, err error) {

	err = PgDB.QueryRow(`
		SELECT
			COUNT(*)
		FROM
			device 
		WHERE
			fk_wallet = $1 
		AND
			mode <> 'DV_DELETED'
	`, walletId).Scan(&recCnt)

	return recCnt, errors.Wrap(err, "db/pg_device/GetDeviceRecCnt")
}

// DeleteDevice

func (*deviceInterface) GetWalletIdOfDevice(dvId int64) (dvWalletId int64, err error) {

	err = PgDB.QueryRow(
		`SELECT
			fk_wallet
		FROM
			device 
		WHERE	
			id = $1 
			ORDER BY id DESC 
			LIMIT 1  
		;`, dvId).Scan(&dvWalletId)
	return dvWalletId, errors.Wrap(err, "db/pg_device/GetWalletIdOfDevice")
}

func (*deviceInterface) GetDevWalletIdByEui(devEui string) (walletId int64, err error) {

	dvId, err := PgDevice.GetDeviceIdByDevEui(devEui)
	if err != nil {
		return 0, errors.Wrap(err, "db/pg_device/GetDevWalletIdByEui")
	}

	walletId, err = PgDevice.GetWalletIdOfDevice(dvId)
	if err != nil {
		return 0, errors.Wrap(err, "db/pg_device/GetDevWalletIdByEui")
	}
	return walletId, err
}

func (*deviceInterface) GetDevWalletIdAndModeByEui(devEui string) (dvWalletId int64, dvMode types.DeviceMode, err error) {

	err = PgDB.QueryRow(
		`SELECT
			fk_wallet, mode
		FROM
			device 
		WHERE	
			dev_eui = $1 
		AND
			mode <> 'DV_DELETED'
		ORDER BY id DESC 
		LIMIT 1  
		;`, devEui).Scan(&dvWalletId, &dvMode)
	return dvWalletId, dvMode, errors.Wrap(err, "db/pg_device/GetDevWalletIdAndModeByEui")

}
