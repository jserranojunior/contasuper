package models

import (
    "gorm.io/gorm"
)

// Transaction struct export
type Transaction struct {
    gorm.Model
    UserId        uint        `gorm:"not null;" json:"user_id"`
    UserRef       User        `gorm:"foreignKey:UserId"`
    Type          string      `gorm:"size:50;not null;" json:"type"`   // Ex: 'deposito', 'saque', 'rendimento'
    Amount        float64     `gorm:"type:decimal(18,2);not null;" json:"amount"` // Valor da transação
    Status        string      `gorm:"size:30;default:'pendente';" json:"status"` // 'pendente', 'aprovado', 'rejeitado'
    AdminId       *uint       `json:"admin_id"` // Quem aprovou (opcional, nulo se estiver pendente)
    AdminRef      User        `gorm:"foreignKey:AdminId"`
}