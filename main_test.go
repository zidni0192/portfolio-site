package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainPrintsScaffoldMessage(t *testing.T) {
	originalStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = writer
	t.Cleanup(func() { os.Stdout = originalStdout })

	run()

	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	os.Stdout = originalStdout
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(output, []byte("portfolio-site scaffold\n")) {
		t.Fatalf("unexpected output %q", output)
	}
}
