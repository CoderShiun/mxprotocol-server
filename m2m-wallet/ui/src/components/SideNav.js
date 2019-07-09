import React, { Component } from "react";
import { Link, withRouter } from "react-router-dom";

import { withStyles } from "@material-ui/core/styles";
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

import Divider from '@material-ui/core/Divider';
import AutocompleteSelect from "./AutocompleteSelect";

import Settings from "mdi-material-ui/Settings";

import AccessPoint from "mdi-material-ui/AccessPoint";
import Repeat from "mdi-material-ui/Repeat";
import CalendarCheckOutline from "mdi-material-ui/CalendarCheckOutline";
import CreditCard from "mdi-material-ui/CreditCard";
import Domain from "mdi-material-ui/Domain";
import RadioTower from "mdi-material-ui/RadioTower";

import WithdrawStore from "../stores/WithdrawStore"
import SessionStore from "../stores/SessionStore"
import PageNextOutline from "mdi-material-ui/PageNextOutline";
import PagePreviousOutline from "mdi-material-ui/PagePreviousOutline";
import styles from "./SideNavStyle";



const LinkToLora = ({children, ...otherProps}) => 
<a href={`appserver:8080`} {...otherProps}>{children}</a>;

class SideNav extends Component {
  constructor() {
    super();

    this.state = {
      open: true,
      organization: {},
      cacheCounter: 0,
    };

    this.onChange = this.onChange.bind(this);
    this.getOrganizationOption = this.getOrganizationOption.bind(this);
    this.getOrganizationOptions = this.getOrganizationOptions.bind(this);
  }

  componentDidMount() {
    //const {match: { params: { organizationID }}} = this.props;
    this.setState({
      organization: {
        id: SessionStore.getOrganizationID()
      }
    })
  }

  componentDidUpdate(prevProps) {
    if (this.props === prevProps) {
      return;
    }

    /* const {match: { params: { organizationID }} } = this.props;
    this.setState({
      organization: {
        id: organizationID
      }
    }); */
  }

  onChange(e) {
    SessionStore.setOrganizationID(e.target.value);
    this.setState({
      organization: {
        id: SessionStore.getOrganizationID()
      }
    })
    //console.log('this.props', this.props.location.pathname);
    
    this.props.history.push(`/withdraw/${e.target.value}`);
  }

  getOrganizationFromLocation() {
    /* const organizationRe = /\/organizations\/(\d+)/g;
    const match = organizationRe.exec(this.props.history.location.pathname);

    if (match !== null && (this.state.organization === null || this.state.organization.id !== match[1])) {
      SessionStore.setOrganizationID(match[1]);
    } */
  }

  getOrganizationOption(id, callbackFunc) {
    WithdrawStore.getWithdrawFee("Ether", resp => {
      const option = resp.userProfile.organizations[0];
      console.log('option', option);
      callbackFunc({label: option.organizationName, value: option.organizationID});
    }); 
  }

  getOrganizationOptions(search, callbackFunc) {
    WithdrawStore.getWithdrawFee("Ether", resp => {
      //dummy data
      resp.userProfile.organizations[0].organizationName = 'lora';
      resp.userProfile.organizations[0].organizationID = '1';
      resp.userProfile.organizations.push({organizationName: 'mxp',organizationID: '2' });
      const options = resp.userProfile.organizations.map((o, i) => { 
        return {label: o.organizationName, value: o.organizationID}});
      callbackFunc(options);
    });
  }

  render() {
    let organizationID = "";
    
    if (this.state.organization !== null) {
      organizationID = this.state.organization.id;
    }
    
    return(
      <Drawer
        variant="persistent"
        anchor="left"
        open={this.props.open}
        classes={{paper: this.props.classes.drawerPaper}}
      >
        {this.state.organization && <List className={this.props.classes.static}>
          {/* <ListItem button component={Link} to={`/withdraw/${this.state.organization.id}`}> */}
          <Divider />
          <div>
          <AutocompleteSelect
            id="organizationID"
            margin="none"
            value={organizationID}
            onChange={this.onChange}
            //getOption={this.getOrganizationOption}
            getOptions={this.getOrganizationOptions}
            className={this.props.classes.select}
            triggerReload={this.state.cacheCounter}
          />
        </div>
          <ListItem button component={Link} to={`/withdraw/${this.state.organization.id}`}>
            <ListItemIcon className={this.props.classes.iconStyle}>
              <PagePreviousOutline />
            </ListItemIcon>
            <ListItemText primary="Withdraw" />
          </ListItem>
          <ListItem button component={Link} to={`/topup/${this.state.organization.id}`}>
            <ListItemIcon>
              <PageNextOutline />
            </ListItemIcon>
            <ListItemText primary="Topup" />
          </ListItem>
          <ListItem button component={Link} to={`/history/${this.state.organization.id}`}>
            <ListItemIcon>
              <CalendarCheckOutline />
            </ListItemIcon>
            <ListItemText primary="History" />
          </ListItem>
          <ListItem button component={Link} to={`/modify-account/${this.state.organization.id}`}>
            <ListItemIcon>
              <CreditCard />
            </ListItemIcon>
            <ListItemText primary="ETH Account" />
          </ListItem>
          
              <List className={this.props.classes.card}>
                {/* <ListItem button  onClick={this.handleOpenLora}> */}
                <ListItem button component={LinkToLora} className={this.props.classes.static}>  
                  <ListItemIcon>
                    <img src="/logo/logo.png" className="iconStyle" alt="LoRa Server" />
                  </ListItemIcon>
                </ListItem>
                <Divider />
                {/* <ListItem button >
                  <ListItemText primary="Super Node" />
                  <ListItemIcon>
                    <RadioTower />
                  </ListItemIcon>
                </ListItem>
                <ListItem button >
                  <ListItemText primary="Organization" />
                  <ListItemIcon>
                    <Domain />
                  </ListItemIcon>
                </ListItem> */}
                <Divider />
                <ListItem>
                  <ListItemText primary="Powered by" />
                  <ListItemIcon>
                    <img src="/logo/mxc_logo.png" className="iconStyle" alt="LoRa Server" />
                  </ListItemIcon>
                </ListItem>
                {/* <ListItem button onClick={this.handleOpenM2M} >
                  <ListItemText primary="Change Account" />
                  <ListItemIcon>
                    <Repeat />
                  </ListItemIcon>
                </ListItem> */}
              </List>
        </List>}
      </Drawer>
    );
  }
}

export default withRouter(withStyles(styles)(SideNav));
