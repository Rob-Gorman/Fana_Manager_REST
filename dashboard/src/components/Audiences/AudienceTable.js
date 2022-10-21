import { AudienceRow } from "./AudienceRow"
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Table from "@mui/material/Table";
import TableContainer from "@mui/material/TableContainer";

export const AudienceTable = ({ audiences }) => {
  return (
    <TableContainer>
      <Table stickyHeader>
        <TableHead>
          <TableRow>
            <TableCell style={{width: 300}}>Audience Name</TableCell>
            <TableCell style={{width: 200}}>Audience Key</TableCell>
            <TableCell>Created</TableCell>
            <TableCell>Last Updated</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {audiences.map(audience => {
            return (<AudienceRow key={audience.key} audience={audience} />)
          })}
        </TableBody>
      </Table>
    </TableContainer>
  )
}