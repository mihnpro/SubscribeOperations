package operations

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"test_task/db"
	"test_task/models/request"
	"test_task/models/response"
)

// GetAllSubInfo godoc
// @Summary Получить все подписки
// @Description Получает список всех подписок
// @Tags subscriptions
// @Accept json
// @Produce json
// @Success 200 {array} responseExamples.GetSudscriptionExample
// @Failure 404 {object} responseExamples.SubNotFoundExample "Подписка не найдена"
// @Failure 500 {object} responseExamples.SomeErrorGetingAllSubsExample "Внутренняя ошибка сервера"
// @Router /getAllSubs [get]
func GetAllSubsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	subscription, err := db.GetAllSubOper(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting all users:", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	var subResp []response.GetSudscription
	for _, sub := range subscription {
		subResp = append(subResp, response.GetSudscription{
			ServiceName: sub.ServiceName,
			Price:       sub.Price,
			UserID:      sub.UserID,
			StartDate:   sub.StartDate,
			EndDate:     sub.EndDate,
		})
	}

	if len(subResp) == 0 {
		log.Println("No subscriptions found")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Subscription not found"})
	}

	if err = json.NewEncoder(w).Encode(subResp); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// CreateSubInfo godoc
// @Summary Создать новую подписку
// @Description Создает новую подписку. Если не указывать user_id он сгенерируется автоматически.
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body requestExamples.CreateSubscriptionReqExample true "Данные подписки"
// @Success 201 {object} responseExamples.SubCreatedExample "Подписка создана"
// @Failure 400 {object} responseExamples.EmptyDateOrInvalidExample "Пустая дата начала и дата окончания подписки или дата начала больше даты окончания"
// @Failure 500 {object} responseExamples.SomeServerErrorsDuriingCreatingSubExample "Внутренняя ошибка сервера"
// @Router /createSub [post]
func CreateSubHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var subscription request.CreateSubscriptionReq
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error decoding request body:", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if subscription.StartDate.IsZero() || subscription.EndDate.IsZero() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "start_date and end_date are required"})
		return
	}

	if subscription.EndDate.Before(subscription.StartDate) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "end_date cannot be before start_date"})
		return
	}

	err := db.CreateSubOper(subscription, r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error creating subscription:", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Println("Subscription created successfully")
	json.NewEncoder(w).Encode(map[string]string{"message": "Subscription created"})
}

// UpdateSubInfo godoc
// @Summary Обновить подписку
// @Description Обновляет существующую подписку
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body requestExamples.UpdateSubscriptionReqExample true "Обновленные данные подписки"
// @Success 200 {object} responseExamples.SubUpdatedExample "Подписка обновлена"
// @Failure 400 {object} responseExamples.EmptyDateOrInvalidExample "Пустая дата начала и дата окончания подписки или дата начала больше даты окончания"
// @Failure 404 {object} responseExamples.SubNotFoundExample "Подписка не найдена"
// @Failure 500 {object} responseExamples.SomeServerErrorsDuriingUpdatingSubExample "Внутренняя ошибка сервера"
// @Router /updateSub [put]
func UpdateSubHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var subscription request.UpdateSubscriptionReq
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error decoding request body:", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if subscription.StartDate.IsZero() || subscription.EndDate.IsZero() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "start_date and end_date are required"})
		return
	}

	if subscription.EndDate.Before(subscription.StartDate) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "end_date cannot be before start_date"})
		return
	}

	err := db.UpdateSubOper(subscription, r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error updating subscription:", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Println("Subscription updated successfully")
	json.NewEncoder(w).Encode(map[string]string{"message": "Subscription updated"})
}

// DeleteSubInfo godoc
// @Summary Удалить подписку
// @Description Удаляет подписку по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body requestExamples.DeleteSubscriptionReqExample true "Данные для удаления"
// @Success 200 {object} responseExamples.SubDeletedExample "Подписка удалена"
// @Failure 404 {object} responseExamples.SubNotFoundExample "Подписка не найдена"
// @Failure 500 {object} responseExamples.SomeServerErrorsDuriingDeletingSubExample "Внутренняя ошибка сервера"
// @Router /deleteSub [delete]
func DeleteSubHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var subscription request.DeleteSubscriptionrReq
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error decoding request body:", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	err := db.DeleteSubOper(subscription, r.Context())

	if err != nil {
		if err.Error() == "no rows deleted" {
			w.WriteHeader(http.StatusNotFound)
			log.Println("Subscription not found:", err)
			json.NewEncoder(w).Encode(map[string]string{"error": "Subscription not found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error deleting subscription:", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Println("Subscription deleted successfully")
	json.NewEncoder(w).Encode(map[string]string{"message": "Subscription deleted"})
}

// GetSubHandler godoc
// @Summary Получить подписку по ID
// @Description Получает информацию о конкретной подписке
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscriptionID path string true "ID подписки"
// @Success 200 {object} responseExamples.GetSudscriptionExample
// @Failure 400 {object} responseExamples.EmptySubOrInvalidIDExample "Не указан ID подписки или ID подписки не является числом"
// @Failure 404 {object} responseExamples.SubNotFoundExample "Подписка не найдена"
// @Failure 500 {object} responseExamples.SomeServerErrorsDuringGettingSubExample "Внутренняя ошибка сервера при получении подписки"
// @Router /getSub/{subscriptionID} [get]
func GetSubHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	subIDStr := r.PathValue("subscriptionID")

	if subIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Subscription ID is required")
		json.NewEncoder(w).Encode(map[string]string{"error": "Subscription ID is required"})
		return
	}

	subID, err := strconv.ParseUint(subIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid subscription ID:", subIDStr)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid subscription ID"})
		return
	}

	sub, err := db.GetSubOperBySubOperID(uint(subID), r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting operation by ID:", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if sub == nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Subscription not found:", subID)
		json.NewEncoder(w).Encode(map[string]string{"error": "Subscription not found"})
		return
	}

	response := response.GetSudscription{
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate,
		EndDate:     sub.EndDate,
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// GetFullSubPriceByPeriod godoc
// @Summary Получить суммарную стоимость подписок за период
// @Description Подсчитывает общую стоимость и количество подписок пользователя за указанный период
// @Tags analytics
// @Accept json
// @Produce json
// @Param request body requestExamples.GetFullPriceByPeriodReqExample true "Параметры периода и пользователя"
// @Success 200 {object} responseExamples.GetFullPriceByPeriodExample
// @Failure 400 {object} responseExamples.EmptyDateOrInvalidExample "Пустая дата начала и дата окончания подписки или дата начала больше даты окончания"
// @Failure 404 {object} responseExamples.SubNotFoundExample "Подписка не найдена"
// @Failure 500 {object} responseExamples.SomeServerErrorsDuringGettingSubExample "Внутренняя ошибка сервера"
// @Router /getFullSubPriceByPeriod [post]
func GetFullSubPriceByPeriod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userReq request.GetFullPriceByPeriodReq

	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error decoding request body:", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if userReq.StartDate.IsZero() || userReq.EndDate.IsZero() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "start_date and end_date are required"})
		return
	}

	if userReq.EndDate.Before(userReq.StartDate) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "end_date cannot be before start_date"})
		return
	}

	subInfo, err := db.GetFullSubPriceByPeriodOper(userReq, r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting operation by ID:", err)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if subInfo == nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Subscription not found")
		json.NewEncoder(w).Encode(map[string]string{"error": "Subscription not found"})
		return
	}

	resp := response.GetFullPriceByPeriodResp{
		UserID:    subInfo.UserID,
		TotalCost: subInfo.TotalCost,
		Count:     subInfo.Count,
	}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
