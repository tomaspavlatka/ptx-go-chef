package decorators

import (
	"fmt"
	"github.com/tomaspavlatka/ptx-go-chef/handlers/easypay"
)

func ToKins(audits []easypay.KinAudit) {
	fmt.Println(headerStyle.Render("KINS"))

	for _, audit := range audits {
		fmt.Println(headerStyle.Render(translateType(audit.AuditType) + " at " + ToDateWithAgeDetailed(audit.CreatedAt)))
		fmt.Println("- Txid         :", audit.Txid)
		fmt.Println("- Gate         :", audit.Gate)
		fmt.Println("- Seat         :", audit.Seat)
		fmt.Println("- Created by   :", audit.CreatedBy)
		fmt.Println("-------------- :")
		fmt.Println("- Type         :", audit.Type)
		fmt.Println("- Value        :", audit.Value)
		fmt.Println()
	}
}
