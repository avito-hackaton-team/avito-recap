package recap

const (
	badgeBronze = "bronze"
	badgeSilver = "silver"
	badgeGold   = "gold"
)

type badge struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Level       string `json:"level"`
}

type badgeRule struct {
	threshold   int64
	code        string
	title       string
	description string
	level       string
}

var (
	purchaseBadges = []badgeRule{
		{10, "buyer_gold", "Знаток покупок", "10 и больше покупок за год", badgeGold},
		{5, "buyer_silver", "Уверенный покупатель", "5 и больше покупок за год", badgeSilver},
		{1, "buyer_bronze", "Первая покупка", "Есть закрытые покупки за год", badgeBronze},
	}

	saleBadges = []badgeRule{
		{10, "seller_gold", "Мастер продаж", "10 и больше продаж за год", badgeGold},
		{5, "seller_silver", "Опытный продавец", "5 и больше продаж за год", badgeSilver},
		{1, "seller_bronze", "Первая продажа", "Есть закрытые продажи за год", badgeBronze},
	}
)

func awardBadge(rules []badgeRule, value int64) *badge {
	for _, rule := range rules {
		if value >= rule.threshold {
			return &badge{
				Code:        rule.code,
				Title:       rule.title,
				Description: rule.description,
				Level:       rule.level,
			}
		}
	}

	return nil
}
