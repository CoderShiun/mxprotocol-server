import React, { Component } from "react";
import { Link, withRouter } from "react-router-dom";

import { withStyles } from "@material-ui/core/styles";
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

import Card from '@material-ui/core/Card';
import CardContent from "@material-ui/core/CardContent";

import Settings from "mdi-material-ui/Settings";

import AccessPoint from "mdi-material-ui/AccessPoint";
import Repeat from "mdi-material-ui/Repeat";
import CalendarCheckOutline from "mdi-material-ui/CalendarCheckOutline";
import CreditCard from "mdi-material-ui/CreditCard";
import Domain from "mdi-material-ui/Domain";
import RadioTower from "mdi-material-ui/RadioTower";

import ArrowExpandLeft from "mdi-material-ui/ArrowExpandLeft";
import ArrowExpandRight from "mdi-material-ui/ArrowExpandRight";
import theme from "../theme";


const styles = {
  drawerPaper: {
    position: "fixed",
    width: 270,
    paddingTop: theme.spacing.unit * 9,
    backgroundColor: '#09006E',
    left: 'inherit',
    color: '#FFFFFF',
  },
  select: {
    paddingTop: theme.spacing.unit,
    paddingLeft: theme.spacing.unit * 3,
    paddingRight: theme.spacing.unit * 3,
    paddingBottom: theme.spacing.unit * 1,
  },
  card: {
    width: '96%',
    height: 250,
    position: 'absolute',
    bottom: 5,
    backgroundColor: '#09006E',
    color: '#FFFFFF',
  },
  static: {
    position: 'static'
  },
  iconStyle: {
    color: theme.palette.common.white,
  }
};

const LinkToLora = ({children, ...otherProps}) => 
<a href={`http://localhost:3002`} {...otherProps}>{children}</a>;

class SideNav extends Component {
  constructor() {
    super();

    this.state = {
      open: true,
      organization: {id:"1", name:"loraserver"},
      cacheCounter: 0,
    };


    /* this.onChange = this.onChange.bind(this);
    this.getOrganizationOption = this.getOrganizationOption.bind(this);
    this.getOrganizationOptions = this.getOrganizationOptions.bind(this);
    this.getOrganizationFromLocation = this.getOrganizationFromLocation.bind(this); */
  }

  componentDidMount() {
    /* SessionStore.on("organization.change", () => {
      OrganizationStore.get(SessionStore.getOrganizationID(), resp => {
        this.setState({
          organization: resp.organization,
        });
      });
    });

    OrganizationStore.on("create", () => {
      this.setState({
        cacheCounter: this.state.cacheCounter + 1,
      });
    });

    OrganizationStore.on("change", (org) => {
      if (this.state.organization !== null && this.state.organization.id === org.id) {
        this.setState({
          organization: org,
        });
      }
    });

    OrganizationStore.on("delete", id => {
      if (this.state.organization !== null && this.state.organization.id === id) {
        this.setState({
          organization: null,
        });
      }

      this.setState({
        cacheCounter: this.state.cacheCounter + 1,
      });
    });

    if (SessionStore.getOrganizationID() !== null) {
      OrganizationStore.get(SessionStore.getOrganizationID(), resp => {
        this.setState({
          organization: resp.organization,
        });
      });
    }

    this.getOrganizationFromLocation(); */
  }

  componentDidUpdate(prevProps) {
    if (this.props === prevProps) {
      return;
    }

    //this.getOrganizationFromLocation();
  }

  onChange(e) {
    //this.props.history.push(`/organizations/${e.target.value}/applications`);
  }

  getOrganizationFromLocation() {
    /* const organizationRe = /\/organizations\/(\d+)/g;
    const match = organizationRe.exec(this.props.history.location.pathname);

    if (match !== null && (this.state.organization === null || this.state.organization.id !== match[1])) {
      SessionStore.setOrganizationID(match[1]);
    } */
  }

  /* getOrganizationOption(id, callbackFunc) {
    OrganizationStore.get(id, resp => {
      callbackFunc({label: resp.organization.name, value: resp.organization.id});
    });
  }

  getOrganizationOptions(search, callbackFunc) {
    OrganizationStore.list(search, 10, 0, resp => {
      const options = resp.result.map((o, i) => {return {label: o.name, value: o.id}});
      callbackFunc(options);
    });
  } */

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
          <ListItem button component={Link} to={`/withdraw`}>
            <ListItemIcon className={this.props.classes.iconStyle}>
              <ArrowExpandLeft />
            </ListItemIcon>
            <ListItemText primary="Withdraw" />
          </ListItem>
          <ListItem button component={Link} to={`/topup`}>
            <ListItemIcon>
              <ArrowExpandRight />
            </ListItemIcon>
            <ListItemText primary="Topup" />
          </ListItem>
          <ListItem button component={Link} to={`/history`}>
            <ListItemIcon>
              <CalendarCheckOutline />
            </ListItemIcon>
            <ListItemText primary="History" />
          </ListItem>
          <ListItem button component={Link} to={`/modify-account`}>
            <ListItemIcon>
              <CreditCard />
            </ListItemIcon>
            <ListItemText primary="ModifyEthAccount" />
          </ListItem>
          <Card className={this.props.classes.card}>
            <CardContent>
              <List className={this.props.classes.static}>
                {/* <ListItem button  onClick={this.handleOpenLora}> */}
                <ListItem button component={LinkToLora}>  
                  <ListItemIcon>
                    <AccessPoint />
                  </ListItemIcon>
                  <ListItemText primary="LoRa Server" />
                </ListItem>
                <ListItem button >
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
                </ListItem>
                <ListItem button  onClick={this.handleOpenLora}>
                  <ListItemText primary="Account name" />
                  <ListItemIcon>
                    <Settings />
                  </ListItemIcon>
                </ListItem>
                <ListItem button onClick={this.handleOpenM2M} >
                  <ListItemText primary="Change Account" />
                  <ListItemIcon>
                    <Repeat />
                  </ListItemIcon>
                </ListItem>
              </List>
            </CardContent>
          </Card>
        </List>}
      </Drawer>
    );
  }
}

export default withRouter(withStyles(styles)(SideNav));