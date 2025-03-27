package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Ticket struct {
	ID        string
	Title     string
	Code      string
	CreatedAt int64
	UpdatedAt int64
	ExpiredAt int64
}

func (tkt *Ticket) GenerateCode() {
	data := fmt.Sprintf("%d", time.Now().UnixNano())
	hash := md5.Sum([]byte(data))
	tkt.Code = strings.ToUpper(hex.EncodeToString(hash[:]))
}
