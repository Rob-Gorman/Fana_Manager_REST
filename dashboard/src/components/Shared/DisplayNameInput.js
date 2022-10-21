import TextField from "@mui/material/TextField"

export const DisplayNameInput = ({ currentVal, currentValSetter }) => {
  return (
    <TextField required 
      id="outlined-basic" 
      label="Title" 
      variant="outlined" 
      value={currentVal}
      onChange={(e) => currentValSetter(e.target.value)} 
    />
  )
}