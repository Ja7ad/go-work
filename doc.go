// Package go_work
// # go-work
// An example true use of go work feature
//
// ### create go.work
//
//	$ go work init
//
// ### create root go.mod
//
//	$ go mod init github.com/Ja7ad/go-work
//
// ### add module to work
//
//	$ go work use mod1
//
// ### init go.mod in modules
//
//	$ cd mod1
//	$ go mod init github.com/Ja7ad/go-work/mod1
//
// ### sync module in work
//
//	$ go work sync
//
// ### create tag for each modules
//
// for tagging each module just need use this pattern `{module name}/{version with prefix v}`
//
//	$ git tag mod1/v1.0.0
//	$ git tag mod2/v1.0.0
//	$ git tag mod3/v1.0.0
//
// ### how to add pkg.go.dev
//
// - first push modules with tags `git push origin main --tags`
// - add git address to pkg.go.dev `https://pkg.go.dev/github.com/Ja7ad/go-work`
package go_work
