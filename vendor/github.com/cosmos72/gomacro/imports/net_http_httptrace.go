// this file was generated by gomacro command: import _b "net/http/httptrace"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"net/http/httptrace"
)

// reflection: allow interpreted code to import "net/http/httptrace"
func init() {
	Packages["net/http/httptrace"] = Package{
	Binds: map[string]Value{
		"ContextClientTrace":	ValueOf(httptrace.ContextClientTrace),
		"WithClientTrace":	ValueOf(httptrace.WithClientTrace),
	}, Types: map[string]Type{
		"ClientTrace":	TypeOf((*httptrace.ClientTrace)(nil)).Elem(),
		"DNSDoneInfo":	TypeOf((*httptrace.DNSDoneInfo)(nil)).Elem(),
		"DNSStartInfo":	TypeOf((*httptrace.DNSStartInfo)(nil)).Elem(),
		"GotConnInfo":	TypeOf((*httptrace.GotConnInfo)(nil)).Elem(),
		"WroteRequestInfo":	TypeOf((*httptrace.WroteRequestInfo)(nil)).Elem(),
	}, 
	}
}
