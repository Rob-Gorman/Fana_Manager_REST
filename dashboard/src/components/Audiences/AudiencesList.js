import apiClient from '../../lib/apiClient';
import { initializationErrorMessage } from '../../lib/messages';
import { useEffect, useState, useCallback} from 'react';
import { AudienceTable } from './AudienceTable';
import { CreateAudienceModal } from './CreateAudienceModal';
import { SuccessAlert } from '../SuccessAlert';
import { SearchBox } from '../Shared/SearchBox';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';

export const AudiencesList = () => {
  const [ready, setReady] = useState(false);
  const [audiences, setAudiences] = useState([]);
  const [displayedAudiences, setDisplayedAudiences] = useState([]);
  const [formOpen, setFormOpen] = useState(false);
  const [audienceCreated, setAudienceCreated] = useState(false);

  const fetchAudiences = async () => {
    const a = await apiClient.getAudiences();
    setAudiences(a);
    return a;
  }

  useEffect(() => {
    const initialize = async () => {
      try {
        const a = await fetchAudiences();
        setDisplayedAudiences(a);
        setReady(true)
      } catch (e) {
        alert(initializationErrorMessage)
      }
    }

    initialize();
  }, [])

  const searchFilterCriteria = useCallback((searchText) => {
    return (audience) => {
      return (audience.displayName.toLowerCase().includes(searchText) ||
              audience.key.toLowerCase().includes(searchText))
    }
  }, []);

  if (!ready) {
    return <>Loading...</>
  }

  return (
    <Grid container spacing={2}>
      <Grid item xs={12}>
        <Typography variant="h3">Audiences</Typography>
      </Grid>
      {audienceCreated && (<SuccessAlert text="Audience has been created." successStateSetter={setAudienceCreated} />)}
      <Grid item xs={8}>
        <SearchBox entities={audiences} displayedSetter={setDisplayedAudiences} filterCriteria={searchFilterCriteria} />
      </Grid>
      <Grid item container xs={3} direction="column" alignItems="flex-end" justify="flex-end">
        <Button variant="outlined" onClick={() => setFormOpen(true)}>Create audience</Button>
      </Grid>
      {formOpen && (<CreateAudienceModal isOpen={formOpen} setFormOpen={setFormOpen} refreshAudiences={fetchAudiences} successStateSetter={setAudienceCreated} />)}
      <AudienceTable audiences={displayedAudiences} />
    </Grid>
  )
}

