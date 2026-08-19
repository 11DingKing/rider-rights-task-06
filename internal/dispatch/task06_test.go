package dispatch

import (
	"context"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"riderguard/internal/domain"
)

func TestActiveDefaultRuleFallback(t *testing.T) {
	when := time.Date(2026, 8, 19, 12, 0, 0, 0, time.UTC)
	clk := clock.NewMock()
	clk.Set(when)
	item := &domain.RightsCase{ID: "case-06", Category: "心理辅导", RegisteredAt: when}
	rule := &domain.Rule{Version: 2, Name: "default", IsDefault: true, LeadDepartment: "骑手之家", EffectiveFrom: when.Add(-time.Hour), Status: domain.RuleStatusActive}
	ref, err := NewAdjudicator(clk).Adjudicate(context.Background(), item, []*domain.Rule{rule})
	if err != nil || ref.LeadDepartment != "骑手之家" {
		t.Fatalf("default rule was not used: ref=%+v err=%v", ref, err)
	}
}
