package main

import "github.com/99designs/gqlgen/graphql"

// central file to the whole api gateway
// with the help of server we can call the account product and order client

type Server struct {
	accountClient *account.Client
	catalogClient *catalog.Client
	orderClient   *order.Client
}

func NewGraphQLserver(accountUrl, catalogUrl, orderUrl string) (*Server, error) {

	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	catalogClient, err := catalog.NewClient(catalogUrl)

	if err != nil {
		accountClient.Close()
		return nil, err
	}

	orderClient, err := order.NewClient(orderUrl)

	if err != nil {
		accountClient.close()
		catalogClient.close()
		return nil, err
	}

	return &Server{
		accountClient,
		catalogClient,
		orderClient,
	}, nil

}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}

}

func (s *Server) Account() AccountResolver {
	return &accountResolver{
		server: s,
	}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
