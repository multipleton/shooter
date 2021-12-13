package weapon

type Weapon struct {
	Name           string  `json:"name"`
	Damage         uint16  `json:"damage"`
	MagazineSize   uint16  `json:"magazineSize"`
	Ammo           uint16  `json:"ammo"`
	ReloadSpeed    float32 `json:"reloadSpeed"`    // seconds
	Accuracy       uint16  `json:"accuracy"`       // num/10 hits
	FireRate       uint16  `json:"fireRate"`       // bullets per minute
	FireRange      float32 `json:"fireRange"`      // meters
	BulletSpeed    float32 `json:"bulletSpeed"`    // meters per second
	MobilityDebuff float32 `json:"mobilityDebuff"` // percents debuff
}
