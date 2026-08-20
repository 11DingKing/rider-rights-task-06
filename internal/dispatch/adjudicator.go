package dispatch

import (
	"context"
	"fmt"
	"sort"

	"github.com/google/uuid"

	"riderguard/internal/domain"
)

type Adjudicator struct {
	clock domain.Clock
}

func NewAdjudicator(clock domain.Clock) *Adjudicator {
	return &Adjudicator{clock: clock}
}

func (a *Adjudicator) Adjudicate(ctx context.Context, item *domain.RightsCase, rules []*domain.Rule) (*domain.Referral, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// 专项规则优先匹配；默认规则不参与竞争，仅收集当前生效的默认规则用于兜底，
	// 避免默认规则在专项规则命中时抢先生效。
	var matched []*domain.Rule
	var defaults []*domain.Rule
	for _, rule := range rules {
		if !rule.IsActiveAt(item.RegisteredAt) {
			continue
		}
		if rule.IsDefault {
			defaults = append(defaults, rule)
			continue
		}
		if rule.Matches(item) {
			matched = append(matched, rule)
		}
	}

	// 仅在无专项规则命中时启用默认规则兜底：诉求既未命中类别也未命中关键词时，
	// 由当前生效的默认规则接收。
	if len(matched) == 0 {
		matched = defaults
	}

	if len(matched) == 0 {
		return nil, fmt.Errorf("no matching dispatch rule for item %s: %w", item.ID, domain.ErrNoMatchingRule)
	}

	sort.Slice(matched, func(i, j int) bool {
		if matched[i].Priority != matched[j].Priority {
			return matched[i].Priority > matched[j].Priority
		}
		return matched[i].Version > matched[j].Version
	})

	chosen := matched[0]
	now := a.clock.Now()

	return &domain.Referral{
		ID:             uuid.NewString(),
		ItemID:         item.ID,
		LeadDepartment: chosen.LeadDepartment,
		CoDepartments:  chosen.CoDepartments,
		RuleVersion:    chosen.Version,
		AdjudicatedAt:  now,
		AdjudicatedBy:  "system",
		IsCurrent:      true,
		DataVersion:    domain.DataVersion,
	}, nil
}

func (a *Adjudicator) SelectEscalationLead(currentLead string, level int) string {
	return domain.EscalationDepartmentName(currentLead, level)
}
