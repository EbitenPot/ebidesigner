unit Mainform;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ExtCtrls,
  ActnList, ComCtrls;

type

  { TMainform }

  TMainform = class(TForm)
    Titlename: TEdit;
    Label1: TLabel;
    StartButton: TButton;
    Button2: TButton;
    GroupBox1: TGroupBox;
    GroupBox2: TGroupBox;
    LeftBox: TGroupBox;
    MiddleBox: TGroupBox;
    PageControl: TPageControl;
    PageControl1: TPageControl;
    RightBox: TGroupBox;
    TabSheet1: TTabSheet;
    TabSheet2: TTabSheet;
    TabSheet3: TTabSheet;
    TabSheet4: TTabSheet;
    TabSheet5: TTabSheet;
    TabSheet6: TTabSheet;
    procedure StartButtonClick(Sender: TObject);
    procedure FormCreate(Sender: TObject);
    procedure PageControl1Change(Sender: TObject);
  private

  public

  end;

var
  Mainform: TMainform;

implementation

{$R *.lfm}

{ TMainform }

procedure TMainform.FormCreate(Sender: TObject);
begin

end;

procedure TMainform.PageControl1Change(Sender: TObject);
begin

end;

procedure TMainform.StartButtonClick(Sender: TObject);
begin

end;

end.

