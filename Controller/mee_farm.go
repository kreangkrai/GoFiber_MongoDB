package Controller

import (
	"context"
	"log"
	"time"

	"github.com/kriangkrai/GoFiber/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReadDoc(device string) []Models.DataModel {

	ctx, cancelFindOne := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"device": device}
	SingleResult, errFind := mg.Db.Collection("mee_farm").Find(ctx, filter)
	if errFind != nil {
		panic(errFind)
	}
	cancelFindOne()

	datas := []Models.DataModel{}
	defer SingleResult.Close(ctx)
	for SingleResult.Next(ctx) {
		var episode Models.DataModel
		if err := SingleResult.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		datas = append(datas, episode)
	}

	return datas
}

func ReadDocAll() []Models.DataModel {

	ctx, cancelFindOne := context.WithTimeout(context.Background(), 10*time.Second)
	SingleResult, errFind := mg.Db.Collection("mee_farm").Find(ctx, bson.D{{}})
	if errFind != nil {
		panic(errFind)
	}
	cancelFindOne()
	datas := []Models.DataModel{}
	defer SingleResult.Close(ctx)
	for SingleResult.Next(ctx) {
		var episode Models.DataModel
		if err := SingleResult.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		datas = append(datas, episode)
	}

	return datas
}

func InsertDoc(data Models.DataModel) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := mg.Db.Collection("mee_farm").InsertOne(ctx, Models.DataModel{ID: primitive.NewObjectID(), Device: data.Device, Date: time.Now().Local().String(), Value: data.Value})

	if err != nil {
		return nil, err
	}
	return res, nil
}
func UpdateDoc(data Models.DataModel) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{{Key: "device", Value: data.Device}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "value", Value: data.Value},
		{Key: "date", Value: time.Now().Local().String()},
	}}}
	res, err := mg.Db.Collection("mee_farm").UpdateMany(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return res, nil
}

//Delete
func DeleteDoc(device string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := mg.Db.Collection("mee_farm").DeleteMany(ctx, bson.M{"device": device})

	if err != nil {
		return nil, err
	}
	return res, nil
}
