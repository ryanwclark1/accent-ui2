// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func NavBarSoundsGreetings() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><div><button type=\"button\" class=\"text-gray-700 hover:text-indigo-600 hover:bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold\" aria-controls=\"sub-menu-1\" aria-expanded=\"false\"><svg class=\"h-6 w-6 shrink-0 text-gray-400 group-hover:text-indigo-600\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" aria-hidden=\"false\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"m9 9 10.5-3m0 6.553v3.75a2.25 2.25 0 0 1-1.632 2.163l-1.32.377a1.803 1.803 0 1 1-.99-3.467l2.31-.66a2.25 2.25 0 0 0 1.632-2.163Zm0 0V2.25L9 5.25v10.303m0 0v3.75a2.25 2.25 0 0 1-1.632 2.163l-1.32.377a1.803 1.803 0 0 1-.99-3.467l2.31-.66A2.25 2.25 0 0 0 9 15.553Z\"></path></svg> Sounds<!-- Expanded: \"rotate-90 text-gray-500\", Collapsed: \"text-gray-400\" --><svg class=\"text-gray-400 ml-auto h-5 w-5 shrink-0\" viewBox=\"0 0 20 20\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.21 14.77a.75.75 0 01.02-1.06L11.168 10 7.23 6.29a.75.75 0 111.04-1.08l4.5 4.25a.75.75 0 010 1.08l-4.5 4.25a.75.75 0 01-1.06-.02z\" clip-rule=\"evenodd\"></path></svg></button><!-- Expandable link section, show/hide based on state. --><ul class=\"mt-1 px-2\" id=\"sub-menu-1\" x-show=\"open\" style=\"display: none;\"><li><!-- 44px --><a href=\"#\" class=\"hover:bg-gray-50 block rounded-md py-2 pr-2 pl-9 text-sm leading-6 text-gray-700\">Engineering</a></li><li><!-- 44px --><a href=\"#\" class=\"hover:bg-gray-50 block rounded-md py-2 pr-2 pl-9 text-sm leading-6 text-gray-700\">Human Resources</a></li><li><!-- 44px --><a href=\"#\" class=\"hover:bg-gray-50 block rounded-md py-2 pr-2 pl-9 text-sm leading-6 text-gray-700\">Customer Success</a></li></ul></div></li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
