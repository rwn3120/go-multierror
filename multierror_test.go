package multierror

import (
    "testing"
    "fmt"
    "errors"
)

func TestNewEmpty(t *testing.T) {
    fmt.Println(t.Name(), "... running")

    me := New()
    if me == nil {
        t.Fatal("me is nil")
    }
    if me.Error() != "" {
        t.Fatal("error is not empty")
    }
    if me.ErrorOrNil() != nil {
        t.Fatal("ErrorOrNil is not nil")
    }
    if me.ErrorsOrNil() != nil {
        t.Fatal("ErrorsOrNil is not nil")
    }
    if me.Size() != 0 {
        t.Fatal("Size is not 0")
    }
}

func TestNew1(t *testing.T) {
    fmt.Println(t.Name(), "... running")

    err := errors.New("error1")
    me := New(err)
    if me == nil {
        t.Fatal("me is nil")
    }
    if me.Size() != 1 {
        t.Fatal("Size is not 1")
    }
    if me.Error() != err.Error() {
        t.Fatalf("unexpected error: %s (expected %s)\n", me.Error(), err.Error())
    }
    if me.ErrorOrNil() == nil {
        t.Fatal("ErrorOrNil is nil")
    }
    if me.ErrorOrNil().Error() != err.Error() {
        t.Fatal("ErrorOrNil is nil")
    }
    if me.ErrorsOrNil() == nil {
        t.Fatal("ErrorsOrNil is nil")
    }
    if me.ErrorsOrNil()[0] == nil {
        t.Fatal("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[0].Error() != err.Error() {
        t.Fatalf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err.Error())
    }
}


func TestNew2(t *testing.T) {
    fmt.Println(t.Name(), "... running")

    err1 := errors.New("error1")
    err2 := errors.New("error2")
    me := New(err1, err2)
    if me == nil {
        t.Fatal("me is nil")
    }
    if me.Size() != 2 {
        t.Fatal("Size is not 2")
    }
    expected := fmt.Sprintf("%s\n%s", err1.Error(), err2.Error())
    if me.Error() != expected {
        t.Fatalf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorOrNil() == nil {
        t.Fatal("ErrorOrNil is nil")
    }
    if me.ErrorOrNil().Error() != expected {
        t.Fatalf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorsOrNil() == nil {
        t.Fatal("ErrorsOrNil is nil")
    }
    if me.ErrorsOrNil()[0] == nil {
        t.Fatal("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[0].Error() != err1.Error() {
        t.Fatalf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err1.Error())
    }
    if me.ErrorsOrNil()[1] == nil {
        t.Fatal("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[1].Error() != err2.Error() {
        t.Fatalf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err2.Error())
    }
}

func TestNew3(t *testing.T) {
    fmt.Println(t.Name(), "... running")

    err1 := errors.New("error1")
    err2 := errors.New("error2")
    me := New()
    me.Add(err1, err2)
    if me == nil {
        t.Fatal("me is nil")
    }
    if me.Size() != 2 {
        t.Fatal("Size is not 2")
    }
    expected := fmt.Sprintf("%s\n%s", err1.Error(), err2.Error())
    if me.Error() != expected {
        t.Fatalf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorOrNil() == nil {
        t.Fatal("ErrorOrNil is nil")
    }
    if me.ErrorOrNil().Error() != expected {
        t.Fatalf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorsOrNil() == nil {
        t.Fatal("ErrorsOrNil is nil")
    }
    if me.ErrorsOrNil()[0] == nil {
        t.Fatal("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[0].Error() != err1.Error() {
        t.Fatalf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err1.Error())
    }
    if me.ErrorsOrNil()[1] == nil {
        t.Fatal("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[1].Error() != err2.Error() {
        t.Fatalf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err2.Error())
    }
}

func TestNew4(t *testing.T) {
    fmt.Println(t.Name(), "... running")

    err1 := errors.New("error1")
    err2 := errors.New("error2")
    me := New()
    me.Add(err1)
    me.Add(err2)
    if me == nil {
        t.Fatal("me is nil")
    }
    if me.Size() != 2 {
        t.Fatal("Size is not 2")
    }
    expected := fmt.Sprintf("%s\n%s", err1.Error(), err2.Error())
    if me.Error() != expected {
        t.Fatalf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorOrNil() == nil {
        t.Fatal("ErrorOrNil is nil")
    }
    if me.ErrorOrNil().Error() != expected {
        t.Fatalf("unexpected error: %s (expected %s)\n", me.Error(), expected)
    }
    if me.ErrorsOrNil() == nil {
        t.Fatal("ErrorsOrNil is nil")
    }
    if me.ErrorsOrNil()[0] == nil {
        t.Fatal("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[0].Error() != err1.Error() {
        t.Fatal("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err1.Error())
    }
    if me.ErrorsOrNil()[1] == nil {
        t.Fatal("ErrorsOrNil[0] is nil")
    }
    if me.ErrorsOrNil()[1].Error() != err2.Error() {
        t.Fatalf("unexpected ErrorsOrNil[0] error: %s (expected %s)\n", me.Error(), err2.Error())
    }
}