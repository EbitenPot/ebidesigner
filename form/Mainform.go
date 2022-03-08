// 由res2go IDE插件自动生成，不要编辑。
package form

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TMainform struct {
    *vcl.TForm
    LeftBox      *vcl.TGroupBox
    RightBox     *vcl.TGroupBox
    MiddleBox    *vcl.TGroupBox
    PageControl  *vcl.TPageControl
    TabSheet1    *vcl.TTabSheet
    TabSheet2    *vcl.TTabSheet
    TabSheet3    *vcl.TTabSheet
    GroupBox1    *vcl.TGroupBox
    StartButton  *vcl.TButton
    PageControl1 *vcl.TPageControl
    TabSheet4    *vcl.TTabSheet
    GroupBox2    *vcl.TGroupBox
    Button2      *vcl.TButton `events:"OnStartButtonClick"`
    TabSheet5    *vcl.TTabSheet
    TabSheet6    *vcl.TTabSheet
    Titlename    *vcl.TEdit
    Label1       *vcl.TLabel

    //::private::
    TMainformFields
}

var Mainform *TMainform




// vcl.Application.CreateForm(&Mainform)

func NewMainform(owner vcl.IComponent) (root *TMainform)  {
    vcl.CreateResForm(owner, &root)
    return
}

//go:embed resources/Mainform.gfm
var mainformBytes []byte

// 注册Form资源  
var _ = vcl.RegisterFormResource(Mainform, &mainformBytes)
