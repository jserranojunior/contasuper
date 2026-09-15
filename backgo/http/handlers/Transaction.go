package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jserranojunior/sefro/backgo/models"
)

// Struct para validar os dados recebidos na requisição
type DepositInput struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

// CreateDepositHandler lida com a solicitação de depósito feita pelo cliente
func CreateDepositHandler(c *gin.Context) {
	// 1. Recupera o ID do usuário autenticado pelo middleware JWT
	userIDVal, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Usuário não autenticado",
		})
		return
	}
	userID := userIDVal.(uint)

	// 2. Faz o bind e valida o JSON enviado pelo app
	var input DepositInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Valor inválido ou ausente. O valor deve ser maior que zero.",
		})
		return
	}

	// 3. Monta e salva o registro da transação no banco de dados
	transaction := models.Transaction{
		UserId: userID,
		Type:   "deposito",
		Amount: input.Amount,
		Status: "pendente", // Fica pendente até o caixa confirmar
	}

	if err := DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Não foi possível registrar a solicitação de depósito no banco de dados",
		})
		return
	}

	// 4. Retorna a resposta de sucesso para o aplicativo
	c.JSON(http.StatusCreated, gin.H{
		"message": "Solicitação de depósito enviada com sucesso! Dirija-se ao caixa para entregar o valor.",
		"user_id": userID,
		"amount":  input.Amount,
		"status":  "pendente",
	})
}

func GetUserTransactionsHandler(c *gin.Context) {
	// 1. Recupera o ID do usuário autenticado pelo middleware JWT
	userIDVal, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Usuário não autenticado",
		})
		return
	}
	userID := userIDVal.(uint)

	// 2. Busca as transações do usuário no banco de dados ordenadas da mais recente para a mais antiga
	var transactions []models.Transaction
	if err := DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Não foi possível buscar as transações no banco de dados",
		})
		return
	}

	// 3. Retorna a lista de transações com sucesso para o aplicativo/frontend
	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}