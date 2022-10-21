import { useState, useEffect } from "react"
import { SingleCondition } from "./SingleCondition"
import { ConditionBuilder } from "./ConditionBuilder"
import _ from "lodash"
import Typography from "@mui/material/Typography"
import Stack from "@mui/material/Stack"
import Select from "@mui/material/Select"
import Box from "@mui/material/Box"
import MenuItem from "@mui/material/MenuItem"
import Divider from "@mui/material/Divider"
import Button from "@mui/material/Button"
import List from "@mui/material/List"

export const SingleViewConditions = ({ conditions, combination, pendingChanges, setPendingChanges, submitConditionEdit }) => {
  const [temporaryConditions, setTemporaryConditions] = useState(conditions.slice());
  const [temporaryCombination, setTemporaryCombination] = useState(combination);

  useEffect(() => {
    if (!_.isEqual(conditions, temporaryConditions) || temporaryCombination !== combination) {
      setPendingChanges(true);
    } else {
      setPendingChanges(false);
    }
  }, [temporaryConditions, conditions, temporaryCombination, combination, setPendingChanges])

  const removeCondition = (removedConditionIdx) => {
    const updatedConditions = temporaryConditions.filter((c, idx) => {
      if (idx === removedConditionIdx) {
        return false;
      }
      return true;
    })

    setTemporaryConditions(updatedConditions);
  }

  const addCondition = (newCondition) => {
    setTemporaryConditions(temporaryConditions.concat(newCondition));
  }

  const handleSubmit = () => {
    const condsWithoutAttKey = temporaryConditions.map(c => {
      const { attribute, ...otherFields } = c;
      return otherFields;
    })

    const patchedAudience = {
      combine: temporaryCombination,
      conditions: condsWithoutAttKey
    }

    submitConditionEdit(patchedAudience);
  }

  return (
    <Box>
      <Typography variant="h4">Conditions</Typography>
        <Stack direction="row">
          <Typography variant="body1">User must meet</Typography>
          <Select
            variant="standard"
            value={temporaryCombination}
            style={{ marginLeft: 6, marginRight: 6}}
            onChange={(e) => setTemporaryCombination(e.target.value)}
            >
              <MenuItem value="ANY">ANY</MenuItem>
              <MenuItem value="ALL">ALL</MenuItem>
          </Select>
          <Typography variant="body1">of the conditions to qualify for this audience</Typography>
        </Stack>
        <Stack
          container="true"
          divider={<Divider orientation="vertical" flexItem />}
          spacing={10}
          direction="row"
        >
          <Stack>
            <List style={{ width: 350 }}>
              {temporaryConditions.map((condition, idx) => {
                return (<SingleCondition key={idx} idx={idx} condition={condition} handleRemove={removeCondition} />)
              })}
            </List>
            <Button disabled={!pendingChanges} variant="outlined" onClick={handleSubmit}>Save Conditions</Button>
          </Stack>
          <Stack style={{ width: 350 }}>
          <ConditionBuilder closable={false} handleSaveCondition={addCondition} />
          </Stack>
        </Stack>
    </Box>
  )
}