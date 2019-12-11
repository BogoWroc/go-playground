package kyu8

func combat(health, damage float64) float64 {
	switch {
	case damage > health:
		return 0
	default:
		return health - damage
	}
}
