package mapper

import (
	"time"

	"aurora.com/aurora-backend/internal/features/user/domain"
	"aurora.com/aurora-backend/internal/features/user/dto"
)

func mapFinancialProfileFromDTO(req *dto.CreateUserRequest) *domain.FinancialProfile {
	if req.FinancialProfile == nil {
		return nil
	}

	now := time.Now()
	expiresAt := now.AddDate(1, 0, 0)

	return &domain.FinancialProfile{
		BillingAddress: domain.BillingAddress{
			Street:  req.FinancialProfile.BillingAddress.Street,
			City:    req.FinancialProfile.BillingAddress.City,
			State:   req.FinancialProfile.BillingAddress.State,
			ZipCode: req.FinancialProfile.BillingAddress.ZipCode,
			Country: "BR",
		},
		Subscription: domain.Subscription{
			PlanID:    req.FinancialProfile.SubscriptionPlan.PlanID,
			StartDate:  now,
			EndDate: expiresAt,
			Active:    true,
		},
		PaymentMethods: "pending_setup",
	}
}