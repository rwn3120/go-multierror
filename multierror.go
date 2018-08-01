package multierror

import (
    "bytes"
    "errors"
    "fmt"
)

type MultiError struct {
    errors []error
}

func New(errors ...error) *MultiError {
    me := &MultiError{}
    for _, err := range errors {
        me.Add(err)
    }
    return me
}

func NewError(format string, args ...interface{}) *MultiError {
    err := errors.New(fmt.Sprintf(format, args...))
    return New(err)
}

func (me *MultiError) Add(err error) *MultiError {
    if err != nil {
        me.errors = append(me.errors, err)
    }
    return me
}

func (me *MultiError) Error() string {
    err := me.ErrorOrNil()
    if err != nil {
        return err.Error()
    }
    return ""
}

func (me *MultiError) ErrorsOrNil() []error {
    if len(me.errors) == 0 {
        return nil
    } else {
        return me.errors
    }
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