package handlers

import (
	"fmt"
	"mampu-go-api/models"
	"mampu-go-api/repositories"

	"github.com/gofiber/fiber/v2"
)

type WalletHandler struct {
	userRepo  repositories.UserRepository
	transRepo repositories.TransactionRepository
}

func NewWalletHandler(userRepo repositories.UserRepository, transRepo repositories.TransactionRepository) *WalletHandler {
	return &WalletHandler{
		userRepo:  userRepo,
		transRepo: transRepo,
	}
}

func (h *WalletHandler) GetBalance(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": "ID pengguna tidak valid"})
	}

	user, err := h.userRepo.GetByID(uint(userId))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"success": false, "message": "Pengguna tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"success": true, "balance": user.Balance})
}

func (h *WalletHandler) Withdraw(c *fiber.Ctx) error {
	type WithdrawRequest struct {
		UserID uint    `json:"user_id"`
		Amount float64 `json:"amount"`
	}

	var req WithdrawRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": "Permintaan tidak valid"})
	}

	user, err := h.userRepo.GetByID(req.UserID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"success": false, "message": "Pengguna tidak ditemukan"})
	}

	if user.Balance < req.Amount {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": "Saldo tidak cukup. Sisa saldo Anda " + fmt.Sprintf("%.2f", user.Balance)})
	}

	user.Balance -= req.Amount
	err = h.userRepo.Update(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "message": "Gagal memperbarui saldo"})
	}

	transaction := &models.Transaction{
		UserID: req.UserID,
		Amount: req.Amount,
		Type:   "withdraw",
	}
	err = h.transRepo.Create(transaction)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "message": "Gagal mencatat transaksi"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Penarikan berhasil. Sisa saldo Anda " + fmt.Sprintf("%.2f", user.Balance), "new_balance": user.Balance})
}
