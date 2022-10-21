import { Link } from "react-router-dom"
import moment from "moment";
import TableRow from "@mui/material/TableRow"
import TableCell from "@mui/material/TableCell";
import MUILink from '@mui/material/Link';

export const LogRow = ({ log }) => {
  const link = `/${log.type}/${log.id}`;

  return (
    <TableRow>
      <TableCell>{log.type}</TableCell>
      <TableCell>
        <Link to={link}>
          <MUILink underline="hover" component="span">{log.key}</MUILink>
        </Link>
      </TableCell>
      <TableCell>{log.action}</TableCell>
      <TableCell>{moment(log.created_at).format('MMMM Do YYYY, h:mm:ss a')}</TableCell>
    </TableRow>
  )
}