import Alert from "@mui/material/Alert";
import AlertTitle from "@mui/material/AlertTitle";

export const WarningAlert = ({ text }) => {
  return (
    <Alert 
      severity="warning"  
      style={{ width: 300, position: 'absolute', marginLeft: 'auto', marginRight: 'auto', left: 0, right: 0, textAlign: 'center'}}>
      <AlertTitle>Heads up!</AlertTitle>
      {text}
    </Alert>
  )
}