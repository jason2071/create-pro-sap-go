package main

import (
	"create-pro-sap-go/database"
	"create-pro-sap-go/models"
	"fmt"
)

func main() {
	// กำหนด DSN (Data Source Name)
	db := database.ConnectDatabase()
	defer database.CloseDatabase(db)

	var promotionQueues []models.PromotionQueue
	if err := db.Where("status IN (?)", []string{"new", "error"}).Find(&promotionQueues).Error; err != nil {
		fmt.Println("Error retrieving data:", err)
	} else {
		for _, pq := range promotionQueues {
			fmt.Println("")
			fmt.Printf("ID: %d\n", pq.ID)
			fmt.Printf("RefID: %s\n", pq.RefID)
			fmt.Printf("GroupCode: %s\n", pq.GroupCode)
			fmt.Printf("GroupNo: %s\n", pq.GroupNo)
			fmt.Printf("PromotionCode: %s\n", pq.PromotionCode)
			fmt.Printf("Filename: %s\n", pq.Filename)
			fmt.Printf("Status: %s\n", pq.Status)
			fmt.Printf("InstUser: %d\n", pq.InstUser)
			fmt.Printf("CreatedAt: %s\n", pq.CreatedAt)
			fmt.Printf("UpdatedAt: %s\n", pq.UpdatedAt)
			fmt.Println("")
		}

		fmt.Println("Total records:", len(promotionQueues))
	}
}
