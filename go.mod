module example.com/demo

go 1.16

require (
	example.com/greetings2 v0.0.0-00010101000000-000000000000 // indirect
	rsc.io/quote v1.5.2
)

replace example.com/greetings2 => ./greetings2
