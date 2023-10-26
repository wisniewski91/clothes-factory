package clients

func (s *ClientsService) Find(id int) (*ClientRecord, error) {
	result, err := s.repository.Fetch(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ClientsService) FindAll() (*[]ClientRecord, error) {
	result, err := s.repository.FetchAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ClientsService) FindPage(page int, limit int) (*[]ClientRecord, error) {
	result, err := s.repository.FetchPage(page, limit)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ClientsService) Save(data *ClientRecord) (*ClientRecord, error) {
	result, err := s.repository.Update(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ClientsService) Create(data *ClientRecord) (*ClientRecord, error) {
	result, err := s.repository.Insert(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ClientsService) Remove(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
