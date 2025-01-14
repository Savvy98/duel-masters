package dm06

import (
	"duel-masters/game/civ"
	"duel-masters/game/cnd"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
	"fmt"
)

func MightyBanditAceOfThieves(c *match.Card) {

	c.Name = "Mighty Bandit, Ace of Thieves"
	c.Power = 2000
	c.Civ = civ.Nature
	c.Family = family.BeastFolk
	c.ManaCost = 3
	c.ManaRequirement = []string{civ.Nature}
	c.TapAbility = true

	c.Use(fx.Creature, fx.When(fx.TapAbility, func(card *match.Card, ctx *match.Context) {

		ctx.Match.Chat("Server", fmt.Sprintf("%s activated %s's tap ability", card.Player.Username(), card.Name))
		creatures := match.Search(card.Player, ctx.Match, card.Player, match.BATTLEZONE, "Select 1 creature from your battlezone that will gain +5000 Power", 1, 1, false)
		for _, creature := range creatures {

			creature.AddCondition(cnd.PowerAmplifier, 5000, card.ID)
			ctx.Match.Chat("Server", fmt.Sprintf("%s was given +5000 power by %s until end of turn", creature.Name, card.Name))

			card.Tapped = true
		}
	}))
}
