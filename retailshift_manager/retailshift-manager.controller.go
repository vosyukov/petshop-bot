package retailshift_manager

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"path"
)

type AutoGenerated struct {
	AuditContext struct {
		Meta struct {
			Type string `json:"type"`
			Href string `json:"href"`
		} `json:"meta"`
		UID    string `json:"uid"`
		Moment string `json:"moment"`
	} `json:"auditContext"`
	Events []struct {
		Meta struct {
			Type string `json:"type"`
			Href string `json:"href"`
		} `json:"meta"`
		Action    string `json:"action"`
		AccountID string `json:"accountId"`
	} `json:"events"`
}

func OnHandleCreateRetailshiftController(c *gin.Context) {
	var obj AutoGenerated
	if errA := c.ShouldBind(&obj); errA == nil {
		u, _ := url.Parse(obj.Events[0].Meta.Href)
		var retailshiftId = path.Base(u.Path)

		OnHandleCreateRetailshift(retailshiftId)

		c.Status(http.StatusOK)
	}

}

func OnHandleUpdateRetailshiftController(c *gin.Context) {
	var obj AutoGenerated
	if errA := c.ShouldBind(&obj); errA == nil {

		u, _ := url.Parse(obj.Events[0].Meta.Href)
		var retailshiftId = path.Base(u.Path)

		OnHandleUpdateRetailshift(retailshiftId)

		c.Status(http.StatusOK)
	}
}
