import { attrTypeMapper } from '../../lib/formConstants';
import { Link } from 'react-router-dom';
import moment from "moment";
import MUILink from '@mui/material/Link';
import TableRow from "@mui/material/TableRow"
import TableCell from "@mui/material/TableCell";

export const AttributeRow = ({ attribute }) => {
  const link = "/attributes/" + attribute.id;
  return (
    <TableRow>
      <TableCell>
      <Link to={link}>
        <MUILink underline="hover" component={'span'}>
        {attribute.key}
      </MUILink >
      </Link>
      </TableCell>
      <TableCell>{attrTypeMapper[attribute.attrType]}</TableCell>
      <TableCell>{moment(attribute.CreatedAt).format("MMM Do YY")}</TableCell>
    </TableRow>
  )
}