package smt

type Social int

const(
  BUTTON_SOCIAL_FACEBOOK Social = iota
  BUTTON_SOCIAL_TWITTER
  BUTTON_SOCIAL_GOOGLE
  BUTTON_SOCIAL_VK
  BUTTON_SOCIAL_LINKEDIN
  BUTTON_SOCIAL_INSTAGRAM
  BUTTON_SOCIAL_YOUTUBE
)

var socials = [...]string{
  "facebook",
  "twitter",
  "google",
  "vk",
  "linkedin",
  "instagram",
  "youtube",
}
func (m Social) String() string {
  return socials[ m ]
}

func (m Social) Content() string {
  switch m {
  case 0:
    return `<i class="facebook icon"></i>Facebook`
  case 1:
    return `<i class="twitter icon"></i>Twitter`
  case 2:
    return `<i class="google plus icon"></i>Google Plus`
  case 3:
    return `<i class="vk icon"></i>VK`
  case 4:
    return `<i class="linkedin icon"></i>LinkedIn`
  case 5:
    return `<i class="instagram icon"></i>Instagram`
  case 6:
    return `<i class="youtube icon"></i>YouTube`
  }
  return ``
}