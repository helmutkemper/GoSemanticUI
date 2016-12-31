package smtStep

import (
  "smt/smtIcon"
  "def"
  "smt/html"
)

type Step struct {
  Title         string
  Description   string
  Icon          smtIcon.Icon
  States        def.A
  Link          string
}

func ( StepAStt *Step ) Get() string {

  classTitle := html.Attribute{
    Name: "class",
    Values: def.String{ "title" },
  }

  divTitle := html.Tag{
    Name: "div",
    Attributes: html.Attributes{ classTitle },
    Value: def.String{ StepAStt.Title },
  }

  classDescription := html.Attribute{
    Name: "class",
    Values: def.String{ "description" },
  }

  divDescription := html.Tag{
    Name: "div",
    Attributes: html.Attributes{ classDescription },
    Value: def.String{ StepAStt.Description },
  }

  classContent := html.Attribute{
    Name: "class",
    Values: def.String{ "content" },
  }

  divContent := html.Tag{
    Name: "div",
    Attributes: html.Attributes{ classContent },
    Value: def.String{ divTitle.GetClosed(), divDescription.GetClosed() },
  }

  classStep := html.Attribute{
    Name: "class",
  }
  classStep.Values.Add( StepAStt.States, "step" )

  divStep := html.Tag{
    Name: "div",
    Attributes: html.Attributes{ classStep },
    Value: def.String{ StepAStt.Icon.Get(), divContent.GetClosed() },
  }

  return divStep.GetClosed()
}