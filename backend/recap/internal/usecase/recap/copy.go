package recap

import (
	"fmt"
	"math"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

const (
	titleActiveDays       = "Количество активных дней"
	titleViews            = "Просмотры объявлений"
	titleFavorites        = "Добавления в избранное"
	titleFavoriteCategory = "Любимая категория"
	titlePurchases        = "Покупки за год"
	titleSales            = "Продажи за год"
	titleMessages         = "Сообщения"
	titleInterests        = "Как менялись интересы"
	titleFinal            = "Это был ваш год"
	titleArchetype        = "Ваш тип на площадке"
)

const daysInYear = 365

var ordinals = map[int64]string{
	2: "второй", 3: "третий", 4: "четвёртый", 5: "пятый",
	6: "шестой", 7: "седьмой", 8: "восьмой", 9: "девятый", 10: "десятый",
}

func plural(count int64, one, few, many string) string {
	if count < 0 {
		count = -count
	}

	switch {
	case count%100 >= 11 && count%100 <= 14:
		return many
	case count%10 == 1:
		return one
	case count%10 >= 2 && count%10 <= 4:
		return few
	default:
		return many
	}
}

func everyNthDay(days int64) int64 {
	if days <= 0 {
		return 0
	}

	return int64(math.Round(daysInYear / float64(days)))
}

func activeDaysHeadline(days int64) string {
	switch {
	case days <= 0:
		return ""
	case days >= 330:
		return "Вы заходили почти каждый день!"
	case days >= daysInYear/2:
		return "Вы были с Авито больше половины года!"
	}

	every := everyNthDay(days)

	if ordinal, known := ordinals[every]; known {
		return fmt.Sprintf("Вы были с Авито почти каждый %s день!", ordinal)
	}

	return fmt.Sprintf("Вы заглядывали раз в %d %s", every, plural(every, "день", "дня", "дней"))
}

func viewsHeadline(views int64) string {
	if views <= 0 {
		return ""
	}

	return fmt.Sprintf("%s вы посмотрели в этом году",
		plural(views, "объявление", "объявления", "объявлений"))
}

func favoritesHeadline(favorites int64) string {
	if favorites <= 0 {
		return ""
	}

	return fmt.Sprintf("%s ваше внимание",
		plural(favorites, "товар сохранил", "товара сохранили", "товаров сохранили"))
}

func purchasesHeadline(purchases int64) string {
	if purchases <= 0 {
		return ""
	}

	return fmt.Sprintf("%s вы совершили в этом году",
		plural(purchases, "покупку", "покупки", "покупок"))
}

func salesHeadline(sales int64) string {
	if sales <= 0 {
		return ""
	}

	return fmt.Sprintf("%s в этом году",
		plural(sales, "успешная продажа", "успешные продажи", "успешных продаж"))
}

func messagesHeadline(messages int64) string {
	if messages <= 0 {
		return ""
	}

	return fmt.Sprintf("%s отправлено",
		plural(messages, "сообщение", "сообщения", "сообщений"))
}

func activeDaysStatLabel(days int64) string {
	return plural(days, "активный день", "активных дня", "активных дней")
}

func viewsStatLabel(views int64) string {
	return plural(views, "объявление просмотрено", "объявления просмотрено", "объявлений просмотрено")
}

func favoritesStatLabel(favorites int64) string {
	return plural(favorites, "товар в избранном", "товара в избранном", "товаров в избранном")
}

func messagesStatLabel(messages int64) string {
	return plural(messages, "сообщение отправлено", "сообщения отправлено", "сообщений отправлено")
}

func seasonsStatLabel(seasons int64) string {
	return plural(seasons, "сезон интересов", "сезона интересов", "сезонов интересов")
}

func categoryHeadline(category entity.CategoryScore, share int32) string {
	return fmt.Sprintf("%s забрала %d%% вашей активности", category.Title, share)
}
