package main

import (
	"database/sql"
	"testing"
)

func Test_calculatePopulation_empty(t *testing.T) {
	// ここにテストを書いていく
	cities := []City{}
	got := calculatePopulation(cities)
	want := map[string]int{}
	// 長さが0になっているかどうかを確認する
	if len(got) != 0 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}
func Test_calculatePopulation_single(t *testing.T) {
	cities := []City{
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 100,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 200,
				Valid: true,
			},
		},
	}
	// ここにテストを書いていく
	got := calculatePopulation(cities)
	want := map[string]int64{
		"JPN": 300,
	}
	// 長さが0になっているかどうかを確認する
	if len(got) != 1 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	if got["JPN"] != 300 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}
func Test_calculatePopulation_multi(t *testing.T) {
	cities := []City{
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 100,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 200,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 300,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "USA",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 70,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "USA",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 900,
				Valid: true,
			},
		},
	}
	// ここにテストを書いていく
	got := calculatePopulation(cities)
	want := map[string]int64{
		"JPN": 300,
	}
	// 長さが0になっているかどうかを確認する
	if len(got) != 2 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	if got["JPN"] != 600 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	if got["USA"] != 970 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}
func Test_calculatePopulation_withinvalid(t *testing.T) {
	cities := []City{
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 100,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 200,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 300,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "USA",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 70,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "",
				Valid:  false,
			},
			Population: sql.NullInt64{
				Int64: 300,
				Valid: true,
			},
		},
	}
	// ここにテストを書いていく
	got := calculatePopulation(cities)
	want := map[string]int64{
		"JPN": 300,
	}
	// 長さが0になっているかどうかを確認する
	if len(got) != 2 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	if got["JPN"] != 600 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
	if got["USA"] != 70 {
		t.Errorf("calculatePopulation(%v) = %v, want %v", cities, got, want)
	}
}
