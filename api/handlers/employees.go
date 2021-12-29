package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mospyx/data_accounting/pkg/models"
	"net/http"
	"strconv"
)

//Employees handlers

// GetEmployees /api/companies/{company_id}/employees [get]
func GetEmployees(c *gin.Context) {
	cmpIDstr := c.Param("company_id")
	cmpID, err := strconv.Atoi(cmpIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	eList, err := models.GetEmployeeListByCompanyID(uint(cmpID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, eList)
}

// CreateEmployee /api/companies/{company_id}/employees [post]
func CreateEmployee(c *gin.Context) {
	cmpIDstr := c.Param("company_id")
	cmpID, err := strconv.Atoi(cmpIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	e := models.Employee{}

	if err = c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = e.Create(uint(cmpID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, e)
}

//Employee handlers

// GetEmployee /api/companies/{company_id}/employees/{employee_id} [get]
func GetEmployee(c *gin.Context) {
	cmpIDstr := c.Param("company_id")
	cmpID, err := strconv.Atoi(cmpIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	eIDstr := c.Param("employee_id")
	eID, err := strconv.Atoi(eIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	e, err := models.GetEmployeeByIDAndCompanyID(uint(eID), uint(cmpID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, e)
}

// UpdateEmployee /api/companies/{company_id}/employees/{employee_id} [put]
func UpdateEmployee(c *gin.Context) {
	cmpIDstr := c.Param("company_id")
	cmpID, err := strconv.Atoi(cmpIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	eIDstr := c.Param("employee_id")
	eID, err := strconv.Atoi(eIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	e, err := models.GetEmployeeByIDAndCompanyID(uint(eID), uint(cmpID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err = c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = e.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, e)
}

// DeleteEmployee /api/companies/{company_id}/employees/{employee_id} [delete]
func DeleteEmployee(c *gin.Context) {
	cmpIDstr := c.Param("company_id")
	cmpID, err := strconv.Atoi(cmpIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	eIDstr := c.Param("employee_id")
	eID, err := strconv.Atoi(eIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	e, err := models.GetEmployeeByIDAndCompanyID(uint(eID), uint(cmpID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err = e.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, e)
}
