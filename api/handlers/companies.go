package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mospyx/data_accounting/pkg/models"
	"net/http"
	"strconv"
)

//Companies handlers

// GetCompanies /api/companies [get]
func GetCompanies(c *gin.Context) {
	cmpList, err := models.GetCompanyList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, cmpList)
}

// CreateCompany /api/companies [post]
func CreateCompany(c *gin.Context) {
	cmp := models.Company{}

	if err := c.ShouldBindJSON(&cmp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cmp.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, cmp)
}

//Company handlers

// GetCompany /api/companies/{company_id} [get]
func GetCompany(c *gin.Context) {
	cmpIDstr := c.Param("company_id")
	cmpID, err := strconv.Atoi(cmpIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	cmp, err := models.GetCompany(uint(cmpID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, cmp)
}

// DeleteCompany /api/companies/{company_id} [delete]
func DeleteCompany(c *gin.Context) {
	cmpIDstr := c.Param("company_id")
	cmpID, err := strconv.Atoi(cmpIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	cmp, err := models.GetCompany(uint(cmpID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	cp := cmp.CompanyProfile

	if err = cp.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err = cmp.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, cmp)
}

//CompanyProfile handlers

// GetCompanyProfile /api/companies/{company_id}/company_profile [get]
func GetCompanyProfile(c *gin.Context) {
	cmpIDstr := c.Param("company_id")
	cmpID, err := strconv.Atoi(cmpIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	cmp, err := models.GetCompany(uint(cmpID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	cp := cmp.CompanyProfile

	c.JSON(http.StatusOK, cp)
}

// UpdateCompanyProfile /api/companies/{company_id}/company_profile [put]
func UpdateCompanyProfile(c *gin.Context) {
	cmpIDstr := c.Param("company_id")
	cmpID, err := strconv.Atoi(cmpIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	cmp, err := models.GetCompany(uint(cmpID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	cp := cmp.CompanyProfile

	if err = c.ShouldBindJSON(&cp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = cp.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, cp)
}
