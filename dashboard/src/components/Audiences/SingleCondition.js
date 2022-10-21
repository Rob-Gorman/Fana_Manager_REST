import { operators } from '../../lib/formConstants';
import Grid from '@mui/material/Grid';
import ListItem from '@mui/material/ListItem';
import Paper from '@mui/material/Paper';
import IconButton from '@mui/material/IconButton';
import DeleteIcon from '@mui/icons-material/Delete';

export const SingleCondition = ({condition, handleRemove, idx}) => {
  const operatorString = (condition.negate ? "NOT " : "") + operators[condition.operator];
  const conditionString = (
    <>
      <b>{condition.attribute}</b>&nbsp;
      {operatorString}&nbsp;
      <b>{condition.vals}</b>
    </>
  )
  return (
    <Grid item xs={5}>
      <Paper style={{ overflow: 'scroll' }} elevation={3}>
        <ListItem divider
          secondaryAction={
            <IconButton edge="end" aria-label="delete" onClick={() => handleRemove(idx)} >
              <DeleteIcon/>
            </IconButton>
          }>
          {conditionString}
        </ListItem>
      </Paper>
    </Grid>)
}