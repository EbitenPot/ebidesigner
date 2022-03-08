unit ebitenform;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, Menus;

type

  { TEbitenForm }

  TEbitenForm = class(TForm)
    procedure FormClose(Sender: TObject; var CloseAction: TCloseAction);
    procedure FormCreate(Sender: TObject);
    procedure FormShow(Sender: TObject);
  private

  public

  end;

var
  EbitenForm: TEbitenForm;

implementation

{$R *.lfm}

{ TEbitenForm }

procedure TEbitenForm.FormClose(Sender: TObject; var CloseAction: TCloseAction);
begin

end;

procedure TEbitenForm.FormCreate(Sender: TObject);
begin

end;

procedure TEbitenForm.FormShow(Sender: TObject);
begin

end;

end.

