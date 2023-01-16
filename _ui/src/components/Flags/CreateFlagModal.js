import apiClient from '../../lib/apiClient';
import { duplicateErrorMessage, generalErrorMessage } from '../../lib/messages';
import { smallModalStyle } from '../../utils/modalStyle';
import { useState, useEffect } from 'react';
import { KeyInput } from '../Shared/KeyInput';
import { DisplayNameInput } from '../Shared/DisplayNameInput';
import Modal from '@mui/material/Modal';
import Fade from '@mui/material/Fade';
import Backdrop from '@mui/material/Backdrop';
import Box from '@mui/material/Box'
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import FormControl from '@mui/material/FormControl';
import FormHelperText from '@mui/material/FormHelperText';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Chip from '@mui/material/Chip';

export const CreateFlagModal = ({isOpen, setFormOpen, refreshFlags, successStateSetter}) => {
  const [audiences, setAudiences] = useState([]);
  const [displayName, setDisplayName] = useState('');
  const [flagKey, setFlagKey] = useState('');
  const [selectedAudiences, setSelectedAudiences] = useState([]);

  useEffect(() => {
    const fetchAudiences = async () => {
      const a = await apiClient.getAudiences();
      setAudiences(a);
    }
    fetchAudiences();
  }, [])

  const handleSubmit = async (e) => {
    const submission = {
      key: flagKey,
      displayName,
      audiences: selectedAudiences
    }
    try {
      await apiClient.createFlag(submission);
      setFormOpen(false);
      refreshFlags();
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
      style={{ overflow: 'scroll' }}
      open={isOpen}
      onClose={() => setFormOpen(false)}
      closeAfterTransition
      BackdropComponent={Backdrop}
      BackdropProps={{
        timeout: 500,
      }}
    >
    <Fade in={isOpen}>
      <Box sx={smallModalStyle}>
        <Stack container="true" spacing={2}>
          <Typography variant="h5">Create a new flag</Typography>
          <DisplayNameInput currentVal={displayName} currentValSetter={setDisplayName} />
          <KeyInput currentVal={flagKey} currentValSetter={setFlagKey} />
          <Typography variant="h6">Targeted Audiences</Typography>
          <FormControl sx={{ minWidth: 300 }}>
            <InputLabel id="audience-dropdown-label">Select Audiences</InputLabel>
            <Select multiple
              labelId="audience-dropdown-label"
              value={selectedAudiences}
              onChange={(e) => setSelectedAudiences(e.target.value)}
              input={<OutlinedInput label="Select Audiences" />}
              renderValue={(selected) => (
                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                  {selected.map((value) => (
                    <Chip key={value} label={value} />
                  ))}
                </Box>
              )}>
              {audiences.map((a) => (
                <MenuItem key={a.key} value={a.key}>{a.displayName}</MenuItem>
              ))}
            </Select>
            <FormHelperText>This flag will serve to ANY selected audience</FormHelperText>
          </FormControl>
          <Button disabled={displayName.length === 0 || flagKey.length === 0} variant="outlined" onClick={handleSubmit}>Create</Button>
        </Stack>
      </Box>
    </Fade>
  </Modal>
  )
}