import apiClient from '../../lib/apiClient';
import { useEffect, useState, useCallback } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import { FlagAudience } from './FlagAudience';
import { FlagStatusToggle } from './FlagStatusToggle';
import { AddAudienceToFlag } from './AddAudienceToFlag';
import _ from 'lodash';
import {
  deletedEntityMessageCreator,
  generalErrorMessage,
  initializationErrorMessage,
} from '../../lib/messages';
import { SuccessAlert } from '../SuccessAlert';
import { WarningAlert } from '../WarningAlert';
import { EntityNotFoundPage } from '../EntityNotFoundPage';
import { DisplayName } from '../Shared/DisplayName';
import DeleteIcon from '@mui/icons-material/Delete';
import Divider from '@mui/material/Divider';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import InputLabel from '@mui/material/InputLabel';
import Box from '@mui/material/Box';
import List from '@mui/material/List';
import Stack from '@mui/material/Stack';

export const Flag = () => {
  const flagId = useParams().id;
  const history = useHistory();
  const [ready, setReady] = useState(false);
  const [flag, setFlag] = useState();
  const [temporaryAudiences, setTemporaryAudiences] = useState([]);
  const [allAudiences, setAllAudiences] = useState([]);
  const [pendingChanges, setPendingChanges] = useState(false);
  const [titleUpdated, setTitleUpdated] = useState(false);
  const [audiencesUpdated, setAudiencesUpdated] = useState(false);
  const [flagToggled, setFlagToggled] = useState(false);
  const [loadError, setLoadError] = useState(false);

  const closeAllAlerts = () => {
    setFlagToggled(false);
    setAudiencesUpdated(false);
    setTitleUpdated(false);
  };

  const removeAudience = (audienceKey) => {
    const updatedAudiences = temporaryAudiences.filter((a) => {
      if (a.key === audienceKey) {
        return false;
      }

      return true;
    });

    setTemporaryAudiences(updatedAudiences);
  };

  const addAudience = (audienceKey) => {
    const addedAudience = allAudiences.find((a) => a.key === audienceKey);
    const updatedAudiences = temporaryAudiences.concat(addedAudience);
    setTemporaryAudiences(updatedAudiences);
  };

  const submitAudienceEdit = async () => {
    const patchedFlag = {
      audiences: temporaryAudiences.map((a) => a.key),
    };

    try {
      await apiClient.editFlag(flag.id, patchedFlag);
      let f = await fetchFlag();
      setTemporaryAudiences(f.audiences);
      fetchAudiences();
      closeAllAlerts();
      setAudiencesUpdated(true);
    } catch (e) {
      alert(generalErrorMessage);
    }
  };

  const submitDisplayNameEdit = async (newDisplayName) => {
    const patchedFlag = {
      displayName: newDisplayName,
    };

    try {
      await apiClient.editFlag(flag.id, patchedFlag);
      fetchFlag();
      closeAllAlerts();
      setTitleUpdated(true);
    } catch (e) {
      alert(generalErrorMessage);
    }
  };

  const handleDelete = async () => {
    if (flag.status) {
      alert('You must turn the flag off before deleting it')
    } else {
      if (window.confirm('Are you sure you want to delete this flag?')) {
        try {
          await apiClient.deleteFlag(flag.id);
          history.push('/flags');
          alert(deletedEntityMessageCreator('flag', flag.key));
        } catch (e) {
          alert(generalErrorMessage);
        }
      }
    }
  };

  const fetchFlag = useCallback(async () => {
    const f = await apiClient.getFlag(flagId);
    setFlag(f);
    return f;
  }, [flagId]);

  const fetchAudiences = async () => {
    const a = await apiClient.getAudiences();
    setAllAudiences(a);
  };

  useEffect(() => {
    const initialize = async () => {
      try {
        const f = await fetchFlag();
        setTemporaryAudiences(f.audiences);
        fetchAudiences();
        setReady(true);
      } catch (e) {
        if (e.response.status === 404) {
          setLoadError(true);
        } else {
          alert(initializationErrorMessage);
        }
      }
    };
    initialize();
  }, [fetchFlag]);

  useEffect(() => {
    if (ready) {
      const tempKeys = temporaryAudiences.map((a) => a.key).sort();
      const appliedKeys = flag.audiences.map((a) => a.key).sort();
      if (!_.isEqual(tempKeys, appliedKeys)) {
        setPendingChanges(true);
      } else {
        setPendingChanges(false);
      }
    }
  }, [temporaryAudiences, ready, flag?.audiences]);

  if (loadError) {
    return <EntityNotFoundPage />;
  }

  if (!ready) {
    return <>Loading...</>;
  }

  return (
    <Box container="true" spacing={1}>
      {titleUpdated && (
        <SuccessAlert
          text="Title has been updated."
          successStateSetter={setTitleUpdated}
        />
      )}
      {audiencesUpdated && (
        <SuccessAlert
          text="Targeted audiences have been updated."
          successStateSetter={setAudiencesUpdated}
        />
      )}
      {flagToggled && (
        <SuccessAlert
          text="Flag has been toggled."
          successStateSetter={setFlagToggled}
        />
      )}
      {pendingChanges && (
        <WarningAlert text="Changes are not saved until you click on 'Save Audiences'." />
      )}
      <Stack container="true" spacing={2}>
        <Typography variant="h3">Flag Details</Typography>
        <Stack>
          <Typography variant="caption">Title</Typography>
          <Stack direction="row" justifyContent="space-between">
            <DisplayName entity={flag} submitDisplayNameEdit={submitDisplayNameEdit} />
            <Button variant="outlined" onClick={handleDelete} startIcon={<DeleteIcon />} color="error">
              Delete flag
            </Button>
          </Stack>
        </Stack>
        <Stack>
          <Typography variant="caption">Key</Typography>
          <Typography variant="subtitle1">{flag.key}</Typography>
        </Stack>
        <Stack>
          <InputLabel id="flag-toggle-label">Enabled</InputLabel>
          <FlagStatusToggle
            successStateSetter={setFlagToggled}
            flag={flag}
            labelId="flag-toggle-label"
            refreshFlags={fetchFlag}
          />
        </Stack>
        <Divider variant="middle" />
        <Typography variant="h4">Targeted Audiences</Typography>
        <Typography variant="subtitle2">
          This flag will serve to ANY targeted audience
        </Typography>
        <Stack
          container="true"
          divider={<Divider orientation="vertical" flexItem />}
          spacing={10}
          direction="row"
        >
          <Stack>
            <List style={{ width: 350 }}>
              {temporaryAudiences.map((audience) => {
                return (
                  <FlagAudience
                    key={audience.key}
                    audience={audience}
                    removeAudience={removeAudience}
                  />
                );
              })}
            </List>
            <Button
              disabled={!pendingChanges}
              variant="outlined"
              onClick={submitAudienceEdit}
            >
              Save Audiences
            </Button>
          </Stack>
          <AddAudienceToFlag
            allAudiences={allAudiences}
            addAudience={addAudience}
            currentAppliedAudiences={temporaryAudiences}
          />
        </Stack>
      </Stack>
    </Box>
  );
};
