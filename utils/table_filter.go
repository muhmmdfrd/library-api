package utils

import (
	"errors"
	"library-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseFilter(c *gin.Context) (models.TableFilter, error) {
	var filter models.TableFilter
	query := c.Request.URL.Query()

	keyword := query.Get("keyword")
	size, err := strconv.Atoi(query.Get("size"))
	if err != nil {
		return filter, errors.New("unable to parsing query parameter: size")
	}

	index, err := strconv.Atoi(query.Get("index"))
	if err != nil {
		return filter, errors.New("unable to parsing query parameter: index")
	}

	if index <= 0 {
		index = 1
	}

    filter = models.TableFilter{
			Keyword: keyword,
			Index:   index,
			Size:    size,
	}

	return filter, nil
}