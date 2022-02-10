package usecase

import (
	"polynomial/model"
	usecase "polynomial/usecase/interface"
	"regexp"
	"strconv"
)

var DATASET = []string{
	"1",
	"x",
	"8",
	"17",
	"y",
	"z",
	"78",
	"113",
}

var AROUND = 3 // around of third degree polynomial

var IsLetter = regexp.MustCompile(`^[a-z]+$`).MatchString

type polynomialUsecase struct {
}

func NewPolynomialUsecase() usecase.PolynomialUsecaseInterface {
	return &polynomialUsecase{}
}

// หาลำดับพหุนาม
func (p polynomialUsecase) Calculate(datasetReq []int, around int) (datasetRes []int) {
	for i, v := range datasetReq {
		if i+1 != len(datasetReq) {
			r := datasetReq[i+1] - v
			datasetRes = append(datasetRes, r)
		}
	}

	around += 1

	if !p.CheckDataset(datasetRes) && len(datasetRes) > 1 && around < 3 {
		return p.Calculate(datasetRes, around)
	}

	return
}

func (p polynomialUsecase) CheckDataset(datasetReq []int) bool {
	if len(datasetReq) <= 1 {
		return false
	}

	for i, v := range datasetReq {
		if i+1 != len(datasetReq) {
			if v != datasetReq[i+1] {
				return false
			}
		}
	}

	return true
}

func (p polynomialUsecase) RadomArg() (datasetRes []int) {
	// find variable
	indexVariable := make(map[int]int)
	for di, dv := range DATASET {
		isVar := IsLetter(dv)
		if isVar {
			indexVariable[di] = di
		}
	}

	var expectArg [][]int

	for di, _ := range DATASET {
		if di == indexVariable[di] && di != 0 {
			i, j := 1, 1
			if IsLetter(DATASET[di-i]) {
				i += 1
			}
			if IsLetter(DATASET[di+j]) {
				j += 1
			}
			fVar := DATASET[di-i]
			lVar := DATASET[di+j]
			iFVar, _ := strconv.Atoi(fVar)
			iLVar, _ := strconv.Atoi(lVar)

			var expectArr []int
			for fArr := iFVar; fArr <= iLVar; fArr++ {
				expectArr = append(expectArr, fArr)
			}
			expectArg = append(expectArg, expectArr)
		}
	}

	// x argument
	for x := 0; x < len(expectArg[0]); x++ {
		// y argument
		for y := 0; y < len(expectArg[1]); y++ {
			// z argument
			for z := 0; z < len(expectArg[2]); z++ {
				prepareDataset := model.CalculateRequest{}
				prepareDataset.X = expectArg[0][x]
				prepareDataset.Y = expectArg[1][y]
				prepareDataset.Z = expectArg[2][z]
				dataset := p.PrepareDataset(prepareDataset)
				calculate := p.Calculate(dataset, 0)
				checkDataset := p.CheckDataset(calculate)
				if checkDataset {
					return dataset
				}
			}
		}
	}
	return
}

func (p polynomialUsecase) PrepareDataset(request model.CalculateRequest) (preDataset []int) {
	for _, v := range DATASET {
		switch v {
		case "x":
			preDataset = append(preDataset, request.X)
		case "y":
			preDataset = append(preDataset, request.Y)
		case "z":
			preDataset = append(preDataset, request.Z)
		default:
			i, _ := strconv.Atoi(v)
			preDataset = append(preDataset, i)
		}
	}
	return
}
