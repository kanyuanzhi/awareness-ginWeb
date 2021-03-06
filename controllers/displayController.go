package controllers

import (
	"TouchAll-Web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SingleEquipmentDisplay(c *gin.Context) {
	equipmentIDStr := c.DefaultQuery("equipmentID", "0")
	equipmentID, _ := strconv.Atoi(equipmentIDStr)
	equipmentGroupsMap := utils.GetEquipmentGroupsMap()
	equipmentTypesMap := utils.GetEquipmentTypesMap()
	if equipment, has := utils.FindEquipmentByID(equipmentID); has {
		equipment.EquipmentTypeName, _ = equipmentTypesMap[equipment.EquipmentType]
		equipment.EquipmentGroupName, _ = equipmentGroupsMap[equipment.EquipmentGroup]
		c.HTML(http.StatusOK, "display/equipment.html", gin.H{
			"title":        "单个设备状态展示",
			"equipment":    equipment,
			"equipmentIDs": []int{equipmentID},
		})
	} else {
		c.String(http.StatusOK, "没有id为%d的设备，请确认设备ID！", equipmentID)
	}
}

func AllEquipmentsDisplay(c *gin.Context) {
	equipments := utils.FindAllEquipmentIDs()
	equipmentIDs := make([]int, 0)
	for _, equipment := range equipments {
		equipmentIDs = append(equipmentIDs, equipment.EquipmentID)
	}

	c.HTML(http.StatusOK, "display/allEquipments.html", gin.H{
		"title":        "所有设备状态展示",
		"equipments":   equipments,
		"equipmentIDs": equipmentIDs,
	})
}

func CameraDisplay(c *gin.Context) {
	cameras := utils.FindAllCameras()
	if cameras == nil {
		c.HTML(http.StatusOK, "display/camera.html", gin.H{
			"title": "监控页面",
		})
	}
	c.HTML(http.StatusOK, "display/camera.html", gin.H{
		"title":   "监控页面",
		"cameras": cameras,
	})
}

func AICameraDisplay(c *gin.Context) {
	aiCameras := utils.FindAllAICameras()
	if aiCameras == nil {
		c.HTML(http.StatusOK, "display/aiCamera.html", gin.H{
			"title": "智能监控页面",
		})
	}
	c.HTML(http.StatusOK, "display/aiCamera.html", gin.H{
		"title":     "智能监控页面",
		"aiCameras": aiCameras,
	})
}

func SingleSensorDisplay(c *gin.Context) {
	sensorIDStr := c.DefaultQuery("sensorID", "0")
	sensorID, _ := strconv.Atoi(sensorIDStr)
	if sensor, has := utils.FindSensorByID(sensorID); has {
		c.HTML(http.StatusOK, "display/sensor.html", gin.H{
			"title":     "单个环境监测展示",
			"sensor":    sensor,
			"sensorIDs": []int{sensorID},
		})
	} else {
		c.String(http.StatusOK, "没有id为%d的设备，请确认设备ID！", sensorID)
	}
}

func AllSensorsDisplay(c *gin.Context) {
	sensors := utils.FindAllSensors()
	sensorIDs := make([]int, 0)
	for _, sensor := range sensors {
		sensorIDs = append(sensorIDs, sensor.SensorID)
	}
	c.HTML(http.StatusOK, "display/allSensors.html", gin.H{
		"title":     "环境监测页面",
		"sensors":   sensors,
		"sensorIDs": sensorIDs,
	})
}
