import apiClient from '../../lib/apiClient';
import { initializationErrorMessage } from '../../lib/messages';
import { useEffect, useState, useCallback } from 'react';
import { CreateAttributeModal } from "./CreateAttributeModal"
import { AttributeTable } from './AttributeTable';
import { SuccessAlert } from "../SuccessAlert";
import { SearchBox } from '../Shared/SearchBox';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';

export const AttributesList = () => {
  const [ready, setReady] = useState(false);
  const [attributes, setAttributes] = useState([]);
  const [displayedAttributes, setDisplayedAttributes] = useState([]);
  const [formOpen, setFormOpen] = useState(false);
  const [attributeCreated, setAttributeCreated] = useState(false);

  const fetchAttributes = async () => {
    const a = await apiClient.getAttributes();
    setAttributes(a);
    return a;
  }

  useEffect(() => {
    const initialize = async () => {
      try {
        const a = await fetchAttributes();
        setDisplayedAttributes(a);
        setReady(true)
      } catch (e) {
        alert(initializationErrorMessage)
      }
    }
    initialize();
  }, [])


  const searchFilterCriteria = useCallback((searchText) => {
    return (attribute) => {
      return (attribute.attrType.toLowerCase().includes(searchText) ||
              attribute.key.toLowerCase().includes(searchText))
    }
  }, [])

  if (!ready) {
    return <>Loading...</>
  }

  return (
    <Grid container spacing={2}>
      <Grid item xs={12}>
        <Typography variant="h3">Attributes</Typography>
      </Grid>
      {attributeCreated && (<SuccessAlert text="Attribute has been created." successStateSetter={setAttributeCreated} />)}
      <Grid item xs={8}>
        <SearchBox entities={attributes} displayedSetter={setDisplayedAttributes} filterCriteria={searchFilterCriteria} />
      </Grid>
      <Grid item container xs={3} direction="column" alignItems="flex-end" justify="flex-end">
        <Button variant="outlined" onClick={() => setFormOpen(true)}>Create attribute</Button>
      </Grid>
      {formOpen && (<CreateAttributeModal successStateSetter={setAttributeCreated} isOpen={formOpen} setFormOpen={setFormOpen} refreshAtts={fetchAttributes} />)}
      <AttributeTable attributes={displayedAttributes} />
    </Grid>
  );
};