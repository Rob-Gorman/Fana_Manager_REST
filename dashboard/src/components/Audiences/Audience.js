import apiClient from "../../lib/apiClient";
import { deletedEntityMessageCreator, generalErrorMessage, initializationErrorMessage } from "../../lib/messages";
import { useCallback, useEffect, useState } from "react";
import { useHistory, useParams } from "react-router-dom"
import { SuccessAlert } from "../SuccessAlert";
import { WarningAlert } from "../WarningAlert";
import { EntityNotFoundPage } from "../EntityNotFoundPage";
import { DisplayName } from "../Shared/DisplayName";
import { SingleViewConditions } from "./SingleViewConditions";
import { RelatedEntityList } from "../Shared/RelatedEntityList";
import DeleteIcon from '@mui/icons-material/Delete';
import List from "@mui/material/List";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Stack from "@mui/material/Stack";
import Divider from "@mui/material/Divider";
export const Audience = () => {
  const audienceId = useParams().id;
  const history = useHistory();
  const [ready, setReady] = useState(false);
  const [audience, setAudience] = useState();
  const [pendingChanges, setPendingChanges] = useState(false);
  const [titleUpdated, setTitleUpdated] = useState(false);
  const [conditionsUpdated, setConditionsUpdated] = useState(false);
  const [loadError, setLoadError] = useState(false);

  const closeAllAlerts = () => {
    setConditionsUpdated(false);
    setTitleUpdated(false);
  }

  const handleDelete = async () => {
    const flagCount = audience.flags.length;
    if (flagCount > 0) {
      alert(`This audience is being used in ${flagCount} flag${flagCount > 1 ? 's' : ''}. Please remove before deleting.`)
    } else {
      if (window.confirm('Are you sure you want to delete this audience?')) {
        try {
          await apiClient.deleteAudience(audience.id);
          history.push("/audiences")
          alert(deletedEntityMessageCreator('audience', audience.key))
        } catch (e) {
          alert(generalErrorMessage);
        }
      }
    }
  }

  const submitConditionEdit = async (patchedAudience) => {
    try {
      await apiClient.editAudience(audience.id, patchedAudience);
      fetchAudience();
      closeAllAlerts();
      setConditionsUpdated(true);
    } catch(e) {
      alert(generalErrorMessage);
    }
  }

  const submitDisplayNameEdit = async (newDisplayName) => {
    const patchedAudience = {
      displayName: newDisplayName,
    }

    try {
      await apiClient.editAudience(audience.id, patchedAudience);
      fetchAudience();
      closeAllAlerts();
      setTitleUpdated(true);
    } catch(e) {
      alert(generalErrorMessage)
    }
  }

  const fetchAudience = useCallback(async () => {
    const a = await apiClient.getAudience(audienceId);
    setAudience(a);
    return a;
  }, [audienceId])

  useEffect(() => {
    const initialize = async () => {
      try {
        const a = await fetchAudience()
        setReady(true);
      } catch (e) {
        if (e.response.status === 404) {
          setLoadError(true);
        } else {
          alert(initializationErrorMessage)
        }
      }
    }

    initialize();
  }, [fetchAudience])

  if (loadError) {
    return <EntityNotFoundPage />
  }

  if (!ready) {
    return <>Loading...</>
  }

  return (
    <Box container="true" spacing={1}>
      {titleUpdated && (<SuccessAlert text="Title has been updated." successStateSetter={setTitleUpdated} />)}
      {conditionsUpdated && (<SuccessAlert text="Conditions have been updated." successStateSetter={setConditionsUpdated} />)}
      {pendingChanges && (<WarningAlert text="Changes are not saved until you click on 'Save Conditions'." />)}
      <Stack container="true" spacing={2}>
        <Typography variant="h3">Audience Details</Typography>
        <Stack>
          <Typography variant="caption">Title</Typography>
          <Stack direction="row" justifyContent="space-between">
            <DisplayName entity={audience} submitDisplayNameEdit={submitDisplayNameEdit} />
            <Button variant="outlined" onClick={handleDelete} startIcon={<DeleteIcon />} color="error">
              Delete audience
            </Button>
          </Stack>
        </Stack>
        <Stack>
          <Typography variant="caption">Key</Typography>
          <Typography variant="subtitle1">{audience.key}</Typography>
        </Stack>
        <Divider variant="middle" />
        <SingleViewConditions conditions={audience.conditions} combination={audience.combine} pendingChanges={pendingChanges} setPendingChanges={setPendingChanges} submitConditionEdit={submitConditionEdit} />
        <Divider variant="middle" />
        <Stack>
          <Typography variant="h4">Related Flags</Typography>
          <Typography variant="subtitle2">
          List of flags that reference this audience{' '}
        </Typography>
        <List style={{ width: 350 }}>
          {audience.flags.map(flag => 
            (<RelatedEntityList key={flag.id} entity={flag} entityName={'flags'} />)
            )}
        </List>
        </Stack>
      </Stack>
    </Box>
  )
}