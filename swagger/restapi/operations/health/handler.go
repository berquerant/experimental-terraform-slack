package health

import (
	"github.com/berquerant/terraform-slack/swagger/models"
	"github.com/go-openapi/runtime/middleware"
)

func HandleCheckAlive(_ CheckAliveParams, _ *models.Principal) middleware.Responder {
	payload := &models.HealthModel{
		BaseModel: models.NewBaseModel(true),
	}
	return NewCheckAliveOK().WithPayload(payload)
}
