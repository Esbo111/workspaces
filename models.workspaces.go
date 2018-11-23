package main

import "errors"

type workspaces struct {
	ID     int    `json:"id"`
	Name   string `json:"name" `
	IsChos bool   `json:"is_choosable"`
}

var spacesList = []workspaces{
	workspaces{ID: 1, Name: "workspace1", IsChos: true},
	workspaces{ID: 2, Name: "workspace2", IsChos: true},
}

func getAllWorkspaces() []workspaces {
	return spacesList
}

func getWorkSpaceByID(id int) (*workspaces, error) {
	for _, a := range spacesList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Workspace not found")
}
