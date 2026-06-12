package dao

import (
	"context"
	"market/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type KlineDao struct {
	db *mongo.Database
}

func NewKlineDao(db *mongo.Database) *KlineDao {
	return &KlineDao{
		db: db,
	}
}

func (k *KlineDao) FindBySymbol(ctx context.Context, symbol, period string, count int64) (list []*model.Kline, err error) {
	// 按照时间 降序排序
	mk := model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	cur, err := collection.Find(ctx, bson.D{{}}, &options.FindOptions{
		Limit: &count,
		Sort:  bson.D{{"time", -1}},
	})
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return
}
func (k *KlineDao) FindBySymbolTime(ctx context.Context, symbol, period string, from, end int64) (list []*model.Kline, err error) {

	// 按照时间 降序排序
	mk := model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	cur, err := collection.Find(ctx, bson.D{{"time", bson.D{{"$gte", from}, {"$lte", end}}}}, &options.FindOptions{
		Sort: bson.D{{"time", -1}},
	})
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return
}
