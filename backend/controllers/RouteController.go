package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/abdianysyah/backend/database"
	"github.com/abdianysyah/backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllRoute(c *gin.Context)  {
	search := c.Query("search")

	var routes []models.Route

	query := database.DB.Model(&models.Route{})

	if search != "" {
		query = query.Where(
			"origin LIKE ? OR destination LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		)
		
	}

	query.Find(&routes)

	c.JSON(http.StatusOK, gin.H{
		"data" : routes,
	})
}


func CreateRoute(c *gin.Context)  {
	var routes models.Route

	if err := c.ShouldBindJSON(&routes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	originCity, err1 := getOrCreateCity(routes.Origin)
	destCity, err2 := getOrCreateCity(routes.Destination)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Kota tidak ditemukan",
		})
		return
	}

	routes.OriginLat = originCity.Lat
	routes.OriginLng = originCity.Lng
	routes.DestLat = destCity.Lat
	routes.DestLng = destCity.Lng

	if err := database.DB.Create(&routes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Gagal menyimpan rute",
		})		
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Rute berhasil ditambahkan!",
		"data" : routes,
	})
}

func GetRouteByID(c *gin.Context)  {
	id := c.Param("id")

	var routes models.Route

	if err := database.DB.First(&routes, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Rute tidak ditemukan!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : routes,
	})
}


func UpdateRoute(c *gin.Context)  {
	id := c.Param("id")

	var routes models.Route

	if err := database.DB.First(&routes, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "Rute tidak ditemukan!",
		})
		return
	}

	var input models.Route
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	routes.Origin = input.Origin
	routes.Destination = input.Destination
	routes.Distance = input.Distance

	originCity, err1 := getOrCreateCity(routes.Origin)
	destCity, err2 := getOrCreateCity(routes.Destination)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Kota tidak ditemukan",
		})
		return
	}

	routes.OriginLat = originCity.Lat
	routes.OriginLng = originCity.Lng
	routes.DestLat = destCity.Lat
	routes.DestLng = destCity.Lng

	database.DB.Save(&routes)

	c.JSON(http.StatusOK, gin.H{
		"message" : "Rute berhasil diupdate!",
		"rute" : routes,
	})
}

func DeleteRoute(c *gin.Context)  {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Route{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Gagal menghapus rute",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Rute berhasil dihapus!",
	})
}

// Update, nomatim mati!
func getCoordinates(city string) (float64, float64) {
	url := "https://photon.komoot.io/api/?q=" + city + "&limit=1"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "NgeBus App")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	features := result["features"].([]interface{})
	if len(features) == 0 {
		return 0, 0
	}

	coords := features[0].(map[string]interface{})["geometry"].(map[string]interface{})["coordinates"].([]interface{})

	lng := coords[0].(float64)
	lat := coords[1].(float64)

	return lat, lng
}

func getOrCreateCity(cityName string) (models.City, error)  {
	var city models.City

	err := database.DB.Where("name = ?", cityName).First(&city).Error
	if err == nil {
		return city, nil
	}

	lat, lng := getCoordinates(cityName)

	if lat == 0 {
		return city, errors.New("Kota tidak ditemukan")
	}

	city = models.City{
		Name: cityName,
		Lat: lat,
		Lng: lng,
	}

	database.DB.Create(&city)

	return city, nil
}