// 由res2go IDE插件自动生成，不要编辑。
package form

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TMainform struct {
    *vcl.TForm
    LeftBox      *vcl.TGroupBox
    PageControl  *vcl.TPageControl
    TabSheet1    *vcl.TTabSheet
    GroupBox1    *vcl.TGroupBox
    StartButton  *vcl.TButton
    Label1       *vcl.TLabel
    Label2       *vcl.TLabel
    TabSheet2    *vcl.TTabSheet
    TabSheet3    *vcl.TTabSheet
    RightBox     *vcl.TGroupBox
    PageControl1 *vcl.TPageControl
    TabSheet4    *vcl.TTabSheet
    GroupBox2    *vcl.TGroupBox
    Button2      *vcl.TButton `events:"OnStartButtonClick"`
    TabSheet5    *vcl.TTabSheet
    TabSheet6    *vcl.TTabSheet
    MiddleBox    *vcl.TGroupBox
    MainMenuUp   *vcl.TMainMenu
    MenuItem3    *vcl.TMenuItem
    MenuItem4    *vcl.TMenuItem
    MenuItem5    *vcl.TMenuItem
    MenuItem6    *vcl.TMenuItem
    MenuItem7    *vcl.TMenuItem
    MenuItem8    *vcl.TMenuItem
    MenuItem9    *vcl.TMenuItem
    MenuItem10   *vcl.TMenuItem
    MenuItem11   *vcl.TMenuItem
    MenuItem12   *vcl.TMenuItem
    MenuItem13   *vcl.TMenuItem
    MenuItem14   *vcl.TMenuItem
    MenuItem15   *vcl.TMenuItem
    MenuItem2    *vcl.TMenuItem
    MenuItem16   *vcl.TMenuItem
    MenuItem17   *vcl.TMenuItem
    MenuItem18   *vcl.TMenuItem
    MenuItem19   *vcl.TMenuItem
    MenuItem20   *vcl.TMenuItem
    MenuItem21   *vcl.TMenuItem
    MenuItem22   *vcl.TMenuItem
    MenuItem23   *vcl.TMenuItem
    MenuItem24   *vcl.TMenuItem
    MenuItem25   *vcl.TMenuItem
    MenuItem26   *vcl.TMenuItem
    MenuItem27   *vcl.TMenuItem
    MenuItem28   *vcl.TMenuItem
    MenuItem29   *vcl.TMenuItem
    MenuItem30   *vcl.TMenuItem
    MenuItem31   *vcl.TMenuItem
    MenuItem32   *vcl.TMenuItem
    MenuItem33   *vcl.TMenuItem
    MenuItem34   *vcl.TMenuItem
    MenuItem35   *vcl.TMenuItem
    MenuItem36   *vcl.TMenuItem
    MenuItem37   *vcl.TMenuItem
    MenuItem38   *vcl.TMenuItem
    MenuItem39   *vcl.TMenuItem
    MenuItem40   *vcl.TMenuItem
    MenuItem41   *vcl.TMenuItem
    GroupBox3    *vcl.TGroupBox
    Xxxxxx       *vcl.TScrollBar
    Yyyyyyy      *vcl.TScrollBar

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
