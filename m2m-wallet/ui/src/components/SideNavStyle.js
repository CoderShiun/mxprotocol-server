import theme from "../theme";

const SideNavStyle = {
    drawerPaper: {
      position: "fixed",
      width: 270,
      paddingTop: theme.spacing.unit * 9,
      backgroundColor: '#1C1478',
      left: 'inherit',
      color: '#F9FAFC',
      fontSize: 14,
    },
    select: {
      paddingTop: theme.spacing.unit,
      paddingLeft: theme.spacing.unit * 3,
      paddingRight: theme.spacing.unit * 3,
      paddingBottom: theme.spacing.unit * 1,
    },
    card: { // lora server options
      width: '96%',
      height: 250,
      position: 'absolute',
      bottom: 5,
      backgroundColor: '#1C1478',
      color: '#F9FAFC',
      fontSize: 14,
    },
    static: {
      position: 'static'
    },
    iconStyle: {
      color: theme.palette.common.white,
    },
    divider: {
      padding: 0,
      color: '#1C1478',
      width: '100%',
    },
  };
  
export default SideNavStyle;
