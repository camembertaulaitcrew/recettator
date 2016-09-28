package ingredients

import (
	"fmt"
	"math/rand"
)

type SecondaryIngredient struct {
	name          string
	gender        string
	quantity      string
	isMultiple    bool
	isUncountable bool
	isPowder      bool
	isCitrus      bool
	isSpice       bool
	isByPiece     bool
	isSpreadable  bool
	rand          *rand.Rand
}

func NewSecondaryIngredient(name string, gender string, isMultiple bool, rnd *rand.Rand) *SecondaryIngredient {
	ingredient := SecondaryIngredient{
		name:       name,
		gender:     gender,
		isMultiple: isMultiple,
		rand:       rnd,
		/*
			isMultiple:
			isUncountable:
			isPowder:
			isCitrus:
			isSpice:
			isByPiece:
			isSpreadable:
		*/
	}
	return &ingredient
}

func (i *SecondaryIngredient) prepare() {
	switch {
	case i.isUncountable:
		switch {
		case i.isMultiple:
			i.quantity = "des "
			break
		case beginsWithVoyel(i.name):
			i.quantity = "de l'"
			break
		case i.gender == "male":
			i.quantity = "du "
			break
		case i.gender == "female":
			i.quantity = "de la "
			break
		}
		break
	case i.isPowder:
		value := (i.rand.Intn(50) + 1) * 10
		switch {
		case value == 1 && !beginsWithVoyel(i.name):
			i.quantity = fmt.Sprintf("%d gramme de ", value)
			break
		case value == 1 && beginsWithVoyel(i.name):
			i.quantity = fmt.Sprintf("%d gramme d'", value)
			break
		case !beginsWithVoyel(i.name):
			i.quantity = fmt.Sprintf("%d grammes de ", value)
			break
		case beginsWithVoyel(i.name):
			i.quantity = fmt.Sprintf("%d grammes d'", value)
			break
		}
		break
	case i.isByPiece:
		break
	case i.isSpice:
		break
	case i.isSpreadable:
		break
	case i.isCitrus:
		break
	}
}

func (i SecondaryIngredient) Kind() string { return "secondary-ingredient" }
func (i SecondaryIngredient) Name() string { return i.name }
func (i SecondaryIngredient) NameAndQuantity() string {
	if i.quantity == "" {
		i.prepare()
	}
	return fmt.Sprintf("%s%s", i.quantity, i.name)
}
func (i SecondaryIngredient) GetGender() string { return i.gender }
func (i SecondaryIngredient) IsMultiple() bool  { return i.isMultiple }
func (i SecondaryIngredient) TitlePart(left Ingredient) string {
	// FIXME: implement
	return ""
}

func (i SecondaryIngredient) ToMap() map[string]interface{} {
	ret := make(map[string]interface{}, 0)
	ret["name"] = i.name
	ret["kind"] = i.Kind()
	ret["name-and-quantity"] = i.NameAndQuantity()
	ret["quantity"] = i.quantity
	ret["is-multiple"] = i.isMultiple
	ret["gender"] = i.gender
	ret["is-by-piece"] = i.isByPiece
	ret["is-uncountable"] = i.isUncountable
	ret["is-powder"] = i.isPowder
	ret["is-citrus"] = i.isCitrus
	ret["is-spice"] = i.isSpice
	ret["is-spreadable"] = i.isSpreadable
	return ret
}

func (i *SecondaryIngredient) SetIsByPiece() *SecondaryIngredient {
	i.isByPiece = true
	return i
}
func (i *SecondaryIngredient) SetIsSpreadable() *SecondaryIngredient {
	i.isSpreadable = true
	return i
}
func (i *SecondaryIngredient) SetIsPowder() *SecondaryIngredient {
	i.isPowder = true
	return i
}
func (i *SecondaryIngredient) SetIsUncountable() *SecondaryIngredient {
	i.isUncountable = true
	return i
}
func (i *SecondaryIngredient) SetIsSpice() *SecondaryIngredient {
	i.isSpice = true
	return i
}
func (i *SecondaryIngredient) SetIsCitrus() *SecondaryIngredient {
	i.isCitrus = true
	return i
}

//, uncountable, powder, citrus, spice, byPiece, spreadable bool) SecondaryIngredient {
