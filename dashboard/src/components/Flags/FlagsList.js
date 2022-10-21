import apiClient from '../../lib/apiClient';
import { initializationErrorMessage } from '../../lib/messages';
import { SuccessAlert } from '../SuccessAlert';
import { SearchBox } from '../Shared/SearchBox';
import { useCallback, useEffect, useState } from 'react';
import { CreateFlagModal } from './CreateFlagModal';
import { FlagTable } from './FlagTable';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';

export const FlagsList = () => {
  const [ready, setReady] = useState(false);
  const [flags, setFlags] = useState([]);
  const [displayedFlags, setDisplayedFlags] = useState([]);
  const [formOpen, setFormOpen] = useState(false);
  const [enabledOnly, setEnabledOnly] = useState(false);
  const [flagCreated, setFlagCreated] = useState(false);
  const [flagToggled, setFlagToggled] = useState(false);

  const fetchFlags = async () => {
    const f = await apiClient.getFlags();
    setFlags(f);
    return f;
  }

  useEffect(() => {
    const initialize = async () => {
      try {
        const f = await fetchFlags();
        setDisplayedFlags(f);
        setReady(true)
      } catch (e) {
        alert(initializationErrorMessage)
      }
    }

    initialize();
  }, [])

  const searchFilterCriteria = useCallback((searchText) => {
  return (flag) => {
    if (enabledOnly && !flag.status) {
      return false
    }
    return (flag.displayName.toLowerCase().includes(searchText) ||
            flag.key.toLowerCase().includes(searchText))
    }
  }, [enabledOnly])

  if (!ready) {
    return <>Loading...</>
  }

  return (
    <Grid container spacing={2}>
      <Grid item xs={12}>
        <Typography variant="h3">Flags</Typography>
      </Grid>
      {flagCreated && (<SuccessAlert text="Flag has been created." successStateSetter={setFlagCreated} />)}
      {flagToggled && (<SuccessAlert text="Flag has been toggled." successStateSetter={setFlagToggled} />)}
      <Grid item xs={4}>
        <SearchBox entities={flags} displayedSetter={setDisplayedFlags} filterCriteria={searchFilterCriteria} />
      </Grid>
      <Grid item xs={4}>
        <FormControlLabel control={
          <Checkbox checked={enabledOnly} onChange={() => setEnabledOnly(!enabledOnly)} />
        } label="Show Enabled Only" />
      </Grid>
      <Grid item container xs={3} direction="column" alignItems="flex-end" justify="flex-end">
        <Button variant="outlined" onClick={() => setFormOpen(true)}>Create flag</Button>
      </Grid>
      {formOpen && (<CreateFlagModal successStateSetter={setFlagCreated} isOpen={formOpen} setFormOpen={setFormOpen} refreshFlags={fetchFlags} />)}
      <FlagTable flags={displayedFlags} refreshFlags={fetchFlags} successStateSetter={setFlagToggled} />
    </Grid>
  )
}