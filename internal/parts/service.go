package parts

func (s *PartsService) Find(id int) (*PartRecord, error) {
	result, err := s.repository.Fetch(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PartsService) FindAll() (*[]PartRecord, error) {
	result, err := s.repository.FetchAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PartsService) FindPage(page int, limit int) (*[]PartRecord, error) {
	result, err := s.repository.FetchPage(page, limit)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PartsService) Save(data *PartRecord) (*PartRecord, error) {
	result, err := s.repository.Update(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PartsService) Create(data *PartRecord) (*PartRecord, error) {
	result, err := s.repository.Insert(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PartsService) Remove(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}