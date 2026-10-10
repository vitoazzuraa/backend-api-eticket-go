package service

func ValidateEvent(name, venue string, price, quota int) map[string]string {
	errs := map[string]string{}

	if name == "" {
		errs["name"] = "nama wajib diisi"
	}
	if venue == "" {
		errs["venue"] = "venue wajib diisi"
	}
	if price <= 0 {
		errs["price"] = "harga harus lebih dari 0"
	}
	if quota <= 0 {
		errs["quota"] = "kuota harus lebih dari 0"
	}

	return errs
}
