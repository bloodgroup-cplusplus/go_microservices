package main

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
		return nil, err
	}

	orderClient, err := order.NewClient(orderUrl)

	if err != nil {
		return nil, err
	}

}
