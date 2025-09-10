package controller

import (
	"go_mail_sender/constant"
	"go_mail_sender/entity"
	"go_mail_sender/repository"
	"go_mail_sender/service"
	"go_mail_sender/usecase"

	"net/http"

	"github.com/labstack/echo"
)

type MailController struct {
	Mysql *repository.MysqlCon
}

func (msc *MailController) SendMail(c echo.Context) error {
	var mailReq entity.Mail
	if err := c.Bind(&mailReq); err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Code:    http.StatusBadRequest,
			Message: constant.INVALID_REQUEST,
			Error:   err.Error(),
		})
	}

	// call sender service
	service := service.MailService{}
	log, err := service.Send(mailReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Code:    http.StatusBadRequest,
			Message: constant.INVALID_REQUEST,
			Error:   err.Error(),
		})
	}

	mailUsecase := usecase.MailUseCase{
		Mysql: msc.Mysql,
	}

	err = mailUsecase.CreateMailLogs(log)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Code:    http.StatusBadRequest,
			Message: constant.MAIL_SEND_FAILED,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, entity.Response{
		Data: entity.MailSendResponse{
			Status:    log.Status,
			Delivered: log.Delivered,
			Error:     log.ErrorText,
		},
	})

}

func (msc *MailController) GetMailLogs(c echo.Context) error {

	mailUsecase := usecase.MailUseCase{
		Mysql: msc.Mysql,
	}

	logs, err := mailUsecase.GetMailLogs()
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Code:    http.StatusBadRequest,
			Message: constant.MAIL_FETCH_FAILED,
			Error:   err.Error(),
		})
	}
	// last 100 days logs

	return c.JSON(http.StatusAccepted, entity.Response{
		Code:    http.StatusAccepted,
		Message: constant.MAIL_FETCH_SUCCESSFULLY,
		Data:    logs,
	})

}
