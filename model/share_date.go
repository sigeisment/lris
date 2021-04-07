// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    sharedData, err := UnmarshalSharedData(bytes)
//    bytes, err = sharedData.Marshal()

package model

import "bytes"
import "errors"
import "encoding/json"

func UnmarshalSharedData(data []byte) (SharedData, error) {
	var r SharedData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SharedData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SharedData struct {
	Config                      Config                 `json:"config"`
	CountryCode                 string                 `json:"country_code"`
	LanguageCode                string                 `json:"language_code"`
	Locale                      string                 `json:"locale"`
	EntryData                   EntryData              `json:"entry_data"`
	Hostname                    string                 `json:"hostname"`
	IsWhitelistedCrawlBot       bool                   `json:"is_whitelisted_crawl_bot"`
	ConnectionQualityRating     string                 `json:"connection_quality_rating"`
	DeploymentStage             string                 `json:"deployment_stage"`
	Platform                    string                 `json:"platform"`
	Nonce                       string                 `json:"nonce"`
	MidPct                      float64                `json:"mid_pct"`
	ZeroData                    ZeroData               `json:"zero_data"`
	CacheSchemaVersion          int64                  `json:"cache_schema_version"`
	ServerChecks                ServerChecks           `json:"server_checks"`
	Knobx                       map[string]*Knobx      `json:"knobx"`
	ToCache                     ToCache                `json:"to_cache"`
	DeviceID                    string                 `json:"device_id"`
	BrowserPushPubKey           string                 `json:"browser_push_pub_key"`
	Encryption                  Encryption             `json:"encryption"`
	IsDev                       bool                   `json:"is_dev"`
	SignalCollectionConfig      SignalCollectionConfig `json:"signal_collection_config"`
	ShouldShowConsentDialog     bool                   `json:"should_show_consent_dialog"`
	UseLoggedOut3PConsentDialog bool                   `json:"use_logged_out_3p_consent_dialog"`
	RolloutHash                 string                 `json:"rollout_hash"`
	BundleVariant               string                 `json:"bundle_variant"`
	FrontendEnv                 string                 `json:"frontend_env"`
}

type Config struct {
	CSRFToken string `json:"csrf_token"`
	Viewer    Viewer `json:"viewer"`
	ViewerID  string `json:"viewerId"`
}

type Viewer struct {
	Biography             string      `json:"biography"`
	CategoryName          interface{} `json:"category_name"`
	ExternalURL           interface{} `json:"external_url"`
	Fbid                  string      `json:"fbid"`
	FullName              string      `json:"full_name"`
	HasPhoneNumber        bool        `json:"has_phone_number"`
	HasProfilePic         bool        `json:"has_profile_pic"`
	HasTabbedInbox        bool        `json:"has_tabbed_inbox"`
	ID                    string      `json:"id"`
	IsJoinedRecently      bool        `json:"is_joined_recently"`
	IsPrivate             bool        `json:"is_private"`
	IsProfessionalAccount bool        `json:"is_professional_account"`
	ProfilePicURL         string      `json:"profile_pic_url"`
	ProfilePicURLHD       string      `json:"profile_pic_url_hd"`
	ShouldShowCategory    bool        `json:"should_show_category"`
	Username              string      `json:"username"`
	BadgeCount            string      `json:"badge_count"`
}

type Encryption struct {
	KeyID     string `json:"key_id"`
	PublicKey string `json:"public_key"`
	Version   string `json:"version"`
}

type EntryData struct {
	ProfilePage []ProfilePage `json:"ProfilePage"`
}

type ProfilePage struct {
	LoggingPageID           string      `json:"logging_page_id"`
	ShowSuggestedProfiles   bool        `json:"show_suggested_profiles"`
	ShowFollowDialog        bool        `json:"show_follow_dialog"`
	Graphql                 Graphql     `json:"graphql"`
	ToastContentOnLoad      interface{} `json:"toast_content_on_load"`
	ShowViewShop            bool        `json:"show_view_shop"`
	ProfilePicEditSyncProps interface{} `json:"profile_pic_edit_sync_props"`
}

type QueryPage struct {
	Data Graphql `json:"data"`
}

func UnmarshalQueryPage(data []byte) (QueryPage, error) {
	var r QueryPage
	err := json.Unmarshal(data, &r)
	return r, err
}

type Graphql struct {
	User GraphqlUser `json:"user"`
}

type GraphqlUser struct {
	Biography                string                      `json:"biography"`
	BlockedByViewer          bool                        `json:"blocked_by_viewer"`
	RestrictedByViewer       bool                        `json:"restricted_by_viewer"`
	CountryBlock             bool                        `json:"country_block"`
	ExternalURL              interface{}                 `json:"external_url"`
	ExternalURLLinkshimmed   interface{}                 `json:"external_url_linkshimmed"`
	EdgeFollowedBy           EdgeFollowClass             `json:"edge_followed_by"`
	Fbid                     string                      `json:"fbid"`
	FollowedByViewer         bool                        `json:"followed_by_viewer"`
	EdgeFollow               EdgeFollowClass             `json:"edge_follow"`
	FollowsViewer            bool                        `json:"follows_viewer"`
	FullName                 string                      `json:"full_name"`
	HasArEffects             bool                        `json:"has_ar_effects"`
	HasClips                 bool                        `json:"has_clips"`
	HasGuides                bool                        `json:"has_guides"`
	HasChannel               bool                        `json:"has_channel"`
	HasBlockedViewer         bool                        `json:"has_blocked_viewer"`
	HighlightReelCount       int64                       `json:"highlight_reel_count"`
	HasRequestedViewer       bool                        `json:"has_requested_viewer"`
	ID                       string                      `json:"id"`
	IsBusinessAccount        bool                        `json:"is_business_account"`
	IsJoinedRecently         bool                        `json:"is_joined_recently"`
	BusinessCategoryName     interface{}                 `json:"business_category_name"`
	OverallCategoryName      interface{}                 `json:"overall_category_name"`
	CategoryEnum             interface{}                 `json:"category_enum"`
	CategoryName             string                      `json:"category_name"`
	IsPrivate                bool                        `json:"is_private"`
	IsVerified               bool                        `json:"is_verified"`
	EdgeMutualFollowedBy     EdgeMutualFollowedBy        `json:"edge_mutual_followed_by"`
	ProfilePicURL            string                      `json:"profile_pic_url"`
	ProfilePicURLHD          string                      `json:"profile_pic_url_hd"`
	RequestedByViewer        bool                        `json:"requested_by_viewer"`
	ShouldShowCategory       bool                        `json:"should_show_category"`
	Username                 Username                    `json:"username"`
	ConnectedFbPage          interface{}                 `json:"connected_fb_page"`
	EdgeFelixVideoTimeline   EdgeFelixVideoTimelineClass `json:"edge_felix_video_timeline"`
	EdgeOwnerToTimelineMedia EdgeOwnerToTimelineMedia    `json:"edge_owner_to_timeline_media"`
	EdgeSavedMedia           EdgeFelixVideoTimelineClass `json:"edge_saved_media"`
	EdgeMediaCollections     EdgeFelixVideoTimelineClass `json:"edge_media_collections"`
}

type EdgeFelixVideoTimelineClass struct {
	Count    int64                        `json:"count"`
	PageInfo PageInfo                     `json:"page_info"`
	Edges    []EdgeFelixVideoTimelineEdge `json:"edges"`
}

type EdgeFelixVideoTimelineEdge struct {
	Node PurpleNode `json:"node"`
}

type PurpleNode struct {
	Typename               string              `json:"__typename"`
	ID                     string              `json:"id"`
	Shortcode              string              `json:"shortcode"`
	Dimensions             Dimensions          `json:"dimensions"`
	DisplayURL             string              `json:"display_url"`
	EdgeMediaToTaggedUser  EdgeMediaTo         `json:"edge_media_to_tagged_user"`
	FactCheckOverallRating interface{}         `json:"fact_check_overall_rating"`
	FactCheckInformation   interface{}         `json:"fact_check_information"`
	GatingInfo             interface{}         `json:"gating_info"`
	SharingFrictionInfo    SharingFrictionInfo `json:"sharing_friction_info"`
	MediaOverlayInfo       interface{}         `json:"media_overlay_info"`
	MediaPreview           string              `json:"media_preview"`
	Owner                  Owner               `json:"owner"`
	IsVideo                bool                `json:"is_video"`
	AccessibilityCaption   interface{}         `json:"accessibility_caption"`
	DashInfo               DashInfo            `json:"dash_info"`
	HasAudio               bool                `json:"has_audio"`
	TrackingToken          string              `json:"tracking_token"`
	VideoURL               string              `json:"video_url"`
	VideoViewCount         int64               `json:"video_view_count"`
	EdgeMediaToCaption     EdgeMediaTo         `json:"edge_media_to_caption"`
	EdgeMediaToComment     EdgeFollowClass     `json:"edge_media_to_comment"`
	CommentsDisabled       bool                `json:"comments_disabled"`
	TakenAtTimestamp       int64               `json:"taken_at_timestamp"`
	EdgeLikedBy            EdgeFollowClass     `json:"edge_liked_by"`
	EdgeMediaPreviewLike   EdgeFollowClass     `json:"edge_media_preview_like"`
	Location               interface{}         `json:"location"`
	ThumbnailSrc           string              `json:"thumbnail_src"`
	ThumbnailResources     []ThumbnailResource `json:"thumbnail_resources"`
	FelixProfileGridCrop   interface{}         `json:"felix_profile_grid_crop"`
	EncodingStatus         interface{}         `json:"encoding_status"`
	IsPublished            bool                `json:"is_published"`
	ProductType            string              `json:"product_type"`
	Title                  string              `json:"title"`
	VideoDuration          float64             `json:"video_duration"`
}

type DashInfo struct {
	IsDashEligible    bool   `json:"is_dash_eligible"`
	VideoDashManifest string `json:"video_dash_manifest"`
	NumberOfQualities int64  `json:"number_of_qualities"`
}

type Dimensions struct {
	Height int64 `json:"height"`
	Width  int64 `json:"width"`
}

type EdgeFollowClass struct {
	Count int64 `json:"count"`
}

type EdgeMediaTo struct {
	Edges []EdgeMediaToCaptionEdge `json:"edges"`
}

type EdgeMediaToCaptionEdge struct {
	Node FluffyNode `json:"node"`
}

type FluffyNode struct {
	Text string `json:"text"`
}

type Owner struct {
	ID       string   `json:"id"`
	Username Username `json:"username"`
}

type SharingFrictionInfo struct {
	ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
	BloksAppURL               interface{} `json:"bloks_app_url"`
}

type ThumbnailResource struct {
	Src          string `json:"src"`
	ConfigWidth  int64  `json:"config_width"`
	ConfigHeight int64  `json:"config_height"`
}

type PageInfo struct {
	HasNextPage bool    `json:"has_next_page"`
	EndCursor   *string `json:"end_cursor"`
}

type EdgeMutualFollowedBy struct {
	Count int64         `json:"count"`
	Edges []interface{} `json:"edges"`
}

type EdgeOwnerToTimelineMedia struct {
	Count    int64                          `json:"count"`
	PageInfo PageInfo                       `json:"page_info"`
	Edges    []EdgeOwnerToTimelineMediaEdge `json:"edges"`
}

type EdgeOwnerToTimelineMediaEdge struct {
	Node TentacledNode `json:"node"`
}

type TentacledNode struct {
	Typename               Typename               `json:"__typename"`
	ID                     string                 `json:"id"`
	Shortcode              string                 `json:"shortcode"`
	Dimensions             Dimensions             `json:"dimensions"`
	DisplayURL             string                 `json:"display_url"`
	EdgeMediaToTaggedUser  EdgeMediaToTaggedUser  `json:"edge_media_to_tagged_user"`
	FactCheckOverallRating interface{}            `json:"fact_check_overall_rating"`
	FactCheckInformation   interface{}            `json:"fact_check_information"`
	GatingInfo             interface{}            `json:"gating_info"`
	SharingFrictionInfo    SharingFrictionInfo    `json:"sharing_friction_info"`
	MediaOverlayInfo       interface{}            `json:"media_overlay_info"`
	MediaPreview           *string                `json:"media_preview"`
	Owner                  Owner                  `json:"owner"`
	IsVideo                bool                   `json:"is_video"`
	AccessibilityCaption   string                 `json:"accessibility_caption"`
	EdgeMediaToCaption     EdgeMediaTo            `json:"edge_media_to_caption"`
	EdgeMediaToComment     EdgeFollowClass        `json:"edge_media_to_comment"`
	CommentsDisabled       bool                   `json:"comments_disabled"`
	TakenAtTimestamp       int64                  `json:"taken_at_timestamp"`
	EdgeLikedBy            EdgeFollowClass        `json:"edge_liked_by"`
	EdgeMediaPreviewLike   EdgeFollowClass        `json:"edge_media_preview_like"`
	Location               interface{}            `json:"location"`
	ThumbnailSrc           string                 `json:"thumbnail_src"`
	ThumbnailResources     []ThumbnailResource    `json:"thumbnail_resources"`
	EdgeSidecarToChildren  *EdgeSidecarToChildren `json:"edge_sidecar_to_children,omitempty"`
}

type EdgeMediaToTaggedUser struct {
	Edges []PurpleEdge `json:"edges"`
}

type PurpleEdge struct {
	Node StickyNode `json:"node"`
}

type StickyNode struct {
	User NodeUser `json:"user"`
	X    float64  `json:"x"`
	Y    float64  `json:"y"`
}

type NodeUser struct {
	FullName      string `json:"full_name"`
	ID            string `json:"id"`
	IsVerified    bool   `json:"is_verified"`
	ProfilePicURL string `json:"profile_pic_url"`
	Username      string `json:"username"`
}

type EdgeSidecarToChildren struct {
	Edges []EdgeSidecarToChildrenEdge `json:"edges"`
}

type EdgeSidecarToChildrenEdge struct {
	Node IndigoNode `json:"node"`
}

type IndigoNode struct {
	Typename               Typename              `json:"__typename"`
	ID                     string                `json:"id"`
	Shortcode              string                `json:"shortcode"`
	Dimensions             Dimensions            `json:"dimensions"`
	DisplayURL             string                `json:"display_url"`
	EdgeMediaToTaggedUser  EdgeMediaToTaggedUser `json:"edge_media_to_tagged_user"`
	FactCheckOverallRating interface{}           `json:"fact_check_overall_rating"`
	FactCheckInformation   interface{}           `json:"fact_check_information"`
	GatingInfo             interface{}           `json:"gating_info"`
	SharingFrictionInfo    SharingFrictionInfo   `json:"sharing_friction_info"`
	MediaOverlayInfo       interface{}           `json:"media_overlay_info"`
	MediaPreview           string                `json:"media_preview"`
	Owner                  Owner                 `json:"owner"`
	IsVideo                bool                  `json:"is_video"`
	AccessibilityCaption   string                `json:"accessibility_caption"`
}

type ServerChecks struct {
	Hfe bool `json:"hfe"`
}

type SignalCollectionConfig struct {
	BBS int64       `json:"bbs"`
	Ctw interface{} `json:"ctw"`
	Dbs int64       `json:"dbs"`
	Fd  int64       `json:"fd"`
	Hbc Hbc         `json:"hbc"`
	I   int64       `json:"i"`
	Rt  int64       `json:"rt"`
	Sbs int64       `json:"sbs"`
	Sc  Sc          `json:"sc"`
	Sid int64       `json:"sid"`
}

type Hbc struct {
	Hbbi  int64  `json:"hbbi"`
	Hbcbc int64  `json:"hbcbc"`
	Hbi   int64  `json:"hbi"`
	Hbv   string `json:"hbv"`
	Hbvbc int64  `json:"hbvbc"`
}

type Sc struct {
	C [][]int64 `json:"c"`
	T int64     `json:"t"`
}

type ToCache struct {
	Gatekeepers    map[string]bool `json:"gatekeepers"`
	Qe             Qe              `json:"qe"`
	ProbablyHasApp bool            `json:"probably_has_app"`
	Cb             bool            `json:"cb"`
}

type Qe struct {
	The0                          The0                          `json:"0"`
	The12                         The12                         `json:"12"`
	The13                         The100                        `json:"13"`
	The16                         The100                        `json:"16"`
	The22                         The22                         `json:"22"`
	The23                         The101                        `json:"23"`
	The25                         The109                        `json:"25"`
	The26                         The26                         `json:"26"`
	The28                         The100                        `json:"28"`
	The31                         The109                        `json:"31"`
	The33                         The109                        `json:"33"`
	The34                         The100                        `json:"34"`
	The36                         The101                        `json:"36"`
	The37                         The100                        `json:"37"`
	The39                         The101                        `json:"39"`
	The41                         The41                         `json:"41"`
	The42                         The100                        `json:"42"`
	The43                         The101                        `json:"43"`
	The44                         The44                         `json:"44"`
	The45                         The45                         `json:"45"`
	The46                         The100                        `json:"46"`
	The47                         The101                        `json:"47"`
	The49                         The100                        `json:"49"`
	The50                         The100                        `json:"50"`
	The54                         The100                        `json:"54"`
	The58                         The22                         `json:"58"`
	The59                         The100                        `json:"59"`
	The62                         The100                        `json:"62"`
	The67                         The101                        `json:"67"`
	The69                         The100                        `json:"69"`
	The71                         The71                         `json:"71"`
	The72                         The101                        `json:"72"`
	The73                         The100                        `json:"73"`
	The74                         The101                        `json:"74"`
	The75                         The100                        `json:"75"`
	The77                         The159                        `json:"77"`
	The80                         The101                        `json:"80"`
	The84                         The101                        `json:"84"`
	The85                         The124                        `json:"85"`
	The87                         The100                        `json:"87"`
	The93                         The100                        `json:"93"`
	The95                         The109                        `json:"95"`
	The98                         The159                        `json:"98"`
	The100                        The100                        `json:"100"`
	The101                        The101                        `json:"101"`
	The108                        The101                        `json:"108"`
	The109                        The109                        `json:"109"`
	The111                        The101                        `json:"111"`
	The113                        The101                        `json:"113"`
	The118                        The101                        `json:"118"`
	The119                        The100                        `json:"119"`
	The121                        The100                        `json:"121"`
	The122                        The100                        `json:"122"`
	The123                        The101                        `json:"123"`
	The124                        The124                        `json:"124"`
	The125                        The100                        `json:"125"`
	The127                        The101                        `json:"127"`
	The128                        The101                        `json:"128"`
	The129                        The101                        `json:"129"`
	The131                        The101                        `json:"131"`
	The132                        The100                        `json:"132"`
	The135                        The101                        `json:"135"`
	The137                        The109                        `json:"137"`
	The141                        The101                        `json:"141"`
	The142                        The100                        `json:"142"`
	The143                        The101                        `json:"143"`
	The146                        The101                        `json:"146"`
	The148                        The101                        `json:"148"`
	The149                        The100                        `json:"149"`
	The152                        The101                        `json:"152"`
	The154                        The154                        `json:"154"`
	The155                        The109                        `json:"155"`
	The156                        The100                        `json:"156"`
	The158                        The109                        `json:"158"`
	The159                        The159                        `json:"159"`
	The160                        The101                        `json:"160"`
	The162                        The109                        `json:"162"`
	The163                        The109                        `json:"163"`
	The164                        The101                        `json:"164"`
	The165                        The101                        `json:"165"`
	The166                        The154                        `json:"166"`
	The167                        The101                        `json:"167"`
	The168                        The101                        `json:"168"`
	The170                        The100                        `json:"170"`
	The171                        The100                        `json:"171"`
	AppUpsell                     AppUpsell                     `json:"app_upsell"`
	IglAppUpsell                  AppUpsell                     `json:"igl_app_upsell"`
	Notif                         AppUpsell                     `json:"notif"`
	Onetaplogin                   AppUpsell                     `json:"onetaplogin"`
	FelixClearFbCookie            FelixClearFbCookie            `json:"felix_clear_fb_cookie"`
	FelixCreationDurationLimits   FelixCreationDurationLimits   `json:"felix_creation_duration_limits"`
	FelixCreationFbCrossposting   FelixCreationFbCrossposting   `json:"felix_creation_fb_crossposting"`
	FelixCreationFbCrosspostingV2 FelixCreationFbCrosspostingV2 `json:"felix_creation_fb_crossposting_v2"`
	FelixCreationValidation       FelixCreationValidation       `json:"felix_creation_validation"`
	PostOptions                   PostOptions                   `json:"post_options"`
	StickerTray                   AppUpsell                     `json:"sticker_tray"`
	WebSentry                     WebSentry                     `json:"web_sentry"`
}

type AppUpsell struct {
	G string   `json:"g"`
	P ZeroData `json:"p"`
}

type ZeroData struct {
}

type FelixClearFbCookie struct {
	G string              `json:"g"`
	P FelixClearFbCookieP `json:"p"`
}

type FelixClearFbCookieP struct {
	IsEnabled string `json:"is_enabled"`
	Blacklist string `json:"blacklist"`
}

type FelixCreationDurationLimits struct {
	G string                       `json:"g"`
	P FelixCreationDurationLimitsP `json:"p"`
}

type FelixCreationDurationLimitsP struct {
	MaximumLengthSeconds string `json:"maximum_length_seconds"`
	MinimumLengthSeconds string `json:"minimum_length_seconds"`
}

type FelixCreationFbCrossposting struct {
	G string                       `json:"g"`
	P FelixCreationFbCrosspostingP `json:"p"`
}

type FelixCreationFbCrosspostingP struct {
	IsEnabled string `json:"is_enabled"`
}

type FelixCreationFbCrosspostingV2 struct {
	G string                         `json:"g"`
	P FelixCreationFbCrosspostingV2P `json:"p"`
}

type FelixCreationFbCrosspostingV2P struct {
	IsEnabled      string `json:"is_enabled"`
	DisplayVersion string `json:"display_version"`
}

type FelixCreationValidation struct {
	G string                   `json:"g"`
	P FelixCreationValidationP `json:"p"`
}

type FelixCreationValidationP struct {
	EditVideoControls                  string `json:"edit_video_controls"`
	DescriptionMaximumLength           string `json:"description_maximum_length"`
	MaxVideoSizeInBytes                string `json:"max_video_size_in_bytes"`
	MinimumLengthForFeedPreviewSeconds string `json:"minimum_length_for_feed_preview_seconds"`
	TitleMaximumLength                 string `json:"title_maximum_length"`
	ValidCoverMIMETypes                string `json:"valid_cover_mime_types"`
	ValidVideoExtensions               string `json:"valid_video_extensions"`
	ValidVideoMIMETypes                string `json:"valid_video_mime_types"`
}

type PostOptions struct {
	G string       `json:"g"`
	P PostOptionsP `json:"p"`
}

type PostOptionsP struct {
	UseRefactor     string `json:"use_refactor"`
	EnableIgtvEmbed string `json:"enable_igtv_embed"`
}

type The0 struct {
	P   The0_P   `json:"p"`
	L   ZeroData `json:"l"`
	Qex bool     `json:"qex"`
}

type The0_P struct {
	The9 bool `json:"9"`
}

type The100 struct {
	P   LClass   `json:"p"`
	L   ZeroData `json:"l"`
	Qex bool     `json:"qex"`
}

type LClass struct {
	The0 bool `json:"0"`
}

type The101 struct {
	P   map[string]bool `json:"p"`
	L   ZeroData        `json:"l"`
	Qex bool            `json:"qex"`
}

type The109 struct {
	P   ZeroData `json:"p"`
	L   ZeroData `json:"l"`
	Qex bool     `json:"qex"`
}

type The12 struct {
	P   The12_P  `json:"p"`
	L   ZeroData `json:"l"`
	Qex bool     `json:"qex"`
}

type The12_P struct {
	The0 int64 `json:"0"`
}

type The124 struct {
	P   map[string]*The124_P `json:"p"`
	L   ZeroData             `json:"l"`
	Qex bool                 `json:"qex"`
}

type The154 struct {
	P   LClass `json:"p"`
	L   LClass `json:"l"`
	Qex bool   `json:"qex"`
}

type The159 struct {
	P   The159_P `json:"p"`
	L   ZeroData `json:"l"`
	Qex bool     `json:"qex"`
}

type The159_P struct {
	The1 bool `json:"1"`
}

type The22 struct {
	P   map[string]*The22_P `json:"p"`
	L   ZeroData            `json:"l"`
	Qex bool                `json:"qex"`
}

type The26 struct {
	P   The26_P  `json:"p"`
	L   ZeroData `json:"l"`
	Qex bool     `json:"qex"`
}

type The26_P struct {
	The0 string `json:"0"`
}

type The41 struct {
	P   The41_P  `json:"p"`
	L   ZeroData `json:"l"`
	Qex bool     `json:"qex"`
}

type The41_P struct {
	The3 bool `json:"3"`
}

type The44 struct {
	P   map[string]*The44_P `json:"p"`
	L   ZeroData            `json:"l"`
	Qex bool                `json:"qex"`
}

type The45 struct {
	P   map[string]*The45_P `json:"p"`
	L   ZeroData            `json:"l"`
	Qex bool                `json:"qex"`
}

type The71 struct {
	P   The71_P  `json:"p"`
	L   ZeroData `json:"l"`
	Qex bool     `json:"qex"`
}

type The71_P struct {
	The1 string `json:"1"`
}

type WebSentry struct {
	G string     `json:"g"`
	P WebSentryP `json:"p"`
}

type WebSentryP struct {
	ShowFeedback string `json:"show_feedback"`
}

type Username string

const (
	BingbingFan Username = "bingbing_fan"
)

type Typename string

const (
	GraphImage   Typename = "GraphImage"
	GraphSidecar Typename = "GraphSidecar"
)

type Knobx struct {
	Bool    *bool
	Integer *int64
}

func (x *Knobx) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, &x.Bool, nil, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Knobx) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, x.Bool, nil, false, nil, false, nil, false, nil, false, nil, false)
}

type The124_P struct {
	Bool   *bool
	String *string
}

func (x *The124_P) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *The124_P) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

type The22_P struct {
	Bool   *bool
	Double *float64
}

func (x *The22_P) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, &x.Double, &x.Bool, nil, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *The22_P) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, x.Bool, nil, false, nil, false, nil, false, nil, false, nil, false)
}

type The44_P struct {
	Double *float64
	String *string
}

func (x *The44_P) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, &x.Double, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *The44_P) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

type The45_P struct {
	Bool    *bool
	Integer *int64
	String  *string
}

func (x *The45_P) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, &x.Bool, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *The45_P) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, x.Bool, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
