package providers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	user_models "github.com/tnqbao/gau_user_service/models"
	"gorm.io/gorm"
)

func GenCaptchaCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var codes [6]byte
	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + r.Intn(10))
	}

	return string(codes[:])
}

func UpdateUserBooleanField(fieldName string, newValue bool, c *gin.Context) error {
	tokenId, exists := c.Get("user_id")
	if !exists {
		return fmt.Errorf("user_id not found")
	}

	tokenIdUint, ok := tokenId.(uint)
	if !ok {
		return fmt.Errorf("invalid user_id format")
	}

	db := c.MustGet("db").(*gorm.DB)

	var userInfor user_models.UserInformation

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&userInfor, "user_id = ?", tokenIdUint).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("user not found")
			}
			return fmt.Errorf("error fetching user: %v", err)
		}

		if err := tx.Model(&userInfor).Update(fieldName, newValue).Error; err != nil {
			return fmt.Errorf("error updating field %s: %v", fieldName, err)
		}

		return nil
	})

	return err
}
