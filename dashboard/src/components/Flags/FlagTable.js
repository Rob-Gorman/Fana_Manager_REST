import { FlagRow } from "./FlagRow";
import TableContainer from "@mui/material/TableContainer";
import Table from "@mui/material/Table";
import TableHead from "@mui/material/TableHead";
import TableCell from "@mui/material/TableCell";
import TableBody from "@mui/material/TableBody"
import TableRow from "@mui/material/TableRow";

export const FlagTable = ({ flags, refreshFlags, successStateSetter }) => {
  return (
    <TableContainer>
      <Table stickyHeader>
        <TableHead>
          <TableRow>
            <TableCell style={{width: 300}}>Flag Name</TableCell>
            <TableCell style={{width: 200}}>Flag Key</TableCell>
            <TableCell>Enabled</TableCell>
            <TableCell>Created</TableCell>
            <TableCell>Last Updated</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {flags.map(flag => {
            return (<FlagRow successStateSetter={successStateSetter} key={flag.key} flag={flag} refreshFlags={refreshFlags} />)
          })}
        </TableBody>
      </Table>
    </TableContainer>
  )
}