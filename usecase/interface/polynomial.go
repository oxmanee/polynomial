package usecase

import "polynomial/model"

type PolynomialUsecaseInterface interface {
	Calculate(datasetReq []int, around int) (datasetRes []int)
	CheckDataset(datasetReq []int) bool
	RadomArg() (datasetRes []int)
	PrepareDataset(request model.CalculateRequest) (preDataset []int)
}
