import { Link } from 'react-router-dom';
import MUILink from '@mui/material/Link';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';

export const RelatedEntityList = ({ entity, entityName }) => {
  const link = `/${entityName}/` + entity.id;
  return (
    <ListItem divider>
      <ListItemText>
        <Link to={link} >
        <MUILink underline="hover" component={'span'}>
          {entity.displayName}
        </MUILink>
          </Link>
      </ListItemText>
    </ListItem>
  );
};
