package smtModal

type Button struct {
  Color       string
  Text        string
  Icon        string
}

type Image struct {
  Url         string
  Size        string
}

type Icon struct {
  Name        string
  Size        string
}

type ModalProfile struct {
  Title       string
  Image       Image
  Header      string
  Text        string
  Ok          Button
  Cancel      Button
}

func ( ModalAStt ModalProfile ) Get () string {
  return "./Libraries/src/smt/smtModal/modalProfile.tpl"
}

type ModalBasic struct {
  Title       string
  Icon        Icon
  Description string
  Ok          Button
  Cancel      Button
}

func ( ModalAStt ModalBasic ) Get () string {
  return "./Libraries/src/smt/smtModal/modalBasic.tpl"
}