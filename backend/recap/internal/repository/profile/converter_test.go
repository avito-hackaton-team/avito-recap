package profile

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
)

func TestNullableTextValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value pgtype.Text
		want  *string
	}{
		{name: "null", value: pgtype.Text{}, want: nil},
		{name: "empty string", value: pgtype.Text{String: "", Valid: true}, want: stringPointer("")},
		{name: "non-empty string", value: pgtype.Text{String: "profile hint", Valid: true}, want: stringPointer("profile hint")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := nullableTextValue(tt.value)
			if tt.want == nil {
				if got != nil {
					t.Fatalf("nullableTextValue() = %q, want nil", *got)
				}

				return
			}

			if got == nil {
				t.Fatal("nullableTextValue() = nil, want non-nil")
			}
			if *got != *tt.want {
				t.Fatalf("nullableTextValue() = %q, want %q", *got, *tt.want)
			}
		})
	}
}

func stringPointer(value string) *string {
	return &value
}
