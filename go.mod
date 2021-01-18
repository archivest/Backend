module backend

go 1.15

require internal/verses v1.0.0
replace internal/verses => ./internal/verses

require internal/commands v1.0.0
replace internal/commands => ./internal/commands