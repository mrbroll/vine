package recipe

const (
	Count      MeasuredUnit = ""
	Inch       MeasuredUnit = "in"
	Cup        MeasuredUnit = "c"
	Tablespoon MeasuredUnit = "tbsp"
	Teaspoon   MeasuredUnit = "tsp"
)

var (
	UnitAliases = map[string]MeasuredUnit{
		"":            Count,
		"count":       Count,
		"c":           Cup,
		"cup":         Cup,
		"cups":        Cup,
		"tbsp":        Tablespoon,
		"tbsps":       Tablespoon,
		"tablespoon":  Tablespoon,
		"tablespoons": Tablespoon,
		"tsp":         Teaspoon,
		"tsps":        Teaspoon,
		"teaspoon":    Teaspoon,
		"teaspoons":   Teaspoon,
	}
)

// Recipe represents a recipe.
type Recipe struct {
	Uid          string         `json:"uid,omitempty"`
	Name         string         `json:"name,omitempty"`
	Ingredients  []*Ingredient  `json:"ingredient,omitempty"`
	Instructions []*Instruction `json:"instruction,omitempty"`
}

// Ingredient is a single ingredient of a recipe.
type Ingredient struct {
	Uid         string       `json:"uid,omitempty"`
	Name        string       `json:"name,omitempty"`
	Measurement *Measurement `json:"measurement,omitempty"`
	Preparation string       `json:"preparation,omitempty"`
}

// Instruction is a single instruction for preparing a recipe.
type Instruction struct {
	Uid   string `json:"uid,omitempty"`
	Order int    `json:"order,omitempty"`
	Text  string `json:"text,omitempty"`
}

// Measurement describes the quantity of an ingredient in a recipe.
type Measurement struct {
	Value float64      `json:"value,omitempty"`
	Unit  MeasuredUnit `json:"unit,omitempty"`
}

type MeasuredUnit string
