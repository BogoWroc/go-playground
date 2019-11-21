package capstone

type Planets []string

func (planets Planets) Terraform() {
	for i, value := range planets {
		planets[i] = "New " + value
	}
}

func Terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}
