package clients

import (
	"fmt"

	"github.com/wisniewski91/clothes-factory/internal/utils"
)

func (r *MysqlClientsRepository) Fetch(id int) (*ClientRecord, error) {
	if id <= 0 {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ?", tableName)
	var client ClientRecord
	var address utils.AddressData
	err := r.DB.QueryRow(query, id).Scan(&client.Id, &client.Name, &address.Country, &address.City, &address.District, &address.Street, &address.PostalCode, &address.HouseNumber, &address.FlatNumber, &client.Status, &client.UpdatedAt, &client.CreatedAt)
	if err != nil {
		return nil, err
	}
	client.AddressData = address
	return &client, nil
}

func (r *MysqlClientsRepository) FetchAll() (*[]ClientRecord, error) {
	query := fmt.Sprintf("SELECT * FROM `%s`", tableName)
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var clients []ClientRecord

	for rows.Next() {
		var address utils.AddressData
		client := ClientRecord{}
		err = rows.Scan(&client.Id, &client.Name, &address.Country, &address.City, &address.District, &address.Street, &address.PostalCode, &address.HouseNumber, &address.FlatNumber, &client.Status, &client.UpdatedAt, &client.CreatedAt)
		if err != nil {
			return nil, err
		}
		client.AddressData = address
		clients = append(clients, client)
	}
	return &clients, nil
}

func (r *MysqlClientsRepository) FetchPage(page int, limit int) (*[]ClientRecord, error) {
	query := fmt.Sprintf("SELECT * FROM `%s` LIMIT %d OFFSET %d", tableName, limit, page*limit)
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var clients []ClientRecord

	for rows.Next() {
		var address utils.AddressData
		client := ClientRecord{}
		err = rows.Scan(&client.Id, &client.Name, &address.Country, &address.City, &address.District, &address.Street, &address.PostalCode, &address.HouseNumber, &address.FlatNumber, &client.Status, &client.UpdatedAt, &client.CreatedAt)
		if err != nil {
			return nil, err
		}
		client.AddressData = address
		clients = append(clients, client)
	}
	return &clients, nil
}

func (r *MysqlClientsRepository) Update(data *ClientRecord) (*ClientRecord, error) {
	query := fmt.Sprintf("UPDATE `%s` SET `name` = ?, `country` = ?,`city` = ?, `district` = ?, `street` = ?, `postal_code` = ?, `house_number` = ?, `flat_number` = ? WHERE `id` = ?", tableName)
	_, err := r.DB.Exec(query, data.Name, data.AddressData.Country, data.AddressData.City, data.AddressData.District, data.AddressData.Street, data.AddressData.PostalCode, data.AddressData.HouseNumber, data.AddressData.FlatNumber, data.Id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *MysqlClientsRepository) Insert(data *ClientRecord) (*ClientRecord, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`name`, `country`,`city`,`district`,`street`,`postal_code`,`house_number`,`flat_number`, `status`, `updated_at`, `created_at`) VALUES (?,?,?,?,?,?,?,?, 4, Now(), Now())", tableName)
	_, err := r.DB.Exec(query, data.Name, data.AddressData.Country, data.AddressData.City, data.AddressData.District, data.AddressData.Street, data.AddressData.PostalCode, data.AddressData.HouseNumber, data.AddressData.FlatNumber)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *MysqlClientsRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM `%s` WHERE `id` = ?", tableName)
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
