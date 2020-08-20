package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/mokichi13/dynamo/internal/model"
)

type Repository interface {
	Put(model.Model) error
}

type repo struct {
	db *dynamo.DB
}

func (r repo) Put(m model.Model) error {
	//write object
	if err := r.db.Table("object").Put(m.Object).Run(); err != nil {
		return err
	}

	//write details
	for _, v := range m.Details {
		if err := r.db.Table("detail").Put(v).Run(); err != nil {
			return err
		}
	}
	return nil
}

func NewRepository() (Repository, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region:   aws.String(endpoints.ApNortheast1RegionID),
			Endpoint: aws.String("http://localhost:8000"),
		},
	)
	if err != nil {
		return nil, err
	}

	db := dynamo.New(sess, &aws.Config{
		DisableSSL: aws.Bool(false),
	})

	return &repo{
		db: db,
	}, nil
}
