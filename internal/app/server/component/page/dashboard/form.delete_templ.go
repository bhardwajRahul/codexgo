// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package dashboard

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

const (
	DeleteFormTagID = "delete"
)

func DeleteFormInit(formTagID string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_DeleteFormInit_b376`,
		Function: `function __templ_DeleteFormInit_b376(formTagID){$(` + "`" + `#${formTagID}` + "`" + `)
        .form({
            on: "blur",
            inline: true,
            preventLeaving: true,
            keyboardShortcuts: false,
            fields: {
                Password: {
                    rules: [
                        {
                            type: "notEmpty",
                            prompt: "{name} is required to delete the account"
                        },
                        {
                            type: "size[8..64]"
                        },
                        {
                            type: "regExp[/^.*[^0-9].*$/]",
                            prompt: "{name} cannot be only numbers"
                        }
                    ]
                },
                ConfirmPassword: {
                    rules: [
                        {
                            type: "match[Password]"
                        }
                    ]
                }
            }
        })
        .api({
            action: "user_delete", 
            method: "DELETE",
            onSuccess: function(response, element, xhr) {
                $.toast({
                    class: "success",
                    message: response.Message,
                    showProgress: "top"
                });

                 _.delay(function() {
                    Storage.Clear();
                    window.location.replace("/");
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
		Call:       templ.SafeScript(`__templ_DeleteFormInit_b376`, formTagID),
		CallInline: templ.SafeScriptInline(`__templ_DeleteFormInit_b376`, formTagID),
	}
}

func DeleteFormShow() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_DeleteFormShow_81f8`,
		Function: `function __templ_DeleteFormShow_81f8(){$(".ui.mini.modal").modal("show");
}`,
		Call:       templ.SafeScript(`__templ_DeleteFormShow_81f8`),
		CallInline: templ.SafeScriptInline(`__templ_DeleteFormShow_81f8`),
	}
}

func DeleteForm() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"ui mini inverted modal\"><div class=\"header\">Delete Your Account</div><div class=\"content\"><form id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(DeleteFormTagID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/app/server/component/page/dashboard/form.delete.templ`, Line: 73, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\" class=\"ui inverted form\"><div class=\"required field\"><label>Password</label><div class=\"ui inverted transparent left icon input\"><i class=\"lock icon\"></i> <input type=\"password\" placeholder=\"Password\" name=\"Password\"></div></div><div class=\"required field\"><label>Confirm Password</label><div class=\"ui inverted transparent left icon input\"><i class=\"redo icon\"></i> <input type=\"password\" placeholder=\"Confirm Password\" name=\"ConfirmPassword\"></div></div><div class=\"ui divider\"></div><button class=\"ui fluid positive submit button\">Approve</button></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = DeleteFormInit(DeleteFormTagID).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
