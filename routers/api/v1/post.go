package v1

import (
	"blog/models"
	"blog/packages/error"
	"blog/packages/settings"
	"blog/packages/utils"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

func GetPost(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id cannot be less then 0")

	code := error.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistPostByID(id) {
			data = models.GetPost(id)
			code = error.SUCCESS
		} else {
			code = error.ERROR_NOT_EXIST_POST
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

func GetPosts(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("state can only be choose from 0 and 1")
	}

	tagId := -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("id cannot be less then 1")
	}

	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		code = error.SUCCESS

		data["lists"] = models.GetPosts(utils.GetPage(c), settings.PageSize, maps)
		data["total"] = models.GetPostNum(maps)

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

func AddPost(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("id cannot be less then 1")
	valid.Required(title, "title").Message("title cannot be empty")
	valid.Required(desc, "desc").Message("desc cannot be empty")
	valid.Required(content, "content").Message("content cannot be empty")
	valid.Required(createdBy, "created_by").Message("created_by cannot be empty")
	valid.Range(state, 0, 1, "state").Message("state can only be choose from 0 and 1")

	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddPost(data)
			code = error.SUCCESS
		} else {
			code = error.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

func EditPost(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("state can only be choose from 0 and 1")
	}

	valid.Min(id, 1, "id").Message("id cannot be less then 1")
	valid.MaxSize(title, 100, "title").Message("max length for title is 100")
	valid.MaxSize(desc, 255, "desc").Message("max length for desc is 100")
	valid.MaxSize(content, 65535, "content").Message("max length for content is 65535")
	valid.Required(modifiedBy, "modified_by").Message("modified_by cannot be empty")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("max length for modified_by is 100")

	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistPostByID(id) {
			if models.ExistTagByID(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				models.EditPost(id, data)
				code = error.SUCCESS
			} else {
				code = error.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = error.ERROR_NOT_EXIST_POST
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeletePost(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id cannot be less then 1")

	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistPostByID(id) {
			models.DeletePost(id)
			code = error.SUCCESS
		} else {
			code = error.ERROR_NOT_EXIST_POST
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]string),
	})
}
