package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/HF-RapidResponse/geotools/volunteerd/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *Server) zipCodeTestCalls(ctx context.Context) error {
	zipCodeCount, err := models.ZipCodes().Count(ctx, s.db)
	if err != nil {
		return err
	}

	zipCodes, err := models.ZipCodes(
		qm.Where(`zcta5ce10 = ?`, *flagZipCode),
	).All(ctx, s.db)
	if err != nil {
		return err
	}
	if *flagVerbose {
		fmt.Println("zip codes:", zipCodeCount)
		fmt.Println("zip code query:", *flagZipCode)
		for _, zipCode := range zipCodes {
			fmt.Println("zip code:", zipCode)
		}
	}
	cds, err := models.CongressionalDistricts(
		qm.Where(`ST_Intersects(geom, (select geom from zip_codes where zcta5ce10 = ?))`, *flagZipCode),
	).All(ctx, s.db)
	if err != nil {
		return err
	}
	m := json.NewEncoder(os.Stdout)
	for _, cd := range cds {
		m.Encode(cd)
	}
	return nil
}
