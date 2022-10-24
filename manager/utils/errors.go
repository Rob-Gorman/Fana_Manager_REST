package utils

import (
	"errors"
	"fmt"
	"net/http"
)

func EnvVarError(key string) error {
	return fmt.Errorf("no %s variable found in environment", key)
}

func UnmarshalError(err error) error {
	return composeError("error unmarshalling request body", err)
}

func MarshalError(err error) error {
	return composeError("error marshalling request body", err)
}

func DuplicateError(err error) error {
	return composeError("cannot create resource with duplicate keys", err)
}

func DBConnError(err error) error {
	return composeError("cannot connect to DB at URI", err)
}

func RedisConnErr(err error) error {
	return composeError("couldn't reach redis server", err)
}

func RedisPublishErr(err error) error {
	return composeError("error trying to publish to redis", err)
}

func composeError(msg string, err error) error {
	text := fmt.Sprintf("%s: %v", msg, err.Error())
	return errors.New(text)
}

func NotFoundErr(relation string, id int) (int, error) {
	return http.StatusNotFound, fmt.Errorf("%s id %d not found",relation, id)
}

func UnprocessableErr(msg string) (int, error) {
	return http.StatusUnprocessableEntity, fmt.Errorf(msg)
}

func InternalErr(msg string) (int, error) {
	return http.StatusInternalServerError, errors.New(msg)
}