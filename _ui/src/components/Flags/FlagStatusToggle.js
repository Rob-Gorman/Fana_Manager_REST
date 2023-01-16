import apiClient from "../../lib/apiClient";
import { generalErrorMessage } from "../../lib/messages";
import { Switch } from "@mui/material";

export const FlagStatusToggle = ({ flag, refreshFlags, successStateSetter }) => {
  const toggleFlagStatus = async () => {
    try {
      await apiClient.toggleFlag(flag.id, { status: !flag.status })
      refreshFlags();
      successStateSetter(true);
    } catch(e) {
      alert(generalErrorMessage)
    }
  }

  return (
    <Switch checked={flag.status} onChange={toggleFlagStatus} />
  )
}