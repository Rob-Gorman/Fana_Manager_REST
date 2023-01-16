import { useHistory } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Stack from "@mui/material/Stack";
import Button from "@mui/material/Button";

export const EntityNotFoundPage = () => {
  const history = useHistory();

  return (
    <Stack>
      <Typography variant="h4">This entity was not found. It may have been deleted.</Typography>
      <Button variant="outlined" onClick={() => history.goBack()}>Go Back</Button>
    </Stack>
  )
}