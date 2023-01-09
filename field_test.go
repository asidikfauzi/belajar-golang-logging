package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "fauzi").Info("Hello World")

	logger.WithField("username", "adelia").
		WithField("name", "Adelia Risky").
		Info("Hello World")

}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "asidikfauzi",
		"name":     "achmad sidik fauzi",
	}).Info("Hello World")
}
