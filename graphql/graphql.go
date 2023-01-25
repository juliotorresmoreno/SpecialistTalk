package graphql

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo/v4"
)

/**
* Esta funcion se deja pero no se implementara, Graphql resulta ser una poderosa
* herramienta pero otorga ciertas consideraciones respecto a seguridad, rendimiento
* y aprovisionamiento correcto de cache y otros recursos.
*
* Se conserva el codigo por si se decide exponer algunos endpoints por medio de este
* medio en el futuro
 */

type GraphQLController struct {
	Schema *graphql.Schema
}

func NewGraphQLController() *GraphQLController {
	return new(GraphQLController)
}

func (u *GraphQLController) AppendQuery(fields graphql.Fields) {
	for key, field := range fields {
		rootQueryFields[key] = field
	}
}

func (u *GraphQLController) AppendMutation(fields graphql.Fields) {
	for key, field := range fields {
		rootMutationFields[key] = field
	}
}

func (u *GraphQLController) AppendSubscription(fields graphql.Fields) {
	for key, field := range fields {
		rootSubscriptionFields[key] = field
	}
}

func (u *GraphQLController) AttachGraphQL(g *echo.Group) {
	// Schema
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: rootQueryFields}
	rootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: rootMutationFields}
	var rootSubscription *graphql.Object
	if len(rootSubscriptionFields) > 0 {
		rootSubscription = graphql.NewObject(
			graphql.ObjectConfig{
				Name:   "RootSubscription",
				Fields: rootSubscriptionFields,
			},
		)
	}

	schemaConfig := graphql.SchemaConfig{
		Query:        graphql.NewObject(rootQuery),
		Mutation:     graphql.NewObject(rootMutation),
		Subscription: rootSubscription,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	u.Schema = &schema

	h := handler.New(&handler.Config{
		Schema:   u.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	graphqlHandler := echo.WrapHandler(h)
	g.GET("", graphqlHandler)
	// g.POST("", graphqlHandler)

	g.POST("", u.Post)
}

type PostPayload struct {
	Query string `json:"query"`
}

func (u GraphQLController) Post(c echo.Context) error {
	payload := &PostPayload{}

	err := c.Bind(payload)
	if err != nil {
		c.Error(err)
	}

	params := graphql.Params{Schema: *u.Schema, RequestString: payload.Query}
	r := graphql.Do(params)

	return c.JSON(200, r)
}
