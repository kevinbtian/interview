module github.com/kevinbtian/interview/kong

go 1.20

replace github.com/kevinbtian/interview/kong/handler => ./handler

replace github.com/kevinbtian/interview/kong/database => ./database

require github.com/lib/pq v1.10.8 // indirect
