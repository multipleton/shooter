package weapon

type Weapon struct {
	Name           string
	Damage         uint16
	MagazineSize   uint16
	Ammo           uint16
	ReloadSpeed    float32 // seconds
	Accuracy       uint16  // num/10 hits
	FireRate       uint16  // bullets per minute
	FireRange      float32 // meters
	BulletSpeed    float32 // meters per second
	MobilityDebuff float32 // percents debuff
}
