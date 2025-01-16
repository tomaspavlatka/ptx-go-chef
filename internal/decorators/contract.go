package decorators

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
)

func ToContract(c easypay.Contract) {
  fmt.Println(headerStyle.Render("CONTRACT"))
	fmt.Println(headerStyle.Render("ID:" + c.Id + ", S:" + c.Status + ", V:" + strconv.Itoa(c.Version)))
	fmt.Println("- Name         :", c.Name)
	fmt.Println("- Investment   :", ToMoney(c.Investment))
	fmt.Println("- Down payment :", ToMoney(c.DownPayment))
	fmt.Println("- Duration     :", ToDuration(c.DurationMonths))
	fmt.Println("- Monthly      :", ToMoney(c.MonthlyInstallment))
	fmt.Println("- Interest     :", toInterestRate(c.NominalInterestRate))
	fmt.Println("- Total Credit :", ToMoney(c.TotalCreditAmount))
	fmt.Println("- Access token :", c.AccessToken)
	fmt.Println("- Expires at   :", ToDateWithAge(&c.AccessTokenExpiresAt))
	fmt.Println("- Reviewed     :", toReviewed(c.ReviewedAt, c.ReviewedBy))
	fmt.Println("- Created at   :", ToDateWithAge(c.CreatedAt))
	fmt.Println("- Updated at   :", ToDateWithAge(c.UpdatedAt))
}

func toInterestRate(i easypay.InterestRate) string {
	return fmt.Sprintf("%.02f %%", i.Metadata.Percentage)
}

func toReviewed(reviewedAt *time.Time, reviewedBy string) string {
	if reviewedAt == nil {
		return "---"
	}

	return ToDateWithAge(reviewedAt) + " ~ by " + reviewedBy
}

func ToContractAudits(audits []easypay.ContractAudit) {
	fmt.Println(headerStyle.Render("AUDITS"))

	var (
		status             string
		interestRate       int
		investment         int
		durationMonths     int
		monthlyInstallment int
		accessToken        string
		reviewedBy         string
		reviewedAt         *time.Time
		companyId          string
		name               string
		currency           string
		applicantId        string
		downPayment        int = -1
	)

	for _, audit := range audits {
		fmt.Println(headerStyle.Render(translateType(audit.AuditType) + " at " + ToDateWithAgeDetailed(audit.CreatedAt)))
		fmt.Println("- Txid         :", audit.Txid)
		fmt.Println("- Gate         :", audit.Gate)
		fmt.Println("- Seat         :", audit.Seat)
		fmt.Println("- Created by   :", audit.CreatedBy)
		fmt.Println("-------------- :")

		if newName, changed := gotChanged(name, &audit.Name); changed {
			name = newName
			fmt.Println("- Name         :", name)
		}

		if newCompanyId, changed := gotChanged(companyId, &audit.CompanyId); changed {
			companyId = newCompanyId
			fmt.Println("- Company Id   :", companyId)
		}

		if newApplicantId, changed := gotChanged(applicantId, audit.ApplicantId); changed {
			applicantId = newApplicantId
			fmt.Println("- Applicant Id   :", applicantId)
		}

		if newStatus, changed := gotChanged(status, &audit.Status); changed {
			status = newStatus
			fmt.Println("- Status       :", status)
		}

		if newCurrency, changed := gotChanged(currency, &audit.Currency); changed {
			currency = newCurrency
			fmt.Println("- Currency     :", currency)
		}

		if newInvestment, changed := gotChanged(investment, &audit.Investment); changed {
			investment = newInvestment
			fmt.Println("- Investment   :", ToMoneyFromCentAmount(investment, audit.Currency))
		}

		if newDownPayment, changed := gotChanged(downPayment, &audit.DownPayment); changed {
			downPayment = newDownPayment
			fmt.Println("- Down payment :", ToMoneyFromCentAmount(downPayment, audit.Currency))
		}

		if newDurationMonths, changed := gotChanged(durationMonths, &audit.DurationMonths); changed {
			durationMonths = newDurationMonths
			fmt.Println("- Duration     :", ToDuration(durationMonths))
		}

		if newMonthlyInstallment, changed := gotChanged(monthlyInstallment, &audit.MonthlyInstallment); changed {
			monthlyInstallment = newMonthlyInstallment
			fmt.Println("- Monthly      :", ToMoneyFromCentAmount(monthlyInstallment, audit.Currency))
		}

		if newInterestRate, changed := gotChanged(interestRate, audit.InterestRate); changed {
			interestRate = newInterestRate
			fmt.Println("- Interest     :", toInterestRate(
				easypay.InterestRate{
					Metadata: easypay.InterestRateMeta{Percentage: float64(interestRate) / 10000.0},
				},
			))
		}

		if newAccessToken, changed := gotChanged(accessToken, &audit.AccessToken); changed {
			accessToken = newAccessToken
			fmt.Println("- Access Token :", accessToken)
		}

		if newReviewedBy, changed := gotChanged(reviewedBy, &audit.ReviewedBy); changed {
			reviewedBy = newReviewedBy
			fmt.Println("- Reviewed by  :", reviewedBy)
		}

		if newReviewedAt, changed := gotChanged(reviewedAt, &audit.ReviewedAt); changed {
			reviewedAt = newReviewedAt
			fmt.Println("- Reviewed at  :", ToDateWithAgeDetailed(reviewedAt))
		}
		fmt.Println()
	}
}

