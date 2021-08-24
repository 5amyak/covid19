package repository

import (
	"context"
	"time"

	"github.com/covid19/domain"
	"github.com/covid19/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Cases *mongo.Collection

func UpdateCaseCount(regionalCase *domain.RegionalCase, lastRefreshedOn time.Time) {
	document := bson.D{
		{"state", regionalCase.State},
		{"confirmed", regionalCase.Confirmed},
		{"discharged", regionalCase.Discharged},
		{"deaths", regionalCase.Deaths},
		{"lastRefreshedOn", lastRefreshedOn},
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"state", regionalCase.State}}
	update := bson.D{{"$set", document}}

	_, err := Cases.UpdateOne(context.TODO(), filter, update, opts)
	helpers.CheckErr(err)
}

func FetchAggregatedCaseCount() map[string]interface{} {
	var err error
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", 1},
			{"totalConfirmed", bson.D{
				{"$sum", "$confirmed"},
			}},
			{"totalDeaths", bson.D{
				{"$sum", "$deaths"},
			}},
			{"totalDischarged", bson.D{
				{"$sum", "$discharged"},
			}},
		}},
	}

	cursor, err := Cases.Aggregate(context.TODO(), mongo.Pipeline{groupStage})
	helpers.CheckErr(err)
	var results []map[string]interface{}
	err = cursor.All(context.TODO(), &results)
	helpers.CheckErr(err)
	return results[0]
}

func FetchCaseCount(state string) map[string]interface{} {
	filter := bson.D{{"state", state}}

	var result map[string]interface{}
	err := Cases.FindOne(context.TODO(), filter).Decode(&result)
	helpers.CheckErr(err)

	return result
}
