import { useState } from "react";
import Stack from "@mui/material/Stack";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

export const DisplayName = ({ entity, submitDisplayNameEdit }) => {
  const [temporaryDisplayName, setTemporaryDisplayName] = useState(entity.displayName)
  const [editingDisplayName, setEditingDisplayName] = useState(false);

  const handleSubmit = async () => {
    await submitDisplayNameEdit(temporaryDisplayName);
    setEditingDisplayName(false);
  }
  
  return (
    editingDisplayName ? (
      <Stack direction="row" spacing={2}>
        <TextField
          id="outlined-basic"
          label="Edit title"
          variant="outlined"
          value={temporaryDisplayName}
          onChange={(e) => setTemporaryDisplayName(e.target.value)}
        />
        <Button variant="outlined" disabled={temporaryDisplayName.trim().length === 0} onClick={handleSubmit}>Save</Button>
        <Button variant="outlined" color="error" onClick={() => setEditingDisplayName(false)}>Cancel</Button>
      </Stack>
    ) : (
      <Stack direction="row" spacing={2}>
        <Typography variant="subtitle1">{entity.displayName}</Typography>
        <Button variant="outlined" onClick={() => setEditingDisplayName(true)}>Edit</Button>
      </Stack>
    )
  )
}