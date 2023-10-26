package jobtypes

import (
	"fmt"
)

func (r *MysqlJobtypesRepository) Fetch(id int) (*JobtypeRecord, error) {
	if id <= 0 {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ?", tableName)
	var jobType JobtypeRecord

	err := r.DB.QueryRow(query, id).Scan(&jobType.Id, &jobType.Name, &jobType.Status, &jobType.UpdatedAt, &jobType.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &jobType, nil
}

func (r *MysqlJobtypesRepository) FetchAll() (*[]JobtypeRecord, error) {
	query := fmt.Sprintf("SELECT * FROM `%s`", tableName)
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var jobTypes []JobtypeRecord

	for rows.Next() {
		c := JobtypeRecord{}
		err = rows.Scan(&c.Id, &c.Name, &c.Status, &c.UpdatedAt, &c.CreatedAt)
		if err != nil {
			return nil, err
		}

		jobTypes = append(jobTypes, c)
	}
	return &jobTypes, nil
}

func (r *MysqlJobtypesRepository) FetchPage(page int, limit int) (*[]JobtypeRecord, error) {
	query := fmt.Sprintf("SELECT * FROM `%s` LIMIT %d OFFSET %d", tableName, limit, page*limit)
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var jobTypes []JobtypeRecord

	for rows.Next() {
		c := JobtypeRecord{}
		err = rows.Scan(&c.Id, &c.Name, &c.Status, &c.UpdatedAt, &c.CreatedAt)
		if err != nil {
			return nil, err
		}

		jobTypes = append(jobTypes, c)
	}
	return &jobTypes, nil
}

func (r *MysqlJobtypesRepository) Update(data *JobtypeRecord) (*JobtypeRecord, error) {
	query := fmt.Sprintf("UPDATE `%s` SET `name` = ? WHERE `id` = ?", tableName)
	_, err := r.DB.Exec(query, data.Name, data.Id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *MysqlJobtypesRepository) Insert(data *JobtypeRecord) (*JobtypeRecord, error) {
	query := fmt.Sprintf("INSERT INTO `%s` (`name`, `status`, `updated_at`, `created_at`) VALUES (?, 4, Now(), Now())", tableName)
	_, err := r.DB.Exec(query, data.Name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *MysqlJobtypesRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM `%s` WHERE `id` = ?", tableName)
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
