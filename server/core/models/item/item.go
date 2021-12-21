package item

type Item struct {
	Name        string  `json:"name"`
	SpawnChance float32 `json:"spawnChance"` // percents
}
