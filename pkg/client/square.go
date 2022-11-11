package square

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
)

func RetrieveDummyData() SearchCatalogObjectsResponse {
	file, err := os.ReadFile("debug/square_data.json")
	if err != nil {
		log.Fatal("Unable to open json file.")
	}
	data := SearchCatalogObjectsResponse{}
	json.Unmarshal([]byte(file), &data)
	return data
}

func GetLocalCategoryItems() []RelatedObject {
	var categoryNames []RelatedObject
	allObjs := RetrieveDummyData()
	for _, obj := range allObjs.RelatedObjects {
		if obj.Type == "CATEGORY" {
			categoryNames = append(categoryNames, obj)
		}
	}
	return categoryNames
}

func GetLocalItemsByCategoryId(categoryId string) []CatalogObject {
	var catelogObjects []CatalogObject
	allObjs := RetrieveDummyData()
	for _, obj := range allObjs.Objects {
		if obj.ItemData.CategoryID == categoryId && obj.ItemData.EcomAvailable && obj.ItemData.EcomVisibility == "VISIBLE" {
			catelogObjects = append(catelogObjects, obj)
		}
	}
	return catelogObjects
}

func GetLocalCategoryNames() []string {
	var categoryNames []string
	allObjs := RetrieveDummyData()
	for _, obj := range allObjs.RelatedObjects {
		if obj.Type == "CATEGORY" {
			categoryNames = append(categoryNames, obj.CategoryData.Name)
		}
	}
	return categoryNames
}

func GetCategoryNames() ([]CategoryObject, error) {
	var categoryObjectsResp CategoryObjectResponse
	resp, err := http.Get("http://localhost:8080/categories")
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&categoryObjectsResp)
	if err != nil {
		panic(err)
	}
	categoryObjects := sortCategoriesByName(categoryObjectsResp.Objects)
	return categoryObjects, nil
}

func sortCategoriesByName(categoryItems []CategoryObject) []CategoryObject {
	sort.Slice(categoryItems, func(i, j int) bool {
		return categoryItems[i].CategoryData.Name < categoryItems[j].CategoryData.Name
	})
	return categoryItems
}

func sortCatalogItemsByName(categoryItems []CatalogItem) []CatalogItem {
	sort.Slice(categoryItems, func(i, j int) bool {
		return categoryItems[i].ItemData.Name < categoryItems[j].ItemData.Name
	})
	return categoryItems
}

func getOptionsFromCatalogItem(item CatalogItem) []string {
	var options []string
	for _, option := range item.ItemData.Variations {

		options = append(options, option.ID)
	}
	return options
}

func getModifiersFromCatalogItem(item CatalogItem) []string {
	var modifiers []string
	for _, modifier := range item.ItemData.ModifierListInfo {
		if modifier.Enabled {
			modifiers = append(modifiers, modifier.ModifierListID)
		}
	}
	return modifiers
}

func filterCatalogItemByVisibility(visibility string, categoryItems []CatalogItem) []CatalogItem {
	var filteredItems []CatalogItem
	for _, item := range categoryItems {
		if item.ItemData.EcomVisibility == visibility {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}

func GetItemsByCategoryId(categoryId string) ([]CatalogItem, error) {
	var categoryItemsResp SearchCatalogItemsResponse
	req, err := http.NewRequest("GET", "http://localhost:8080/categoryitems", nil)
	if err != nil {
		os.Exit(1)
	}
	req.URL.RawQuery = url.Values{
		"id": {categoryId},
	}.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&categoryItemsResp)

	if err != nil {
		panic(err)
	}
	filteredItems := filterCatalogItemByVisibility("VISIBLE", categoryItemsResp.Items)
	sortedItems := sortCatalogItemsByName(filteredItems)
	return sortedItems, nil
}

func GetItemOptions(item CatalogItem) ([]Variation, error) {
	options := getOptionsFromCatalogItem(item)
	opts, err := GetVariationsByIds(options)
	if err != nil {
		log.Fatal(err)
	}
	return opts, nil
}

func GetItemModifiers(item CatalogItem) ([]ModifierList, error) {
	modifiers := getModifiersFromCatalogItem(item)
	mods, err := GetModifierListByIds(modifiers)
	if err != nil {
		log.Fatal(err)
	}
	return mods, nil
}
func GetCatalogObjectsByIds(ids []string) ([]CatalogObject, error) {
	batchRequest := BatchRetrieveCatalogObjectsRequest{
		ObjectIds: ids,
	}
	var batchObjectResponse BatchRetrieveCatalogObjectsResponse
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(batchRequest)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "http://localhost:8080/catalogobjects", &buf)
	if err != nil {
		os.Exit(1)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&batchObjectResponse)

	if err != nil {
		panic(err)
	}
	return batchObjectResponse.Objects, nil
}

func GetModifierListByIds(ids []string) ([]ModifierList, error) {
	batchRequest := BatchRetrieveCatalogObjectsRequest{
		ObjectIds: ids,
	}
	var modifierListResponse ModifierListResponse
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(batchRequest)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "http://localhost:8080/catalogobjects", &buf)
	if err != nil {
		os.Exit(1)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&modifierListResponse)

	if err != nil {
		panic(err)
	}
	return modifierListResponse.Objects, nil
}

func GetVariationsByIds(ids []string) ([]Variation, error) {
	batchRequest := BatchRetrieveCatalogObjectsRequest{
		ObjectIds: ids,
	}
	var variationResponse VariationResponse
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(batchRequest)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "http://localhost:8080/catalogobjects", &buf)
	if err != nil {
		os.Exit(1)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&variationResponse)

	if err != nil {
		panic(err)
	}
	return variationResponse.Objects, nil
}
