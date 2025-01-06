package decorators

import (
	"fmt"

	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
)

func ToContract(c easypay.Contract) {
	fmt.Println("ID:", c.Id)
	fmt.Println("- Investment   :", ToMoney(c.Investment))
	fmt.Println("- Down payment :", ToMoney(c.DownPayment))
	fmt.Println("- Duration     :", ToDuration(c.DurationMonths))
	fmt.Println("- Monthly      :", ToMoney(c.MonthlyInstallment))
	fmt.Println("- Interest     :", toInterestRate(c.InterestRate))
	fmt.Println("- Version      :", c.Version)
	fmt.Println("- Status       :", c.Status)
	fmt.Println("- Created      :", ToDateWithAge(c.CreatedAt))
	fmt.Println("- Updated      :", ToDateWithAge(c.UpdatedAt))
	fmt.Println()
}

func toInterestRate(i easypay.InterestRate) string {
	return fmt.Sprintf("%.02f %%", i.Metadata.Percentage)
}
