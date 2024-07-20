package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(conn *sql.DB) ProductRepository {
	return ProductRepository{
		connection: conn,
	}
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT " +
		"id, name, price " +
		"FROM product " +
		"WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product
	err = query.QueryRow(id).Scan(&product.ID, &product.Name, &product.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	return &product, err
}

func (pr *ProductRepository) UpdateProductById(id int, product model.Product) (*model.Product, error) {
	query, err := pr.connection.Prepare("UPDATE product SET name = $1, price = $2 WHERE id = $3 RETURNING id, name, price")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer query.Close()

	err = query.QueryRow(product.Name, product.Price, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepository) CreateProduct(p model.Product) (int, error) {
	query, err := pr.connection.Prepare("INSERT INTO " +
		"product(name, price) " +
		"VALUES ($1, $2) RETURNING ID")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var id int
	err = query.QueryRow(p.Name, p.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, err
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price FROM product"

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product
	for rows.Next() {
		err = rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}
	rows.Close()

	return productList, err
}
