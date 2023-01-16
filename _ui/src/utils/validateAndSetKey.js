const validateAndSetKey = (value, keySetter, errorSetter) => {
  const re = /^\w+$/
  if (!re.test(value) && value.length > 0) {
    errorSetter(true)
  } else {
    keySetter(value);
    errorSetter(false);
  }
}

export default validateAndSetKey;