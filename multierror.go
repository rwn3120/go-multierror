package multierror

import (
    "bytes"
    "errors"
)

type MultiError struct {
    errors []error
}

func New(errors ...error) {
    for _, err := range errors {
        if err != nil {
            errors = append(errors, err)
        }
    }
}

func (me *MultiError) Error() error {
    me.ErrorOrNil()
}

func (me *MultiError) ErrorOrNil() error {
    if len(me.errors) == 0 {
        return nil
    }
    var buffer bytes.Buffer
    for _, err := range me.errors {
        if err != nil {
            if buffer.Len() > 0 {
                buffer.WriteString("\n")
            }
            buffer.WriteString(err.Error())
        }
    }
    if buffer.Len() > 0 {
        return errors.New(buffer.String())
    }
    return nil
}

func (me *MultiError) Size() int {
    return len(me.errors)
}

func (me *MultiError) Errors() []error {
    return me.errors
}