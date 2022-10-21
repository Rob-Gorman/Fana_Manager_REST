import { Link } from "react-router-dom"
import moment from "moment";
import TableRow from "@mui/material/TableRow"
import TableCell from "@mui/material/TableCell";
import MUILink from '@mui/material/Link';

export const AudienceRow = ({ audience }) => {
  const link = "/audiences/" + audience.id;
  return (
    <TableRow>
      <TableCell>
        <Link to={link}>
          <MUILink underline="hover" component="span">{audience.displayName}</MUILink>
        </Link>
      </TableCell>
      <TableCell>{audience.key}</TableCell>
      <TableCell>{moment(audience.CreatedAt).format("MMM Do YY")}</TableCell>
      <TableCell>{moment(audience.UpdatedAt).format("MMM Do YY")}</TableCell>
    </TableRow>
  )
}