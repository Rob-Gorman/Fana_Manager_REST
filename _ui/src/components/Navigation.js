import { Link } from "react-router-dom"
import darkLogo from '../images/logo.png';
import whiteLogo from '../images/logo_white.png';
import ListItemButton from "@mui/material/ListItemButton";
import Drawer from "@mui/material/Drawer";
import List from "@mui/material/List"
import ListItemText from "@mui/material/ListItemText";
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import FlagIcon from "@mui/icons-material/Flag";
import GroupsIcon from "@mui/icons-material/Groups";
import AttributionIcon from "@mui/icons-material/Attribution";
import SettingsIcon from "@mui/icons-material/Settings";
import AccessTimeIcon from "@mui/icons-material/AccessTime";
import Toolbar from "@mui/material/Toolbar";
import MUILink from '@mui/material/Link';
import Switch from "@mui/material/Switch";
import InputLabel from "@mui/material/InputLabel";
import Stack from '@mui/material/Stack';

export const Navigation = ({ darkModeOn, darkModeToggle }) => {
  const drawerWidth = 200;
  const pages = [
    { name: 'Flags', url: '/flags', icon: <FlagIcon />}, 
    { name: 'Audiences', url: '/audiences', icon: <GroupsIcon />}, 
    { name: 'Attributes', url: '/attributes', icon: <AttributionIcon />},
    { name: 'History', url: '/history', icon: <AccessTimeIcon />},
    { name: 'Settings', url: '/settings', icon: <SettingsIcon />}
  ];

  let logo = darkModeOn ? whiteLogo : darkLogo;
  
  return (
    <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
          },
        }}
        variant="permanent"
        anchor="left"
      >
        <img src={logo} alt="logo" style={{ marginTop: '25px', width: '85%', marginLeft: 'auto', marginRight: 'auto' }}/>
        <Toolbar/>
        <List>
          {pages.map(page => {
            return (
              <MUILink component={Link} key={page.name} to={page.url} underline="none">
                <ListItem disablePadding>
                  <ListItemButton>
                    <ListItemIcon>
                      {page.icon}
                    </ListItemIcon>
                    <ListItemText primary={page.name} />
                  </ListItemButton>
                </ListItem>
              </MUILink>
            )
          })}
        <Stack direction="column" alignItems="center">
          <InputLabel>Dark Mode</InputLabel>
          <Switch checked={darkModeOn} onChange={darkModeToggle} />
        </Stack>
      </List>
    </Drawer>
  )
}