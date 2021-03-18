package qheap1

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestQHeap1(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`5
		1 4
		1 9
		3
		2 4
		3`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	QHeap1(in)
}

func TestQHeap2(t *testing.T) {
	in, _ := ioutil.TempFile(os.TempDir(), "")
	defer func() { _ = in.Close() }()

	testcase :=
		`22
		1 286789035
		1 255653921
		1 274310529
		1 494521015
		3
		2 255653921
		2 286789035
		3
		1 236295092
		1 254828111
		2 254828111
		1 465995753
		1 85886315
		1 7959587
		1 20842598
		2 7959587
		3
		1 -51159108
		3
		2 -51159108
		3
		1 789534713`
	if _, err := io.WriteString(in, testcase); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		t.Error(err)
		t.FailNow()
	}
	QHeap1(in)
	// nolint:gocritic
	/*
		255653921
		274310529
		20842598
		-51159108
		20842598
	*/
}
