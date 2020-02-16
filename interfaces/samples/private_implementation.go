package main

func main() {
	repo := customer.NewRepo(db)
	c := repo.Get(1)
}
