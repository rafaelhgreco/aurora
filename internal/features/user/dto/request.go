package dto

type CreateUserRequest struct {
	Name             string   `json:"name" binding:"required"`
	Email            string   `json:"email" binding:"required,email"`
	Password         string   `json:"password" binding:"required,min=8"`
	Type             string   `json:"type" binding:"required,oneof=COMMON COLLABORATOR ADMIN FINANCIAL"`
	Permissions      []string `json:"permissions"`
	TeamID           string   `json:"teamId"`
	FinancialProfile *struct {
		BillingAddress *struct {
			Street  string `json:"street"`
			City    string `json:"city"`
			State   string `json:"state"`
			ZipCode string `json:"zipCode"`
			Country string `json:"country"`
		} `json:"billingAddress"`
		SubscriptionPlan *struct {
			PlanID string `json:"planId"`
		} `json:"subscriptionPlan"`
	} `json:"financialProfile"`
}

// UpdateUserRequest permanece o mesmo
type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty" binding:"omitempty,email"`
}

// DeleteUserRequest permanece o mesmo
type DeleteUserRequest struct {
	ID string `json:"id" binding:"required"`
}