import apiClient from '../lib/apiClient';
import { initializationErrorMessage } from '../lib/messages';
import { useEffect, useState } from 'react';
import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';
import Divider from '@mui/material/Divider';

export const Settings = () => {
  const [serverSdkKey, setServerSdkKey] = useState({});
  const [clientSdkKey, setClientSdkKey] = useState({})
  const [ready, setReady] = useState(false);

  useEffect(() => {
    const fetchSdkKeys = async () => {
      try {
        const keys = await apiClient.getSdkKey();
        for (let key of keys) {
          if (key.type === 'client') {
            setClientSdkKey(key);
          } else if (key.type === 'server') {
            setServerSdkKey(key);
          }
        }
        setReady(true);
      } catch (e) {
        alert(initializationErrorMessage)
      }
    }
    fetchSdkKeys()
  }, [])

  const regenerateKey = async (key) => {
    const accept = window.confirm('This will invalidate your current SDK key. Are you sure you want to regenerate?');
    if (accept) {
      const newKey = await apiClient.regenSdkKey(key.id, key.type);
      if (newKey.type === 'client') {
        setClientSdkKey(newKey);
      } else if (newKey.type === 'server') {
        setServerSdkKey(newKey);
      }
      alert('New SDK Key issued.')
    }
  }

  const copyKeyToClipboard = (keyString) => {
    navigator.clipboard.writeText(keyString);
  }

  if (!ready) {
    return <>Loading...</>
  }

  return (
    <Stack container="true" spacing={2}>
      <Typography variant="h4">Settings</Typography>
      <Stack spacing={1}>
        <Typography variant="h6">React Client SDK Key</Typography>
        <Typography variant="subtitle1">Use this in your React app</Typography>
        <Stack direction="row" spacing={2}>
          <Typography variant="subtitle1">{clientSdkKey.key}</Typography>
          <Button variant="outlined" onClick={() => copyKeyToClipboard(clientSdkKey.key)}>Copy</Button>
        </Stack>
        <Button variant="contained" onClick={() => regenerateKey(clientSdkKey)}>Regenerate Client SDK Key</Button>
        <Divider />
        <Typography variant="h6">Node Server SDK Key</Typography>
        <Typography variant="subtitle1">Use this in your Node app</Typography>
        <Stack direction="row" spacing={2}>
          <Typography variant="subtitle1">{serverSdkKey.key}</Typography>
          <Button variant="outlined" onClick={() => copyKeyToClipboard(serverSdkKey.key)}>Copy</Button>
        </Stack>
        <Button variant="contained" onClick={() => regenerateKey(serverSdkKey)}>Regenerate Server SDK Key</Button>
      </Stack>
    </Stack>
  )
}