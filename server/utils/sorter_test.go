package utils

import (
	"reflect"
	"testing"

	"main.go/model"
)

func TestSortByPrice(t *testing.T) {
	flights := []model.Flight{
		{BokingId: "B2", Price: 300},
		{BokingId: "B3", Price: 150},
		{BokingId: "B1", Price: 450},
	}

	sorted := SortByPrice(append([]model.Flight(nil), flights...))

	if len(sorted) != 3 {
		t.Fatalf("expected 3 flights, got %d", len(sorted))
	}

	prices := []float64{sorted[0].Price, sorted[1].Price, sorted[2].Price}
	expectedPrices := []float64{150, 300, 450}
	if !reflect.DeepEqual(prices, expectedPrices) {
		t.Fatalf("expected prices %v, got %v", expectedPrices, prices)
	}

	ids := []string{sorted[0].BokingId, sorted[1].BokingId, sorted[2].BokingId}
	expectedIDs := []string{"B3", "B2", "B1"}
	if !reflect.DeepEqual(ids, expectedIDs) {
		t.Fatalf("expected booking order %v, got %v", expectedIDs, ids)
	}
}

