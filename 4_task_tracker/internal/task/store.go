package task

import (
	"encoding/json"
	"os"
)

const filepath = "tasks.json"

type Store struct {
	Tasks []Task
}

func Load() (*Store, error) {
	data, err := os.ReadFile(filepath)
	if os.IsNotExist(err) {
		return &Store{Tasks: []Task{}}, nil
	}

	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return &Store{Tasks: tasks}, nil
}

func (s *Store) Save() error {
	data, err := json.MarshalIndent(s.Tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, data, 0644)
}

func (s *Store) Add(t Task) {
	s.Tasks = append(s.Tasks, t)
}

func (s *Store) Update(id int, description string) bool {
	for i := range s.Tasks {
		if s.Tasks[i].ID == id {
			s.Tasks[i].Description = description
			return true
		}
	}
	return false
}

func (s *Store) Delete(id int) bool {
	for i, t := range s.Tasks {
		if t.ID == id {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (s *Store) SetStatus(id int, status Status) bool {
	for i := range s.Tasks {
		if s.Tasks[i].ID == id {
			s.Tasks[i].Status = status
			return true
		}
	}
	return false
}

func (s *Store) ByStatus(status Status) []Task {
	var out []Task
	for _, t := range s.Tasks {
		if t.Status == status {
			out = append(out, t)
		}
	}
	return out
}