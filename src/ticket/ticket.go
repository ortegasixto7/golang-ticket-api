package ticket

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID        uuid.UUID
	Title     string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	ExpiresAt *time.Time
}

func (tkt *Ticket) GenerateCode() {
	data := fmt.Sprintf("%d", time.Now().UnixNano())
	hash := md5.Sum([]byte(data))
	tkt.Code = strings.ToUpper(hex.EncodeToString(hash[:]))
}
