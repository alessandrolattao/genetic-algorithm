package main

import "fmt"

// Antarctic research station resupply via shipping container.
// Container capacity: 10 tonnes. Choose from 64 items to maximize mission value.
// Values are intentionally decoupled from weights to create real tradeoffs:
// some light items are critical, some heavy items are low priority, and vice versa.

const MaxWeight = 10_000 // 10,000 kg

type Item struct {
	Name   string
	Weight int // kg
	Value  int
}

var items = []Item{
	// Power & Energy - heavy but value varies wildly
	{"Diesel Generator 50kW", 1200, 90},  // heavy, meh value
	{"Diesel Generator 20kW", 680, 150},   // lighter, better value
	{"Diesel Generator 10kW", 420, 85},    // trap: light but low value
	{"Solar Panel Array A", 450, 200},     // great value for weight
	{"Solar Panel Array B", 380, 60},      // trap: similar weight, bad value
	{"Wind Turbine 5kW", 800, 70},         // heavy and not worth it
	{"Wind Turbine 2kW", 520, 170},        // decent tradeoff
	{"Battery Bank 100kWh", 1100, 300},    // expensive but very valuable
	{"Battery Bank 50kWh", 620, 280},      // lighter, almost as valuable
	{"Fuel Drums (1000L)", 850, 40},       // heavy, low value
	{"Fuel Drums (500L)", 440, 35},        // still bad
	{"Electrical Cable Spools", 320, 110}, // ok

	// Shelter - habitat modules are traps (huge weight)
	{"Prefab Habitat Module A", 3500, 350}, // trap: eats 35% capacity for mediocre value
	{"Prefab Habitat Module B", 2800, 400}, // slightly better trap
	{"Insulation Panels (set)", 900, 130},  // meh
	{"Emergency Shelter Tent", 280, 190},   // light, high value
	{"Welding Station", 380, 45},           // bad deal
	{"Construction Tools", 310, 160},       // good

	// Vehicles - mixed bag
	{"Snowmobile A", 350, 55},           // not great
	{"Snowmobile B", 380, 180},          // same-ish weight, way better
	{"ATV", 500, 75},                    // bad ratio
	{"Spare Parts Kit A", 220, 210},     // light, very valuable
	{"Spare Parts Kit B", 180, 50},      // trap: lighter but worthless
	{"Cargo Sled", 150, 140},            // good deal

	// Science - high value but tricky weights
	{"Weather Station Pro", 180, 250},     // excellent
	{"Weather Station Basic", 90, 240},    // even better!
	{"Seismograph Array", 250, 60},        // bad
	{"Ice Core Drill A", 450, 270},        // good but heavy
	{"Ice Core Drill B", 380, 100},        // trap: similar weight, worse
	{"Laboratory Equipment A", 600, 290},  // very valuable but heavy
	{"Laboratory Equipment B", 420, 285},  // lighter, almost same value!
	{"Sample Freezers (-80C)", 400, 70},   // not worth it
	{"Satellite Uplink Pro", 160, 310},    // best ratio in the game
	{"Satellite Uplink Basic", 90, 180},   // also great
	{"Field Instruments Kit", 140, 165},   // solid

	// Food - generally good value, hard to choose between them
	{"MRE Pallet (90 days)", 800, 260},      // decent but heavy
	{"MRE Pallet (45 days)", 420, 255},      // almost same value, half weight!
	{"Freeze-Dried Food (6mo)", 400, 175},   // ok
	{"Freeze-Dried Food (3mo)", 220, 170},   // lighter, almost same value
	{"Water Purification A", 300, 230},      // good
	{"Water Purification B", 180, 220},      // lighter, nearly same value
	{"Hydroponic Grow Kit", 350, 45},        // bad
	{"Cooking Equipment", 210, 155},         // decent

	// Medical - small items with huge value differences
	{"Medical Bay Full", 500, 295},       // heavy but very valuable
	{"Medical Bay Basic", 300, 120},      // trap: not much cheaper, way less value
	{"Surgical Kit", 120, 260},           // light, very high value
	{"Pharmacy Supplies A", 180, 245},    // great
	{"Pharmacy Supplies B", 100, 80},     // trap: lighter but bad value
	{"Fire Suppression A", 400, 65},      // not worth it
	{"Fire Suppression B", 250, 60},      // also bad
	{"Rescue Equipment", 350, 135},       // meh

	// Communications - light items, value all over the place
	{"HF Radio Station", 140, 195},       // good
	{"Starlink Terminal", 60, 275},        // amazing ratio
	{"Server Rack Full", 400, 55},         // bad
	{"Server Rack Compact", 200, 50},      // also bad
	{"Laptops (10 units)", 80, 185},       // great
	{"Networking Equipment", 110, 40},     // bad

	// Clothing & Personal - lots of similar-looking choices
	{"Extreme Cold Gear (10p)", 280, 200}, // good
	{"Extreme Cold Gear (5p)", 150, 195},  // lighter, almost same value!
	{"Sleeping Systems (10p)", 220, 145},  // ok
	{"Sleeping Systems (5p)", 120, 140},   // lighter, almost same
	{"Personal Hygiene (6mo)", 100, 90},   // meh
	{"Recreation Equipment", 140, 30},     // waste of space
	{"Laundry Equipment", 280, 35},        // terrible
	{"Morale & Comfort Pack", 90, 160},    // surprisingly good
}

const GenomeSize = 64

func fitness(genome uint64) int {
	totalWeight := 0
	totalValue := 0

	for i, item := range items {
		if isBitSet(genome, i) {
			totalWeight += item.Weight
			totalValue += item.Value
		}
	}

	if totalWeight > MaxWeight {
		return 0
	}

	return totalValue
}

func printSolution(best Individual, gen int, completed bool) {
	fmt.Printf("=== Generation %d | Fitness: %d ===\n", gen, best.Fitness)

	totalWeight := 0
	selected := 0
	for i, item := range items {
		if isBitSet(best.Genome, i) {
			totalWeight += item.Weight
			selected++
			if completed {
				fmt.Printf("  [%2d] %-28s %5d kg  val:%3d  (ratio:%.2f)\n",
					i, item.Name, item.Weight, item.Value,
					float64(item.Value)/float64(item.Weight))
			}
		}
	}
	fmt.Printf("  --- %d/%d items | %d kg / %d kg (%.1f%% full) ---\n\n",
		selected, GenomeSize, totalWeight, MaxWeight,
		float64(totalWeight)/float64(MaxWeight)*100)
}
