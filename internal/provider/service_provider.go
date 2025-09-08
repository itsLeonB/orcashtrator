package provider

import (
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/service"
)

type Services struct {
	Auth           service.AuthService
	Profile        service.ProfileService
	Friendship     service.FriendshipService
	TransferMethod service.TransferMethodService
	Debt           service.DebtService
	FriendDetails  service.FriendDetailsService
	GroupExpense   service.GroupExpenseService
	ExpenseItem    service.ExpenseItemService
	OtherFee       service.OtherFeeService
	ExpenseBill    service.ExpenseBillService
}

func ProvideServices(clients *Clients, logger ezutil.Logger) *Services {
	if clients == nil {
		panic("clients cannot be nil")
	}
	if logger == nil {
		panic("logger cannot be nil")
	}

	authService := service.NewAuthService(clients.Auth)

	profileService := service.NewProfileService(clients.Profile)

	friendshipService := service.NewFriendshipService(
		clients.Friendship,
	)

	transferMethodService := service.NewTransferMethodService(clients.TransferMethod)

	debtService := service.NewDebtService(
		clients.Debt,
		friendshipService,
	)

	friendDetailsService := service.NewFriendDetailsService(clients.Friendship, debtService)

	groupExpenseService := service.NewGroupExpenseService(
		friendshipService,
		debtService,
		profileService,
		clients.GroupExpense,
	)

	expenseItemSvc := service.NewExpenseItemService(
		profileService,
		clients.ExpenseItem,
	)

	otherFeeSvc := service.NewOtherFeeService(
		profileService,
		clients.OtherFee,
	)

	expenseBillService := service.NewExpenseBillService(
		logger,
		friendshipService,
		clients.ExpenseBill,
	)

	return &Services{
		authService,
		profileService,
		friendshipService,
		transferMethodService,
		debtService,
		friendDetailsService,
		groupExpenseService,
		expenseItemSvc,
		otherFeeSvc,
		expenseBillService,
	}
}
