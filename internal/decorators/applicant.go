package decorators

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
)

func ToApplicant(a easypay.Applicant) {
	fmt.Println(headerStyle.Render("APPLICANT"))
	fmt.Println(headerStyle.Render("ID:" + a.Id + ", S:" + a.Status + ", V:" + strconv.Itoa(a.Version)))
	fmt.Println("- First Name   :", a.FirstName)
	fmt.Println("- Last Name    :", a.LastName)
	fmt.Println("- Phone        :", a.Phone)
	fmt.Println("- E-mail       :", a.Email)
	fmt.Println("- Date birth   :", ToDateWithYearAndAge(a.DateOfBirth))
	fmt.Println("- Created at   :", ToDateWithAge(a.CreatedAt))
	fmt.Println("- Updated at   :", ToDateWithAge(a.UpdatedAt))
}

func ToApplicantAudits(audits []easypay.ApplicantAudit) {
	fmt.Println(headerStyle.Render("AUDITS"))

	var (
		status      string
		applicantId string
		companyId   string
		firstName   string
		lastName    string
		email       string
		phone       string
		dateOfBirth *time.Time
	)

	for _, audit := range audits {
		fmt.Println(headerStyle.Render(translateType(audit.AuditType) + " at " + ToDateWithAgeDetailed(audit.CreatedAt)))
		fmt.Println("- Txid         :", audit.Txid)
		fmt.Println("- Gate         :", audit.Gate)
		fmt.Println("- Seat         :", audit.Seat)
		fmt.Println("- Created by   :", audit.CreatedBy)
		fmt.Println("-------------- :")

		if newApplicantId, changed := gotChanged(applicantId, &audit.ApplicantId); changed {
			applicantId = newApplicantId
      fmt.Println("- Applicant Id :", applicantId)
		}

		if newCompanyId, changed := gotChanged(companyId, &audit.CompanyId); changed {
			companyId = newCompanyId
			fmt.Println("- Company Id   :", companyId)
		}

		if newFirstName, changed := gotChanged(firstName, &audit.FirstName); changed {
			firstName = newFirstName
			fmt.Println("- First Name   :", firstName)
		}

		if newLastName, changed := gotChanged(lastName, &audit.LastName); changed {
			lastName = newLastName
			fmt.Println("- Last Name    :", lastName)
		}

		if newEmail, changed := gotChanged(email, &audit.Email); changed {
			email = newEmail
			fmt.Println("- Email        :", email)
		}

		if newPhone, changed := gotChanged(phone, &audit.Phone); changed {
			phone = newPhone
			fmt.Println("- Phone        :", phone)
		}

		if newStatus, changed := gotChanged(status, &audit.Status); changed {
			status = newStatus
			fmt.Println("- Status       :", status)
		}

		if newDateOfBirth, changed := gotChanged(dateOfBirth, &audit.DateOfBirth); changed {
			dateOfBirth = newDateOfBirth
			fmt.Println("- Date birth   :", ToDateWithAgeDetailed(dateOfBirth))
		}
		fmt.Println()
	}
}
