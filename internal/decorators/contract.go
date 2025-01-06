package decorators

import (
	"fmt"
	"time"

	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
)

func ToContract(c easypay.Contract) {
  fmt.Println("ID:", c.Id, "| S:", c.Status, "| V:", c.Version)
	fmt.Println("- Investment   :", ToMoney(c.Investment))
	fmt.Println("- Down payment :", ToMoney(c.DownPayment))
	fmt.Println("- Duration     :", ToDuration(c.DurationMonths))
	fmt.Println("- Monthly      :", ToMoney(c.MonthlyInstallment))
	fmt.Println("- Interest     :", toInterestRate(c.InterestRate))
  fmt.Println("- Access token :", c.AccessToken)
  fmt.Println("- Expires at   :", ToDateWithAge(&c.AccessTokenExpiresAt))
	fmt.Println("- Reviewed     :", toReviewed(c.ReviewedAt, c.ReviewedBy))
	fmt.Println("- Created at   :", ToDateWithAge(c.CreatedAt))
	fmt.Println("- Updated at   :", ToDateWithAge(c.UpdatedAt))
	fmt.Println()
}

func toInterestRate(i easypay.InterestRate) string {
	return fmt.Sprintf("%.02f %%", i.Metadata.Percentage)
}


func toReviewed(reviewedAt *time.Time, reviewedBy string) string {
	if reviewedAt == nil {
		return "---"
	}
  
  return ToDateWithAge(reviewedAt) + " ~ by " + reviewedBy;
}
