package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ProcessRequestBody(c *gin.Context, expectedKeys []string) (map[string]interface{}, error) {
	/* _, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	defer c.Request.Body.Close() */

	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		return nil, err
	}

	for _, key := range expectedKeys {
		if _, exists := data[key]; !exists {
			return nil, fmt.Errorf("Property '%s' not found in request body", key)
		}
	}

	return data, nil
}
