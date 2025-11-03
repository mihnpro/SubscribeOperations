package db

import (
	"context"
	"errors"
	"log"
	"test_task/models/database"
	"test_task/models/request"

	"github.com/gofrs/uuid"
)

func GetAllSubOper(ctx context.Context) ([]database.SubscriptionDB, error) {
	db := GetDB()

	var subscriptions []database.SubscriptionDB

	query := "select service_name, price, user_id, start_date, end_date from subscriptions order by service_name"

	err := db.SelectContext(ctx, &subscriptions, query)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}

	return subscriptions, nil
}

func CreateSubOper(sub request.CreateSubscriptionReq, ctx context.Context) error {
	db := GetDB()

	dbSub := database.SubscriptionDB{
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate,
		EndDate:     sub.EndDate,
	}

	if dbSub.UserID == uuid.Nil {
		newUUID, err := uuid.NewV4()
		if err != nil {
			log.Printf("Error generating UUID: %v", err)
			return err
		}

		dbSub.UserID = newUUID
	}

	exists, err := CheckIfSubAlreadyExists(dbSub, ctx)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("subscription alredy exists")
	}

	query := "insert into subscriptions (service_name, price, user_id, start_date, end_date) values ($1, $2, $3, $4, $5)"
	res, err := db.ExecContext(ctx, query, dbSub.ServiceName, dbSub.Price, dbSub.UserID, dbSub.StartDate, dbSub.EndDate)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}

	count, errWithRowsAffected := res.RowsAffected()
	if errWithRowsAffected != nil {
		return errWithRowsAffected
	}

	if count == 0 {
		return errors.New("no rows inserted")
	}

	return nil
}

func CheckIfSubAlreadyExists(sub database.SubscriptionDB, ctx context.Context) (bool, error) {
	db := GetDB()

	query := "select count(*) from subscriptions where service_name = $1 and user_id = $2 and start_date = $3 and end_date = $4"

	var count int
	err := db.GetContext(ctx, &count, query, sub.ServiceName, sub.UserID, sub.StartDate, sub.EndDate)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return false, err
	}

	return count > 0, nil

}
func UpdateSubOper(sub request.UpdateSubscriptionReq, ctx context.Context) error {
	db := GetDB()

	subDb := database.SubscriptionDB{
		SubscriptionID: sub.SubscriptionID,
		ServiceName:    sub.ServiceName,
		Price:          sub.Price,
		StartDate:      sub.StartDate,
		EndDate:        sub.EndDate,
	}

	query := "update subscriptions set service_name = $1, price = $2, start_date = $3, end_date = $4 where subscription_id = $5"
	res, err := db.ExecContext(ctx, query, subDb.ServiceName, subDb.Price, subDb.StartDate, subDb.EndDate, subDb.SubscriptionID)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}
	count, errWithRowsAffected := res.RowsAffected()
	if errWithRowsAffected != nil {
		return errWithRowsAffected
	}

	if count == 0 {
		return errors.New("subscription not updated")
	}

	return nil
}

func DeleteSubOper(sub request.DeleteSubscriptionrReq, ctx context.Context) error {
	db := GetDB()

	query := "delete from subscriptions where subscription_id = $1"
	res, err := db.ExecContext(ctx, query, sub.SubscriptionID)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}
	count, errWithRowsAffected := res.RowsAffected()
	if errWithRowsAffected != nil {
		return errWithRowsAffected
	}

	if count == 0 {
		return errors.New("no rows deleted")
	}
	return nil
}

func GetSubOperBySubOperID(subOperID uint, ctx context.Context) (*database.SubscriptionDB, error) {

	db := GetDB()
	query := "select service_name, price, user_id, start_date, end_date from subscriptions where subscription_id = $1"

	var subInfo database.SubscriptionDB
	err := db.GetContext(ctx, &subInfo, query, subOperID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return &database.SubscriptionDB{}, err
	}

	return &subInfo, nil
}

func GetFullSubPriceByPeriodOper(subInfo request.GetFullPriceByPeriodReq, ctx context.Context) (*database.FullPriceByPeriodDbResp, error) {
	db := GetDB()

	subInfoDb := database.SubscriptionDB{
		UserID:    subInfo.UserID,
		StartDate: subInfo.StartDate,
		EndDate:   subInfo.EndDate,
	}

	var fullPriceDBresp database.FullPriceByPeriodDbResp

	query := "select coalesce(sum(price), 0) as total_cost, count(*) as count from subscriptions where user_id = $1 and start_date <= $2 and end_date >= $3"

	errWithQueryExec := db.GetContext(ctx, &fullPriceDBresp, query, subInfoDb.UserID, subInfoDb.StartDate, subInfoDb.EndDate)
	if errWithQueryExec != nil {
		log.Printf("Error executing query: %v", errWithQueryExec)
		return nil, errWithQueryExec
	}

	if fullPriceDBresp.Count == 0 {
		return nil, errors.New("no rows found")
	}

	fullPriceDBresp.UserID = subInfoDb.UserID

	return &fullPriceDBresp, nil

}
