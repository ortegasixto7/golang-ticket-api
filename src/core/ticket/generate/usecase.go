package generate

import (
	"fmt"

	"github.com/ortegasixto7/golang-ticket/src/models"
)

func Execute(req *GenerateRequest) error {
	tkt := models.Ticket{}
	tkt.GenerateCode()
	fmt.Println(tkt)
	return nil
}
