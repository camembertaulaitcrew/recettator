package recettator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/camembertaulaitcrew/recettator/pkg/ingredients"
	"github.com/moul/advanced-ssh-config/pkg/templates"
)

type Settings struct {
	MainIngredients      uint64
	SecondaryIngredients uint64
	Steps                uint64
	// vegan
}

type Recettator struct {
	// components
	title  string
	people uint64
	steps  Steps
	pool   *ingredients.IngredientsPool

	// settings
	settings Settings

	// internal
	seed  int64
	ready bool
	rnd   *rand.Rand // global random, used to add ingredients and steps
}

func New(seed int64) Recettator {
	rnd := rand.New(rand.NewSource(seed))
	return Recettator{
		seed:  seed,
		steps: make(Steps, 0),
		pool:  ingredients.NewPool(rnd),
		rnd:   rnd,
	}
}

func (r *Recettator) applyDefaults() {
	if r.settings.MainIngredients == 0 {
		qty := r.rnd.Intn(4) - 1
		if qty > 0 {
			r.settings.MainIngredients = uint64(qty)
		}
	}
}

func (r *Recettator) pickItems() {
	for i := uint64(0); i < r.settings.MainIngredients; i++ {
		r.pool.MainIngredients.Pick()
	}
	for i := uint64(0); i < r.settings.SecondaryIngredients; i++ {
		r.pool.SecondaryIngredients.Pick()
	}
}

func (r *Recettator) isValid() error {
	if r.settings.MainIngredients+r.settings.SecondaryIngredients < 1 {
		return fmt.Errorf("not enough ingredients.")
	}
	return nil
}

func (r *Recettator) prepare() {
	if r.ready {
		return
	}

	// pick items
	r.applyDefaults()
	if err := r.isValid(); err != nil {
		logrus.Fatalf("Invalid recette: %v", err)
	}
	r.pickItems()

	// compute fields
	titleParts := []string{}
	//var left ingredients.Ingredients
	//for _, ingredient := range r.pool.MainIngredients.Picked {
	//	titleParts = append(titleParts, ingredient.GetTitleParts(left))
	//	left = ingredient
	//}
	r.title = strings.Join(titleParts, " ")
	r.people = uint64(r.rnd.Intn(4) + 1)

	r.ready = true
}

func (r *Recettator) Seed() int64                        { return r.seed }
func (r *Recettator) Settings() Settings                 { return r.settings }
func (r *Recettator) Title() string                      { r.prepare(); return r.title }
func (r *Recettator) People() uint64                     { r.prepare(); return r.people }
func (r *Recettator) Pool() *ingredients.IngredientsPool { r.prepare(); return r.pool }
func (r *Recettator) Steps() Steps                       { r.prepare(); return r.steps }

func (r *Recettator) SetSettings(settings Settings) {
	r.settings = settings
}

func (r *Recettator) Markdown() (string, error) {
	var buff bytes.Buffer
	tmpl, err := templates.New(strings.TrimSpace(`
# {{ .Title }}

Pour {{ .People }} personnes.

## Ingrédients

{{ range .Pool.MainIngredients.Picked }}* {{ .NameAndQuantity }}
{{ end }}{{ range .Pool.SecondaryIngredients.Picked }}* {{ .NameAndQuantity }}
{{ end }}
## Etapes

{{ .Steps }}
`))
	if err != nil {
		return "", err
	}

	if err := tmpl.Execute(&buff, r); err != nil {
		return "", err
	}

	return buff.String(), nil
}

func (r *Recettator) JSON() string {
	export := make(map[string]interface{}, 0)
	r.prepare()
	export["seed"] = r.seed
	export["title"] = r.title
	export["steps"] = r.steps
	export["people"] = r.people
	export["settings"] = r.settings
	export["pool"] = map[string][]ingredients.IngredientMap{
		"main-ingredients":      r.pool.MainIngredients.Picked.ToMap(),
		"secondary-ingredients": r.pool.SecondaryIngredients.Picked.ToMap(),
	}
	output, _ := json.MarshalIndent(export, "", "  ")
	return string(output)
}
