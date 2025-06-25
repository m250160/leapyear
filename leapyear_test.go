package main

import (
    "testing"
)

func TestLeapYear01(t *testing.T) {
    expected := leapyear(1)
    actual := false
    if expected != actual {
        t.Errorf("Error")
    }
}

func TestLeapYear02(t *testing.T) {
    expected := leapyear(2)
    actual := false
    if expected != actual {
        t.Errorf("Error")
    }
}
