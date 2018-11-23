package main

import "testing"

func TestAllWorkspaces(t *testing.T) {
	wlist := getAllWorkspaces()

	if len(wlist) != len(spacesList) {
		t.Fail()
	}

	for i, v := range wlist {
		if v.Name != spacesList[i].Name ||
			v.ID != spacesList[i].ID ||
			v.IsChos != spacesList[i].IsChos {
			t.Fail()
			break
		}
	}
}
