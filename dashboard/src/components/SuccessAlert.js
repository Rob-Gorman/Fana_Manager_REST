import Alert from "@mui/material/Alert";
import AlertTitle from "@mui/material/AlertTitle";

export const SuccessAlert = ({ text, successStateSetter }) => {
  return (
    <Alert 
      severity="success" 
      onClose={() => successStateSetter(false)} 
      style={{ width: 300, position: 'absolute', marginLeft: 'auto', marginRight: 'auto', left: 0, right: 0, textAlign: 'center'}}>
      <AlertTitle>Nice!</AlertTitle>
      {text}
    </Alert>
  )
}