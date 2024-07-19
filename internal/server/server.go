package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kahuri1/birthday-notifier/internal/entity"
	"github.com/kahuri1/birthday-notifier/pkg/server"
)

type UserService interface {
	Create(ctx context.Context, user entity.User) error
}

type SubscriptionService interface {
	Find(userID int) ([]entity.Subscription, error)
}

type Server struct {
	UserService         UserService
	SubscriptionService SubscriptionService
}

func NewServer(userService UserService, service SubscriptionService) *Server {
	return &Server{
		UserService:         userService,
		SubscriptionService: service,
	}
}

func (s *Server) UserCreateV1(c *gin.Context) {
	var user server.PostUserCreateV1RequestBody
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, server.BadRequestResponse{
			Status: http.StatusBadRequest,
			Title:  err.Error(),
		})
		return
	}

	err := s.UserService.Create(c.Request.Context(), entity.User{
		Email:     user.Email,
		Name:      user.Name,
		Birthdate: user.Birthdate,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, server.BadRequestResponse{
			Status: http.StatusBadRequest,
			Title:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (s *Server) UserSubscribeV1(c *gin.Context, id int64) {
	var subscribe server.PostUserSubscribeV1RequestBody
	if err := c.ShouldBindJSON(&subscribe); err != nil {
		c.JSON(http.StatusBadRequest, server.BadRequestResponse{
			Status: http.StatusBadRequest,
			Title:  err.Error(),
		})
		return
	}
}

func (s *Server) UserSubscriptionsV1(c *gin.Context, id int64) {
	subscriptions, _ := s.SubscriptionService.Find(int(id))

	responseSubscriptions := make([]server.Subscription, 0, len(subscriptions))
	for _, subscription := range subscriptions {
		responseSubscriptions = append(responseSubscriptions, server.Subscription{
			Id:     subscription.SubscriberID,
			UserId: subscription.UserID,
		})
	}

	response := server.UserSubscriptionsV1Response{
		Subscriptions: responseSubscriptions,
	}

	c.JSON(http.StatusOK, response)
}
