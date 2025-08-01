// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.924
package dashboard

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/bastean/codexgo/v4/internal/app/server/component/layout"
	"github.com/bastean/codexgo/v4/internal/app/server/component/scripts"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
)

func PageInit() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_PageInit_cff3`,
		Function: `function __templ_PageInit_cff3(){$(".ui.container")
        .transition("fade in", "3s")
    ;

    $(".ui.menu .right .dropdown")
        .dropdown()
    ;

    $(".ui.menu .right .header .icon")
        .popup()
    ;

    $(".ui.blue.nag")
        .nag({
            key: "account-confirmation",
            value: true
        })
    ;

    $(".ui.green.nag")
        .nag({
            key: "account-confirmed",
            value: true
        })
    ;
}`,
		Call:       templ.SafeScript(`__templ_PageInit_cff3`),
		CallInline: templ.SafeScriptInline(`__templ_PageInit_cff3`),
	}
}

func Logout() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_Logout_55ac`,
		Function: `function __templ_Logout_55ac(){Storage.ClearSession();
    window.location.replace("/");
}`,
		Call:       templ.SafeScript(`__templ_Logout_55ac`),
		CallInline: templ.SafeScriptInline(`__templ_Logout_55ac`),
	}
}

func Page(attributes *read.ResponseAttributes) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"ui centered grid container\" style=\"height: 100%;\"><div class=\"sixteen wide column\"><div class=\"ui inverted secondary menu\"><div class=\"header item\"><a href=\"https://github.com/bastean/codexgo\" rel=\"noopener noreferrer\" target=\"_blank\"><img class=\"ui small image\" src=\"/public/static/assets/logo.png\"></a></div><div class=\"right menu\"><h4 class=\"ui inverted header\"><img class=\"ui avatar image\" src=\"/public/static/assets/favicon-512.png\"><div class=\"content\"><div class=\"ui inverted dropdown\"><div class=\"text\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(attributes.Username)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/app/server/component/page/dashboard/page.dashboard.templ`, Line: 59, Col: 48}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div><i class=\"dropdown icon\"></i><div class=\"menu\"><div class=\"header\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			switch attributes.Verified {
			case true:
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<i class=\"inverted green check circle icon\" data-variation=\"green\" data-content=\"Verified\"></i> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			default:
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<i class=\"inverted red times circle icon\" data-variation=\"red\" data-content=\"Unverified\"></i> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "Account</div><div class=\"divider\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, DeleteFormShow())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<div class=\"item\" onclick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 templ.ComponentScript = DeleteFormShow()
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\"><i class=\"trash icon\"></i> Delete</div><div class=\"divider\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, Logout())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<div class=\"item\" onclick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 templ.ComponentScript = Logout()
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var5.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\"><i class=\"sign out alt icon\"></i> Sign out</div></div></div></div></h4></div></div></div><div class=\"sixteen wide mobile eight wide computer column\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = UpdateForm(attributes.Email, attributes.Username).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = DeleteForm().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			switch attributes.Verified {
			case true:
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<div class=\"ui looping pulsating transition green bottom fixed nag\"><div class=\"title\"><i class=\"check circle icon\"></i> Account Confirmed</div><i class=\"close icon\"></i></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			default:
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<div class=\"ui looping pulsating transition blue bottom fixed nag\"><div class=\"title\"><i class=\"check circle icon\"></i> Account Confirmation</div><div>Link sent. Please check your inbox</div><i class=\"close icon\"></i></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = PageInit().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = layout.Index(scripts.Head{}, scripts.Body{}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
