package main

import (
    "testing"
)

func TestAddZeros(t *testing.T) {
    t.Run("should_return_0000_if_n_4", func(t *testing.T) {
        got := addZeros(4)
        want := "0000"
        if got != want {
            t.Errorf("got '%s',want '%s'", got, want)
        }
    })
}

func TestReverseString(t *testing.T) {
    s := "12345"
    got := reverseString(s, 0, len(s)-1)
    want := "54321"
    if got != want {
        t.Errorf("got '%s',want '%s'", got, want)
    }
}

func TestAddString(t *testing.T) {
    s1 := "12345"
    s2 := "12345"
    got := addStr(s1, s2)
    want := "24690"
    if got != want {
        t.Errorf("got '%s',want '%s'", got, want)
    }
}

func TestRotateList(t *testing.T) {

}


