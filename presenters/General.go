package presenters

//This package

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//method to parse strucs to bson
func StructToBson(structObject interface{}) (bson.M, error) {
	var bsonToReturn bson.M
	data, err := bson.Marshal(structObject)
	if err != nil {
		return bsonToReturn, err
	}

	errUnmarshaling := bson.Unmarshal(data, &bsonToReturn)
	if err != nil {
		return bsonToReturn, errUnmarshaling
	}
	return bsonToReturn, nil
}

//method to parse array strucs to bson array
func ArrayStructToBson(array, outArray interface{}) error {
	inStructArrData, err := bson.Marshal(array)
	if err != nil {
		return err
	}
	// kind 4 for array
	raw := bson.Raw{Kind: 4, Data: inStructArrData}

	return raw.Unmarshal(outArray)
}

//method to get timenow
func GetTimeNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//method to push our response in correctly format
func ReturnHttpPayload(object interface{}, w http.ResponseWriter) {
	payload, _ := json.Marshal(object)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "limit, includes, skip, sort, filter, select, self, destination, Accept, Accept-Language, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token,access_token, sectionId, level")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(payload))
}

//methos to push and error return
func ReturnHttpError(err error, w http.ResponseWriter, httpErrType int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "limit, includes, skip, sort, filter, select, destination, self, Accept, Accept-Language, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token, access_token, AuthorizationBroker, sectionId, level")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpErrType)
	w.Write([]byte("{\"success\":false,\"error\": \"" + err.Error() + "\"}"))
}

//method to get all request Values
func GetRequestValue(key string, r *http.Request) string {
	params := mux.Vars(r)
	return params[key]
}

//class and methos to instance new objects into collections in MongoDB
const (
	InstanceStatusActive   string = "ACTIVE"
	InstanceStatusInactive string = "INACTIVE"
	InstanceStatusDeleted  string = "DELETED"
)

type Instance struct {
	Status     string `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt  int64  `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ModifiedAt int64  `json:"modifiedAt,omitempty" bson:"modifiedAt,omitempty"`
}

func CreateInstance() Instance {
	var instance Instance
	instance.Status = InstanceStatusActive
	instance.ModifiedAt = GetTimeNow()
	instance.CreatedAt = GetTimeNow()

	return instance
}
