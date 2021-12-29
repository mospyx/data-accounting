package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mospyx/data_accounting/pkg/models"
	"net/http"
)

//Companies handlers

func GetCompanies(c *gin.Context) {

}

// CreateCompany /api/companies [post]
func CreateCompany(c *gin.Context) {
	cmp := models.Company{}

	if err := c.ShouldBindJSON(&cmp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cmp.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, cmp)
}

//Company handlers

func GetCompany(c *gin.Context) {

}

func UpdateCompany(c *gin.Context) {

}

func DeleteCompany(c *gin.Context) {

}

//CompanyProfiles handlers

func GetCompanyProfiles(c *gin.Context) {

}

//CompanyProfile handlers

func GetCompanyProfile(c *gin.Context) {

}

func UpdateCompanyProfile(c *gin.Context) {

}

func DeleteCompanyProfile(c *gin.Context) {

}
