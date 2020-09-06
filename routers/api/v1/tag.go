package v1

import (
	"blog/models"
	"blog/packages/error"
	"blog/packages/settings"
	"blog/packages/utils"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetTags(c *gin.Context) {
	// c.DefaultQuery support default value.
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := error.SUCCESS

	data["lists"] = models.GetTags(utils.GetPage(c), settings.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("name cannot be empty")
	valid.MaxSize(name, 100, "name").Message("max length for name is 100")
	valid.Required(createdBy, "created_by").Message("created_by cannot be empty")
	valid.MaxSize(createdBy, 100, "created_by").Message("max length for created_by is 100")
	valid.Range(state, 0, 1, "state").Message("state can only be choose from 0 and 1")

	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = error.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = error.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("state can only be choose from 0 and 1")
	}

	valid.Required(id, "id").Message("id cannot be empty")
	valid.Required(modifiedBy, "modified_by").Message("modified_by cannot be empty")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("max length for modified_by is 100")
	valid.MaxSize(name, 100, "name").Message("max length for name is 100")

	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		code = error.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = error.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id cannot be less then 1")

	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		code = error.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = error.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]string),
	})
}
