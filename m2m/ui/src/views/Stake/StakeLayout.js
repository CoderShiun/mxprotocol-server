import React, { Component } from "react";
import { Link } from "react-router-dom";
import Grid from "@material-ui/core/Grid";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import WalletStore from "../../stores/WalletStore.js";
import GatewayStore from "../../stores/GatewayStore.js";
import Button from "@material-ui/core/Button";
import StakeStore from "../../stores/StakeStore";
import ExtLink from "../../components/ExtLink";
import Typography from '@material-ui/core/Typography';
//import WithdrawBalanceInfo from "./WithdrawBalanceInfo";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";
import styles from "./StakeStyle"
import { DISMISS, STAKE_DESCRIPTION, LEARN_MORE } from "../../util/Messages"
import { EXT_URL_STAKE } from "../../util/Data"

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

class StakeLayout extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
      isFirst: true,
      dismissOn: true
    };
  }

  loadData = async () => {
    StakeStore.getStakingHistory(this.props.match.params.organizationID, 0, 1, data => {
      let amount = 0;
      let isFirst = true;
      if( data.stakingHist.length > 0){
        this.props.history.push(`/stake/${this.props.match.params.organizationID}/set-stake`);
      }
      this.setState({
        amount,
        isFirst
      })
    }); 
  }

  componentWillMount(){
    this.loadData();
  }

  componentDidMount() {
  }

  componentDidUpdate(oldProps) {
    if (this.props === oldProps) {
      return;
    }
    this.loadData();
  }
  
  dismissOn = () => {
    this.setState({
      dismissOn: false
    });
  }
  
  onSubmit = (e, apiWithdrawReqRequest) => {
    e.preventDefault();
  }

  render() {
    return (
      <Grid container spacing={24} className={this.props.classes.backgroundColor}>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
              {/* <TitleBar className={this.props.classes.padding}>
                <TitleBarTitle title="Stake" />
              </TitleBar> */}    
              {/* <Divider light={true}/> */}
              <div className={this.props.classes.between}>
              <TitleBar>
                <TitleBarTitle title="Stake" />
                {/* <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                <TitleBarTitle component={Link} to="#" title="/" className={this.props.classes.link}/>
                <TitleBarTitle component={Link} to="#" title="Devices" className={this.props.classes.link}/> */}
              </TitleBar>
              <Button color="primary.main" component={Link} to={`/stake/${this.props.match.params.organizationID}/set-stake`} /* onClick={this.handleOpenAXS} */ type="button" disabled={false}>SET STAKE</Button>
              {/* <TitleBarButton
                label="SET STAKE"
                color="primary"
                to={`/stake/${this.props.match.params.organizationID}/set-stake`}
                classes={this.props.classes}
              /> */}
              </div>
          </div>
        </Grid>
        <Grid item xs={12} className={this.props.classes.divider}>
          <Grid item xs={6}>
          {this.state.dismissOn && <div className={this.props.classes.infoBox}>
                  <p>{STAKE_DESCRIPTION}</p>
                  <div className={this.props.classes.between}>
                    <ExtLink dismissOn={this.dismissOn} for={'local'} context={DISMISS} />&nbsp;&nbsp;&nbsp;
                    <ExtLink to={EXT_URL_STAKE} context={LEARN_MORE} />
                  </div>
                </div>}
          </Grid>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(StakeLayout));