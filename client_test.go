package ctago_test

import (
	"os"
	"testing"

	ctago "github.com/jackmerrill/cta-go"
)

func TestGetArrivals(t *testing.T) {
	client := ctago.NewClient(os.Getenv("API_KEY"))

	arrivals, err := client.Arrivals.Get(ctago.Stations["Monroe (Red)"])

	if err != nil {
		t.Error(err)
	}

	if arrivals == nil {
		t.Error("Arrivals was nil")
	}
}
