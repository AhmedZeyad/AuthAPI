package initializer
import 	"github.com/AhmedZeyad/AuthAPI/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}