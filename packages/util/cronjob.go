package util

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"xyz/modules/installment/repository"
	userRepository "xyz/modules/user/repository"
)

type CronService struct {
	InstallmentRepo repository.InstallmentQueryRepositoryInterface
	UserRepo     userRepository.UserQueryRepositoryInterface
}

func (c *CronService) Start() {
	scheduler := cron.New()

	// Run every day at 8 AM
	scheduler.AddFunc("0 9,12,15 * * *", func() {
		c.checkInstallmentsDueSoon()
	})

	scheduler.Start()
	fmt.Println("Cronjob started...")
}


func (c *CronService) checkInstallmentsDueSoon() {
	installments, err := c.InstallmentRepo.GetInstallmentsNearDueDateWithoutIds()
	if err != nil {
		fmt.Println("Error fetching installments:", err)
		return
	}

	if len(installments) == 0 {
		fmt.Println("No installments due soon.")
		return
	}

	for _, inst := range installments {
		id := inst.Transaction.Consumer.User_ID
		dueDate := inst.DueDate.Format("2006-01-02 15:04:05")
	    userEmail, _ := c.UserRepo.GetUserByID(id)
		email := userEmail.Email
		fmt.Printf("Reminder: %v has an installment due on %s\n", email, dueDate)
		SendEmailNotificationForReminderPayment(email, inst.DueDate, inst.PaymentAmount, inst.Transaction.AssetName, inst.Transaction.ContractNumber, inst.Transaction.TransactionDate)
	}
}
