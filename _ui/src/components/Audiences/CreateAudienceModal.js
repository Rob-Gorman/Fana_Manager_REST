import apiClient from '../../lib/apiClient';
import { duplicateErrorMessage, generalErrorMessage } from '../../lib/messages';
import { bigModalStyle } from '../../utils/modalStyle';
import { useState, useEffect } from 'react';
import { SingleCondition } from './SingleCondition';
import { KeyInput } from '../Shared/KeyInput';
import { DisplayNameInput } from '../Shared/DisplayNameInput';
import { ConditionBuilder } from './ConditionBuilder';
import Modal from '@mui/material/Modal';
import Fade from '@mui/material/Fade';
import Backdrop from '@mui/material/Backdrop';
import Box from '@mui/material/Box'
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Grid from '@mui/material/Grid';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';

export const CreateAudienceModal = ({ isOpen, setFormOpen, refreshAudiences, successStateSetter }) => {
  const [displayName, setDisplayName] = useState('');
  const [audienceKey, setAudienceKey] = useState('');
  const [combination, setCombination] = useState('ANY');
  const [conditionFieldActive, setConditionFieldActive] = useState(false);
  const [conditions, setConditions] = useState([]);
  const [readyToSubmit, setReadyToSubmit] = useState(false);

  useEffect(() => {
    if (displayName.trim().length === 0 ||
        audienceKey.trim().length === 0 ||
        conditionFieldActive) {
        setReadyToSubmit(false);
    } else {
      setReadyToSubmit(true);
    }
  }, [displayName, audienceKey, conditionFieldActive, setReadyToSubmit])

  const closeConditionForm = () => {
    setConditionFieldActive(false);
  }

  const removeCondition = (idx) => {
    const copy = conditions.slice();
    copy.splice(idx, 1);
    setConditions(copy);
  }

  const handleSaveCondition = (condition) => {
    setConditions(conditions.concat(condition));
  }

  const handleSubmit = async (e) => {
    const condsWithoutAttKey = conditions.map(c => {
      const { attribute, ...otherFields } = c;
      return otherFields;
    })

    const submission = {
      displayName,
      key: audienceKey,
      combine: combination,
      conditions: condsWithoutAttKey
    }

    try {
      await apiClient.createAudience(submission);
      setFormOpen(false);
      refreshAudiences();
      successStateSetter(true);
    } catch (e) {
      if (e.response.status === 422) {
        alert(duplicateErrorMessage);
      } else {
        alert(generalErrorMessage);
      }
    }
  };

  return (
    <Modal
      style={{ overflowY: 'scroll', height: '100%', display: 'block' }}
      open={isOpen}
      onClose={() => setFormOpen(false)}
      closeAfterTransition
      BackdropComponent={Backdrop}
      BackdropProps={{
        timeout: 500,
      }}
    >
      <Fade in={isOpen}>
        <Box sx={bigModalStyle}>
          <Typography variant="h5">Create a new audience</Typography>
          <Stack container="true" spacing={2} direction="row">
            <Stack style={{ width: "50%"}}>
              <DisplayNameInput currentVal={displayName} currentValSetter={setDisplayName} />
            </Stack>
            <KeyInput currentVal={audienceKey} currentValSetter={setAudienceKey} />
          </Stack>
          <Stack spacing={2}>
            <Typography variant="h6">Conditions</Typography>
            <Stack direction="row">
              <Typography variant="body1">User must meet</Typography>
                <Select
                variant="standard"
                value={combination}
                style={{ marginLeft: 6, marginRight: 6}}
                onChange={(e) => setCombination(e.target.value)}
                >
                  <MenuItem value="ANY">ANY</MenuItem>
                  <MenuItem value="ALL">ALL</MenuItem>
                </Select>
              <Typography variant="body1">of the conditions to qualify for this audience</Typography>
            </Stack>
            <Grid container spacing={2}>
              {conditions.map((c, idx) => {
                return (
                  <SingleCondition key={idx} idx={idx} condition={c} handleRemove={removeCondition} />
                )
              })}
            </Grid>
            {conditionFieldActive ? 
              <ConditionBuilder handleSaveCondition={handleSaveCondition} closable={true} closeConditionForm={closeConditionForm} /> :
              <Button variant="outlined" onClick={() => setConditionFieldActive(true)}>Add Condition</Button>
            }
            <Button disabled={!readyToSubmit} variant="outlined" onClick={handleSubmit}>Create Audience</Button>
          </Stack>
        </Box>
      </Fade>
    </Modal>
  )
}