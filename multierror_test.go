package multierror

import (
    "testing"
    "fmt"
    "errors"
)

func TestNewEmpty(t *testing.T) {
    fmt.Println(t.Name())
    me := New()
    if me == nil {
        t.Error("me is nil")
    }
    if me.Error() != "" {
        t.Error("error is not empty")
    }
    if me.ErrorOrNil() != nil {
        t.Error("ErrorOrNil is not nil")
    }
    if me.ErrorsOrNil() != nil {
        t.Error("ErrorsOrNil is not nil")
    }
    if me.Size() != 0 {
        t.Error("Size is not 0")
    }
}

func TestNew1(t *testing.T) {
    fmt.Println(t.Name())
    err := errors.New("error1")
    me := New(err)
    if me == nil {
        t.Error("me is nil")
    }
    if me.Size() != 1 {
        t.Error("Size is not 1")
    }
    if me.Error() != err.Error() {
        t.Errorf("unexpected error: %s (expected %s)\n", me.Error(), err.Error())
    }
    if me.ErrorOrNil() == nil {
        t.Error("ErrorOrNil is nil")
    }
    if me.ErrorOrNil().Error() != err.Error() {
        t.Error("ErrorOrNil is nil")
    }
    if me.ErrorsOrNil() == nil {
        t.Error("ErrorsOrNil is nil")
    }
    if me.ErrorsOrNil()[0] == nil {
        t.Error("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[0].Error() != err.Error() {
        t.Errorf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err.Error())
    }
}


func TestNew2(t *testing.T) {
    fmt.Println(t.Name())
    err1 := errors.New("error1")
    err2 := errors.New("error2")
    me := New(err1, err2)
    if me == nil {
        t.Error("me is nil")
    }
    if me.Size() != 2 {
        t.Error("Size is not 2")
    }
    expected := fmt.Sprintf("%s\n%s", err1.Error(), err2.Error())
    if me.Error() != expected {
        t.Errorf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorOrNil() == nil {
        t.Error("ErrorOrNil is nil")
    }
    if me.ErrorOrNil().Error() != expected {
        t.Errorf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorsOrNil() == nil {
        t.Error("ErrorsOrNil is nil")
    }
    if me.ErrorsOrNil()[0] == nil {
        t.Error("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[0].Error() != err1.Error() {
        t.Errorf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err1.Error())
    }
    if me.ErrorsOrNil()[1] == nil {
        t.Error("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[1].Error() != err2.Error() {
        t.Errorf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err2.Error())
    }
}

func TestNew3(t *testing.T) {
    fmt.Println(t.Name())
    err1 := errors.New("error1")
    err2 := errors.New("error2")
    me := New()
    me.Add(err1, err2)
    if me == nil {
        t.Error("me is nil")
    }
    if me.Size() != 2 {
        t.Error("Size is not 2")
    }
    expected := fmt.Sprintf("%s\n%s", err1.Error(), err2.Error())
    if me.Error() != expected {
        t.Errorf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorOrNil() == nil {
        t.Error("ErrorOrNil is nil")
    }
    if me.ErrorOrNil().Error() != expected {
        t.Errorf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorsOrNil() == nil {
        t.Error("ErrorsOrNil is nil")
    }
    if me.ErrorsOrNil()[0] == nil {
        t.Error("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[0].Error() != err1.Error() {
        t.Errorf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err1.Error())
    }
    if me.ErrorsOrNil()[1] == nil {
        t.Error("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[1].Error() != err2.Error() {
        t.Errorf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err2.Error())
    }
}

func TestNew4(t *testing.T) {
    fmt.Println(t.Name())
    err1 := errors.New("error1")
    err2 := errors.New("error2")
    me := New()
    me.Add(err1)
    me.Add(err2)
    if me == nil {
        t.Error("me is nil")
    }
    if me.Size() != 2 {
        t.Error("Size is not 2")
    }
    expected := fmt.Sprintf("%s\n%s", err1.Error(), err2.Error())
    if me.Error() != expected {
        t.Errorf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorOrNil() == nil {
        t.Error("ErrorOrNil is nil")
    }
    if me.ErrorOrNil().Error() != expected {
        t.Errorf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorsOrNil() == nil {
        t.Error("ErrorsOrNil is nil")
    }
    if me.ErrorsOrNil()[0] == nil {
        t.Error("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[0].Error() != err1.Error() {
        t.Errorf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err1.Error())
    }
    if me.ErrorsOrNil()[1] == nil {
        t.Error("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[1].Error() != err2.Error() {
        t.Errorf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err2.Error())
    }
}