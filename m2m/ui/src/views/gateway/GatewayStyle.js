import theme from "../../theme";

const GatewayStyles = {
  title: {
    color: '#FFFFFF',
    fontSize: 14,
    padding: 6,
  },
  pos: {
    marginBottom: 12,
    color: '#FFFFFF',
    textAlign: 'right',
  },
  between: {
    display: 'flex',
    justifyContent:'spaceBetween'
  },
  flex: {
    display: 'flex',
    flexDirection: 'column'
  },
  navText: {
    fontSize: 14,
  },
  TitleBar: {
    width: '100%',
    light: true,
    display: 'flex',
    flexDirection: 'column',
    padding: '0px 0px 50px 0px'
  },
  divider: {
    padding: 0,
    color: '#FFFFFF',
    width: '100%',
  },
  padding: {
    padding: 0,
  },
  link: {
    textDecoration: "none",
    fontWeight: "bold",
    fontSize: 12,
    color: theme.palette.textSecondary.main,
    opacity: 0.7,
      "&:hover": {
        opacity: 1,
      }
  },    
};
  
export default GatewayStyles;
