package converter

import (
	"encoding/json"
	"fmt"

	"github.com/avito-hackaton-team/avito-recap/backend/recap/generated/recapapi"
	"github.com/avito-hackaton-team/avito-recap/backend/recap/internal/entity"
)

func ConvertEntityProfileToAPIProfile(profile entity.Profile) recapapi.Profile {
	result := recapapi.Profile{
		ID:      recapapi.UUID(profile.ID),
		Name:    profile.Name,
		Surname: profile.Surname,
	}

	if profile.Hint != "" {
		result.Hint = recapapi.NewOptString(profile.Hint)
	}

	return result
}

func ConvertEntityRecapToAPIRecap(recap entity.Recap) (recapapi.Recap, error) {
	slides, err := ConvertRawMessageToAPISlides(recap.Slides)
	if err != nil {
		return recapapi.Recap{}, err
	}

	return recapapi.Recap{
		ID:          recapapi.UUID(recap.ID),
		ProfileId:   recapapi.UUID(recap.UserID),
		Year:        ConvertIntToAPIYear(recap.Year),
		Status:      recapapi.RecapStatusReady,
		Archetype:   ConvertEntityArchetypeToAPIArchetype(recap.Archetype),
		Slides:      slides,
		GeneratedAt: recap.GeneratedAt,
	}, nil
}

func ConvertRawMessageToAPISlides(raw json.RawMessage) ([]recapapi.Slide, error) {
	if len(raw) == 0 {
		return nil, nil
	}

	var slides []recapapi.Slide
	if err := json.Unmarshal(raw, &slides); err != nil {
		return nil, fmt.Errorf("decode slides: %w", err)
	}

	return slides, nil
}

func ConvertEntityArchetypeToAPIArchetype(archetype entity.Archetype) recapapi.Archetype {
	reasons := make([]recapapi.ArchetypeReason, 0, len(archetype.Reasons))

	for _, reason := range archetype.Reasons {
		reasons = append(reasons, recapapi.ArchetypeReason{
			Metric:      recapapi.MetricCode(reason.Metric),
			Value:       reason.Value,
			Explanation: reason.Explanation,
		})
	}

	return recapapi.Archetype{
		Code:        recapapi.ArchetypeCode(archetype.UserArchetype),
		Title:       archetype.Title,
		Description: archetype.Description,
		Reasons:     reasons,
	}
}

func ConvertIntToAPIYear(year int) recapapi.Year {
	return recapapi.Year(year)
}
