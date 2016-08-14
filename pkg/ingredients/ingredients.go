package ingredients

import (
	"fmt"
	"math/rand"
)

func (i *PoolCategory) append(ingredient Ingredient) {
	i.Availables = append(i.Availables, ingredient)
}

type Ingredient interface {
	Name() string
	// Gender() string
	// Quantity() string

	NameAndQuantity() string
}

type Ingredients []Ingredient

type IngredientsPool struct {
	rand                 *rand.Rand
	MainIngredients      PoolCategory
	SecondaryIngredients PoolCategory
}

type PoolCategory struct {
	Availables []Ingredient
	Picked     []Ingredient
}

func (i *PoolCategory) Pick() Ingredient {
	// FIXME: shuffle
	i.Picked = append(i.Picked, i.Availables[0])
	i.Availables = i.Availables[1:]
	return i.Availables[0]
}

type StandardMainIngredient struct {
	name     string
	quantity string

	Gender   string
	Multiple bool
}

func NewMainIngredient(name, gender string, multiple bool) StandardMainIngredient {
	return StandardMainIngredient{
		name:     name,
		quantity: "42",

		Gender:   gender,
		Multiple: multiple,
	}
}

func (i StandardMainIngredient) Name() string { return i.name }
func (i StandardMainIngredient) NameAndQuantity() string {
	return fmt.Sprintf("%s %s", i.quantity, i.name)
}

func NewPool(rnd *rand.Rand) *IngredientsPool {
	var pool IngredientsPool
	pool.rand = rnd
	pool.MainIngredients.append(NewMainIngredient("agneau", "male", false))
	pool.MainIngredients.append(NewMainIngredient("autruche", "female", false))
	pool.MainIngredients.append(NewMainIngredient("canard", "male", false))
	pool.MainIngredients.append(NewMainIngredient("carpe", "female", false))
	pool.MainIngredients.append(NewMainIngredient("cheval", "male", false))
	pool.MainIngredients.append(NewMainIngredient("chips", "female", true))
	pool.MainIngredients.append(NewMainIngredient("dinde", "female", false))
	pool.MainIngredients.append(NewMainIngredient("foie d'oie", "male", false))
	pool.MainIngredients.append(NewMainIngredient("foie gras", "male", false))
	pool.MainIngredients.append(NewMainIngredient("jambon", "male", false))
	pool.MainIngredients.append(NewMainIngredient("lardons", "male", true))
	pool.MainIngredients.append(NewMainIngredient("lièvre", "male", false))
	pool.MainIngredients.append(NewMainIngredient("lotte", "female", false))
	pool.MainIngredients.append(NewMainIngredient("nems", "male", true))
	pool.MainIngredients.append(NewMainIngredient("oie", "female", false))
	pool.MainIngredients.append(NewMainIngredient("poney", "male", false))
	pool.MainIngredients.append(NewMainIngredient("poulet", "male", false))
	pool.MainIngredients.append(NewMainIngredient("requin", "male", false))
	pool.MainIngredients.append(NewMainIngredient("saucisse", "female", false))
	pool.MainIngredients.append(NewMainIngredient("saucisses Knacki®", "female", true))
	pool.MainIngredients.append(NewMainIngredient("surimi", "male", false))
	pool.MainIngredients.append(NewMainIngredient("veau", "male", false))
	// pool.MainIngredients.append(NewMainIngredient("", "", false))
	return &pool
}