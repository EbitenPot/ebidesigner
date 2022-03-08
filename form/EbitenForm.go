// 由res2go IDE插件自动生成，不要编辑。
package form

import (
	_ "embed"
	"github.com/ying32/govcl/vcl"
)

type TEbitenForm struct {
	*vcl.TForm

	//::private::
	TEbitenFormFields
}

var EbitenForm *TEbitenForm

// vcl.Application.CreateForm(&EbitenForm)

func NewEbitenForm(owner vcl.IComponent) (root *TEbitenForm) {
	vcl.CreateResForm(owner, &root)
	return
}

//go:embed resources/EbitenForm.gfm
var ebitenFormBytes []byte

// 注册Form资源
var _ = vcl.RegisterFormResource(EbitenForm, &ebitenFormBytes)
