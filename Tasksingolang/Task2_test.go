package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard) // suppress log output during tests
	os.Exit(m.Run())
}

func TestComplexWork(t *testing.T) {
	complexWork()
	// no assertions needed since the function doesn't return anything
	fmt.Println("success")
}

func TestSimpleWork(t *testing.T) {
	simpleWork()
	// no assertions needed since the function doesn't return anything
}

func TestTestComplex(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	testComplex()

	expectedOutput := "Starting complex working.. \nComplex work done!\n"
	if buf.String() != expectedOutput {
		t.Errorf("unexpected output from testComplex: expected %q, got %q", expectedOutput, buf.String())
	}
}

func TestTestSimple(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	

	testSimple()

	expectedOutput := "Starting simple working..\nSimple work done!\n"
	if buf.String() != expectedOutput {
		t.Errorf("unexpected output from testSimple: expected %q, got %q", expectedOutput, buf.String())
	}
}

func TestMainFunction(t *testing.T) {
    var buf bytes.Buffer
    log.SetOutput(&buf)

    // redirect stdout to buffer to capture output
    oldStdout := os.Stdout
    w := bufio.NewWriter(&buf)
	
    defer w.Flush()
    
    main()

    // restore original stdout
    os.Stdout = oldStdout

    // check output for expected execution times
    expectedOutput := "Execution time without goroutines:"
    if !bytes.Contains(buf.Bytes(), []byte(expectedOutput)) {
        t.Errorf("expected output to contain %q, but got %q", expectedOutput, buf.String())
    }
    expectedOutput = "Execution time with goroutines:"
    if !bytes.Contains(buf.Bytes(), []byte(expectedOutput)) {
        t.Errorf("expected output to contain %q, but got %q", expectedOutput, buf.String())
    }
}




