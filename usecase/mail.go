package usecase

import (
	"fmt"
	"go_mail_sender/constant"
	"go_mail_sender/entity"
	"go_mail_sender/repository"
)

type MailUseCase struct {
	Mysql *repository.MysqlCon
}

func (msuc *MailUseCase) CreateMailLogs(log *entity.EmailLog) error {
	if msuc.Mysql.Connection == nil {
		fmt.Println("Database connection failed")
		fmt.Errorf(constant.DB_CONNECTION_FAILED)
	}

	err := msuc.Mysql.Connection.Table(constant.MAIL_LOG_TABLE).Create(log).Error
	if err != nil {
		return fmt.Errorf("user creation failed %v", err)
	}
	return nil
}
func (msuc *MailUseCase) GetMailLogs() ([]entity.EmailLogResponse, error) {
	if msuc.Mysql.Connection == nil {
		return nil, fmt.Errorf(constant.DB_CONNECTION_FAILED)
	}

	var mailLogs []entity.EmailLogResponse

	// Last 100 days logs
	err := msuc.Mysql.Connection.
		Table(constant.MAIL_LOG_TABLE).
		Where("created_at >= NOW() - INTERVAL 100 DAY").
		Find(&mailLogs).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch logs: %v", err)
	}

	return mailLogs, nil
}
