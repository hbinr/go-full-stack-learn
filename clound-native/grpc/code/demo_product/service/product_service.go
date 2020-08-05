package service

import (
	"context"
	"strconv"
)

type ProdService struct {
}

func (p *ProdService) GetProdName(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	result := new(ProdResponse)
	if req.ProdID > 0 && req.ProdID == 40 {
		result.ProdName = "华为 Mate 40"
	} else {
		result.ProdName = "华为 Mate XX"
	}
	return result, nil
}

func (p *ProdService) GetProdNameList(ctx context.Context, req *QueryRequest) (*ProdListResponse, error) {
	var prods []*ProdResponse
	size := req.PageSize

	if size > 0 {
		for i := 0; i < int(size); i++ {
			prodRes := new(ProdResponse)
			prodRes.ProdName = "华为 Mate 4" + strconv.Itoa(i)
			prods = append(prods, prodRes)
		}
	}

	return &ProdListResponse{
		ProdList: prods,
	}, nil
}
