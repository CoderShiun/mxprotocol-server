import React, { Component } from "react";
import { Link } from "react-router-dom";
import Divider from '@material-ui/core/Divider';
import Grid from "@material-ui/core/Grid";
import i18n, { packageNS } from '../../i18n';
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Typography from '@material-ui/core/Typography';
import SessionStore from "../../stores/SessionStore.js";
import DeviceStore from "../../stores/DeviceStore.js";
import WalletStore from "../../stores/WalletStore.js";
import GatewayStore from "../../stores/GatewayStore.js";
import DeviceForm from "./DeviceForm";
import Modal from "../../components/Modal";
//import WithdrawBalanceInfo from "./WithdrawBalanceInfo";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";
import styles from "./DeviceStyle"
import { DV_INACTIVE, DV_FREE_GATEWAYS_LIMITED, DV_WHOLE_NETWORK } from "../../util/Data"

function doIHaveGateway(orgId) {
  return new Promise((resolve, reject) => {
    GatewayStore.getGatewayList(orgId, 0, 1, data => {
      resolve(parseInt(data.count));
    });
  });
}  

function getDlPrice(orgId) {
  return new Promise((resolve, reject) => {
    WalletStore.getDlPrice(orgId, resp => {
      resolve(resp.downLinkPrice);
    });
  });
}

class DeviceLayout extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
      mod: null,
      haveGateway: false,
      downlinkFee: 0
    };
  }

  loadData = async () => {
    try {
      const orgId = this.props.match.params.organizationID;
      this.setState({loading: true})
      let res = await doIHaveGateway(orgId);
      let downlinkFee = await getDlPrice(orgId);
      let haveGateway = (res > 0) ? true : false;

      this.setState({
        downlinkFee,
        haveGateway
      });

      this.setState({loading: false})
    } catch (error) {
      this.setState({loading: false})
      console.error(error);
      this.setState({ error });
    }
  }

  componentDidMount() {
    this.loadData();
  }

  componentDidUpdate(oldProps) {
    if (this.props === oldProps) {
      return;
    }
    this.loadData();
  }
  
  onSubmit = (e, apiWithdrawReqRequest) => {
    e.preventDefault();
  }

  handleCloseModal = () => {
    this.setState({
      modal: null
    })
  }

  onSelectChange = (device) => {
    const { dvId, dvMode } = device;
    //console.log('device', device);
    DeviceStore.setDeviceMode(this.props.match.params.organizationID, dvId, dvMode, data => {
      this.props.history.push(`/device/${this.props.match.params.organizationID}`);
    });
  }

  onSwitchChange = (device, e) => {
    e.preventDefault();
    const { dvId, available } = device;
    //console.log('onSwitchChange', device);
    let mod = DV_FREE_GATEWAYS_LIMITED;
    if(!this.state.haveGateway){
      mod = DV_WHOLE_NETWORK;
    }
    if(!available){
     mod = DV_INACTIVE;   
    }
    //console.log('onSwitchChange', mod);
    DeviceStore.setDeviceMode(this.props.match.params.organizationID, dvId, mod, data => {
    });
  }

  render() {
    return (
      <Grid container spacing={24} className={this.props.classes.backgroundColor}>
        <Grid item xs={12} md={12} lg={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
              <TitleBar className={this.props.classes.padding}>
                <TitleBarTitle title={i18n.t(`${packageNS}:menu.devices.devices`)} />
              </TitleBar>    
              {/* <Divider light={true}/> */}
              <div className={this.props.classes.between}>
              <TitleBar>
                <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                <TitleBarTitle component={Link} to="#" title="/" className={this.props.classes.link}/>
                <TitleBarTitle component={Link} to="#" title={i18n.t(`${packageNS}:menu.devices.devices`)} className={this.props.classes.link}/>
              </TitleBar>
              <div className={this.props.classes.subTitle}>
                {i18n.t(`${packageNS}:menu.devices.downlink_fee_mxc`)} {this.state.downlinkFee} MXC
              </div>
              </div>
          </div>
        </Grid>
        <Grid item xs={12} md={12} lg={6} className={this.props.classes.divider}>
          <DeviceForm
            submitLabel={i18n.t(`${packageNS}:menu.devices.devices`)}
            onSubmit={this.onSubmit}
            downlinkFee={this.state.downlinkFee}
            haveGateway={this.state.haveGateway}
            onSelectChange={this.onSelectChange}
            onSwitchChange={this.onSwitchChange}
          />
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(DeviceLayout));