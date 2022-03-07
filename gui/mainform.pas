unit Mainform;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs;

type

  { TMainform }

  TMainform = class(TForm)
    procedure FormCreate(Sender: TObject);
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

end.

