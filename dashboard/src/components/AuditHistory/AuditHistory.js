import apiClient from "../../lib/apiClient";
import { initializationErrorMessage } from "../../lib/messages";
import { useEffect, useState, useCallback } from "react"
import { LogRow } from "./LogRow";
import { SearchBox } from "../Shared/SearchBox";
import { Typography, Box } from "@mui/material";
import Table from "@mui/material/Table";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableRow from "@mui/material/TableRow";
import TableBody from "@mui/material/TableBody";
import TableHead from "@mui/material/TableHead";
import IconButton from "@mui/material/IconButton";
import SortIcon from "@mui/icons-material/Sort";
import Grid from "@mui/material/Grid";
import TableFooter from "@mui/material/TableFooter";
import TablePagination from "@mui/material/TablePagination";

export const AuditHistory = () => {
  const [logs, setLogs] = useState([]);
  const [displayedLogs, setDisplayedLogs] = useState([]);
  const [ready, setReady] = useState(false);
  const [newestFirst, setNewestFirst] = useState(true);
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);

  useEffect(() => {
    if (displayedLogs.length <= rowsPerPage * page) {
      setPage(0);
    }
  }, [displayedLogs, rowsPerPage, page])

  const emptyRows =
    page > 0 ? Math.max(0, (1 + page) * rowsPerPage - displayedLogs.length) : 0;

  const handleChangePage = (e, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (e) => {
    setRowsPerPage(parseInt(e.target.value, 10));
    setPage(0);
  };

  const sortLogs = useCallback((logs) => {
    if (newestFirst) {
      return logs.sort((a, b) => (a.created_at < b.created_at) ? 1 : -1)
    } else {
      return logs.sort((a, b) => (a.created_at > b.created_at) ? 1 : -1)
    }
  }, [newestFirst])

  const fetchLogs = useCallback(async () => {
    const l = await apiClient.getLogs();
    const flagLogs = l.flagLogs.map(f => {
      return {...f, type: 'flags', logID: `flag${f.logID}`};
    });
    const audienceLogs = l.audienceLogs.map(a => {
      return {...a, type: 'audiences', logID: `audience${a.logID}`};
    });
    const attributeLogs = l.attributeLogs.map(a => {
      return {...a, type: 'attributes', logID: `attribute${a.logID}`};
    })
    const compiledLogs = flagLogs.concat(audienceLogs, attributeLogs);
    setLogs(compiledLogs);

    return compiledLogs;
  }, [])

  useEffect(() => {
    const initialize = async () => {
      try {
        const l = await fetchLogs();
        setDisplayedLogs(l);
        setReady(true);
      } catch (e) {
        alert(initializationErrorMessage)
      }
    }
    initialize();
  }, [fetchLogs])

  const searchFilterCriteria = useCallback((searchText) => {
    return (log) => {
      return (log.type.toLowerCase().includes(searchText) ||
              log.key.toLowerCase().includes(searchText) ||
              log.action.toLowerCase().includes(searchText))
    }
  }, [])

  if (!ready) {
    return (<>Loading...</>)
  }

  return (
    <Box container="true" spacing={1}>
      <Typography variant="h3">Audit History</Typography>
      <Grid item xs={4}>
        <SearchBox entities={logs} displayedSetter={setDisplayedLogs} filterCriteria={searchFilterCriteria} />
      </Grid>
      <TableContainer>
      <Table stickyHeader>
        <TableHead>
          <TableRow>
            <TableCell width={200}>Entity Type</TableCell>
            <TableCell width={200}>Entity Key</TableCell>
            <TableCell width={200}>Event</TableCell>
            <TableCell>
              Date
              <IconButton onClick={() => setNewestFirst(!newestFirst)} >
                <SortIcon />
              </IconButton>
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
        {(rowsPerPage > 0
              ? sortLogs(displayedLogs).slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
              : sortLogs(displayedLogs)
            ).map(log => {
              return (<LogRow key={log.logID} log={log} />)
            })}
        {emptyRows > 0 && (
          <TableRow style={{ height: 53 * emptyRows }}>
            <TableCell colSpan={4} />
          </TableRow>
        )}
        </TableBody>
        <TableFooter>
          <TableRow>
            <TablePagination
              count={displayedLogs.length}
              rowsPerPage={rowsPerPage}
              page={page}
              onPageChange={handleChangePage}
              onRowsPerPageChange={handleChangeRowsPerPage}
              rowsPerPageOptions={[10, 20, 30]}
              labelRowsPerPage={<span>Rows:</span>}
              labelDisplayedRows={({ page }) => {
                return `Page: ${page + 1}`;
              }}
              backIconButtonProps={{
                color: "secondary"
              }}
              nextIconButtonProps={{ color: "secondary" }}
              SelectProps={{
                inputProps: {
                  "aria-label": "page number"
                }
              }}
              showFirstButton={true}
              showLastButton={true}
            />
          </TableRow>
        </TableFooter>
      </Table>
    </TableContainer>
    </Box>
  )
}