package main

func calculatePopulation(cities []City) map[string]int64 {
	result := make(map[string]int64)
	for _, city := range cities {
		if city.CountryCode.Valid {
			// まだmapに存在しなかった場合、初期化する
			if _, ok := result[city.CountryCode.String]; !ok {
				result[city.CountryCode.String] = 0
			}
			result[city.CountryCode.String] += city.Population.Int64
		}
	}
	return result
}
