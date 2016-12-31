package smt

type IconName int

const (
  ICON_NAME_ADD_TO_CALENDAR IconName = iota
  ICON_NAME_ALARM_OUTLINE
  ICON_NAME_ALARM_MUTE_OUTLINE
  ICON_NAME_ALARM_MUTE
  ICON_NAME_ALARM
  ICON_NAME_AT
  ICON_NAME_BROWSER
  ICON_NAME_BUG
  ICON_NAME_CALENDAR_OUTLINE
  ICON_NAME_CALENDAR
  ICON_NAME_CHECKED_CALENDAR
  ICON_NAME_CLOUD
  ICON_NAME_CODE
  ICON_NAME_COMMENT_OUTLINE
  ICON_NAME_COMMENT
  ICON_NAME_COMMENTS_OUTLINE
  ICON_NAME_COMMENTS
  ICON_NAME_COPYRIGHT
  ICON_NAME_CREATIVE_COMMONS
  ICON_NAME_DASHBOARD
  ICON_NAME_DELETE_CALENDAR
  ICON_NAME_EXTERNAL_SQUARE
  ICON_NAME_EXTERNAL
  ICON_NAME_EYEDROPPER
  ICON_NAME_FEED
  ICON_NAME_FIND
  ICON_NAME_HAND_POINTER
  ICON_NAME_HASHTAG
  ICON_NAME_HEARTBEAT
  ICON_NAME_HISTORY
  ICON_NAME_HOURGLASS_EMPTY
  ICON_NAME_HOURGLASS_END
  ICON_NAME_HOURGLASS_FULL
  ICON_NAME_HOURGLASS_HALF
  ICON_NAME_HOURGLASS_START
  ICON_NAME_IDEA
  ICON_NAME_IMAGE
  ICON_NAME_INBOX
  ICON_NAME_INDUSTRY
  ICON_NAME_LAB
  ICON_NAME_MAIL_OUTLINE
  ICON_NAME_MAIL_SQUARE
  ICON_NAME_MAIL
  ICON_NAME_MOUSE_POINTER
  ICON_NAME_OPTIONS
  ICON_NAME_PAINT_BRUSH
  ICON_NAME_PAYMENT
  ICON_NAME_PERCENT
  ICON_NAME_PRIVACY
  ICON_NAME_PROTECT
  ICON_NAME_REGISTERED
  ICON_NAME_REMOVE_FROM_CALENDAR
  ICON_NAME_SEARCH
  ICON_NAME_SETTING
  ICON_NAME_SETTINGS
  ICON_NAME_SHOP
  ICON_NAME_SHOPPING_BAG
  ICON_NAME_SHOPPING_BASKET
  ICON_NAME_SIGNAL
  ICON_NAME_SITEMAP
  ICON_NAME_TAG
  ICON_NAME_TAGS
  ICON_NAME_TASKS
  ICON_NAME_TERMINAL
  ICON_NAME_TEXT_TELEPHONE
  ICON_NAME_TICKET
  ICON_NAME_TRADEMARK
  ICON_NAME_TROPHY
  ICON_NAME_WIFI
  ICON_NAME_ADD_TO_CART
  ICON_NAME_ADD_USER
  ICON_NAME_ADJUST
  ICON_NAME_ARCHIVE
  ICON_NAME_BAN
  ICON_NAME_BOOKMARK
  ICON_NAME_CALL
  ICON_NAME_CALL_SQUARE
  ICON_NAME_CLONE
  ICON_NAME_CLOUD_DOWNLOAD
  ICON_NAME_CLOUD_UPLOAD
  ICON_NAME_TALK
  ICON_NAME_TALK_OUTLINE
  ICON_NAME_COMPRESS
  ICON_NAME_CONFIGURE
  ICON_NAME_DOWNLOAD
  ICON_NAME_EDIT
  ICON_NAME_ERASE
  ICON_NAME_EXCHANGE
  ICON_NAME_EXPAND
  ICON_NAME_EXTERNAL_SHARE
  ICON_NAME_FILTER
  ICON_NAME_HIDE
  ICON_NAME_IN_CART
  ICON_NAME_LOCK
  ICON_NAME_MAIL_FORWARD
  ICON_NAME_OBJECT_GROUP
  ICON_NAME_OBJECT_UNGROUP
  ICON_NAME_PIN
  ICON_NAME_PRINT
  ICON_NAME_RANDOM
  ICON_NAME_RECYCLE
  ICON_NAME_REFRESH
  ICON_NAME_REMOVE_BOOKMARK
  ICON_NAME_REMOVE_USER
  ICON_NAME_REPEAT
  ICON_NAME_REPLY_ALL
  ICON_NAME_REPLY
  ICON_NAME_RETWEET
  ICON_NAME_SEND
  ICON_NAME_SEND_OUTLINE
  ICON_NAME_SHARE_ALTERNATE
  ICON_NAME_SHARE_ALTERNATE_SQUARE
  ICON_NAME_SHARE
  ICON_NAME_SHARE_SQUARE
  ICON_NAME_SIGN_IN
  ICON_NAME_SIGN_OUT
  ICON_NAME_THEME
  ICON_NAME_TRANSLATE
  ICON_NAME_UNDO
  ICON_NAME_UNHIDE
  ICON_NAME_UNLOCK_ALTERNATE
  ICON_NAME_UNLOCK
  ICON_NAME_UPLOAD
  ICON_NAME_WAIT
  ICON_NAME_WIZARD
  ICON_NAME_WRITE
  ICON_NAME_WRITE_SQUARE
  ICON_NAME_ANNOUNCEMENT
  ICON_NAME_BIRTHDAY
  ICON_NAME_HELP_CIRCLE
  ICON_NAME_HELP
  ICON_NAME_INFO_CIRCLE
  ICON_NAME_INFO
  ICON_NAME_WARNING_CIRCLE
  ICON_NAME_WARNING
  ICON_NAME_WARNING_SIGN
  ICON_NAME_CHILD
  ICON_NAME_DOCTOR
  ICON_NAME_HANDICAP
  ICON_NAME_SPY
  ICON_NAME_STUDENT
  ICON_NAME_USER
  ICON_NAME_USERS
  ICON_NAME_FEMALE
  ICON_NAME_GAY
  ICON_NAME_GENDERLESS
  ICON_NAME_HETEROSEXUAL
  ICON_NAME_INTERGENDER
  ICON_NAME_LESBIAN
  ICON_NAME_MALE
  ICON_NAME_MAN
  ICON_NAME_NEUTER
  ICON_NAME_NON_BINARY_TRANSGENDER
  ICON_NAME_OTHER_GENDER_HORIZONTAL
  ICON_NAME_OTHER_GENDER
  ICON_NAME_OTHER_GENDER_VERTICAL
  ICON_NAME_TRANSGENDER
  ICON_NAME_WOMAN
  ICON_NAME_BLOCK_LAYOUT
  ICON_NAME_CROP
  ICON_NAME_GRID_LAYOUT
  ICON_NAME_LIST_LAYOUT
  ICON_NAME_MAXIMIZE
  ICON_NAME_RESIZE_HORIZONTAL
  ICON_NAME_RESIZE_VERTICAL
  ICON_NAME_ZOOM
  ICON_NAME_ZOOM_OUT
  ICON_NAME_ANCHOR
  ICON_NAME_BAR
  ICON_NAME_BOMB
  ICON_NAME_BOOK
  ICON_NAME_BULLSEYE
  ICON_NAME_CALCULATOR
  ICON_NAME_COCKTAIL
  ICON_NAME_DIAMOND
  ICON_NAME_FAX
  ICON_NAME_FIRE_EXTINGUISHER
  ICON_NAME_FIRE
  ICON_NAME_FLAG_CHECKERED
  ICON_NAME_FLAG
  ICON_NAME_FLAG_OUTLINE
  ICON_NAME_GIFT
  ICON_NAME_HAND_LIZARD
  ICON_NAME_HAND_PEACE
  ICON_NAME_HAND_PAPER
  ICON_NAME_HAND_ROCK
  ICON_NAME_HAND_SCISSORS
  ICON_NAME_HAND_SPOCK
  ICON_NAME_LAW
  ICON_NAME_LEAF
  ICON_NAME_LEGAL
  ICON_NAME_LEMON
  ICON_NAME_LIFE_RING
  ICON_NAME_LIGHTNING
  ICON_NAME_MAGNET
  ICON_NAME_MONEY
  ICON_NAME_MOON
  ICON_NAME_PLANE
  ICON_NAME_PUZZLE
  ICON_NAME_ROAD
  ICON_NAME_ROCKET
  ICON_NAME_SHIPPING
  ICON_NAME_SOCCER
  ICON_NAME_STICKY_NOTE
  ICON_NAME_STICKY_NOTE_OUTLINE
  ICON_NAME_SUITCASE
  ICON_NAME_SUN
  ICON_NAME_TRAVEL
  ICON_NAME_TREATMENT
  ICON_NAME_UMBRELLA
  ICON_NAME_WORLD
  ICON_NAME_ASTERISK
  ICON_NAME_CERTIFICATE
  ICON_NAME_CIRCLE
  ICON_NAME_CIRCLE_NOTCHED
  ICON_NAME_CIRCLE_THIN
  ICON_NAME_CROSSHAIRS
  ICON_NAME_CUBE
  ICON_NAME_CUBES
  ICON_NAME_ELLIPSIS_HORIZONTAL
  ICON_NAME_ELLIPSIS_VERTICAL
  ICON_NAME_QUOTE_LEFT
  ICON_NAME_QUOTE_RIGHT
  ICON_NAME_SPINNER
  ICON_NAME_SQUARE
  ICON_NAME_SQUARE_OUTLINE
  ICON_NAME_ADD_CIRCLE
  ICON_NAME_ADD_SQUARE
  ICON_NAME_CHECK_CIRCLE
  ICON_NAME_CHECK_CIRCLE_OUTLINE
  ICON_NAME_CHECK_SQUARE
  ICON_NAME_CHECKMARK_BOX
  ICON_NAME_CHECKMARK
  ICON_NAME_MINUS_CIRCLE
  ICON_NAME_MINUS
  ICON_NAME_MINUS_SQUARE
  ICON_NAME_MINUS_SQUARE_OUTLINE
  ICON_NAME_MOVE
  ICON_NAME_PLUS
  ICON_NAME_PLUS_SQUARE_OUTLINE
  ICON_NAME_RADIO
  ICON_NAME_REMOVE_CIRCLE
  ICON_NAME_REMOVE_CIRCLE_OUTLINE
  ICON_NAME_REMOVE
  ICON_NAME_SELECTED_RADIO
  ICON_NAME_TOGGLE_OFF
  ICON_NAME_TOGGLE_ON
  ICON_NAME_AREA_CHART
  ICON_NAME_BAR_CHART
  ICON_NAME_CAMERA_RETRO
  ICON_NAME_FILM
  ICON_NAME_LINE_CHART
  ICON_NAME_NEWSPAPER
  ICON_NAME_PHOTO
  ICON_NAME_PIE_CHART
  ICON_NAME_SOUND
  ICON_NAME_ANGLE_DOUBLE_DOWN
  ICON_NAME_ANGLE_DOUBLE_LEFT
  ICON_NAME_ANGLE_DOUBLE_RIGHT
  ICON_NAME_ANGLE_DOUBLE_UP
  ICON_NAME_ANGLE_DOWN
  ICON_NAME_ANGLE_LEFT
  ICON_NAME_ANGLE_RIGHT
  ICON_NAME_ANGLE_UP
  ICON_NAME_ARROW_CIRCLE_DOWN
  ICON_NAME_ARROW_CIRCLE_LEFT
  ICON_NAME_ARROW_CIRCLE_OUTLINE_DOWN
  ICON_NAME_ARROW_CIRCLE_OUTLINE_LEFT
  ICON_NAME_ARROW_CIRCLE_OUTLINE_RIGHT
  ICON_NAME_ARROW_CIRCLE_OUTLINE_UP
  ICON_NAME_ARROW_CIRCLE_RIGHT
  ICON_NAME_ARROW_CIRCLE_UP
  ICON_NAME_ARROW_DOWN
  ICON_NAME_ARROW_LEFT
  ICON_NAME_ARROW_RIGHT
  ICON_NAME_ARROW_UP
  ICON_NAME_CARET_DOWN
  ICON_NAME_CARET_LEFT
  ICON_NAME_CARET_RIGHT
  ICON_NAME_CARET_UP
  ICON_NAME_CHEVRON_CIRCLE_DOWN
  ICON_NAME_CHEVRON_CIRCLE_LEFT
  ICON_NAME_CHEVRON_CIRCLE_RIGHT
  ICON_NAME_CHEVRON_CIRCLE_UP
  ICON_NAME_CHEVRON_DOWN
  ICON_NAME_CHEVRON_LEFT
  ICON_NAME_CHEVRON_RIGHT
  ICON_NAME_CHEVRON_UP
  ICON_NAME_LONG_ARROW_DOWN
  ICON_NAME_LONG_ARROW_LEFT
  ICON_NAME_LONG_ARROW_RIGHT
  ICON_NAME_LONG_ARROW_UP
  ICON_NAME_POINTING_DOWN
  ICON_NAME_POINTING_LEFT
  ICON_NAME_POINTING_RIGHT
  ICON_NAME_POINTING_UP
  ICON_NAME_TOGGLE_DOWN
  ICON_NAME_TOGGLE_LEFT
  ICON_NAME_TOGGLE_RIGHT
  ICON_NAME_TOGGLE_UP
  ICON_NAME_MOBILE
  ICON_NAME_TABLET
  ICON_NAME_BATTERY_EMPTY
  ICON_NAME_BATTERY_FULL
  ICON_NAME_BATTERY_LOW
  ICON_NAME_BATTERY_MEDIUM
  ICON_NAME_DESKTOP
  ICON_NAME_DISK_OUTLINE
  ICON_NAME_GAME
  ICON_NAME_HIGH_BATTERY
  ICON_NAME_KEYBOARD
  ICON_NAME_LAPTOP
  ICON_NAME_PLUG
  ICON_NAME_POWER
  ICON_NAME_FILE_ARCHIVE_OUTLINE
  ICON_NAME_FILE_AUDIO_OUTLINE
  ICON_NAME_FILE_CODE_OUTLINE
  ICON_NAME_FILE_EXCEL_OUTLINE
  ICON_NAME_FILE
  ICON_NAME_FILE_IMAGE_OUTLINE
  ICON_NAME_FILE_OUTLINE
  ICON_NAME_FILE_PDF_OUTLINE
  ICON_NAME_FILE_POWERPOINT_OUTLINE
  ICON_NAME_FILE_TEXT
  ICON_NAME_FILE_TEXT_OUTLINE
  ICON_NAME_FILE_VIDEO_OUTLINE
  ICON_NAME_FILE_WORD_OUTLINE
  ICON_NAME_FOLDER
  ICON_NAME_FOLDER_OPEN
  ICON_NAME_FOLDER_OPEN_OUTLINE
  ICON_NAME_FOLDER_OUTLINE
  ICON_NAME_LEVEL_DOWN
  ICON_NAME_LEVEL_UP
  ICON_NAME_TRASH
  ICON_NAME_TRASH_OUTLINE
  ICON_NAME_BARCODE
  ICON_NAME_BLUETOOTH_ALTERNATIVE
  ICON_NAME_BLUETOOTH
  ICON_NAME_CSS3
  ICON_NAME_DATABASE
  ICON_NAME_FORK
  ICON_NAME_HTML5
  ICON_NAME_OPENID
  ICON_NAME_QRCODE
  ICON_NAME_RSS
  ICON_NAME_RSS_SQUARE
  ICON_NAME_SERVER
  ICON_NAME_USB
  ICON_NAME_EMPTY_HEART
  ICON_NAME_EMPTY_STAR
  ICON_NAME_FROWN
  ICON_NAME_HEART
  ICON_NAME_MEH
  ICON_NAME_SMILE
  ICON_NAME_STAR_HALF_EMPTY
  ICON_NAME_STAR_HALF
  ICON_NAME_STAR
  ICON_NAME_THUMBS_DOWN
  ICON_NAME_THUMBS_OUTLINE_DOWN
  ICON_NAME_THUMBS_OUTLINE_UP
  ICON_NAME_THUMBS_UP
  ICON_NAME_BACKWARD
  ICON_NAME_CLOSED_CAPTIONING
  ICON_NAME_EJECT
  ICON_NAME_FAST_BACKWARD
  ICON_NAME_FAST_FORWARD
  ICON_NAME_FORWARD
  ICON_NAME_MUSIC
  ICON_NAME_MUTE
  ICON_NAME_PAUSE_CIRCLE
  ICON_NAME_PAUSE_CIRCLE_OUTLINE
  ICON_NAME_PAUSE
  ICON_NAME_PLAY
  ICON_NAME_RECORD
  ICON_NAME_STEP_BACKWARD
  ICON_NAME_STEP_FORWARD
  ICON_NAME_STOP_CIRCLE
  ICON_NAME_STOP_CIRCLE_OUTLINE
  ICON_NAME_STOP
  ICON_NAME_UNMUTE
  ICON_NAME_VIDEO_PLAY
  ICON_NAME_VIDEO_PLAY_OUTLINE
  ICON_NAME_VOLUME_DOWN
  ICON_NAME_VOLUME_OFF
  ICON_NAME_VOLUME_UP
  ICON_NAME_BICYCLE
  ICON_NAME_BUILDING
  ICON_NAME_BUILDING_OUTLINE
  ICON_NAME_BUS
  ICON_NAME_CAR
  ICON_NAME_COFFEE
  ICON_NAME_COMPASS
  ICON_NAME_EMERGENCY
  ICON_NAME_FIRST_AID
  ICON_NAME_FOOD
  ICON_NAME_H
  ICON_NAME_HOSPITAL
  ICON_NAME_HOTEL
  ICON_NAME_LOCATION_ARROW
  ICON_NAME_MAP
  ICON_NAME_MAP_OUTLINE
  ICON_NAME_MAP_PIN
  ICON_NAME_MAP_SIGNS
  ICON_NAME_MARKER
  ICON_NAME_MILITARY
  ICON_NAME_MOTORCYCLE
  ICON_NAME_PAW
  ICON_NAME_SHIP
  ICON_NAME_SPACE_SHUTTLE
  ICON_NAME_SPOON
  ICON_NAME_STREET_VIEW
  ICON_NAME_SUBWAY
  ICON_NAME_TAXI
  ICON_NAME_TRAIN
  ICON_NAME_TELEVISION
  ICON_NAME_TREE
  ICON_NAME_UNIVERSITY
  ICON_NAME_COLUMNS
  ICON_NAME_SORT_ALPHABET_ASCENDING
  ICON_NAME_SORT_ALPHABET_DESCENDING
  ICON_NAME_SORT_ASCENDING
  ICON_NAME_SORT_CONTENT_ASCENDING
  ICON_NAME_SORT_CONTENT_DESCENDING
  ICON_NAME_SORT_DESCENDING
  ICON_NAME_SORT
  ICON_NAME_SORT_NUMERIC_ASCENDING
  ICON_NAME_SORT_NUMERIC_DESCENDING
  ICON_NAME_TABLE
  ICON_NAME_ALIGN_CENTER
  ICON_NAME_ALIGN_JUSTIFY
  ICON_NAME_ALIGN_LEFT
  ICON_NAME_ALIGN_RIGHT
  ICON_NAME_ATTACH
  ICON_NAME_BOLD
  ICON_NAME_CONTENT
  ICON_NAME_COPY
  ICON_NAME_CUT
  ICON_NAME_FONT
  ICON_NAME_HEADER
  ICON_NAME_INDENT
  ICON_NAME_ITALIC
  ICON_NAME_LINKIFY
  ICON_NAME_LIST
  ICON_NAME_ORDERED_LIST
  ICON_NAME_OUTDENT
  ICON_NAME_PARAGRAPH
  ICON_NAME_PASTE
  ICON_NAME_SAVE
  ICON_NAME_STRIKETHROUGH
  ICON_NAME_SUBSCRIPT
  ICON_NAME_SUPERSCRIPT
  ICON_NAME_TEXT_CURSOR
  ICON_NAME_TEXT_HEIGHT
  ICON_NAME_TEXT_WIDTH
  ICON_NAME_UNDERLINE
  ICON_NAME_UNLINKIFY
  ICON_NAME_UNORDERED_LIST
  ICON_NAME_BITCOIN
  ICON_NAME_DOLLAR
  ICON_NAME_EURO
  ICON_NAME_LIRA
  ICON_NAME_POUND
  ICON_NAME_RUBLE
  ICON_NAME_RUPEE
  ICON_NAME_SHEKEL
  ICON_NAME_WON
  ICON_NAME_YEN
  ICON_NAME_AMERICAN_EXPRESS
  ICON_NAME_CREDIT_CARD_ALTERNATIVE
  ICON_NAME_DINERS_CLUB
  ICON_NAME_DISCOVER
  ICON_NAME_GOOGLE_WALLET
  ICON_NAME_JAPAN_CREDIT_BUREAU
  ICON_NAME_MASTERCARD
  ICON_NAME_PAYPAL_CARD
  ICON_NAME_PAYPAL
  ICON_NAME_STRIPE
  ICON_NAME_VISA
  ICON_NAME_WHEELCHAIR
  ICON_NAME_ASL_INTERPRETING
  ICON_NAME_ASSISTIVE_LISTENING_SYSTEMS
  ICON_NAME_AUDIO_DESCRIPTION
  ICON_NAME_BLIND
  ICON_NAME_BRAILLE
  ICON_NAME_DEAFNESS
  ICON_NAME_LOW_VISION
  ICON_NAME_SIGN_LANGUAGE
  ICON_NAME_UNIVERSAL_ACCESS
  ICON_NAME_VOLUME_CONTROL_PHONE
  ICON_NAME_FIVE_HUNDRED_PX
  ICON_NAME_ADN
  ICON_NAME_AMAZON
  ICON_NAME_ANDROID
  ICON_NAME_ANGELLIST
  ICON_NAME_APPLE
  ICON_NAME_BEHANCE
  ICON_NAME_BEHANCE_SQUARE
  ICON_NAME_BITBUCKET
  ICON_NAME_BITBUCKET_SQUARE
  ICON_NAME_BLACK_TIE
  ICON_NAME_BUYSELLADS
  ICON_NAME_CHROME
  ICON_NAME_CODEPEN
  ICON_NAME_CODIEPIE
  ICON_NAME_CONNECTDEVELOP
  ICON_NAME_CONTAO
  ICON_NAME_DASHCUBE
  ICON_NAME_DELICIOUS
  ICON_NAME_DEVIANTART
  ICON_NAME_DIGG
  ICON_NAME_DRIBBBLE
  ICON_NAME_DROPBOX
  ICON_NAME_DRUPAL
  ICON_NAME_EMPIRE
  ICON_NAME_ENVIRA_GALLERY
  ICON_NAME_EXPEDITEDSSL
  ICON_NAME_FACEBOOK
  ICON_NAME_FACEBOOK_F
  ICON_NAME_FACEBOOK_SQUARE
  ICON_NAME_FIREFOX
  ICON_NAME_FIRST_ORDER
  ICON_NAME_FLICKR
  ICON_NAME_FONT_AWESOME
  ICON_NAME_FONTICONS
  ICON_NAME_FORT_AWESOME
  ICON_NAME_FORUMBEE
  ICON_NAME_FOURSQUARE
  ICON_NAME_GG
  ICON_NAME_GG_CIRCLE
  ICON_NAME_GIT
  ICON_NAME_GIT_SQUARE
  ICON_NAME_GITHUB
  ICON_NAME_GITHUB_ALTERNATE
  ICON_NAME_GITHUB_SQUARE
  ICON_NAME_GITLAB
  ICON_NAME_GITTIP
  ICON_NAME_GLIDE
  ICON_NAME_GLIDE_G
  ICON_NAME_GOOGLE
  ICON_NAME_GOOGLE_PLUS
  ICON_NAME_GOOGLE_PLUS_CIRCLE
  ICON_NAME_GOOGLE_PLUS_SQUARE
  ICON_NAME_HACKER_NEWS
  ICON_NAME_HOUZZ
  ICON_NAME_INSTAGRAM
  ICON_NAME_INTERNET_EXPLORER
  ICON_NAME_IOXHOST
  ICON_NAME_JOOMLA
  ICON_NAME_JSFIDDLE
  ICON_NAME_LASTFM
  ICON_NAME_LASTFM_SQUARE
  ICON_NAME_LEANPUB
  ICON_NAME_LINKEDIN
  ICON_NAME_LINKEDIN_SQUARE
  ICON_NAME_LINUX
  ICON_NAME_MAXCDN
  ICON_NAME_MEANPATH
  ICON_NAME_MEDIUM
  ICON_NAME_MICROSOFT_EDGE
  ICON_NAME_MIXCLOUD
  ICON_NAME_MODX
  ICON_NAME_ODNOKLASSNIKI
  ICON_NAME_ODNOKLASSNIKI_SQUARE
  ICON_NAME_OPENCART
  ICON_NAME_OPERA
  ICON_NAME_OPTINMONSTER
  ICON_NAME_PAGELINES
  ICON_NAME_PIED_PIPER
  ICON_NAME_PIED_PIPER_ALTERNATE
  ICON_NAME_PIED_PIPER_HAT
  ICON_NAME_PINTEREST
  ICON_NAME_PINTEREST_SQUARE
  ICON_NAME_POCKET
  ICON_NAME_PRODUCT_HUNT
  ICON_NAME_QQ
  ICON_NAME_REBEL
  ICON_NAME_REDDIT
  ICON_NAME_REDDIT_ALIEN
  ICON_NAME_REDDIT_SQUARE
  ICON_NAME_RENREN
  ICON_NAME_SAFARI
  ICON_NAME_SCRIBD
  ICON_NAME_SELLSY
  ICON_NAME_SHIRTSINBULK
  ICON_NAME_SIMPLYBUILT
  ICON_NAME_SKYATLAS
  ICON_NAME_SKYPE
  ICON_NAME_SLACK
  ICON_NAME_SLIDESHARE
  ICON_NAME_SNAPCHAT
  ICON_NAME_SNAPCHAT_GHOST
  ICON_NAME_SNAPCHAT_SQUARE
  ICON_NAME_SOUNDCLOUD
  ICON_NAME_SPOTIFY
  ICON_NAME_STACK_EXCHANGE
  ICON_NAME_STACK_OVERFLOW
  ICON_NAME_STEAM
  ICON_NAME_STEAM_SQUARE
  ICON_NAME_STUMBLEUPON
  ICON_NAME_STUMBLEUPON_CIRCLE
  ICON_NAME_TENCENT_WEIBO
  ICON_NAME_THEMEISLE
  ICON_NAME_TRELLO
  ICON_NAME_TRIPADVISOR
  ICON_NAME_TUMBLR
  ICON_NAME_TUMBLR_SQUARE
  ICON_NAME_TWITCH
  ICON_NAME_TWITTER
  ICON_NAME_TWITTER_SQUARE
  ICON_NAME_VIACOIN
  ICON_NAME_VIADEO
  ICON_NAME_VIADEO_SQUARE
  ICON_NAME_VIMEO
  ICON_NAME_VIMEO_SQUARE
  ICON_NAME_VINE
  ICON_NAME_VK
  ICON_NAME_WECHAT
  ICON_NAME_WEIBO
  ICON_NAME_WHATSAPP
  ICON_NAME_WIKIPEDIA
  ICON_NAME_WINDOWS
  ICON_NAME_WORDPRESS
  ICON_NAME_WPBEGINNER
  ICON_NAME_WPFORMS
  ICON_NAME_XING
  ICON_NAME_XING_SQUARE
  ICON_NAME_Y_COMBINATOR
  ICON_NAME_YAHOO
  ICON_NAME_YELP
  ICON_NAME_YOAST
  ICON_NAME_YOUTUBE
  ICON_NAME_YOUTUBE_PLAY
  ICON_NAME_YOUTUBE_SQUARE
  ICON_NAME_DISABLED_USERS
  ICON_NAME_SPINNER_LOADING
  ICON_NAME_NOTCHED_CIRCLE_LOADING
  ICON_NAME_ASTERISK_LOADING
  ICON_NAME_FITTED_HELP
  ICON_NAME_MINI_HOME
  ICON_NAME_TINY_HOME
  ICON_NAME_SMALL_HOME
  ICON_NAME_HOME
  ICON_NAME_LARGE_HOME
  ICON_NAME_BIG_HOME
  ICON_NAME_HUGE_HOME
  ICON_NAME_MASSIVE_HOME
  ICON_NAME_CLOSE_LINK
  ICON_NAME_HELP_LINK
  ICON_NAME_HORIZONTALLY_FLIPPED_CLOUD
  ICON_NAME_VERTICALLY_FLIPPED_CLOUD
  ICON_NAME_CLOCKWISE_ROTATED_CLOUD
  ICON_NAME_COUNTERCLOCKWISE_ROTATED_CLOUD
  ICON_NAME_CIRCULAR_USERS
  ICON_NAME_CIRCULAR_TEAL_USERS
  ICON_NAME_CIRCULAR_INVERTED_USERS
  ICON_NAME_CIRCULAR_INVERTED_TEAL_USERS
  ICON_NAME_BORDERED_USERS
  ICON_NAME_BORDERED_TEAL_USERS
  ICON_NAME_BORDERED_INVERTED_BLACK_USERS
  ICON_NAME_BORDERED_INVERTED_TEAL_USERS
  ICON_NAME_RED_USERS
  ICON_NAME_ORANGE_USERS
  ICON_NAME_YELLOW_USERS
  ICON_NAME_OLIVE_USERS
  ICON_NAME_GREEN_USERS
  ICON_NAME_TEAL_USERS
  ICON_NAME_BLUE_USERS
  ICON_NAME_VIOLET_USERS
  ICON_NAME_PURPLE_USERS
  ICON_NAME_PINK_USERS
  ICON_NAME_BROWN_USERS
  ICON_NAME_GREY_USERS
  ICON_NAME_BLACK_USERS
  ICON_NAME_INVERTED_USERS
  ICON_NAME_INVERTED_RED_USERS
  ICON_NAME_INVERTED_ORANGE_USERS
  ICON_NAME_INVERTED_YELLOW_USERS
  ICON_NAME_INVERTED_OLIVE_USERS
  ICON_NAME_INVERTED_GREEN_USERS
  ICON_NAME_INVERTED_TEAL_USERS
  ICON_NAME_INVERTED_BLUE_USERS
  ICON_NAME_INVERTED_VIOLET_USERS
  ICON_NAME_INVERTED_PURPLE_USERS
  ICON_NAME_INVERTED_PINK_USERS
  ICON_NAME_INVERTED_BROWN_USERS
  ICON_NAME_INVERTED_GREY_USERS
  ICON_NAME_HUGE
  ICON_NAME_BIG_THIN_CIRCLE
  ICON_NAME_BIG_RED_DONT
  ICON_NAME_BLACK_USER
  ICON_NAME_BIG_LOADING_SUN
  ICON_NAME_CORNER_ADD
  ICON_NAME_LARGE
  ICON_NAME_INVERTED_CORNER_ADD
)

