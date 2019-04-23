module bat-pp

go 1.12

require (
	github.com/gpestana/bat-pp/predictor v0.0.0
	github.com/gpestana/bat-pp/scrapper v0.0.0
)

replace github.com/gpestana/bat-pp/predictor v0.0.0 => ./predictor

replace github.com/gpestana/bat-pp/scrapper v0.0.0 => ./scrapper
