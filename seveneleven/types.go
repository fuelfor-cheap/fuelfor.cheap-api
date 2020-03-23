package seveneleven

var (
	fuelTypeToEan = map[string]int{
		"300b4ebb-2b7e-45ba-8724-95c82b425815": 52,
		"91e3012e-76d7-4634-9119-979504b98db":  53,
		"c4f5e04b-4d42-4050-b7f7-ab1481275bf2": 54,
		"f5bf8658-027a-4849-8822-564838bf0a7b": 55,
		"357a46c7-94e7-4d2b-9df3-1ac634d7cfd6": 56,
		"ae2c694e-4594-421b-ac9f-3f9c581724d0": 57,
	}

	// EanToFuelType ...
	EanToFuelType = map[int]string{
		52: "300b4ebb-2b7e-45ba-8724-95c82b425815",
		53: "91e3012e-76d7-4634-9119-979504b98db",
		54: "c4f5e04b-4d42-4050-b7f7-ab1481275bf2",
		55: "f5bf8658-027a-4849-8822-564838bf0a7b",
		56: "357a46c7-94e7-4d2b-9df3-1ac634d7cfd6",
		57: "ae2c694e-4594-421b-ac9f-3f9c581724d0",
	}
)
