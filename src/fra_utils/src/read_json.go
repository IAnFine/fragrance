package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SwaggerInfo struct {
	Title   string
	Version string
}

type SwaggerTag struct {
	Name        string
	Description string
}

type SwaggerApiParameter struct {
	Description string
	Format      string
	In          string
	Name        string
	Required    bool
	Type        string
}

type SwaggerPathInfo struct {
	Consumes    []string
	Description string
	OperationId string
	Parameters  []*SwaggerApiParameter
	Tags        []string
	Summary     string
}

type SwaggerPath struct {
	Post SwaggerPathInfo
	Get  SwaggerPathInfo
}

type SwaggerJson struct {
	Host     string
	Info     *SwaggerInfo
	Produces []string
	Swagger  string
	Tags     []*SwaggerTag
	Paths    map[string]*SwaggerPath
}

/*

  根据参数提供的URL 访问对应地址的swagger数据

  swagger暂定版本为2.0
*/
func CreateAPIDocumentFromSwaggerJson(url string) {

}

func GetSwaggerJson(url string) (*SwaggerJson, error) {
	fmt.Println("starting get swagger data.......")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("An error occurred")
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	fmt.Println("star encode........")

	// 判断http状态码是否正确
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result SwaggerJson
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