var iconNames = [...]string{
  "add to calendar",
  "alarm outline",
  "alarm mute outline",
  "alarm mute",
  "alarm",
  "at",
  "browser",
  "bug",
  "calendar outline",
  "calendar",
  "checked calendar",
  "cloud",
  "code",
  "comment outline",
  "comment",
  "comments outline",
  "comments",
  "copyright",
  "creative commons",
  "dashboard",
  "delete calendar",
  "external square",
  "external",
  "eyedropper",
  "feed",
  "find",
  "hand pointer",
  "hashtag",
  "heartbeat",
  "history",
  "hourglass empty",
  "hourglass end",
  "hourglass full",
  "hourglass half",
  "hourglass start",
  "idea",
  "image",
  "inbox",
  "industry",
  "lab",
  "mail outline",
  "mail square",
  "mail",
  "mouse pointer",
  "options",
  "paint brush",
  "payment",
  "percent",
  "privacy",
  "protect",
  "registered",
  "remove from calendar",
  "search",
  "setting",
  "settings",
  "shop",
  "shopping bag",
  "shopping basket",
  "signal",
  "sitemap",
  "tag",
  "tags",
  "tasks",
  "terminal",
  "text telephone",
  "ticket",
  "trademark",
  "trophy",
  "wifi",
  "add to cart",
  "add user",
  "adjust",
  "archive",
  "ban",
  "bookmark",
  "call",
  "call square",
  "clone",
  "cloud download",
  "cloud upload",
  "talk",
  "talk outline",
  "compress",
  "configure",
  "download",
  "edit",
  "erase",
  "exchange",
  "expand",
  "external share",
  "filter",
  "hide",
  "in cart",
  "lock",
  "mail forward",
  "object group",
  "object ungroup",
  "pin",
  "print",
  "random",
  "recycle",
  "refresh",
  "remove bookmark",
  "remove user",
  "repeat",
  "reply all",
  "reply",
  "retweet",
  "send",
  "send outline",
  "share alternate",
  "share alternate square",
  "share",
  "share square",
  "sign in",
  "sign out",
  "theme",
  "translate",
  "undo",
  "unhide",
  "unlock alternate",
  "unlock",
  "upload",
  "wait",
  "wizard",
  "write",
  "write square",
  "announcement",
  "birthday",
  "help circle",
  "help",
  "info circle",
  "info",
  "warning circle",
  "warning",
  "warning sign",
  "child",
  "doctor",
  "handicap",
  "spy",
  "student",
  "user",
  "users",
  "female",
  "gay",
  "genderless",
  "heterosexual",
  "intergender",
  "lesbian",
  "male",
  "man",
  "neuter",
  "non binary transgender",
  "other gender horizontal",
  "other gender",
  "other gender vertical",
  "transgender",
  "woman",
  "block layout",
  "crop",
  "grid layout",
  "list layout",
  "maximize",
  "resize horizontal",
  "resize vertical",
  "zoom",
  "zoom out",
  "anchor",
  "bar",
  "bomb",
  "book",
  "bullseye",
  "calculator",
  "cocktail",
  "diamond",
  "fax",
  "fire extinguisher",
  "fire",
  "flag checkered",
  "flag",
  "flag outline",
  "gift",
  "hand lizard",
  "hand peace",
  "hand paper",
  "hand rock",
  "hand scissors",
  "hand spock",
  "law",
  "leaf",
  "legal",
  "lemon",
  "life ring",
  "lightning",
  "magnet",
  "money",
  "moon",
  "plane",
  "puzzle",
  "road",
  "rocket",
  "shipping",
  "soccer",
  "sticky note",
  "sticky note outline",
  "suitcase",
  "sun",
  "travel",
  "treatment",
  "umbrella",
  "world",
  "asterisk",
  "certificate",
  "circle",
  "circle notched",
  "circle thin",
  "crosshairs",
  "cube",
  "cubes",
  "ellipsis horizontal",
  "ellipsis vertical",
  "quote left",
  "quote right",
  "spinner",
  "square",
  "square outline",
  "add circle",
  "add square",
  "check circle",
  "check circle outline",
  "check square",
  "checkmark box",
  "checkmark",
  "minus circle",
  "minus",
  "minus square",
  "minus square outline",
  "move",
  "plus",
  "plus square outline",
  "radio",
  "remove circle",
  "remove circle outline",
  "remove",
  "selected radio",
  "toggle off",
  "toggle on",
  "area chart",
  "bar chart",
  "camera retro",
  "film",
  "line chart",
  "newspaper",
  "photo",
  "pie chart",
  "sound",
  "angle double down",
  "angle double left",
  "angle double right",
  "angle double up",
  "angle down",
  "angle left",
  "angle right",
  "angle up",
  "arrow circle down",
  "arrow circle left",
  "arrow circle outline down",
  "arrow circle outline left",
  "arrow circle outline right",
  "arrow circle outline up",
  "arrow circle right",
  "arrow circle up",
  "arrow down",
  "arrow left",
  "arrow right",
  "arrow up",
  "caret down",
  "caret left",
  "caret right",
  "caret up",
  "chevron circle down",
  "chevron circle left",
  "chevron circle right",
  "chevron circle up",
  "chevron down",
  "chevron left",
  "chevron right",
  "chevron up",
  "long arrow down",
  "long arrow left",
  "long arrow right",
  "long arrow up",
  "pointing down",
  "pointing left",
  "pointing right",
  "pointing up",
  "toggle down",
  "toggle left",
  "toggle right",
  "toggle up",
  "mobile",
  "tablet",
  "battery empty",
  "battery full",
  "battery low",
  "battery medium",
  "desktop",
  "disk outline",
  "game",
  "high battery",
  "keyboard",
  "laptop",
  "plug",
  "power",
  "file archive outline",
  "file audio outline",
  "file code outline",
  "file excel outline",
  "file",
  "file image outline",
  "file outline",
  "file pdf outline",
  "file powerpoint outline",
  "file text",
  "file text outline",
  "file video outline",
  "file word outline",
  "folder",
  "folder open",
  "folder open outline",
  "folder outline",
  "level down",
  "level up",
  "trash",
  "trash outline",
  "barcode",
  "bluetooth alternative",
  "bluetooth",
  "css3",
  "database",
  "fork",
  "html5",
  "openid",
  "qrcode",
  "rss",
  "rss square",
  "server",
  "usb",
  "empty heart",
  "empty star",
  "frown",
  "heart",
  "meh",
  "smile",
  "star half empty",
  "star half",
  "star",
  "thumbs down",
  "thumbs outline down",
  "thumbs outline up",
  "thumbs up",
  "backward",
  "closed captioning",
  "eject",
  "fast backward",
  "fast forward",
  "forward",
  "music",
  "mute",
  "pause circle",
  "pause circle outline",
  "pause",
  "play",
  "record",
  "step backward",
  "step forward",
  "stop circle",
  "stop circle outline",
  "stop",
  "unmute",
  "video play",
  "video play outline",
  "volume down",
  "volume off",
  "volume up",
  "bicycle",
  "building",
  "building outline",
  "bus",
  "car",
  "coffee",
  "compass",
  "emergency",
  "first aid",
  "food",
  "h",
  "hospital",
  "hotel",
  "location arrow",
  "map",
  "map outline",
  "map pin",
  "map signs",
  "marker",
  "military",
  "motorcycle",
  "paw",
  "ship",
  "space shuttle",
  "spoon",
  "street view",
  "subway",
  "taxi",
  "train",
  "television",
  "tree",
  "university",
  "columns",
  "sort alphabet ascending",
  "sort alphabet descending",
  "sort ascending",
  "sort content ascending",
  "sort content descending",
  "sort descending",
  "sort",
  "sort numeric ascending",
  "sort numeric descending",
  "table",
  "align center",
  "align justify",
  "align left",
  "align right",
  "attach",
  "bold",
  "content",
  "copy",
  "cut",
  "font",
  "header",
  "indent",
  "italic",
  "linkify",
  "list",
  "ordered list",
  "outdent",
  "paragraph",
  "paste",
  "save",
  "strikethrough",
  "subscript",
  "superscript",
  "text cursor",
  "text height",
  "text width",
  "underline",
  "unlinkify",
  "unordered list",
  "bitcoin",
  "dollar",
  "euro",
  "lira",
  "pound",
  "ruble",
  "rupee",
  "shekel",
  "won",
  "yen",
  "american express",
  "credit card alternative",
  "diners club",
  "discover",
  "google wallet",
  "japan credit bureau",
  "mastercard",
  "paypal card",
  "paypal",
  "stripe",
  "visa",
  "wheelchair",
  "asl interpreting",
  "assistive listening systems",
  "audio description",
  "blind",
  "braille",
  "deafness",
  "low vision",
  "sign language",
  "universal access",
  "volume control phone",
  "500px",
  "adn",
  "amazon",
  "android",
  "angellist",
  "apple",
  "behance",
  "behance square",
  "bitbucket",
  "bitbucket square",
  "black tie",
  "buysellads",
  "chrome",
  "codepen",
  "codiepie",
  "connectdevelop",
  "contao",
  "dashcube",
  "delicious",
  "deviantart",
  "digg",
  "dribbble",
  "dropbox",
  "drupal",
  "empire",
  "envira gallery",
  "expeditedssl",
  "facebook",
  "facebook f",
  "facebook square",
  "firefox",
  "first order",
  "flickr",
  "font awesome",
  "fonticons",
  "fort awesome",
  "forumbee",
  "foursquare",
  "gg",
  "gg circle",
  "git",
  "git square",
  "github",
  "github alternate",
  "github square",
  "gitlab",
  "gittip",
  "glide",
  "glide g",
  "google",
  "google plus",
  "google plus circle",
  "google plus square",
  "hacker news",
  "houzz",
  "instagram",
  "internet explorer",
  "ioxhost",
  "joomla",
  "jsfiddle",
  "lastfm",
  "lastfm square",
  "leanpub",
  "linkedin",
  "linkedin square",
  "linux",
  "maxcdn",
  "meanpath",
  "medium",
  "microsoft edge",
  "mixcloud",
  "modx",
  "odnoklassniki",
  "odnoklassniki square",
  "opencart",
  "opera",
  "optinmonster",
  "pagelines",
  "pied piper",
  "pied piper alternate",
  "pied piper hat",
  "pinterest",
  "pinterest square",
  "pocket",
  "product hunt",
  "qq",
  "rebel",
  "reddit",
  "reddit alien",
  "reddit square",
  "renren",
  "safari",
  "scribd",
  "sellsy",
  "shirtsinbulk",
  "simplybuilt",
  "skyatlas",
  "skype",
  "slack",
  "slideshare",
  "snapchat",
  "snapchat ghost",
  "snapchat square",
  "soundcloud",
  "spotify",
  "stack exchange",
  "stack overflow",
  "steam",
  "steam square",
  "stumbleupon",
  "stumbleupon circle",
  "tencent weibo",
  "themeisle",
  "trello",
  "tripadvisor",
  "tumblr",
  "tumblr square",
  "twitch",
  "twitter",
  "twitter square",
  "viacoin",
  "viadeo",
  "viadeo square",
  "vimeo",
  "vimeo square",
  "vine",
  "vk",
  "wechat",
  "weibo",
  "whatsapp",
  "wikipedia",
  "windows",
  "wordpress",
  "wpbeginner",
  "wpforms",
  "xing",
  "xing square",
  "y combinator",
  "yahoo",
  "yelp",
  "yoast",
  "youtube",
  "youtube play",
  "youtube square",
  "disabled users",
  "spinner loading",
  "notched circle loading",
  "asterisk loading",
  "fitted help",
  "mini home",
  "tiny home",
  "small home",
  "home",
  "large home",
  "big home",
  "huge home",
  "massive home",
  "close link",
  "help link",
  "horizontally flipped cloud",
  "vertically flipped cloud",
  "clockwise rotated cloud",
  "counterclockwise rotated cloud",
  "circular users",
  "circular teal users",
  "circular inverted users",
  "circular inverted teal users",
  "bordered users",
  "bordered teal users",
  "bordered inverted black users",
  "bordered inverted teal users",
  "red users",
  "orange users",
  "yellow users",
  "olive users",
  "green users",
  "teal users",
  "blue users",
  "violet users",
  "purple users",
  "pink users",
  "brown users",
  "grey users",
  "black users",
  "inverted users",
  "inverted red users",
  "inverted orange users",
  "inverted yellow users",
  "inverted olive users",
  "inverted green users",
  "inverted teal users",
  "inverted blue users",
  "inverted violet users",
  "inverted purple users",
  "inverted pink users",
  "inverted brown users",
  "inverted grey users",
  "huge",
  "big thin circle",
  "big red dont",
  "black user",
  "big loading sun",
  "corner add",
  "large",
  "inverted corner add",
}

func (m IconName) String() string {
  return iconNames[ m ]
}