package graphql

import (
	"time"

	"github.com/graphql-go/graphql"
)

var rootSubscriptionFields = graphql.Fields{
	"currentTime": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return time.Now().Format(time.ANSIC), nil
		},
	},
}
