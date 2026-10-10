package service

import "testing"

func TestValidateEvent(t *testing.T) {
	cases := []struct {
		desc     string
		name     string
		venue    string
		price    int
		quota    int
		wantErrs int
	}{
		{"semua isian valid", "Comifuro", "Stage 1", 20000, 100, 0},
		{"nama kosong", "", "Stage 1", 20000, 100, 1},
		{"harga nol", "Comifuro", "Stage 1", 0, 100, 1},
		{"kuota minus", "Comifuro", "Stage 1", 20000, -5, 1},
		{"semua isian salah", "", "", 0, 0, 4},
	}

	for _, tc := range cases {
		errs := ValidateEvent(tc.name, tc.venue, tc.price, tc.quota)

		if len(errs) != tc.wantErrs {
			t.Errorf("%s: harap %d error, dapat %d (%v)",
				tc.desc, tc.wantErrs, len(errs), errs)
		}
	}
}
