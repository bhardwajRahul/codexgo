// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.856
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

const (
	LoginFormTagID = "login-form"
)

func LoginFormInit(formTagID string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_LoginFormInit_7de9`,
		Function: `function __templ_LoginFormInit_7de9(formTagID){$(` + "`" + `#${formTagID}` + "`" + `)
        .form({
            on: "blur",
            inline: true,
            preventLeaving: true,
            keyboardShortcuts: false,
            fields: {
                Login: {
                    rules: [
                        {
                            type: "notEmpty"
                        }
                    ]
                },
                Password: {
                    rules: [
                        {
                            type: "size[8..64]"
                        },
                        {
                            type: "regExp[/^.*[^0-9].*$/]",
                            prompt: "{name} cannot be only numbers"
                        }
                    ]
                }
            }
        })
        .api({
            action: "user_login", 
            method: "POST",
            beforeSend: function(settings) {
                let login = settings.data.Login;

                if (_.includes(login, "@")) {
                    settings.data.Email = login;
                } else {
                    settings.data.Username = login;
                }

                settings.data = JSON.stringify(settings.data);
        
                return settings;
            },
            onSuccess: function(response, element, xhr) {
                $.toast({
                    class: "success",
                    message: response.Message,
                    showProgress: "top"
                });

                _.delay(function() {
                    window.location.replace("/dashboard");
                }, 1000);
            },
            onFailure: function(response, element, xhr) {
                $.toast({
                    class: "error",
                    message: response.Message,
                    showProgress: "top"
                });
            }
        })
    ;
}`,
		Call:       templ.SafeScript(`__templ_LoginFormInit_7de9`, formTagID),
		CallInline: templ.SafeScriptInline(`__templ_LoginFormInit_7de9`, formTagID),
	}
}

func LoginForm() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<form id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(LoginFormTagID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/app/server/component/page/home/form.login.templ`, Line: 74, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\" class=\"ui inverted form\"><h1 class=\"ui dividing inverted header\">Sign in to your account<div class=\"sub header\">Don't have an account? ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, ShowTab(RegisterTabTagID))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<a style=\"cursor: pointer;\" onclick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.ComponentScript = ShowTab(RegisterTabTagID)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\">Sign up</a> | ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, ShowModal(ForgotModalTagID))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<a style=\"cursor: pointer;\" onclick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 templ.ComponentScript = ShowModal(ForgotModalTagID)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\">Forgot Password?</a></div></h1><div class=\"required field\"><label>Email/Username</label><div class=\"ui inverted transparent left icon input\"><i class=\"sign in alt icon\"></i> <input type=\"text\" placeholder=\"Email/Username\" name=\"Login\"></div></div><div class=\"required field\"><label>Password</label><div class=\"ui inverted transparent left icon input\"><i class=\"lock icon\"></i> <input type=\"password\" placeholder=\"Password\" name=\"Password\"></div></div><div class=\"ui divider\"></div><button class=\"ui fluid primary submit button\">Sign in</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = LoginFormInit(LoginFormTagID).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
