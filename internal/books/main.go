package main

func main() {
}

func newApplication(
	bookForUserHandler BookForUserHandler,
) Application {
	return Application{
		Commands: Commands{},
		Queries: Queries{
			BookForUser: bookForUserHandler,
		},
	}
}
