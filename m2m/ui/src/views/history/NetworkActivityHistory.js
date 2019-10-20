import React, { Component } from "react";

import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import moment from 'moment';
import mtz from 'moment-timezone';

import WalletStore from "../../stores/WalletStore";
import TitleBar from "../../components/TitleBar";

import TableCellExtLink from '../../components/TableCellExtLink';
import TitleBarButton from "../../components/TitleBarButton";
import DataTable from "../../components/DataTable";
import LinkVariant from "mdi-material-ui/LinkVariant";
import Admin from "../../components/Admin";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";

const styles = {
  maxW140: {
    maxWidth: 140,
    //backgroundColor: "#0C0270",
    whiteSpace: 'nowrap', 
    overflow: 'hidden',
    textOverflow: 'ellipsis'
  },
  flex:{
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'

  }
};

class NetworkActivityHistory extends Component {
  constructor(props) {
    super(props);
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }

  getPage(limit, offset, callbackFunc) {
    WalletStore.getWalletUsageHist(this.props.organizationID, offset, limit, data => {
        callbackFunc({
            totalCount: parseInt(data.count),
            result: data.walletUsageHis
          });
      }); 
  }
  
  getRow(obj, index) {
    const url = process.env.REACT_APP_ETHERSCAN_ROPSTEN_HOST + `/tx/${obj.txHash}`;
    
    return(
      <TableRow key={index}>
        <TableCell align={'center'} className={this.props.classes.maxW140} >{obj.StartAt.substring(0,19)}</TableCell>
        <TableCell align={'right'} className={this.props.classes.maxW140}>{obj.DlCntDv}</TableCell>
        <TableCell align={'right'} className={this.props.classes.maxW140}>{obj.DlCntDvFree}</TableCell>
        <TableCell align={'right'}>{parseInt(obj.DlCntGw - obj.DlCntGwFree)}</TableCell>
        <TableCell align={'right'}>{obj.Income}</TableCell>
        <TableCell align={'right'}>{obj.Spend}</TableCell>
        <TableCell align={'right'}>{obj.UpdatedBalance}</TableCell>
      </TableRow>
    );
  }

  render() {
    return(
      <Grid container spacing={24}>
        <TitleBar
          buttons={
            <Admin organizationID={this.props.match.params.organizationID}>
              <TitleBarButton
                label="Filter"
                //icon={<Plus />}
              />
            </Admin>
          }
        >
        </TitleBar>
        <Grid item xs={12}>
          <DataTable
            header={
              <TableRow>
                <TableCell align={'center'}>Time</TableCell>
                <TableCell align={'right'}>Pkts Sent</TableCell>
                <TableCell align={'right'}>Free Pkts</TableCell>
                <TableCell align={'right'}>Received</TableCell>
                <TableCell align={'right'}>Income</TableCell>
                <TableCell align={'right'}>Cost</TableCell>
                <TableCell align={'right'}>Balance</TableCell>
              </TableRow>
            }
            getPage={this.getPage}
            getRow={this.getRow}
          />
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(NetworkActivityHistory));