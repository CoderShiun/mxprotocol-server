package appserver

import (
	"context"
	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/db"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/device"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/services/wallet"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/types"
	"time"
)

type M2MServerAPI struct{}

// M2MServerAPI returns a new M2MServerAPI.
func NewM2MServerAPI() *M2MServerAPI {
	return &M2MServerAPI{}
}

func (*M2MServerAPI) AddDeviceInM2MServer(ctx context.Context, req *api.AddDeviceInM2MServerRequest) (*api.AddDeviceInM2MServerResponse, error) {
	log.WithFields(log.Fields{
		"orgId": req.OrgId,
	}).Debug("grpc_api/AddDeviceInM2MServer")

	walletId, err := wallet.GetWalletId(req.OrgId)
	if err != nil {
		return &api.AddDeviceInM2MServerResponse{}, err
	}

	dev := types.Device{}
	dev.DevEui = req.DevProfile.DevEui
	dev.CreatedAt = time.Now()
	dev.ApplicationId = req.DevProfile.ApplicationId
	dev.Name = req.DevProfile.Name
	dev.Mode = types.DV_WHOLE_NETWORK
	dev.FkWallet = walletId

	devId, err := db.Device.InsertDevice(dev)
	if err != nil {
		log.WithError(err).Error("grpc_api/AddDeviceInM2MServer")
		// retry
		go func() {
			for {
				err := device.SyncDeviceProfileFromAppserver(dev, types.ADD)
				if err != nil {
					log.WithError(err).Error("grpc_api/AddDeviceInM2MServer: retry failed")
					time.Sleep(5*time.Second)
					continue
				}
				break;
			}
		}()
		return &api.AddDeviceInM2MServerResponse{}, err
	}

	return &api.AddDeviceInM2MServerResponse{DevId: devId}, nil
}

func (*M2MServerAPI) DeleteDeviceInM2MServer(ctx context.Context, req *api.DeleteDeviceInM2MServerRequest) (*api.DeleteDeviceInM2MServerResponse, error) {
	log.WithFields(log.Fields{
		"dvEui": req.DevEui,
	}).Debug("grpc_api/DeleteDeviceInM2MServer")

	devId, err := db.Device.GetDeviceIdByDevEui(req.DevEui)
	if err != nil {
		log.WithError(err).Error("Get devId from DB error")
		return &api.DeleteDeviceInM2MServerResponse{}, err
	}
	err = db.Device.SetDeviceMode(devId, "DELETED")
	if err != nil {
		log.WithError(err).Error("Set devMode error")
		return &api.DeleteDeviceInM2MServerResponse{}, err
	}

	return &api.DeleteDeviceInM2MServerResponse{Status: true}, nil
}

func (*M2MServerAPI) AddGatewayInM2MServer(ctx context.Context, req *api.AddGatewayInM2MServerRequest) (*api.AddGatewayInM2MServerResponse, error) {
	log.WithFields(log.Fields{
		"orgId": req.OrgId,
	}).Debug("grpc_api/AddGatewayInM2MServer")

	walletId, err := wallet.GetWalletId(req.OrgId)
	if err != nil {
		return &api.AddGatewayInM2MServerResponse{}, err
	}

	gw := types.Gateway{}
	gw.Mac = req.GwProfile.Mac
	gw.CreatedAt = time.Now()
	gw.OrgId = req.OrgId
	gw.Description = req.GwProfile.Description
	gw.Name = req.GwProfile.Name
	gw.Mode = types.GW_WHOLE_NETWORK
	gw.FkWallet = walletId

	gwId, err := db.Gateway.InsertGateway(gw)
	if err != nil {
		log.WithError(err).Error("Insert gateway to DB error")
		return &api.AddGatewayInM2MServerResponse{}, err
	}

	return &api.AddGatewayInM2MServerResponse{GwId: gwId}, nil
}

func (*M2MServerAPI) DeleteGatewayInM2MServer(ctx context.Context, req *api.DeleteGatewayInM2MServerRequest) (*api.DeleteGatewayInM2MServerResponse, error) {
	log.WithFields(log.Fields{
		"gwMac": req.MacAddress,
	}).Debug("grpc_api/DeleteGatewayInM2MServer")

	gwId, err := db.Gateway.GetGatewayIdByMac(req.MacAddress)
	if err != nil {
		log.WithError(err).Error("Get gwId from DB error")
		return &api.DeleteGatewayInM2MServerResponse{}, err
	}
	err = db.Gateway.SetGatewayMode(gwId, "DELETED")
	if err != nil {
		log.WithError(err).Error("Set devMode error")
		return &api.DeleteGatewayInM2MServerResponse{}, err
	}

	return &api.DeleteGatewayInM2MServerResponse{Status: true}, nil
}
