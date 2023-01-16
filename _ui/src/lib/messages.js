export const duplicateErrorMessage = "That key is already in use. Please use a different one.";

export const generalErrorMessage = "Something went wrong. Please try again later.";

export const initializationErrorMessage = "Error initializing. Check the connection to the backend manager and try again.";

export const deletedEntityMessageCreator = (entity, key) => {
  return `The ${entity} ${key} has been successfully deleted.`;
}