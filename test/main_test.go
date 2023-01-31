package test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	start := time.Now()
	code := m.Run()
	fmt.Printf("Tests completed in %s\n", time.Since(start))
	os.Exit(code)
}
