package main

import "context"

func listTaskHandler(s *state, c command) error {
	tasks, err := s.DB.GetTasks(context.Background())
	if err != nil {
		return err
	}
	for _, t := range tasks {
		printTask(t)
	}
	return nil
}

func startTaskHandler(s *state, c command) error {
	return nil
}

func completeTaskHandler(s *state, c command) error {
	return nil
}

func editTaskHandler(s *state, c command) error {
	return nil
}

func allocateTimeHandler(s *state, c command) error {
	return nil
}
