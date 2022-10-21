import { useEffect, useState } from 'react';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';
import Button from '@mui/material/Button';

export const AddAudienceToFlag = ({ addAudience, allAudiences, currentAppliedAudiences }) => {
  const [selectedAudienceKey, setSelectedAudienceKey] = useState('');
  const [selectableAudiences, setSelectableAudiences] = useState([]);

  const handleSubmit = () => {
    addAudience(selectedAudienceKey);
    setSelectedAudienceKey('');
  }

  useEffect(() => {
    const currentAppliedAudienceKeys = currentAppliedAudiences.map(a => a.key)
    const audienceDiffs = allAudiences.filter(a => {
      return !currentAppliedAudienceKeys.includes(a.key);
    })

    setSelectableAudiences(audienceDiffs);
    if (audienceDiffs.length) {
      setSelectedAudienceKey(audienceDiffs[0].key)
    }
  }, [allAudiences, currentAppliedAudiences])

  return (
    <FormControl style={{ minWidth: 300 }}>
      <InputLabel id="add-audience-label">Add Audience</InputLabel>
      <Select
        labelId="add-audience-label"
        value={selectedAudienceKey}
        onChange={(e) => setSelectedAudienceKey(e.target.value)}
        input={<OutlinedInput label="Add Audience" />}
      >
        {selectableAudiences.map(a => {
          return (<MenuItem key={a.key} value={a.key}>{a.displayName}</MenuItem>)
        })}
      </Select>
      <Button disabled={selectableAudiences.length === 0} variant="outlined" onClick={handleSubmit}>Add</Button>
    </FormControl>
  )
}