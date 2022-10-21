import { useState } from "react";
import Stack from "@mui/material/Stack";
import TextField from "@mui/material/TextField";
import validateAndSetKey from '../../utils/validateAndSetKey';
import FormHelperText from "@mui/material/FormHelperText";

export const KeyInput = ({ currentVal, currentValSetter, }) => {
  const [keyError, setKeyError] = useState(false);

  const onKeyInput = (e) => {
    validateAndSetKey(e.target.value, currentValSetter, setKeyError);
  }

  return (
    <Stack>
      <TextField required 
        error={keyError}
        id="outlined-basic" 
        label="Key" 
        variant="outlined"
        value={currentVal}
        onChange={onKeyInput}
        onBlur={() => setKeyError(false)}
      />
    <FormHelperText>Alphanumeric and underscores only. This cannot be changed after creation</FormHelperText>
  </Stack>)
}