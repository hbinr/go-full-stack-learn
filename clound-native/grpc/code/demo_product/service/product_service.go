package service

import context "context"

type ProdService struct {
}

func (p *ProdService) GetProdName(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{
		ProdName: "华为 Mate 40",
	}, nil
}
