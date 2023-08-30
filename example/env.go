package example

import "go.mamad.dev/env-loader"

func main() {
	envLoader.Path = ".env"
	envLoader.Load()
}
