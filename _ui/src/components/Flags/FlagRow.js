import { Link } from "react-router-dom"
import { FlagStatusToggle } from "./FlagStatusToggle";
import moment from "moment";
import TableRow from "@mui/material/TableRow"
import TableCell from "@mui/material/TableCell";
import MUILink from '@mui/material/Link'

export const FlagRow = ({ flag, refreshFlags, successStateSetter }) => {
  const link = "/flags/" + flag.id;

  return (
    <TableRow>
      <TableCell>
        <Link to={link}>
          <MUILink underline="hover" component="span">{flag.displayName}</MUILink>
        </Link>
      </TableCell>
      <TableCell>{flag.key}</TableCell>
      <TableCell>
        <FlagStatusToggle flag={flag} refreshFlags={refreshFlags} successStateSetter={successStateSetter} />
      </TableCell>
      <TableCell>{moment(flag.CreatedAt).format("MMM Do YY")}</TableCell>
      <TableCell>{moment(flag.UpdatedAt).format("MMM Do YY")}</TableCell>
    </TableRow>
  )
}