// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"
import "github.com/nint8835/netenvelope/pkg/database/queries"
import "github.com/nint8835/netenvelope/pkg/server/ui"

func Vlans(vlans []queries.Vlan) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"list-inside list-disc\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, vlan := range vlans {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", vlan.Tag))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/server/ui/pages/vlans.templ`, Line: 11, Col: 40}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if vlan.Name.Valid {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var4 string
					templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(" - %s", vlan.Name.String))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/server/ui/pages/vlans.templ`, Line: 13, Col: 52}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul><div class=\"flex justify-center\"><form action=\"/vlans\" method=\"post\" class=\"flex w-96 max-w-full flex-col\"><div class=\"mb-4 flex flex-col\"><label for=\"tag\">Tag</label> <input type=\"number\" name=\"tag\" class=\"rounded-md p-2 text-zinc-900 outline-none ring-blue-500 transition-all focus:ring-2\" required></div><div class=\"mb-4 flex flex-col\"><label for=\"name\">Name</label> <input type=\"text\" class=\"rounded-md p-2 text-zinc-900 outline-none ring-blue-500 transition-all focus:ring-2\" name=\"name\"></div><button type=\"submit\" class=\"mt-8 rounded-md bg-blue-500 py-2 transition-colors hover:bg-blue-600\">Add</button></form></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = ui.Layout(" - VLANs").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
