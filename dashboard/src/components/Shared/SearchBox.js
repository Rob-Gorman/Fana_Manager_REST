import { useState, useEffect } from "react";
import TextField from '@mui/material/TextField';

export const SearchBox = ({ entities, displayedSetter, filterCriteria }) => {
  const [searchText, setSearchText] = useState('');

  useEffect(() => {
    const lcSearchText = searchText.toLowerCase();
    const filteredEntities = entities.filter(filterCriteria(lcSearchText))
    displayedSetter(filteredEntities);
  }, [searchText, entities, displayedSetter, filterCriteria])

  return (
    <TextField
      id="outlined-basic"
      label="Search"
      variant="outlined"
      value={searchText}
      onChange={(e) => setSearchText(e.target.value)}
    />
  )
}