package predictor

import (
	"fmt"
	sc "github.com/gpestana/bat-pp/scrapper"
	lm "github.com/pa-m/sklearn/linear_model"
)

func LinearRegression(dg []sc.DataPointGrowth, dp []sc.DataPoint) {
	regr := lm.LinearModel{}

	// split data into training set and test set
	trainX, trainY := split(dg, dp)
	fmt.Println(trainX, trainY)

	// train model with linear regression
	//regr.Fit(trainX, trainY)
	//regr.Predict()

	coef := ""
	mean_sq_err := ""
	var_score := ""

	fmt.Println(regr)
	fmt.Printf("Coefficients: %.3f\n", coef)
	fmt.Printf("Mean squared error: %.2f\n", mean_sq_err)
	fmt.Printf("Variance score: %.2f\n", var_score)
}

// split data for training
func split(dg []sc.DataPointGrowth, dp []sc.DataPoint) ([]sc.DataPointGrowth, []sc.DataPoint) {
	//..
	ty := dg
	tx := dp
	return ty, tx
}
