package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/orcashtrator/internal/appconstant"
	"github.com/itsLeonB/orcashtrator/internal/config"
	"github.com/itsLeonB/orcashtrator/internal/delivery/http/handler"
	"github.com/itsLeonB/orcashtrator/internal/provider"
)

func registerRoutes(router *gin.Engine, configs config.Config, logger ezutil.Logger, services *provider.Services) {
	handlers := handler.ProvideHandlers(logger, services)
	middlewares := provideMiddlewares(configs.App, logger, services.Auth)

	router.Use(middlewares.logger, middlewares.cors)

	apiRoutes := router.Group("/api", middlewares.err)

	v1 := apiRoutes.Group("/v1")

	authRoutes := v1.Group("/auth")
	authRoutes.POST("/register", handlers.Auth.HandleRegister())
	authRoutes.POST("/login", handlers.Auth.HandleLogin())

	protectedRoutes := v1.Group("/", middlewares.auth)

	profileRoutes := protectedRoutes.Group("/profile")
	profileRoutes.GET("", handlers.Profile.HandleProfile())
	profileRoutes.PATCH("", handlers.Profile.HandleUpdate())

	friendshipRoutes := protectedRoutes.Group("/friendships")
	friendshipRoutes.POST("", handlers.Friendship.HandleCreateAnonymousFriendship())
	friendshipRoutes.GET("", handlers.Friendship.HandleGetAll())
	friendshipRoutes.GET(fmt.Sprintf("/:%s", appconstant.ContextFriendshipID), handlers.Friendship.HandleGetDetails())

	protectedRoutes.GET("/transfer-methods", handlers.TransferMethod.HandleGetAll())

	protectedRoutes.POST("/debts", handlers.Debt.HandleCreate())
	protectedRoutes.GET("/debts", handlers.Debt.HandleGetAll())

	groupExpenseRoutes := protectedRoutes.Group("/group-expenses")
	groupExpenseRoutes.POST("", handlers.GroupExpense.HandleCreateDraft())
	groupExpenseRoutes.GET("", handlers.GroupExpense.HandleGetAllCreated())
	groupExpenseRoutes.GET(fmt.Sprintf("/:%s", appconstant.ContextGroupExpenseID), handlers.GroupExpense.HandleGetDetails())
	groupExpenseRoutes.PATCH(fmt.Sprintf("/:%s/confirmed", appconstant.ContextGroupExpenseID), handlers.GroupExpense.HandleConfirmDraft())

	expenseItemRoutes := groupExpenseRoutes.Group(fmt.Sprintf("/:%s/items", appconstant.ContextGroupExpenseID))
	expenseItemRoutes.POST("", handlers.ExpenseItem.HandleAdd())
	expenseItemRoutes.GET(fmt.Sprintf("/:%s", appconstant.ContextExpenseItemID), handlers.ExpenseItem.HandleGetDetails())
	expenseItemRoutes.PUT(fmt.Sprintf("/:%s", appconstant.ContextExpenseItemID), handlers.ExpenseItem.HandleUpdate())
	expenseItemRoutes.DELETE(fmt.Sprintf("/:%s", appconstant.ContextExpenseItemID), handlers.ExpenseItem.HandleRemove())

	otherFeeRoutes := groupExpenseRoutes.Group(fmt.Sprintf("/:%s/fees", appconstant.ContextGroupExpenseID))
	otherFeeRoutes.POST("", handlers.OtherFee.HandleAdd())
	otherFeeRoutes.PUT(fmt.Sprintf("/:%s", appconstant.ContextOtherFeeID), handlers.OtherFee.HandleUpdate())
	otherFeeRoutes.DELETE(fmt.Sprintf("/:%s", appconstant.ContextOtherFeeID), handlers.OtherFee.HandleRemove())

	groupExpenseRoutes.GET("/fee-calculation-methods", handlers.OtherFee.HandleGetFeeCalculationMethods())

	expenseBillRoutes := groupExpenseRoutes.Group("/bills")
	expenseBillRoutes.POST("", handlers.ExpenseBill.HandleSave())
	expenseBillRoutes.GET("", handlers.ExpenseBill.HandleGetAllCreated())
	expenseBillRoutes.GET(fmt.Sprintf("/:%s", appconstant.ContextExpenseBillID.String()), handlers.ExpenseBill.HandleGet())
	expenseBillRoutes.DELETE(fmt.Sprintf("/:%s", appconstant.ContextExpenseBillID.String()), handlers.ExpenseBill.HandleDelete())
}
