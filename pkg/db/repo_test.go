package db

import (
	"misoda.fr/easyscan/pkg/domain"
	"testing"
)

func TestGetAll(t *testing.T) {
	t.Fatal("not implemented")
}

func TestGetOne(t *testing.T) {
	docRepo := &DocRepo{}
	i, err := docRepo.GetOne("1")
	if err != nil {
		t.Errorf("failed")
	}
	doc, ok := i.(*domain.Document)
	if !ok {
		t.Errorf("failed")
	}
	if doc.ID == 0 {
		t.Errorf("failed")
	}
}

func TestAddOne(t *testing.T) {
	doc := &domain.Document{
		DocType:      "facture",
		Name:         "facture xelians 042020",
		CreationDate: "2020",
	}

	docRepo := &DocRepo{}

	i, err := docRepo.AddOne(doc)

	if err != nil {
		t.Errorf("failed")
	}

	addedDoc, ok := i.(*domain.Document)

	if !ok {
		t.Errorf("failed")
	}

	if addedDoc.ID == 0 {
		t.Errorf("failed")
	}
}

func TestUpdate(t *testing.T) {
	t.Fatal("not implemented")
}
