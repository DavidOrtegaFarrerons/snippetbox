package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC)
	got := humanDate(tm)
	want := "17 Mar 2022 at 10:15"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
