package main

const (
	SlotCount = 27 // ячеек в одиночном сундуке Minecraft
	MaxStack  = 64 // максимум предметов в одной ячейке
)

// knownItems — предметы, которые можно класть в сундук.
var knownItems = map[string]bool{

	"dirt":        true,
	"cobblestone": true,
	"stone":       true,
	"oak_log":     true,
	"oak_planks":  true,
	"sand":        true,
	"gravel":      true,
	"coal":        true,
	"iron_ore":    true,
	"iron_ingot":  true,
	"gold_ingot":  true,
	"diamond":     true,
	"redstone":    true,
	"lapis":       true,
	"glowstone":   true,
	"torch":       true,
	"stick":       true,
	"wheat":       true,
	"bread":       true,
	"apple":       true,
	"netherite":   true,
	"ender_pearl": true,
	"blaze_rod":   true,
	"totem":       true,
	"nether_star": true,
}

// IsKnownItem — напишите: вернуть true, если name есть в knownItems.
func IsKnownItem(name string) bool {
	_, ok := knownItems[name]
	if ok == true {
		return true
	} else {
		return false
	}
}
