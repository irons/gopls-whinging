module foobar/mymodule

go 1.20

require github.com/drhodes/golorem v0.0.0-20220328165741-da82e5b29246 // indirect

replace github.com/drhodes/golorem v0.0.0-20220328165741-da82e5b29246 => ./external/golorem
