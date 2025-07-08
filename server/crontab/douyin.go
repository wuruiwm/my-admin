package crontab

import (
	"app/global"
	"app/util"
	"encoding/json"
	"fmt"
	"github.com/eddieivan01/nic"
)

type DouyinUserProfile struct {
	Code   int    `json:"code"`
	Router string `json:"router"`
	Data   struct {
		Extra struct {
			FatalItemIds []interface{} `json:"fatal_item_ids"`
			Logid        string        `json:"logid"`
			Now          int64         `json:"now"`
		} `json:"extra"`
		LogPb struct {
			ImprId string `json:"impr_id"`
		} `json:"log_pb"`
		StatusCode int         `json:"status_code"`
		StatusMsg  interface{} `json:"status_msg"`
		User       struct {
			AccountCertInfo string `json:"account_cert_info"`
			AppleAccount    int    `json:"apple_account"`
			Avatar168X168   struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_168x168"`
			Avatar300X300 struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_300x300"`
			AvatarLarger struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_larger"`
			AvatarMedium struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_medium"`
			AvatarThumb struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_thumb"`
			AwemeCount                    int  `json:"aweme_count"`
			AwemeCountCorrectionThreshold int  `json:"aweme_count_correction_threshold"`
			BirthdayHideLevel             int  `json:"birthday_hide_level"`
			CanSetItemCover               bool `json:"can_set_item_cover"`
			CanShowGroupCard              int  `json:"can_show_group_card"`
			CardEntries                   []struct {
				GotoUrl  string `json:"goto_url"`
				IconDark struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
				} `json:"icon_dark"`
				IconLight struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
				} `json:"icon_light"`
				SubTitle string `json:"sub_title"`
				Title    string `json:"title"`
				Type     int    `json:"type"`
			} `json:"card_entries"`
			City            string `json:"city"`
			CloseFriendType int    `json:"close_friend_type"`
			CommerceInfo    struct {
				ChallengeList   interface{} `json:"challenge_list"`
				HeadImageList   interface{} `json:"head_image_list"`
				OfflineInfoList interface{} `json:"offline_info_list"`
				SmartPhoneList  interface{} `json:"smart_phone_list"`
				TaskList        interface{} `json:"task_list"`
			} `json:"commerce_info"`
			CommerceUserInfo struct {
				AdRevenueRits            interface{} `json:"ad_revenue_rits"`
				HasAdsEntry              bool        `json:"has_ads_entry"`
				ShowStarAtlasCooperation bool        `json:"show_star_atlas_cooperation"`
				StarAtlas                int         `json:"star_atlas"`
			} `json:"commerce_user_info"`
			CommerceUserLevel     int    `json:"commerce_user_level"`
			Country               string `json:"country"`
			CoverAndHeadImageInfo struct {
				CoverList        interface{} `json:"cover_list"`
				ProfileCoverList []struct {
					CoverUrl struct {
						Uri     string   `json:"uri"`
						UrlList []string `json:"url_list"`
					} `json:"cover_url"`
					DarkCoverColor  string `json:"dark_cover_color"`
					LightCoverColor string `json:"light_cover_color"`
				} `json:"profile_cover_list"`
			} `json:"cover_and_head_image_info"`
			CoverColour string `json:"cover_colour"`
			CoverUrl    []struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"cover_url"`
			CustomVerify string `json:"custom_verify"`
			District     string `json:"district"`
			DongtaiCount int    `json:"dongtai_count"`
			DynamicCover struct {
			} `json:"dynamic_cover"`
			ElfemojiStatus         int    `json:"elfemoji_status"`
			EnableAiDouble         int    `json:"enable_ai_double"`
			EnableWish             bool   `json:"enable_wish"`
			EnterpriseUserInfo     string `json:"enterprise_user_info"`
			EnterpriseVerifyReason string `json:"enterprise_verify_reason"`
			FamiliarConfidence     int    `json:"familiar_confidence"`
			FavoritePermission     int    `json:"favorite_permission"`
			FavoritingCount        int    `json:"favoriting_count"`
			FollowGuide            bool   `json:"follow_guide"`
			FollowStatus           int    `json:"follow_status"`
			FollowerCount          int    `json:"follower_count"`
			FollowerRequestStatus  int    `json:"follower_request_status"`
			FollowerStatus         int    `json:"follower_status"`
			FollowingCount         int    `json:"following_count"`
			ForwardCount           int    `json:"forward_count"`
			Gender                 int    `json:"gender"`
			GeneralPermission      struct {
				FansPageToast              int  `json:"fans_page_toast"`
				FollowingFollowerListToast int  `json:"following_follower_list_toast"`
				IsHitActiveFansGrayed      bool `json:"is_hit_active_fans_grayed"`
			} `json:"general_permission"`
			HasEAccountRole      bool   `json:"has_e_account_role"`
			HasSubscription      bool   `json:"has_subscription"`
			HideRequestUpdate    int    `json:"hide_request_update"`
			ImPrimaryRoleId      int    `json:"im_primary_role_id"`
			ImRoleIds            []int  `json:"im_role_ids"`
			ImageSendExempt      bool   `json:"image_send_exempt"`
			InsId                string `json:"ins_id"`
			IpLocation           string `json:"ip_location"`
			IsActivityUser       bool   `json:"is_activity_user"`
			IsBan                bool   `json:"is_ban"`
			IsBlock              bool   `json:"is_block"`
			IsBlocked            bool   `json:"is_blocked"`
			IsEffectArtist       bool   `json:"is_effect_artist"`
			IsGovMediaVip        bool   `json:"is_gov_media_vip"`
			IsImOverseaUser      int    `json:"is_im_oversea_user"`
			IsMixUser            bool   `json:"is_mix_user"`
			IsNotShow            bool   `json:"is_not_show"`
			IsSeriesUser         bool   `json:"is_series_user"`
			IsSharingProfileUser int    `json:"is_sharing_profile_user"`
			IsStar               bool   `json:"is_star"`
			LifeStoryBlock       struct {
				LifeStoryBlock bool `json:"life_story_block"`
			} `json:"life_story_block"`
			LiveCommerce      bool `json:"live_commerce"`
			LiveStatus        int  `json:"live_status"`
			MateAddPermission int  `json:"mate_add_permission"`
			MateRelation      struct {
				MateApplyForward int `json:"mate_apply_forward"`
				MateApplyReverse int `json:"mate_apply_reverse"`
				MateStatus       int `json:"mate_status"`
			} `json:"mate_relation"`
			MaxFollowerCount        int    `json:"max_follower_count"`
			MessageChatEntry        bool   `json:"message_chat_entry"`
			MixCount                int    `json:"mix_count"`
			MplatformFollowersCount int    `json:"mplatform_followers_count"`
			NewFriendType           int    `json:"new_friend_type"`
			Nickname                string `json:"nickname"`
			OriginalMusician        struct {
				DiggCount      int `json:"digg_count"`
				MusicCount     int `json:"music_count"`
				MusicUsedCount int `json:"music_used_count"`
			} `json:"original_musician"`
			PersonalTagList []struct {
				LightTagImageUrl string `json:"light_tag_image_url,omitempty"`
				PersonalTagType  int    `json:"personal_tag_type"`
				TagImageUrl      string `json:"tag_image_url,omitempty"`
				Text             string `json:"text"`
			} `json:"personal_tag_list"`
			PigeonDarenStatus  string `json:"pigeon_daren_status"`
			PigeonDarenWarnTag string `json:"pigeon_daren_warn_tag"`
			ProfileMobParams   []struct {
				EventKey  string `json:"event_key"`
				MobParams string `json:"mob_params"`
			} `json:"profile_mob_params"`
			ProfileShow struct {
				IdentifyAuthInfos interface{} `json:"identify_auth_infos"`
			} `json:"profile_show"`
			ProfileTabInfo struct {
				ProfileLandingTab int         `json:"profile_landing_tab"`
				ProfileTabList    interface{} `json:"profile_tab_list"`
				ProfileTabListV2  interface{} `json:"profile_tab_list_v2"`
			} `json:"profile_tab_info"`
			ProfileTabType      int    `json:"profile_tab_type"`
			Province            string `json:"province"`
			PublicCollectsCount int    `json:"public_collects_count"`
			PublishLandingTab   int    `json:"publish_landing_tab"`
			RFansGroupInfo      struct {
			} `json:"r_fans_group_info"`
			RecommendReasonRelation   string `json:"recommend_reason_relation"`
			RecommendUserReasonSource int    `json:"recommend_user_reason_source"`
			RiskNoticeText            string `json:"risk_notice_text"`
			RoomId                    int    `json:"room_id"`
			RoomIdStr                 string `json:"room_id_str"`
			SchoolName                string `json:"school_name"`
			SecUid                    string `json:"sec_uid"`
			Secret                    int    `json:"secret"`
			SeriesCount               int    `json:"series_count"`
			ShareInfo                 struct {
				BoolPersist   int    `json:"bool_persist"`
				LifeShareExt  string `json:"life_share_ext"`
				ShareDesc     string `json:"share_desc"`
				ShareImageUrl struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
				} `json:"share_image_url"`
				ShareQrcodeUrl struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
				} `json:"share_qrcode_url"`
				ShareTitle     string `json:"share_title"`
				ShareUrl       string `json:"share_url"`
				ShareWeiboDesc string `json:"share_weibo_desc"`
			} `json:"share_info"`
			ShortId                string `json:"short_id"`
			ShowFavoriteList       bool   `json:"show_favorite_list"`
			ShowSubscription       bool   `json:"show_subscription"`
			Signature              string `json:"signature"`
			SignatureDisplayLines  int    `json:"signature_display_lines"`
			SignatureLanguage      string `json:"signature_language"`
			SocialRealRelationType int    `json:"social_real_relation_type"`
			SpecialFollowStatus    int    `json:"special_follow_status"`
			StoryTabEmpty          bool   `json:"story_tab_empty"`
			SyncToToutiao          int    `json:"sync_to_toutiao"`
			TabSettings            struct {
				PrivateTab struct {
					PrivateTabStyle int  `json:"private_tab_style"`
					ShowPrivateTab  bool `json:"show_private_tab"`
				} `json:"private_tab"`
			} `json:"tab_settings"`
			TotalFavorited                    int    `json:"total_favorited"`
			TotalFavoritedCorrectionThreshold int    `json:"total_favorited_correction_threshold"`
			TwitterId                         string `json:"twitter_id"`
			TwitterName                       string `json:"twitter_name"`
			Uid                               string `json:"uid"`
			UniqueId                          string `json:"unique_id"`
			UrgeDetail                        struct {
				CtlMap    string `json:"ctl_map"`
				UserUrged int    `json:"user_urged"`
			} `json:"urge_detail"`
			UserAge          int `json:"user_age"`
			UserNotSee       int `json:"user_not_see"`
			UserNotShow      int `json:"user_not_show"`
			VerificationType int `json:"verification_type"`
			VideoCover       struct {
			} `json:"video_cover"`
			VideoIcon struct {
				Height  int           `json:"height"`
				Uri     string        `json:"uri"`
				UrlList []interface{} `json:"url_list"`
				Width   int           `json:"width"`
			} `json:"video_icon"`
			WatchStatus   bool `json:"watch_status"`
			WhiteCoverUrl []struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"white_cover_url"`
			WithCommerceEnterpriseTabEntry bool   `json:"with_commerce_enterprise_tab_entry"`
			WithCommerceEntry              bool   `json:"with_commerce_entry"`
			WithFusionShopEntry            bool   `json:"with_fusion_shop_entry"`
			WithNewGoods                   bool   `json:"with_new_goods"`
			YoutubeChannelId               string `json:"youtube_channel_id"`
			YoutubeChannelTitle            string `json:"youtube_channel_title"`
		} `json:"user"`
	} `json:"data"`
}

func Douyin() {
	util.Logger.Info("抖音cookie检测", util.Map{
		"msg": "开始检测",
	})
	u := fmt.Sprintf("%s/api/douyin/web/handler_user_profile?sec_user_id=%s", global.Config.AdminConfig.Douyin.Url, global.Config.AdminConfig.Douyin.UserId)
	result, err := nic.Get(u, nil)
	if err != nil {
		util.Notice("抖音cookie检测", "请求失败 error:"+err.Error())
		return
	}
	douyinUserProfile := &DouyinUserProfile{}
	err = json.Unmarshal(result.Bytes, douyinUserProfile)
	if err != nil {
		util.Notice("抖音cookie检测", "json反序列化失败 error:"+err.Error()+" body:"+result.Text)
		return
	}
	if douyinUserProfile.Code != 200 {
		util.Notice("抖音cookie检测", "爬取响应异常 body:"+result.Text)
		return
	}
	util.Logger.Info("抖音cookie检测", util.Map{
		"msg": "检测正常",
	})
}
