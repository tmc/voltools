package main

import (
	"context"
	"errors"
	"log"

	"github.com/HF-RapidResponse/geotools/volunteerd/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var ErrZipCodeNotFound = errors.New("zipcode not found")

type ZipInfo struct {
	ZipCode               models.ZipCode               `boil:",bind"`
	CongressionalDistrict models.CongressionalDistrict `boil:",bind"`
	State                 models.State                 `boil:",bind"`
}
type ZipInfo2 struct {
	ZipCode                *models.ZipCode                 `boil:",bind"`
	CongressionalDistricts []*models.CongressionalDistrict `boil:",bind"`
}

func (s *Server) lookupZipInfoOld(ctx context.Context, zip string) (*ZipInfo2, error) {
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
		qm.Where(`ST_Intersects(geom, (select geom from zip_codes where zcta5ce10 = ?))`, zipCode.ZipCode),
	).All(ctx, s.db)
	if err != nil {
		return nil, err
	}
	return &ZipInfo2{
		ZipCode:                zipCode,
		CongressionalDistricts: cds,
	}, nil
}

func (s *Server) lookupZipInfo(ctx context.Context, zip string) (*ZipInfo, error) {
	var zi ZipInfo
	err := models.ZipCodes(
		qm.Where(`zcta5ce10 = ?`, zip),
		qm.LeftOuterJoin("congressional_districts cd on ST_Intersects(cd.geom, zip_codes.geom)"),
		qm.LeftOuterJoin("states s on (cd.statefp = s.statefp)"),
	).Bind(ctx, s.db, &zi)
	return &zi, err
}
