package recap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlural(t *testing.T) {
	t.Parallel()

	tests := []struct {
		count int64
		want  string
	}{
		{1, "день"},
		{2, "дня"},
		{4, "дня"},
		{5, "дней"},
		{11, "дней"},
		{12, "дней"},
		{14, "дней"},
		{21, "день"},
		{22, "дня"},
		{25, "дней"},
		{101, "день"},
		{111, "дней"},
	}

	for _, test := range tests {
		assert.Equalf(t, test.want, plural(test.count, "день", "дня", "дней"),
			"unexpected form for %d", test.count)
	}
}

func TestEveryNthDay(t *testing.T) {
	t.Parallel()

	tests := []struct {
		days int64
		want int64
	}{
		{365, 1},
		{180, 2},
		{128, 3},
		{20, 18},
		{1, 365},
		{0, 0},
		{-5, 0},
	}

	for _, test := range tests {
		assert.Equalf(t, test.want, everyNthDay(test.days), "unexpected rhythm for %d days", test.days)
	}
}

func TestHeadlinesCoverEveryBranch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		texts []string
	}{
		{
			name: "active days",
			texts: []string{
				activeDaysHeadline(340), // almost every day
				activeDaysHeadline(200), // more than half a year
				activeDaysHeadline(128), // every third day, ordinal branch
				activeDaysHeadline(20),  // rare visits, interval branch
			},
		},
		{
			name: "counters",
			texts: []string{
				viewsHeadline(1248),
				favoritesHeadline(57),
				purchasesHeadline(11),
				salesHeadline(8),
				messagesHeadline(342),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for index, text := range test.texts {
				assert.NotEmptyf(t, text, "headline %d produced no text", index)
				assert.NotContainsf(t, text, "%!", "headline %d has a broken format", index)
			}
		})
	}
}

func TestHeadlinesStaySilentWithoutData(t *testing.T) {
	t.Parallel()

	headlines := map[string]func(int64) string{
		"active days": activeDaysHeadline,
		"views":       viewsHeadline,
		"favorites":   favoritesHeadline,
		"purchases":   purchasesHeadline,
		"sales":       salesHeadline,
		"messages":    messagesHeadline,
	}

	for name, headline := range headlines {
		assert.Emptyf(t, headline(0), "%s headline must stay silent on an empty counter", name)
	}
}

func TestStatLabelsAgreeWithTheNumber(t *testing.T) {
	t.Parallel()

	labels := map[string]func(int64) string{
		"active days": activeDaysStatLabel,
		"views":       viewsStatLabel,
		"favorites":   favoritesStatLabel,
		"messages":    messagesStatLabel,
		"seasons":     seasonsStatLabel,
	}

	for name, label := range labels {
		one, few, many := label(1), label(2), label(5)

		require.NotEqualf(t, one, few, "%s label does not decline between one and few", name)
		require.NotEqualf(t, few, many, "%s label does not decline between few and many", name)
		require.NotEqualf(t, one, many, "%s label does not decline between one and many", name)
	}
}
