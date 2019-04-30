package utils

import (
	"encoding/json"
	"reflect"

	log "github.com/sirupsen/logrus"

	"github.com/tidwall/pretty"
)

// PrettyPrint - output formatted json
func PrettyPrint(i interface{}) {
	log.Debug(reflect.TypeOf(i))
	s, _ := json.MarshalIndent(i, "", "\t")
	colored := pretty.Color(s, nil)
	log.Debug(string(colored))
}
