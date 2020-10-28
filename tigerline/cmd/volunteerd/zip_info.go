package main

import (
	"context"
	"errors"
	"log"

	"github.com/HF-RapidResponse/geotools/tigerline/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var ErrZipCodeNotFound = errors.New("zipcode not found")

type ZipInfo struct {
	ZipCode                *models.ZipCode
	CongressionalDistricts []*models.CongressionalDistrict
}

func (s *Server) lookupZipInfo(ctx context.Context, zip string) (*ZipInfo, error) {
	zipCodes, err := models.ZipCodes(
		qm.Where(`zcta5ce10 = ?`, zip),
	).All(ctx, s.db)
	if err != nil {
		return nil, err
	}
	if len(zipCodes) == 0 {
		return nil, ErrZipCodeNotFound
	}
	if len(zipCodes) > 1 {
		// TODO: server should have a logger.
		log.Println("multiple zips found:", zipCodes)
	}
	zipCode := zipCodes[0]
	cds, err := models.CongressionalDistricts(
		qm.Where(`ST_Intersects(geom, (select geom from zip_codes where zcta5ce10 = ?))`, zipCodes),
	).All(ctx, s.db)
	if err != nil {
		return nil, err
	}
	return &ZipInfo{
		ZipCode:                zipCode,
		CongressionalDistricts: cds,
	}, nil
}
