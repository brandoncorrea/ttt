go test
fswatch -o ./ | xargs -n1 -I{} go test
