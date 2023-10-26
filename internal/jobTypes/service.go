package jobtypes

func (s *JobtypesService) Find(id int) (*JobtypeRecord, error) {
	result, err := s.repository.Fetch(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *JobtypesService) FindAll() (*[]JobtypeRecord, error) {
	result, err := s.repository.FetchAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *JobtypesService) FindPage(page int, limit int) (*[]JobtypeRecord, error) {
	result, err := s.repository.FetchPage(page, limit)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *JobtypesService) Save(data *JobtypeRecord) (*JobtypeRecord, error) {
	result, err := s.repository.Update(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *JobtypesService) Create(data *JobtypeRecord) (*JobtypeRecord, error) {
	result, err := s.repository.Insert(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *JobtypesService) Remove(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}