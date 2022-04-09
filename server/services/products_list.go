package service

import repository "github.com/lil-shimon/oms/repositories"

func GetProductsLists() (repository.ProductsLists, error) {
	ps, err := repository.GetProductsLists()
	return ps, err
}
