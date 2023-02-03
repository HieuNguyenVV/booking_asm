package handler

import (
	"booking_asm/api/customer_api/request"
	"booking_asm/api/customer_api/response"
	"booking_asm/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type (
	CustomerHandler interface {
		FindCustomer(c *gin.Context)
		CreateCustomer(c *gin.Context)
		UpdateCustomer(c *gin.Context)
		ChangePassword(c *gin.Context)
	}
	customerHandler struct {
		customerClient pb.CustomerServiceClient
	}
)

func New(customerClient pb.CustomerServiceClient) CustomerHandler {
	return &customerHandler{customerClient: customerClient}
}

func (t *customerHandler) FindCustomer(c *gin.Context) {
	req := request.FindCustomerRequest{}
	req.ID = c.Param("id")
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	if validateErr, ok := err.(validator.ValidationErrors); ok {
	// 		errMessage := make([]string, 0)
	// 		for _, v := range validateErr {
	// 			errMessage = append(errMessage, v.Error())
	// 		}
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 			"status": http.StatusText(http.StatusBadRequest),
	// 			"error":  errMessage,
	// 		})
	// 		return
	// 	}
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"status": http.StatusText(http.StatusBadRequest),
	// 		"error":  err.Error(),
	// 	})
	// 	return
	// }
	resp, err := t.customerClient.FindCustomer(context.Background(), &pb.FindCustomerRequest{
		Id: req.ID,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	dto := response.FindCustomerResponse{
		Name:       resp.CustomerName,
		Address:    resp.Address,
		Phone:      resp.Phone,
		License_id: resp.LicenseId,
		Active:     resp.Active,
		ID:         resp.Id,
		Email:      resp.Email,
		Password:   resp.Password,
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status": http.StatusText(http.StatusOK),
		"value":  dto,
	})
}

func (t *customerHandler) CreateCustomer(c *gin.Context) {
	req := request.CreateCustomerRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		if validateErr, ok := err.(validator.ValidationErrors); ok {
			errMessage := make([]string, 0)
			for _, v := range validateErr {
				errMessage = append(errMessage, v.Error())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessage,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	id := uuid.New()
	resp, err := t.customerClient.CreateCustomer(context.Background(), &pb.CreateCustomerRequest{
		CustomerName: req.Name,
		Address:      req.Phone,
		Phone:        req.Phone,
		LicenseId:    req.License_id,
		Active:       req.Active,
		Id:           id.String(),
		Email:        req.Email,
		Password:     req.Password,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	dto := response.CreateCustomerResponse{
		Name:       resp.CustomerName,
		Address:    resp.Address,
		Phone:      resp.Phone,
		License_id: resp.LicenseId,
		Active:     resp.Active,
		ID:         resp.Id,
		Email:      resp.Email,
		Password:   resp.Password,
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status": http.StatusText(http.StatusOK),
		"value":  dto,
	})
}
func (t *customerHandler) UpdateCustomer(c *gin.Context) {
	req := request.UpdateCustomerRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		if validatorErr, ok := err.(validator.ValidationErrors); ok {
			messageErr := make([]string, 0)
			for _, v := range validatorErr {
				messageErr = append(messageErr, v.Error())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  messageErr,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	resp, err := t.customerClient.UpdateCustomer(context.Background(), &pb.UpdateCustomerRequest{
		CustomerName: req.Name,
		Address:      req.Phone,
		Phone:        req.Phone,
		LicenseId:    req.License_id,
		Active:       req.Active,
		Id:           req.ID,
		Email:        req.Email,
		Password:     req.Password,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	dto := response.UpdateCustomerResponse{
		Name:       resp.CustomerName,
		Address:    resp.Address,
		Phone:      resp.Phone,
		License_id: resp.LicenseId,
		Active:     resp.Active,
		ID:         resp.Id,
		Email:      resp.Email,
		Password:   resp.Password,
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status": http.StatusText(http.StatusOK),
		"value":  dto,
	})
}
func (t *customerHandler) ChangePassword(c *gin.Context) {
	req := request.ChangePasswordRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		if validatorErr, ok := err.(validator.ValidationErrors); ok {
			messageErr := make([]string, 0)
			for _, v := range validatorErr {
				messageErr = append(messageErr, v.Error())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  messageErr,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	resp, err := t.customerClient.ChangePassword(context.Background(), &pb.ChangePasswordRequest{
		Id:       req.ID,
		Password: req.Password,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}
	dto := response.ChangePasswordResponse{
		Name:       resp.CustomerName,
		Address:    resp.Address,
		Phone:      resp.Phone,
		License_id: resp.LicenseId,
		Active:     resp.Active,
		ID:         resp.Id,
		Email:      resp.Email,
		Password:   resp.Password,
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status": http.StatusText(http.StatusOK),
		"value":  dto,
	})
}
