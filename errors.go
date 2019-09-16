package predicate

import "errors"

func badParameter(message string, args ...interface{}) error {
	return errors.New(message)
}

func notFound(message string, args ...interface{}) error {
	return errors.New(message)
}

func wrap(err error, args ...interface{}) error {
	return err
}