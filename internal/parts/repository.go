package parts

import (
	"fmt"
)

func (r *MysqlPartsRepository) Fetch(id int) (*PartRecord, error) {
	if id <= 0 {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ?", tableName)
	var part PartRecord

	err := r.DB.QueryRow(query, id).Scan(&part.Id, &part.Name, &part.Status, &part.UpdatedAt, &part.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &part, nil
}

func (r *MysqlPartsRepository) FetchAll() (*[]PartRecord, error) {
	query := fmt.Sprintf("SELECT * FROM `%s`", tableName)
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var parts []PartRecord

	for rows.Next() {
		c := PartRecord{}
		err = rows.Scan(&c.Id, &c.Name, &c.Status, &c.UpdatedAt, &c.CreatedAt)
		if err != nil {
			return nil, err
		}

		parts = append(parts, c)
	}
	return &parts, nil
}

func (r *MysqlPartsRepository) FetchPage(page int, limit int) (*[]PartRecord, error) {
	query := fmt.Sprintf("SELECT * FROM `%s` LIMIT %d OFFSET %d", tableName, limit, page*limit)
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var parts []PartRecord

	for rows.Next() {
		c := PartRecord{}
		err = rows.Scan(&c.Id, &c.Name, &c.Status, &c.UpdatedAt, &c.CreatedAt)
		if err != nil {
			return nil, err
		}

		parts = append(parts, c)
	}
	return &parts, nil
}

func (r *MysqlPartsRepository) Update(data *PartRecord) (*PartRecord, error) {
	query := fmt.Sprintf("UPDATE `%s` SET `name` = ? WHERE `id` = ?", tableName)
	_, err := r.DB.Exec(query, data.Name, data.Id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *MysqlPartsRepository) Insert(data *PartRecord) (*PartRecord, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`name`, `status`, `updated_at`, `created_at`) VALUES (?, 4, Now(), Now())", tableName)
	_, err := r.DB.Exec(query, data.Name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *MysqlPartsRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM `%s` WHERE `id` = ?", tableName)
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
