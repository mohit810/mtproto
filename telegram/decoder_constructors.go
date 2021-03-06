package telegram

import (
	"fmt"

	"github.com/xelaj/mtproto/serialize"
)

func init() {
	serialize.AddObjectConstructor(GenerateStructByConstructor)
}

func GenerateStructByConstructor(constructorID uint32) (object serialize.TL, isEnum bool, err error) {
	switch constructorID {
	case uint32(0xc331e80a):
		return &InputGameShortName{}, false, nil
	case uint32(0x5e068047):
		return &PageListOrderedItemText{}, false, nil
	case uint32(0xe48f964):
		return &UpdateBotInlineSend{}, false, nil
	case uint32(0x7a700873):
		return &SecureValueErrorFile{}, false, nil
	case uint32(0x176f8ba1):
		return &SendMessageGeoLocationAction{}, false, nil
	case uint32(0x353a686b):
		return &MessageFwdHeader{}, false, nil
	case uint32(0x4679b65f):
		return &AccessPointRule{}, false, nil
	case uint32(0xed8af74d):
		return &ChannelsAdminLogResults{}, false, nil
	case uint32(0x258aff05):
		return &KeyboardButtonUrl{}, false, nil
	case uint32(0xc7fb5e01):
		return &TextSuperscript{}, false, nil
	case uint32(0x98f6ac75):
		return &HelpPromoDataEmpty{}, false, nil
	case uint32(0x5a686d7c):
		return &ChatInviteAlready{}, false, nil
	case uint32(0x1da7158f):
		return &HelpAppUpdateObj{}, false, nil
	case uint32(0x68c13933):
		return &UpdateReadMessagesContents{}, false, nil
	case uint32(0x564fe691):
		return &UpdateLoginToken{}, false, nil
	case uint32(0x96a18d5):
		return &UploadFileObj{}, false, nil
	case uint32(0xa56c2a3e):
		return &UpdatesState{}, false, nil
	case uint32(0x4d5bbe0c):
		return &PrivacyValueAllowUsers{}, false, nil
	case uint32(0x193b4417):
		return &InputNotifyUsers{}, false, nil
	case uint32(0xe10db349):
		return &UpdateChatPinnedMessage{}, false, nil
	case uint32(0x6319d612):
		return &DocumentAttributeSticker{}, false, nil
	case uint32(0xb2ae9b0c):
		return &MessageActionChatDeleteUser{}, false, nil
	case uint32(0xf89cf5e8):
		return &MessageActionChatJoinedByLink{}, false, nil
	case uint32(0x64e475c2):
		return &MessageEntityEmail{}, false, nil
	case uint32(0xe062db83):
		return &InputMessagesFilterContacts{}, false, nil
	case uint32(0x20adaef8):
		return &InputPeerChannel{}, false, nil
	case uint32(0x53909779):
		return &ChannelAdminLogEventActionToggleSlowMode{}, false, nil
	case uint32(0x187fa0ca):
		return &SecureValue{}, false, nil
	case uint32(0x695150d7):
		return &MessageMediaPhoto{}, false, nil
	case uint32(0x9db1bc6d):
		return &PeerUser{}, false, nil
	case uint32(0xe4c123d6):
		return &InputGeoPointEmpty{}, false, nil
	case uint32(0x733f2961):
		return &PeerSettings{}, false, nil
	case uint32(0x3a912d4a):
		return &PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{}, false, nil
	case uint32(0x1c138d15):
		return &ContactsBlockedObj{}, false, nil
	case uint32(0xa1144770):
		return &SecureValueErrorTranslationFile{}, false, nil
	case uint32(0xb055eaee):
		return &MessageActionChannelMigrateFrom{}, false, nil
	case uint32(0xb4a2e88d):
		return &UpdateEncryption{}, false, nil
	case uint32(0xb3134d9d):
		return &ContactsFound{}, false, nil
	case uint32(0xe66fbf7b):
		return &InputMediaDice{}, false, nil
	case uint32(0x9b9240a6):
		return &UpdateBotWebhookJSONQuery{}, false, nil
	case uint32(0xe5d7d19c):
		return &MessagesChatFull{}, false, nil
	case uint32(0x823f649):
		return &MessagesVotesList{}, false, nil
	case uint32(0xdf969c2d):
		return &AuthExportedAuthorization{}, false, nil
	case uint32(0xeae87e42):
		return &ContactsContactsObj{}, false, nil
	case uint32(0xe26f42f1):
		return &UserStatusRecently{}, false, nil
	case uint32(0x58dbcab8):
		return &InputReportReasonSpam{}, false, nil
	case uint32(0x1e22c78d):
		return &InputReportReasonViolence{}, false, nil
	case uint32(0xe0c0c5e5):
		return &PageTableRow{}, false, nil
	case uint32(0x77608b83):
		return &KeyboardButtonRow{}, false, nil
	case uint32(0x7ef0dd87):
		return &InputMessagesFilterUrl{}, false, nil
	case uint32(0x7e6260d7):
		return &TextConcat{}, false, nil
	case uint32(0xfc2e05bc):
		return &ChatInviteExported{}, false, nil
	case uint32(0xe5bbfe1a):
		return &InputMediaPhotoExternal{}, false, nil
	case uint32(0xfa4f0bb5):
		return &InputFileBig{}, false, nil
	case uint32(0x25d6c9c7):
		return &UpdateReadChannelOutbox{}, false, nil
	case uint32(0x35e410a8):
		return &MessagesStickerSetInstallResultArchive{}, false, nil
	case uint32(0xadf44ee3):
		return &InputReportReasonChildAbuse{}, false, nil
	case uint32(0x8317c0c3):
		return &UpdateBotWebhookJSON{}, false, nil
	case uint32(0x5353e5a7):
		return &AuthSentCodeTypeCall{}, false, nil
	case uint32(0x635fe375):
		return &PhoneConnectionWebrtc{}, false, nil
	case uint32(0xbfb9f457):
		return &HelpPassportConfigNotModified{}, false, nil
	case uint32(0xb45c69d1):
		return &MessagesAffectedHistory{}, false, nil
	case uint32(0x2dc173c8):
		return &InputEncryptedFileBigUploaded{}, false, nil
	case uint32(0xc8edce1e):
		return &MessagesMessagesSlice{}, false, nil
	case uint32(0x2e0709a5):
		return &MessagesSavedGifsObj{}, false, nil
	case uint32(0x829d99da):
		return &SecureRequiredTypeObj{}, false, nil
	case uint32(0x19360dc0):
		return &UpdateFolderPeers{}, false, nil
	case uint32(0xe537ced6):
		return &SecureValueErrorSelfie{}, false, nil
	case uint32(0xe9baa668):
		return &FolderPeer{}, false, nil
	case uint32(0x1250abde):
		return &AccountAuthorizations{}, false, nil
	case uint32(0x1e36fded):
		return &InputPhoneCall{}, false, nil
	case uint32(0x56e9f0e4):
		return &InputMessagesFilterPhotoVideo{}, false, nil
	case uint32(0x98dd8936):
		return &PageListOrderedItemBlocks{}, false, nil
	case uint32(0xbad88395):
		return &InputMessageReplyTo{}, false, nil
	case uint32(0x1bfbd823):
		return &UpdateUserStatus{}, false, nil
	case uint32(0xd45ab096):
		return &PasswordKdfAlgoUnknown{}, false, nil
	case uint32(0x2d117597):
		return &InputUserFromMessage{}, false, nil
	case uint32(0x1f2b0afd):
		return &UpdateNewMessage{}, false, nil
	case uint32(0xa20db0e5):
		return &UpdateDeleteMessages{}, false, nil
	case uint32(0xa384b779):
		return &ReceivedNotifyMessage{}, false, nil
	case uint32(0x64bd0306):
		return &InputEncryptedFileUploaded{}, false, nil
	case uint32(0x9d4c17c0):
		return &PhoneConnectionObj{}, false, nil
	case uint32(0x7d748d04):
		return &DataJSON{}, false, nil
	case uint32(0xfa04579d):
		return &MessageEntityMention{}, false, nil
	case uint32(0x34b8621):
		return &TextMarked{}, false, nil
	case uint32(0x3334b0f0):
		return &InputSecureFileUploaded{}, false, nil
	case uint32(0x7da07ec9):
		return &InputPeerSelf{}, false, nil
	case uint32(0x4e90bfd6):
		return &UpdateMessageID{}, false, nil
	case uint32(0x4c81c1ba):
		return &InputPrivacyValueAllowChatParticipants{}, false, nil
	case uint32(0xf0e6672a):
		return &ChannelFull{}, false, nil
	case uint32(0x9e8fa6d3):
		return &MessagesFavedStickersNotModified{}, false, nil
	case uint32(0x51bdb021):
		return &MessageActionChatMigrateTo{}, false, nil
	case uint32(0x561bc879):
		return &ContactBlocked{}, false, nil
	case uint32(0x6724abc4):
		return &TextBold{}, false, nil
	case uint32(0x3dcd7a87):
		return &InputBotInlineMessageText{}, false, nil
	case uint32(0x70abc3fd):
		return &PageBlockTitle{}, false, nil
	case uint32(0x65a0fa4d):
		return &PageBlockCollage{}, false, nil
	case uint32(0x9de7a269):
		return &InputStickerSetID{}, false, nil
	case uint32(0xf7444763):
		return &JsonArray{}, false, nil
	case uint32(0x7c8fe7b6):
		return &PageBlockVideo{}, false, nil
	case uint32(0x21ec5a5f):
		return &SecurePlainEmail{}, false, nil
	case uint32(0x764cf810):
		return &BotInlineMessageMediaAuto{}, false, nil
	case uint32(0x27d69997):
		return &InputPeerPhotoFileLocation{}, false, nil
	case uint32(0x9c14984a):
		return &ThemeSettings{}, false, nil
	case uint32(0x3b6ddad2):
		return &PollAnswerVoters{}, false, nil
	case uint32(0xcac943f2):
		return &WebAuthorization{}, false, nil
	case uint32(0xbfd064ec):
		return &PageBlockHeader{}, false, nil
	case uint32(0xf259a80b):
		return &PageBlockEmbedPost{}, false, nil
	case uint32(0xedcdc05b):
		return &TopPeer{}, false, nil
	case uint32(0x5dab1af4):
		return &ExportedMessageLink{}, false, nil
	case uint32(0xbbf2dda0):
		return &SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000{}, false, nil
	case uint32(0x68e9916):
		return &AuthLoginTokenMigrateTo{}, false, nil
	case uint32(0xce4e82fd):
		return &InputMediaGeoLive{}, false, nil
	case uint32(0xd52f73f7):
		return &SendMessageRecordAudioAction{}, false, nil
	case uint32(0xf18cda44):
		return &UploadFileCdnRedirect{}, false, nil
	case uint32(0x9493ff32):
		return &MessagesSentEncryptedFile{}, false, nil
	case uint32(0x47a971e0):
		return &StatsURL{}, false, nil
	case uint32(0x3b5a3e40):
		return &ChannelAdminLogEvent{}, false, nil
	case uint32(0x18be796b):
		return &PrivacyValueAllowChatParticipants{}, false, nil
	case uint32(0x4a95e84e):
		return &InputNotifyChats{}, false, nil
	case uint32(0x7084a7be):
		return &UpdateContactsReset{}, false, nil
	case uint32(0x900802a1):
		return &ContactsBlockedSlice{}, false, nil
	case uint32(0x80e11a7f):
		return &MessageActionPhoneCall{}, false, nil
	case uint32(0x243e1c66):
		return &SendMessageUploadRoundAction{}, false, nil
	case uint32(0x5e002502):
		return &AuthSentCode{}, false, nil
	case uint32(0xec82e140):
		return &PhonePhoneCall{}, false, nil
	case uint32(0xc982eaba):
		return &CdnPublicKey{}, false, nil
	case uint32(0xe1746d0a):
		return &InputReportReasonOther{}, false, nil
	case uint32(0x25e073fc):
		return &PageListItemBlocks{}, false, nil
	case uint32(0x9e19a1f6):
		return &MessageService{}, false, nil
	case uint32(0x200250ba):
		return &UserEmpty{}, false, nil
	case uint32(0x5b38c6c1):
		return &InputMediaUploadedDocument{}, false, nil
	case uint32(0xa3b54985):
		return &ChannelParticipantsKicked{}, false, nil
	case uint32(0xa4bcc6fe):
		return &UpdatesChannelDifferenceTooLong{}, false, nil
	case uint32(0x9801d2f7):
		return &DocumentAttributeHasStickers{}, false, nil
	case uint32(0x9a3bfd99):
		return &MessagesHighScores{}, false, nil
	case uint32(0x17c6b5f6):
		return &HelpSupport{}, false, nil
	case uint32(0xc642724e):
		return &InputChatUploadedPhoto{}, false, nil
	case uint32(0xdfc2f58e):
		return &ChatInviteObj{}, false, nil
	case uint32(0x7761198):
		return &UpdateChatParticipants{}, false, nil
	case uint32(0x62ba04d9):
		return &UpdateNewChannelMessage{}, false, nil
	case uint32(0xd072acb4):
		return &RestrictionReason{}, false, nil
	case uint32(0x34566b6a):
		return &PageTableCell{}, false, nil
	case uint32(0x15ba6c40):
		return &MessagesDialogsObj{}, false, nil
	case uint32(0x36377430):
		return &MessageUserVoteInputOption{}, false, nil
	case uint32(0xaa1c39f):
		return &InputPaymentCredentialsApplePay{}, false, nil
	case uint32(0xf9d27a5a):
		return &UpdateInlineBotCallbackQuery{}, false, nil
	case uint32(0x666220e9):
		return &SecureValueErrorFiles{}, false, nil
	case uint32(0x86e18161):
		return &Poll{}, false, nil
	case uint32(0x8b73e763):
		return &PrivacyValueDisallowAll{}, false, nil
	case uint32(0x9c4e7e8b):
		return &MessageEntityUnderline{}, false, nil
	case uint32(0xa8fb1981):
		return &UpdatesDifferenceSlice{}, false, nil
	case uint32(0xfd5ec8f5):
		return &SendMessageCancelAction{}, false, nil
	case uint32(0xa927fec5):
		return &MessagesInactiveChats{}, false, nil
	case uint32(0xb4c83b4c):
		return &NotifyUsers{}, false, nil
	case uint32(0x50f41ccf):
		return &KeyboardButtonGame{}, false, nil
	case uint32(0x3a20ecb8):
		return &InputMessagesFilterChatPhotos{}, false, nil
	case uint32(0x3f6d7b68):
		return &JsonNull{}, false, nil
	case uint32(0xb8bc5b0c):
		return &InputNotifyPeerObj{}, false, nil
	case uint32(0x4e5f810d):
		return &PaymentsPaymentResultObj{}, false, nil
	case uint32(0xab7ec0a0):
		return &EncryptedChatEmpty{}, false, nil
	case uint32(0x39f23300):
		return &PageBlockCover{}, false, nil
	case uint32(0xe2d6e436):
		return &ChatParticipantAdmin{}, false, nil
	case uint32(0x86872538):
		return &InputMessagePinned{}, false, nil
	case uint32(0xf351d7ab):
		return &SendMessageUploadAudioAction{}, false, nil
	case uint32(0x98657f0d):
		return &Page{}, false, nil
	case uint32(0xe86602c3):
		return &MessagesAllStickersNotModified{}, false, nil
	case uint32(0xf52ff27f):
		return &InputFileObj{}, false, nil
	case uint32(0x11965f3a):
		return &BotInlineResultObj{}, false, nil
	case uint32(0xb6aef7b0):
		return &MessageActionEmpty{}, false, nil
	case uint32(0xe9763aec):
		return &SendMessageUploadVideoAction{}, false, nil
	case uint32(0x8427bbac):
		return &InputWallPaperNoFile{}, false, nil
	case uint32(0x43ae3dec):
		return &UpdateStickerSets{}, false, nil
	case uint32(0xd82363af):
		return &InputPrivacyValueDisallowChatParticipants{}, false, nil
	case uint32(0x22f3afb3):
		return &MessagesRecentStickersObj{}, false, nil
	case uint32(0xcbf24940):
		return &MessageMediaContact{}, false, nil
	case uint32(0x28703c8):
		return &InputStickerSetAnimatedEmoji{}, false, nil
	case uint32(0x4792929b):
		return &MessageActionScreenshotTaken{}, false, nil
	case uint32(0x87eabb53):
		return &PhoneCallRequested{}, false, nil
	case uint32(0xa9d6db1f):
		return &UrlAuthResultDefault{}, false, nil
	case uint32(0x23734b06):
		return &EncryptedMessageService{}, false, nil
	case uint32(0x3bb3b94a):
		return &InputPhotoObj{}, false, nil
	case uint32(0x9664f57f):
		return &InputMediaEmpty{}, false, nil
	case uint32(0x98a12b4b):
		return &UpdateChannelMessageViews{}, false, nil
	case uint32(0x909c3f94):
		return &PaymentRequestedInfo{}, false, nil
	case uint32(0xb74ba9d2):
		return &ContactsContactsNotModified{}, false, nil
	case uint32(0xe56dbf05):
		return &DialogPeerObj{}, false, nil
	case uint32(0xe9a734fa):
		return &PhotoCachedSize{}, false, nil
	case uint32(0xa229dd06):
		return &UpdateConfig{}, false, nil
	case uint32(0xf37f2f16):
		return &MessagesFavedStickersObj{}, false, nil
	case uint32(0x54b56617):
		return &WebPageAttribute{}, false, nil
	case uint32(0xeb1477e8):
		return &WebPageEmpty{}, false, nil
	case uint32(0xcad181f6):
		return &LangPackStringObj{}, false, nil
	case uint32(0xa437c3ed):
		return &WallPaperObj{}, false, nil
	case uint32(0xd1d34a26):
		return &SendMessageUploadPhotoAction{}, false, nil
	case uint32(0x15051f54):
		return &PhotosPhotosSlice{}, false, nil
	case uint32(0x8ea464b6):
		return &StatsGraphObj{}, false, nil
	case uint32(0x296f104):
		return &GeoPointObj{}, false, nil
	case uint32(0x5d75a138):
		return &UpdatesDifferenceEmpty{}, false, nil
	case uint32(0x417bbf11):
		return &InputBotInlineMessageMediaVenue{}, false, nil
	case uint32(0xa44f3ef6):
		return &PageBlockMap{}, false, nil
	case uint32(0x330b5424):
		return &UpdateReadChannelInbox{}, false, nil
	case uint32(0xc4b9f9bb):
		return &Error{}, false, nil
	case uint32(0x76768bed):
		return &PageBlockDetails{}, false, nil
	case uint32(0xa7332b73):
		return &UpdateUserName{}, false, nil
	case uint32(0x997c454a):
		return &PhoneCallAccepted{}, false, nil
	case uint32(0x7f077ad9):
		return &ContactsResolvedPeer{}, false, nil
	case uint32(0xd31a961e):
		return &Channel{}, false, nil
	case uint32(0x3f7ee58b):
		return &MessageMediaDice{}, false, nil
	case uint32(0xc12622c4):
		return &TextUnderline{}, false, nil
	case uint32(0x1ccb966a):
		return &TextPhone{}, false, nil
	case uint32(0x4b425864):
		return &InputBotInlineMessageGame{}, false, nil
	case uint32(0x263d7c26):
		return &PageBlockBlockquote{}, false, nil
	case uint32(0xedfd405f):
		return &MessagesAllStickersObj{}, false, nil
	case uint32(0x1710f156):
		return &UpdateEncryptedChatTyping{}, false, nil
	case uint32(0x38fe25b7):
		return &UpdateEncryptedMessagesRead{}, false, nil
	case uint32(0x8af40b25):
		return &WallPaperNoFile{}, false, nil
	case uint32(0x2a286531):
		return &InputChannelFromMessage{}, false, nil
	case uint32(0xb4608969):
		return &ChannelParticipantsAdmins{}, false, nil
	case uint32(0xd433ad73):
		return &IpPortObj{}, false, nil
	case uint32(0x1e287d04):
		return &InputMediaUploadedPhoto{}, false, nil
	case uint32(0xb6d45656):
		return &UpdateChannel{}, false, nil
	case uint32(0x1c199183):
		return &AccountWallPapersNotModified{}, false, nil
	case uint32(0x683a5e46):
		return &KeyboardButtonCallback{}, false, nil
	case uint32(0x66afa166):
		return &HelpDeepLinkInfoEmpty{}, false, nil
	case uint32(0x16115a96):
		return &PageBlockRelatedArticles{}, false, nil
	case uint32(0x72f0eaae):
		return &InputDocumentEmpty{}, false, nil
	case uint32(0x94d42ee7):
		return &ChannelMessagesFilterEmpty{}, false, nil
	case uint32(0x94bd38ed):
		return &MessageActionPinMessage{}, false, nil
	case uint32(0xafd93fbb):
		return &KeyboardButtonBuy{}, false, nil
	case uint32(0xb549da53):
		return &InputMessagesFilterRoundVideo{}, false, nil
	case uint32(0xc8d7493e):
		return &ChatParticipantObj{}, false, nil
	case uint32(0x4c43da18):
		return &UpdateUserPinnedMessage{}, false, nil
	case uint32(0x73924be0):
		return &MessageEntityPre{}, false, nil
	case uint32(0x560f8935):
		return &MessagesSentEncryptedMessageObj{}, false, nil
	case uint32(0x36f8c871):
		return &DocumentEmpty{}, false, nil
	case uint32(0xe8025ca2):
		return &MessagesSavedGifsNotModified{}, false, nil
	case uint32(0x4f11bae1):
		return &UserProfilePhotoEmpty{}, false, nil
	case uint32(0x7f676421):
		return &AccountThemesObj{}, false, nil
	case uint32(0x4a70994c):
		return &EncryptedFileObj{}, false, nil
	case uint32(0x5ce14175):
		return &PopularContact{}, false, nil
	case uint32(0xc30aa358):
		return &Invoice{}, false, nil
	case uint32(0x18cb9f78):
		return &HelpInviteText{}, false, nil
	case uint32(0x6ed02538):
		return &MessageEntityUrl{}, false, nil
	case uint32(0xe0277a62):
		return &SecureFileObj{}, false, nil
	case uint32(0xab0f6b1e):
		return &UpdatePhoneCall{}, false, nil
	case uint32(0x28ecf961):
		return &HelpTermsOfServiceUpdateObj{}, false, nil
	case uint32(0xf5235d55):
		return &InputEncryptedFileLocation{}, false, nil
	case uint32(0x36585ea4):
		return &MessagesBotCallbackAnswer{}, false, nil
	case uint32(0x500911e1):
		return &PaymentsPaymentReceipt{}, false, nil
	case uint32(0x3751b49e):
		return &InputMessagesFilterMusic{}, false, nil
	case uint32(0xe7026d0d):
		return &InputMessagesFilterGeo{}, false, nil
	case uint32(0x13567e8a):
		return &PageBlockUnsupported{}, false, nil
	case uint32(0x64ff9fd5):
		return &MessagesChatsObj{}, false, nil
	case uint32(0xeb49081d):
		return &RecentMeUrlChatInvite{}, false, nil
	case uint32(0xf911c994):
		return &Contact{}, false, nil
	case uint32(0xc007cec3):
		return &NotifyChats{}, false, nil
	case uint32(0xdc3d824f):
		return &TextEmpty{}, false, nil
	case uint32(0x629f1980):
		return &AuthLoginTokenObj{}, false, nil
	case uint32(0xef1751b5):
		return &PageBlockChannel{}, false, nil
	case uint32(0x26ffde7d):
		return &UpdateDialogFilter{}, false, nil
	case uint32(0xd95c6154):
		return &MessageActionSecureValuesSent{}, false, nil
	case uint32(0xf568028a):
		return &BankCardOpenUrl{}, false, nil
	case uint32(0x137948a5):
		return &AuthPasswordRecovery{}, false, nil
	case uint32(0xb17f890):
		return &MessagesRecentStickersNotModified{}, false, nil
	case uint32(0x1117dd5f):
		return &GeoPointEmpty{}, false, nil
	case uint32(0xd5676710):
		return &ChannelAdminLogEventActionParticipantToggleAdmin{}, false, nil
	case uint32(0x1b8f4ad1):
		return &PhoneCallWaiting{}, false, nil
	case uint32(0xf392b7f4):
		return &InputContact{}, false, nil
	case uint32(0xf3ae2eed):
		return &HelpUserInfoEmpty{}, false, nil
	case uint32(0x9c974fdf):
		return &UpdateReadHistoryInbox{}, false, nil
	case uint32(0xe40370a3):
		return &UpdateEditMessage{}, false, nil
	case uint32(0xffb62b95):
		return &InputStickerSetEmpty{}, false, nil
	case uint32(0x434bd2af):
		return &ChannelAdminLogEventActionChangePhoto{}, false, nil
	case uint32(0xc7345e6a):
		return &JsonBool{}, false, nil
	case uint32(0x4a27eb2d):
		return &StatsGraphAsync{}, false, nil
	case uint32(0xbedc9822):
		return &StatsGraphError{}, false, nil
	case uint32(0x83e5de54):
		return &MessageEmpty{}, false, nil
	case uint32(0xa676a322):
		return &InputMessageID{}, false, nil
	case uint32(0x5cc761bd):
		return &EmojiKeywordsDifference{}, false, nil
	case uint32(0xfc878fc8):
		return &PhoneCallProtocol{}, false, nil
	case uint32(0x1cc6e91f):
		return &InputSingleMedia{}, false, nil
	case uint32(0xd3680c61):
		return &ContactStatus{}, false, nil
	case uint32(0xe0cdc940):
		return &UpdateBotShippingQuery{}, false, nil
	case uint32(0x871fb939):
		return &UpdateGeoLiveViewed{}, false, nil
	case uint32(0x868a2aa5):
		return &SecureValueErrorReverseSide{}, false, nil
	case uint32(0x9a5c33e5):
		return &AccountPasswordSettings{}, false, nil
	case uint32(0xbc7fc6cd):
		return &FileLocation{}, false, nil
	case uint32(0xc27ac8c7):
		return &BotCommand{}, false, nil
	case uint32(0xfdb19008):
		return &MessageMediaGame{}, false, nil
	case uint32(0x78d4dec1):
		return &UpdateShort{}, false, nil
	case uint32(0x8c7f65e2):
		return &BotInlineMessageText{}, false, nil
	case uint32(0xb8d0afdf):
		return &AccountDaysTTL{}, false, nil
	case uint32(0x50f5c392):
		return &InputMessagesFilterVoice{}, false, nil
	case uint32(0xe844ebff):
		return &MessagesSearchCounter{}, false, nil
	case uint32(0x10b78d29):
		return &KeyboardButtonUrlAuth{}, false, nil
	case uint32(0xb98886cf):
		return &InputUserEmpty{}, false, nil
	case uint32(0xca05d50e):
		return &InputPaymentCredentialsAndroidPay{}, false, nil
	case uint32(0xaca1657b):
		return &UpdateMessagePoll{}, false, nil
	case uint32(0x1c0facaf):
		return &ChannelParticipantBanned{}, false, nil
	case uint32(0xe3309f7f):
		return &HelpTermsOfServiceUpdateEmpty{}, false, nil
	case uint32(0xcdc27a1f):
		return &PaymentSavedCredentials{}, false, nil
	case uint32(0x7438f7e8):
		return &DialogFilter{}, false, nil
	case uint32(0x33f0ea47):
		return &SecureCredentialsEncrypted{}, false, nil
	case uint32(0xb3fb5361):
		return &EmojiLanguage{}, false, nil
	case uint32(0xa8d864a7):
		return &InputBotInlineResultPhoto{}, false, nil
	case uint32(0xe4599bbd):
		return &MessagesStickersObj{}, false, nil
	case uint32(0x69df3769):
		return &ChatInviteEmpty{}, false, nil
	case uint32(0x1e148390):
		return &PageBlockKicker{}, false, nil
	case uint32(0xef02ce6):
		return &DocumentAttributeVideo{}, false, nil
	case uint32(0x18d1cdc2):
		return &BotInlineMessageMediaContact{}, false, nil
	case uint32(0x1e8caaeb):
		return &PostAddress{}, false, nil
	case uint32(0xe831c556):
		return &VideoSize{}, false, nil
	case uint32(0xb60a24a6):
		return &MessagesStickerSet{}, false, nil
	case uint32(0xe630b979):
		return &InputWallPaperObj{}, false, nil
	case uint32(0xee2bb969):
		return &UpdateDraftMessage{}, false, nil
	case uint32(0x2c171f72):
		return &DialogObj{}, false, nil
	case uint32(0x438865b):
		return &InputStickeredMediaDocument{}, false, nil
	case uint32(0x61695cb0):
		return &ChatInvitePeek{}, false, nil
	case uint32(0x6e5f8c22):
		return &UpdateChatParticipantDelete{}, false, nil
	case uint32(0x9375341e):
		return &UpdateSavedGifs{}, false, nil
	case uint32(0x56022f4d):
		return &UpdateLangPack{}, false, nil
	case uint32(0x90110467):
		return &InputPrivacyValueDisallowUsers{}, false, nil
	case uint32(0x58fffcd0):
		return &HighScore{}, false, nil
	case uint32(0xffa0a496):
		return &InputStickerSetItem{}, false, nil
	case uint32(0xcb296bf8):
		return &LabeledPrice{}, false, nil
	case uint32(0x5086cf8):
		return &WallPaperSettings{}, false, nil
	case uint32(0xe73547e1):
		return &UpdateBotCallbackQuery{}, false, nil
	case uint32(0xe8a40bd9):
		return &SecureValueErrorData{}, false, nil
	case uint32(0xb722de65):
		return &BotInlineMessageMediaGeo{}, false, nil
	case uint32(0x7311ca11):
		return &WebPageNotModified{}, false, nil
	case uint32(0xacae0690):
		return &PrivacyValueDisallowChatParticipants{}, false, nil
	case uint32(0x352dca58):
		return &MessageEntityMentionName{}, false, nil
	case uint32(0x761e6af4):
		return &MessageEntityBankCard{}, false, nil
	case uint32(0x86471d92):
		return &SecurePasswordKdfAlgoSHA512{}, false, nil
	case uint32(0xea4b0e5c):
		return &UpdateChatParticipantAdd{}, false, nil
	case uint32(0x1b3f4df7):
		return &UpdateEditChannelMessage{}, false, nil
	case uint32(0xf89777f2):
		return &ChannelAdminLogEventActionParticipantLeave{}, false, nil
	case uint32(0xcb43acde):
		return &StatsAbsValueAndPrev{}, false, nil
	case uint32(0x56e0d474):
		return &MessageMediaGeo{}, false, nil
	case uint32(0x702b65a9):
		return &AccountWallPapersObj{}, false, nil
	case uint32(0x656ac4b):
		return &ChannelParticipantsSearch{}, false, nil
	case uint32(0xcbce2fe0):
		return &StatsPercentValue{}, false, nil
	case uint32(0xad4fc9bd):
		return &MessageInteractionCounters{}, false, nil
	case uint32(0x5fb224d5):
		return &ChatAdminRights{}, false, nil
	case uint32(0x9b69e34b):
		return &MessageEntityPhone{}, false, nil
	case uint32(0x9ba2d800):
		return &ChatEmpty{}, false, nil
	case uint32(0x74ae4240):
		return &UpdatesObj{}, false, nil
	case uint32(0x44747e9a):
		return &AuthAuthorizationSignUpRequired{}, false, nil
	case uint32(0x9c95f7bb):
		return &InputPeerChannelFromMessage{}, false, nil
	case uint32(0x861cc8a0):
		return &InputStickerSetShortName{}, false, nil
	case uint32(0xb5a1ce5a):
		return &MessageActionChatEditTitle{}, false, nil
	case uint32(0x8742ae7f):
		return &PhoneCallObj{}, false, nil
	case uint32(0xdebebe83):
		return &CodeSettings{}, false, nil
	case uint32(0x28f1114):
		return &Theme{}, false, nil
	case uint32(0x826f8b60):
		return &MessageEntityItalic{}, false, nil
	case uint32(0x9f2221c9):
		return &InputWebFileGeoPointLocation{}, false, nil
	case uint32(0xe17e23c):
		return &PhotoSizeEmpty{}, false, nil
	case uint32(0xa3289a6d):
		return &ChannelParticipantSelf{}, false, nil
	case uint32(0x8a86659c):
		return &BotInlineMessageMediaVenue{}, false, nil
	case uint32(0xc586da1c):
		return &WebPagePending{}, false, nil
	case uint32(0xc239d686):
		return &InputWebFileLocationObj{}, false, nil
	case uint32(0xe0b0bc2e):
		return &PhotoStrippedSize{}, false, nil
	case uint32(0x184b35ce):
		return &InputPrivacyValueAllowAll{}, false, nil
	case uint32(0xe9e82c18):
		return &ChannelAdminLogEventActionUpdatePinned{}, false, nil
	case uint32(0xbad07584):
		return &InputDocumentFileLocation{}, false, nil
	case uint32(0xe04232f3):
		return &AutoDownloadSettings{}, false, nil
	case uint32(0x77ebc742):
		return &UserStatusLastMonth{}, false, nil
	case uint32(0xfa0f3ca2):
		return &UpdatePinnedDialogs{}, false, nil
	case uint32(0x37c1011c):
		return &ChatPhotoEmpty{}, false, nil
	case uint32(0xe6dfb825):
		return &ChannelAdminLogEventActionChangeTitle{}, false, nil
	case uint32(0x6c37c15c):
		return &DocumentAttributeImageSize{}, false, nil
	case uint32(0xbc0a57dc):
		return &RecentMeUrlStickerSet{}, false, nil
	case uint32(0x76a6d327):
		return &MessageEntityTextUrl{}, false, nil
	case uint32(0x9bf8bb95):
		return &TextStrike{}, false, nil
	case uint32(0x3380c786):
		return &InputBotInlineMessageMediaAuto{}, false, nil
	case uint32(0x40771900):
		return &UpdateChannelWebPage{}, false, nil
	case uint32(0x55188a2e):
		return &ChannelAdminLogEventActionChangeAbout{}, false, nil
	case uint32(0x488a7337):
		return &MessageActionChatAddUser{}, false, nil
	case uint32(0xad01d61d):
		return &Authorization{}, false, nil
	case uint32(0x75588b3f):
		return &InputClientProxy{}, false, nil
	case uint32(0x38641628):
		return &MessagesStickerSetInstallResultSuccess{}, false, nil
	case uint32(0xb6901959):
		return &UpdateChatParticipantAdmin{}, false, nil
	case uint32(0x5a592a6c):
		return &HelpConfigSimple{}, false, nil
	case uint32(0x3371c354):
		return &MessagesPeerDialogs{}, false, nil
	case uint32(0xafeb712e):
		return &InputChannelObj{}, false, nil
	case uint32(0x869d758f):
		return &SecureValueErrorObj{}, false, nil
	case uint32(0x21e753bc):
		return &UploadWebFile{}, false, nil
	case uint32(0x6410a5d2):
		return &StickerSetCoveredObj{}, false, nil
	case uint32(0xc0e24635):
		return &MessagesDhConfigNotModified{}, false, nil
	case uint32(0x12bcbd9a):
		return &UpdateNewEncryptedMessage{}, false, nil
	case uint32(0x1b7907ae):
		return &ChannelAdminLogEventActionToggleInvites{}, false, nil
	case uint32(0x8f31b327):
		return &MessageActionPaymentSentMe{}, false, nil
	case uint32(0xd83466f3):
		return &InputPhotoLegacyFileLocation{}, false, nil
	case uint32(0xffc86587):
		return &InputMessagesFilterGif{}, false, nil
	case uint32(0x2064674e):
		return &UpdatesChannelDifferenceObj{}, false, nil
	case uint32(0x3dbb5986):
		return &AuthSentCodeTypeApp{}, false, nil
	case uint32(0xdbaeae9):
		return &InputStickerSetThumb{}, false, nil
	case uint32(0xad2e1cd8):
		return &AccountAuthorizationForm{}, false, nil
	case uint32(0x3f56aea3):
		return &PaymentsPaymentForm{}, false, nil
	case uint32(0xfffe1bac):
		return &PrivacyValueAllowContacts{}, false, nil
	case uint32(0x3c2884c1):
		return &TextUrl{}, false, nil
	case uint32(0xebe46819):
		return &UpdateServiceNotification{}, false, nil
	case uint32(0x16bf744e):
		return &SendMessageTypingAction{}, false, nil
	case uint32(0xef7ff916):
		return &StatsMegagroupStats{}, false, nil
	case uint32(0xa32dd600):
		return &MessageMediaWebPage{}, false, nil
	case uint32(0x6c47ac9f):
		return &LangPackStringPluralized{}, false, nil
	case uint32(0x2be0dfa4):
		return &JsonNumber{}, false, nil
	case uint32(0x37982646):
		return &IpPortSecret{}, false, nil
	case uint32(0x8f8c0e4e):
		return &UrlAuthResultAccepted{}, false, nil
	case uint32(0xea107ae4):
		return &ChannelAdminLogEventsFilter{}, false, nil
	case uint32(0x9f120418):
		return &ChatBannedRights{}, false, nil
	case uint32(0x18b7a10d):
		return &DcOption{}, false, nil
	case uint32(0xfcaafeb7):
		return &InputDialogPeerObj{}, false, nil
	case uint32(0x7bf09fc):
		return &UserStatusLastWeek{}, false, nil
	case uint32(0x1abfb575):
		return &InputDocumentObj{}, false, nil
	case uint32(0x63cacf26):
		return &AccountAutoDownloadSettings{}, false, nil
	case uint32(0xf9c8bcc6):
		return &WebDocumentNoProxy{}, false, nil
	case uint32(0x92d33a0e):
		return &UrlAuthResultRequest{}, false, nil
	case uint32(0xdb20b188):
		return &PageBlockDivider{}, false, nil
	case uint32(0x17bae2e6):
		return &InputPeerUserFromMessage{}, false, nil
	case uint32(0x46e1d13d):
		return &RecentMeUrlUnknown{}, false, nil
	case uint32(0xa01b22f9):
		return &RecentMeUrlChat{}, false, nil
	case uint32(0x1d1b1245):
		return &InputAppEvent{}, false, nil
	case uint32(0x84551347):
		return &MessageMediaInvoice{}, false, nil
	case uint32(0x62718a82):
		return &EncryptedChatRequested{}, false, nil
	case uint32(0x9a65ea1f):
		return &UpdateChatUserTyping{}, false, nil
	case uint32(0x65d2b464):
		return &UpdateChannelParticipant{}, false, nil
	case uint32(0xed56c9fc):
		return &AccountWebAuthorizations{}, false, nil
	case uint32(0xd0028438):
		return &ImportedContact{}, false, nil
	case uint32(0xbd507cd1):
		return &InputThemeSettings{}, false, nil
	case uint32(0x95d2ac92):
		return &MessageActionChannelCreate{}, false, nil
	case uint32(0xae30253):
		return &MessageRange{}, false, nil
	case uint32(0x27477b4):
		return &SecureRequiredTypeOneOf{}, false, nil
	case uint32(0x7b8e7de6):
		return &InputPeerUser{}, false, nil
	case uint32(0x54826690):
		return &UpdateBotInlineQuery{}, false, nil
	case uint32(0xe511996d):
		return &UpdateFavedStickers{}, false, nil
	case uint32(0x8e1a1775):
		return &NearestDc{}, false, nil
	case uint32(0xed1ecdb0):
		return &SecureValueHash{}, false, nil
	case uint32(0x9fd40bd8):
		return &NotifyPeerObj{}, false, nil
	case uint32(0x1759c560):
		return &PageBlockPhoto{}, false, nil
	case uint32(0xc6dc0c66):
		return &MessagesFeaturedStickersNotModified{}, false, nil
	case uint32(0xc45a6536):
		return &HelpNoAppUpdate{}, false, nil
	case uint32(0xfae69f56):
		return &MessageActionCustomAction{}, false, nil
	case uint32(0x4bd6e798):
		return &MessageMediaPoll{}, false, nil
	case uint32(0x514519e2):
		return &DialogPeerFolder{}, false, nil
	case uint32(0xf4e096c3):
		return &InputMediaInvoice{}, false, nil
	case uint32(0xb0d1865b):
		return &ChannelParticipantsBots{}, false, nil
	case uint32(0x18f3d0f7):
		return &StatsGroupTopPoster{}, false, nil
	case uint32(0xd8411139):
		return &PaymentsPaymentVerificationNeeded{}, false, nil
	case uint32(0xb637edaf):
		return &StatsDateRangeDays{}, false, nil
	case uint32(0x5a17b5e5):
		return &InputEncryptedFileObj{}, false, nil
	case uint32(0xedb93949):
		return &UserStatusOnline{}, false, nil
	case uint32(0xa8718dc5):
		return &PageBlockEmbed{}, false, nil
	case uint32(0xc10eb2cf):
		return &InputPaymentCredentialsSaved{}, false, nil
	case uint32(0xd66b66c9):
		return &InputPrivacyValueDisallowAll{}, false, nil
	case uint32(0xff544e65):
		return &Folder{}, false, nil
	case uint32(0xaed6dbb2):
		return &MaskCoords{}, false, nil
	case uint32(0x71e094f3):
		return &MessagesDialogsSlice{}, false, nil
	case uint32(0xa28e5559):
		return &MessageUserVoteObj{}, false, nil
	case uint32(0x6e6fe51c):
		return &UpdateDialogPinned{}, false, nil
	case uint32(0x42e047bb):
		return &ChannelAdminLogEventActionDeleteMessage{}, false, nil
	case uint32(0xd1451883):
		return &PaymentsValidatedRequestedInfo{}, false, nil
	case uint32(0xeeca5ce3):
		return &LangPackLanguage{}, false, nil
	case uint32(0xf1749a22):
		return &MessagesStickersNotModified{}, false, nil
	case uint32(0x46560264):
		return &UpdateLangPackTooLong{}, false, nil
	case uint32(0xb16a6c29):
		return &KeyboardButtonRequestPhone{}, false, nil
	case uint32(0x7d6099dd):
		return &SecurePlainPhone{}, false, nil
	case uint32(0xd5b3b9f9):
		return &EmojiKeywordObj{}, false, nil
	case uint32(0x9852f9c6):
		return &DocumentAttributeAudio{}, false, nil
	case uint32(0x98e81d3a):
		return &BotInfo{}, false, nil
	case uint32(0xfbd2c296):
		return &InputFolderPeer{}, false, nil
	case uint32(0xf141b5e1):
		return &InputEncryptedChat{}, false, nil
	case uint32(0x31962a4c):
		return &StatsGroupTopInviter{}, false, nil
	case uint32(0x1ca48f57):
		return &InputChatPhotoEmpty{}, false, nil
	case uint32(0x3354678f):
		return &UpdatePtsChanged{}, false, nil
	case uint32(0x89893b45):
		return &UpdateChannelReadMessagesContents{}, false, nil
	case uint32(0xa03e5b85):
		return &ReplyKeyboardHide{}, false, nil
	case uint32(0xf385c1f6):
		return &LangPackDifference{}, false, nil
	case uint32(0x3c20629f):
		return &InlineBotSwitchPM{}, false, nil
	case uint32(0x88bf9319):
		return &InputBotInlineResultObj{}, false, nil
	case uint32(0x390d5c5e):
		return &AuthLoginTokenSuccess{}, false, nil
	case uint32(0x5108d648):
		return &MessagesFoundStickerSetsObj{}, false, nil
	case uint32(0xcd050916):
		return &AuthAuthorizationObj{}, false, nil
	case uint32(0x8dbc3336):
		return &RecentMeUrlUser{}, false, nil
	case uint32(0x6c3f19b9):
		return &TextFixed{}, false, nil
	case uint32(0x35553762):
		return &TextAnchor{}, false, nil
	case uint32(0x467a0766):
		return &PageBlockParagraph{}, false, nil
	case uint32(0xf3f25f76):
		return &MessageActionContactSignUp{}, false, nil
	case uint32(0xbdf78394):
		return &StatsBroadcastStats{}, false, nil
	case uint32(0x208e68c9):
		return &InputMessageEntityMentionName{}, false, nil
	case uint32(0xbf0693d4):
		return &MessageEntityStrike{}, false, nil
	case uint32(0xfc796b3f):
		return &KeyboardButtonRequestGeoLocation{}, false, nil
	case uint32(0x7a7c17a4):
		return &InputMessagesFilterRoundVoice{}, false, nil
	case uint32(0x1e87342b):
		return &DocumentObj{}, false, nil
	case uint32(0x2661bf09):
		return &UpdatePhoneCallSignalingData{}, false, nil
	case uint32(0x88f27fbc):
		return &SendMessageRecordRoundAction{}, false, nil
	case uint32(0x8dca6aa5):
		return &PhotosPhotosObj{}, false, nil
	case uint32(0x6242c773):
		return &FileHash{}, false, nil
	case uint32(0x9b89f93a):
		return &InputReportReasonCopyright{}, false, nil
	case uint32(0xfb52dc99):
		return &InputMediaDocumentExternal{}, false, nil
	case uint32(0xc000bba2):
		return &AuthSentCodeTypeSms{}, false, nil
	case uint32(0x1b287353):
		return &MessageActionSecureValuesSentMe{}, false, nil
	case uint32(0x77744d4a):
		return &DialogFilterSuggested{}, false, nil
	case uint32(0x780a0310):
		return &HelpTermsOfService{}, false, nil
	case uint32(0x9609a51c):
		return &InputMessagesFilterPhotos{}, false, nil
	case uint32(0xde5a0dd6):
		return &TextEmail{}, false, nil
	case uint32(0xed6a8504):
		return &TextSubscript{}, false, nil
	case uint32(0xbfb5ad8b):
		return &ChannelLocationEmpty{}, false, nil
	case uint32(0x2df5fc0a):
		return &ChannelAdminLogEventActionDefaultBannedRights{}, false, nil
	case uint32(0x5725e40a):
		return &CdnConfig{}, false, nil
	case uint32(0xfb8fe43c):
		return &PaymentsSavedInfo{}, false, nil
	case uint32(0x64600527):
		return &InputDialogPeerFolder{}, false, nil
	case uint32(0x7f3b18ea):
		return &InputPeerEmpty{}, false, nil
	case uint32(0xbe3dfa):
		return &SecureValueErrorFrontSide{}, false, nil
	case uint32(0xe6d83d7e):
		return &ChannelAdminLogEventActionParticipantToggleBan{}, false, nil
	case uint32(0xdfdaabe1):
		return &InputFileLocationObj{}, false, nil
	case uint32(0xcbc7ee28):
		return &InputSecureFileLocation{}, false, nil
	case uint32(0xbad0e5bb):
		return &PeerChat{}, false, nil
	case uint32(0x236df622):
		return &EmojiKeywordDeleted{}, false, nil
	case uint32(0x725b04c3):
		return &UpdatesCombined{}, false, nil
	case uint32(0xc37521c9):
		return &UpdateDeleteChannelMessages{}, false, nil
	case uint32(0x15590068):
		return &DocumentAttributeFilename{}, false, nil
	case uint32(0x95e3fbef):
		return &MessageActionChatDeletePhoto{}, false, nil
	case uint32(0x6ca9c2e9):
		return &PollAnswer{}, false, nil
	case uint32(0xdb21d0a7):
		return &InputSecureValue{}, false, nil
	case uint32(0x1527bcac):
		return &SecureSecretSettings{}, false, nil
	case uint32(0x5367e5be):
		return &InputSecureFileObj{}, false, nil
	case uint32(0xda13538a):
		return &ChatParticipantCreator{}, false, nil
	case uint32(0x70db6837):
		return &UpdateChannelAvailableMessages{}, false, nil
	case uint32(0xb4afcfb0):
		return &UpdatePeerLocated{}, false, nil
	case uint32(0xaf509d20):
		return &PeerNotifySettings{}, false, nil
	case uint32(0x77d01c3b):
		return &ContactsImportedContacts{}, false, nil
	case uint32(0xa99fca4f):
		return &UploadCdnFileObj{}, false, nil
	case uint32(0xa2fa4880):
		return &KeyboardButtonObj{}, false, nil
	case uint32(0x688a30aa):
		return &UpdateNewStickerSet{}, false, nil
	case uint32(0x209b82db):
		return &ChannelLocationObj{}, false, nil
	case uint32(0x183040d3):
		return &ChannelAdminLogEventActionParticipantJoin{}, false, nil
	case uint32(0xdb64fd34):
		return &AccountTmpPassword{}, false, nil
	case uint32(0xea02c27e):
		return &PaymentCharge{}, false, nil
	case uint32(0x6f635b0d):
		return &MessageEntityHashtag{}, false, nil
	case uint32(0x3e11affb):
		return &UpdatesChannelDifferenceEmpty{}, false, nil
	case uint32(0xce0d37b0):
		return &PageBlockAnchor{}, false, nil
	case uint32(0x9cd81144):
		return &MessagesChatsSlice{}, false, nil
	case uint32(0x8e5e9873):
		return &UpdateDcOptions{}, false, nil
	case uint32(0x98592475):
		return &UpdateChannelPinnedMessage{}, false, nil
	case uint32(0xe67f520e):
		return &InputStickerSetDice{}, false, nil
	case uint32(0xf3b7acc9):
		return &InputGeoPointObj{}, false, nil
	case uint32(0x811f854f):
		return &AccountSentEmailCode{}, false, nil
	case uint32(0xa6edbffd):
		return &InputBotInlineMessageMediaContact{}, false, nil
	case uint32(0xbec268ef):
		return &UpdateNotifySettings{}, false, nil
	case uint32(0xee3b272a):
		return &UpdatePrivacy{}, false, nil
	case uint32(0xa26f881b):
		return &ChannelAdminLogEventActionChangeLinkedChat{}, false, nil
	case uint32(0x84d19185):
		return &MessagesAffectedMessages{}, false, nil
	case uint32(0xbd610bc9):
		return &MessageEntityBold{}, false, nil
	case uint32(0x81ccf4f):
		return &TextImage{}, false, nil
	case uint32(0xb1db7c7e):
		return &InputNotifyBroadcasts{}, false, nil
	case uint32(0xb52c939d):
		return &ContactsTopPeersDisabled{}, false, nil
	case uint32(0x9a8ae1e1):
		return &PageBlockOrderedList{}, false, nil
	case uint32(0xdbd4feed):
		return &InputReportReasonGeoIrrelevant{}, false, nil
	case uint32(0xf8ab7dfb):
		return &InputMediaContact{}, false, nil
	case uint32(0xd33f43f3):
		return &InputMediaGame{}, false, nil
	case uint32(0x131cc67f):
		return &InputPrivacyValueAllowUsers{}, false, nil
	case uint32(0x709b2405):
		return &ChannelAdminLogEventActionEditMessage{}, false, nil
	case uint32(0x5f5c95f1):
		return &ChannelAdminLogEventActionTogglePreHistoryHidden{}, false, nil
	case uint32(0x29be5899):
		return &InputTakeoutFileLocation{}, false, nil
	case uint32(0x8aeabec3):
		return &SecureData{}, false, nil
	case uint32(0xf94e5f1):
		return &InputMediaPoll{}, false, nil
	case uint32(0x8216fba3):
		return &UpdateTheme{}, false, nil
	case uint32(0xfb197a65):
		return &PhotoObj{}, false, nil
	case uint32(0x50a04e45):
		return &AccountPrivacyRules{}, false, nil
	case uint32(0xf41eb622):
		return &AccountThemesNotModified{}, false, nil
	case uint32(0x914fbf11):
		return &UpdateShortMessage{}, false, nil
	case uint32(0xb3ba0635):
		return &InputMediaPhoto{}, false, nil
	case uint32(0xa5d72105):
		return &UpdateDialogFilterOrder{}, false, nil
	case uint32(0xf4108aa0):
		return &ReplyKeyboardForceReply{}, false, nil
	case uint32(0xedf17c12):
		return &UserFull{}, false, nil
	case uint32(0x1c570ed1):
		return &WebDocumentObj{}, false, nil
	case uint32(0x28a20571):
		return &MessageEntityCode{}, false, nil
	case uint32(0x20df5d0):
		return &MessageEntityBlockquote{}, false, nil
	case uint32(0xf9c44144):
		return &InputMediaGeoPoint{}, false, nil
	case uint32(0xbb2d201):
		return &UpdateStickerSetsOrder{}, false, nil
	case uint32(0x71bd134c):
		return &DialogFolder{}, false, nil
	case uint32(0xbdf9653b):
		return &Game{}, false, nil
	case uint32(0x1142bd56):
		return &SavedContact{}, false, nil
	case uint32(0xbaafe5e0):
		return &PageBlockAuthorDate{}, false, nil
	case uint32(0xf12bb6e1):
		return &PageBlockSubheader{}, false, nil
	case uint32(0x80ece81a):
		return &UpdateUserBlocked{}, false, nil
	case uint32(0x90866cee):
		return &UpdateDeleteScheduledMessages{}, false, nil
	case uint32(0x17db940b):
		return &BotInlineMediaResult{}, false, nil
	case uint32(0x40699cd0):
		return &MessageActionPaymentSent{}, false, nil
	case uint32(0x40181ffe):
		return &InputPhotoFileLocation{}, false, nil
	case uint32(0x9f84f49e):
		return &MessageMediaUnsupported{}, false, nil
	case uint32(0x9fc00e65):
		return &InputMessagesFilterVideo{}, false, nil
	case uint32(0x4a8537):
		return &SecurePasswordKdfAlgoUnknown{}, false, nil
	case uint32(0x23ab23d2):
		return &InputMediaDocument{}, false, nil
	case uint32(0xc0de1bd9):
		return &JSONObjectValue{}, false, nil
	case uint32(0x8ffa9a1f):
		return &PageBlockSubtitle{}, false, nil
	case uint32(0xca461b5d):
		return &PeerLocatedObj{}, false, nil
	case uint32(0x48a30254):
		return &ReplyInlineMarkup{}, false, nil
	case uint32(0x20212ca8):
		return &PhotosPhoto{}, false, nil
	case uint32(0xe0310d7):
		return &HelpRecentMeUrls{}, false, nil
	case uint32(0x65427b82):
		return &PrivacyValueAllowAll{}, false, nil
	case uint32(0x4f4456d3):
		return &PageBlockPullquote{}, false, nil
	case uint32(0x77bfb61b):
		return &PhotoSizeObj{}, false, nil
	case uint32(0xed18c118):
		return &EncryptedMessageObj{}, false, nil
	case uint32(0xe6b76ae):
		return &ChannelAdminLogEventActionChangeLocation{}, false, nil
	case uint32(0x947ca848):
		return &MessagesBotResults{}, false, nil
	case uint32(0xfff8fdc4):
		return &InputBotInlineResultDocument{}, false, nil
	case uint32(0x6a4ee832):
		return &HelpDeepLinkInfoObj{}, false, nil
	case uint32(0x50ca4de1):
		return &PhoneCallDiscarded{}, false, nil
	case uint32(0xbb92ba95):
		return &MessageEntityUnknown{}, false, nil
	case uint32(0x3bda1bde):
		return &ChatObj{}, false, nil
	case uint32(0x9880f658):
		return &InputCheckPasswordEmpty{}, false, nil
	case uint32(0xf7c1b13f):
		return &InputUserSelf{}, false, nil
	case uint32(0xfc900c2b):
		return &ChatParticipantsForbidden{}, false, nil
	case uint32(0xba52007):
		return &InputPrivacyValueDisallowContacts{}, false, nil
	case uint32(0x1b0c841a):
		return &DraftMessageEmpty{}, false, nil
	case uint32(0xa098d6af):
		return &HelpPassportConfigObj{}, false, nil
	case uint32(0x64199744):
		return &SecureFileEmpty{}, false, nil
	case uint32(0x8c718e87):
		return &MessagesMessagesObj{}, false, nil
	case uint32(0x3bf703dc):
		return &EncryptedChatWaiting{}, false, nil
	case uint32(0x11f1331c):
		return &UpdateShortSentMessage{}, false, nil
	case uint32(0x3417d728):
		return &InputPaymentCredentialsObj{}, false, nil
	case uint32(0xeb0467fb):
		return &UpdateChannelTooLong{}, false, nil
	case uint32(0x3c5693e9):
		return &InputThemeObj{}, false, nil
	case uint32(0xa6638b9a):
		return &MessageActionChatCreate{}, false, nil
	case uint32(0x92a72876):
		return &MessageActionGameScore{}, false, nil
	case uint32(0x3e24e573):
		return &PaymentsBankCardData{}, false, nil
	case uint32(0x4fa417f2):
		return &InputBotInlineResultGame{}, false, nil
	case uint32(0xc1b15d65):
		return &InputBotInlineMessageMediaGeo{}, false, nil
	case uint32(0xee8c1e86):
		return &InputChannelEmpty{}, false, nil
	case uint32(0xf56ee2a8):
		return &ChannelsChannelParticipantsObj{}, false, nil
	case uint32(0xcd77d957):
		return &ChannelMessagesFilterObj{}, false, nil
	case uint32(0x8c05f1c9):
		return &HelpSupportName{}, false, nil
	case uint32(0x80c99768):
		return &InputMessagesFilterPhoneCalls{}, false, nil
	case uint32(0xb390dc08):
		return &PageRelatedArticle{}, false, nil
	case uint32(0xbddde532):
		return &PeerChannel{}, false, nil
	case uint32(0xb92fb6cd):
		return &PageListItemText{}, false, nil
	case uint32(0xd8292816):
		return &InputUserObj{}, false, nil
	case uint32(0x8c703f):
		return &UserStatusOffline{}, false, nil
	case uint32(0x452c0e65):
		return &MessageObj{}, false, nil
	case uint32(0x804361ea):
		return &PageBlockAudio{}, false, nil
	case uint32(0x571d2742):
		return &UpdateReadFeaturedStickers{}, false, nil
	case uint32(0x54c01850):
		return &UpdateChatDefaultBannedRights{}, false, nil
	case uint32(0xbadcc1a3):
		return &PollResults{}, false, nil
	case uint32(0x32c3e77):
		return &InputGameID{}, false, nil
	case uint32(0x744694e0):
		return &TextPlain{}, false, nil
	case uint32(0x2c221edd):
		return &MessagesDhConfigObj{}, false, nil
	case uint32(0x2e59d922):
		return &InputReportReasonPornography{}, false, nil
	case uint32(0x9a422c20):
		return &UpdateRecentStickers{}, false, nil
	case uint32(0xde3f3c79):
		return &ChannelParticipantsRecent{}, false, nil
	case uint32(0xf5890df1):
		return &InputThemeSlug{}, false, nil
	case uint32(0x9fbab604):
		return &MessageActionHistoryClear{}, false, nil
	case uint32(0xd0d9b163):
		return &ChannelsChannelParticipant{}, false, nil
	case uint32(0x72091c80):
		return &InputWallPaperSlug{}, false, nil
	case uint32(0xe8fe0de):
		return &MessageUserVoteMultiple{}, false, nil
	case uint32(0x9d05049):
		return &UserStatusEmpty{}, false, nil
	case uint32(0x12b299d4):
		return &StickerPack{}, false, nil
	case uint32(0xe89c45b2):
		return &WebPageObj{}, false, nil
	case uint32(0xd54b65d):
		return &MessagesFoundStickerSetsNotModified{}, false, nil
	case uint32(0x1cd7bf0d):
		return &InputPhotoEmpty{}, false, nil
	case uint32(0x4fcba9c8):
		return &MessagesArchivedStickers{}, false, nil
	case uint32(0xc1f8e69a):
		return &InputMessagesFilterMyMentions{}, false, nil
	case uint32(0x99262e37):
		return &MessagesChannelMessages{}, false, nil
	case uint32(0xf49ca0):
		return &UpdatesDifferenceObj{}, false, nil
	case uint32(0x4afe8f6d):
		return &UpdatesDifferenceTooLong{}, false, nil
	case uint32(0xc13d1c11):
		return &InputMediaVenue{}, false, nil
	case uint32(0xe16459c3):
		return &UpdateDialogUnreadMark{}, false, nil
	case uint32(0x42f88f2c):
		return &UpdateMessagePollVote{}, false, nil
	case uint32(0x1b7c9db3):
		return &ChatFullObj{}, false, nil
	case uint32(0xab03c6d9):
		return &AuthSentCodeTypeFlashCall{}, false, nil
	case uint32(0x5366c915):
		return &PhoneCallEmpty{}, false, nil
	case uint32(0x9c3d198e):
		return &InputPeerNotifySettings{}, false, nil
	case uint32(0x99c1d49d):
		return &JsonObject{}, false, nil
	case uint32(0x48870999):
		return &PageBlockFooter{}, false, nil
	case uint32(0xe317af7e):
		return &UpdatesTooLong{}, false, nil
	case uint32(0x5c486927):
		return &UpdateUserTyping{}, false, nil
	case uint32(0x7f891213):
		return &UpdateWebPage{}, false, nil
	case uint32(0x2331b22d):
		return &PhotoEmpty{}, false, nil
	case uint32(0xf8ec284b):
		return &PeerSelfLocated{}, false, nil
	case uint32(0xd612e8ef):
		return &NotifyBroadcasts{}, false, nil
	case uint32(0xb71e767a):
		return &JsonString{}, false, nil
	case uint32(0xd912a59c):
		return &TextItalic{}, false, nil
	case uint32(0x8c39793f):
		return &HelpPromoDataObj{}, false, nil
	case uint32(0xc070d93e):
		return &PageBlockPreformatted{}, false, nil
	case uint32(0xd09e07b):
		return &InputPrivacyValueAllowContacts{}, false, nil
	case uint32(0x8f079643):
		return &ChannelAdminLogEventActionStopPoll{}, false, nil
	case uint32(0x9eddf188):
		return &InputMessagesFilterDocument{}, false, nil
	case uint32(0x938458c1):
		return &UserObj{}, false, nil
	case uint32(0xbea2f424):
		return &GlobalPrivacySettings{}, false, nil
	case uint32(0xc23727c9):
		return &AccountPasswordInputSettings{}, false, nil
	case uint32(0x3ded6320):
		return &MessageMediaEmpty{}, false, nil
	case uint32(0x4a992157):
		return &InputStickeredMediaPhoto{}, false, nil
	case uint32(0x8953ad37):
		return &InputChatPhotoObj{}, false, nil
	case uint32(0x13d6dd27):
		return &EncryptedChatDiscarded{}, false, nil
	case uint32(0x31f9590):
		return &PageBlockSlideshow{}, false, nil
	case uint32(0xb6abc341):
		return &MessagesFeaturedStickersObj{}, false, nil
	case uint32(0x12b9417b):
		return &UpdateUserPhone{}, false, nil
	case uint32(0x5d2f3aa9):
		return &UpdateBotPrecheckoutQuery{}, false, nil
	case uint32(0x6cef8ac7):
		return &MessageEntityBotCommand{}, false, nil
	case uint32(0x9cb070d7):
		return &MessageMediaDocument{}, false, nil
	case uint32(0xd02e7fd4):
		return &InputKeyboardButtonUrlAuth{}, false, nil
	case uint32(0xd27ff082):
		return &InputCheckPasswordSRPObj{}, false, nil
	case uint32(0x3504914f):
		return &UpdateDialogFilters{}, false, nil
	case uint32(0xb1c3caa7):
		return &ChannelAdminLogEventActionChangeStickerSet{}, false, nil
	case uint32(0xdd6a8f48):
		return &SendMessageGamePlayAction{}, false, nil
	case uint32(0x3502758c):
		return &ReplyKeyboardMarkup{}, false, nil
	case uint32(0x6f747657):
		return &PageCaption{}, false, nil
	case uint32(0xb6213cdf):
		return &ShippingOption{}, false, nil
	case uint32(0x4dba4501):
		return &AccountTakeout{}, false, nil
	case uint32(0xfb834291):
		return &TopPeerCategoryPeers{}, false, nil
	case uint32(0xc7f49b7):
		return &PrivacyValueDisallowUsers{}, false, nil
	case uint32(0x2979eeb2):
		return &LangPackStringDeleted{}, false, nil
	case uint32(0xbbc7515d):
		return &KeyboardButtonRequestPoll{}, false, nil
	case uint32(0xf0e3e596):
		return &MessagesDialogsNotModified{}, false, nil
	case uint32(0x15ebac1d):
		return &ChannelParticipantObj{}, false, nil
	case uint32(0x628cbc6f):
		return &SendMessageChooseContactAction{}, false, nil
	case uint32(0x95313b0c):
		return &UpdateUserPhoto{}, false, nil
	case uint32(0x11b58939):
		return &DocumentAttributeAnimated{}, false, nil
	case uint32(0xbb6ae88d):
		return &ChannelParticipantsContacts{}, false, nil
	case uint32(0x890c3d89):
		return &InputBotInlineMessageID{}, false, nil
	case uint32(0x4c4e743f):
		return &MessageEntityCashtag{}, false, nil
	case uint32(0xde266ef5):
		return &ContactsTopPeersNotModified{}, false, nil
	case uint32(0xe4e88011):
		return &PageBlockList{}, false, nil
	case uint32(0x69d3ab26):
		return &UserProfilePhotoObj{}, false, nil
	case uint32(0x39a51dfb):
		return &UpdateNewScheduledMessage{}, false, nil
	case uint32(0x6a4afc38):
		return &ChannelAdminLogEventActionChangeUsername{}, false, nil
	case uint32(0xe31c34d8):
		return &ChannelAdminLogEventActionParticipantInvite{}, false, nil
	case uint32(0x7fcb13a8):
		return &MessageActionChatEditPhoto{}, false, nil
	case uint32(0x26b5dde6):
		return &MessagesMessageEditData{}, false, nil
	case uint32(0xf041e250):
		return &ChatOnlines{}, false, nil
	case uint32(0xad2641f8):
		return &AccountPassword{}, false, nil
	case uint32(0x330b4067):
		return &Config{}, false, nil
	case uint32(0xeea8e46e):
		return &UploadCdnFileReuploadNeeded{}, false, nil
	case uint32(0x568a748):
		return &KeyboardButtonSwitchInline{}, false, nil
	case uint32(0x57e2f66c):
		return &InputMessagesFilterEmpty{}, false, nil
	case uint32(0x1837c364):
		return &InputEncryptedFileEmpty{}, false, nil
	case uint32(0x3f460fed):
		return &ChatParticipantsObj{}, false, nil
	case uint32(0x16812688):
		return &UpdateShortChatMessage{}, false, nil
	case uint32(0x6a7e7366):
		return &UpdatePeerSettings{}, false, nil
	case uint32(0x57e28221):
		return &AccountContentSettings{}, false, nil
	case uint32(0x9bed434d):
		return &InputWebDocument{}, false, nil
	case uint32(0x7c3c2609):
		return &MessageMediaGeoLive{}, false, nil
	case uint32(0x1427a5e1):
		return &ChannelParticipantsBanned{}, false, nil
	case uint32(0xa187d66f):
		return &SendMessageRecordVideoAction{}, false, nil
	case uint32(0xa575739d):
		return &EmojiURL{}, false, nil
	case uint32(0x6014f412):
		return &StatsGroupTopAdmin{}, false, nil
	case uint32(0x3407e51b):
		return &StickerSetMultiCovered{}, false, nil
	case uint32(0x289da732):
		return &ChannelForbidden{}, false, nil
	case uint32(0x179be863):
		return &InputPeerChat{}, false, nil
	case uint32(0xabe9affe):
		return &MessageActionBotAllowed{}, false, nil
	case uint32(0x2ec0533f):
		return &MessageMediaVenue{}, false, nil
	case uint32(0x74535f21):
		return &MessagesMessagesNotModified{}, false, nil
	case uint32(0xfa56ce36):
		return &EncryptedChatObj{}, false, nil
	case uint32(0x70b772a8):
		return &ContactsTopPeersObj{}, false, nil
	case uint32(0xbf4dea82):
		return &PageBlockTable{}, false, nil
	case uint32(0x1eb3758):
		return &HelpUserInfoObj{}, false, nil
	case uint32(0xd20b9f3c):
		return &ChatPhotoObj{}, false, nil
	case uint32(0x34636dd8):
		return &SecureValueErrorTranslationFiles{}, false, nil
	case uint32(0x808d15a4):
		return &ChannelParticipantCreator{}, false, nil
	case uint32(0x26ae0971):
		return &ChannelAdminLogEventActionToggleSignatures{}, false, nil
	case uint32(0xc21f497e):
		return &EncryptedFileEmpty{}, false, nil
	case uint32(0xaa0cd9e4):
		return &SendMessageUploadDocumentAction{}, false, nil
	case uint32(0xf888fa1a):
		return &PrivacyValueDisallowContacts{}, false, nil
	case uint32(0x7328bdb):
		return &ChatForbidden{}, false, nil
	case uint32(0xccbebbaf):
		return &ChannelParticipantAdmin{}, false, nil
	case uint32(0xfd8e711f):
		return &DraftMessageObj{}, false, nil
	case uint32(0xeeb46f27):
		return &StickerSet{}, false, nil
	case uint32(0x2f2f21bf):
		return &UpdateReadHistoryOutbox{}, false, nil
	case uint32(0xf0173fe9):
		return &ChannelsChannelParticipantsNotModified{}, false, nil
	case uint32(0xea02ec33):
		return SecureValueTypeTemporaryRegistration, true, nil
	case uint32(0x85e42301):
		return PhoneCallDiscardReasonMissed, true, nil
	case uint32(0xc3a12462):
		return BaseThemeClassic, true, nil
	case uint32(0x3d662b7b):
		return PrivacyKeyPhoneCall, true, nil
	case uint32(0x226ccefb):
		return AuthCodeTypeFlashCall, true, nil
	case uint32(0xfc36954e):
		return SecureValueTypeUtilityBill, true, nil
	case uint32(0xfbd81688):
		return BaseThemeDay, true, nil
	case uint32(0x148677e2):
		return TopPeerCategoryBotsInline, true, nil
	case uint32(0x9d2a81e3):
		return SecureValueTypePersonalDetails, true, nil
	case uint32(0x99a48f23):
		return SecureValueTypeInternalPassport, true, nil
	case uint32(0x89137c0d):
		return SecureValueTypeBankStatement, true, nil
	case uint32(0x7efe0e):
		return StorageFileJpeg, true, nil
	case uint32(0xaa963b05):
		return StorageFileUnknown, true, nil
	case uint32(0x4b09ebbc):
		return StorageFileMov, true, nil
	case uint32(0x57adc690):
		return PhoneCallDiscardReasonHangup, true, nil
	case uint32(0xfaf7e8c9):
		return PhoneCallDiscardReasonBusy, true, nil
	case uint32(0xa8406ca9):
		return TopPeerCategoryForwardUsers, true, nil
	case uint32(0x741cd3e3):
		return AuthCodeTypeCall, true, nil
	case uint32(0x528a0677):
		return StorageFileMp3, true, nil
	case uint32(0x6d5f77ee):
		return BaseThemeTinted, true, nil
	case uint32(0x72a3158c):
		return AuthCodeTypeSms, true, nil
	case uint32(0xb7b31ea8):
		return BaseThemeNight, true, nil
	case uint32(0x39491cc8):
		return PrivacyKeyPhoneP2P, true, nil
	case uint32(0x96151fed):
		return PrivacyKeyProfilePhoto, true, nil
	case uint32(0x637b7ed):
		return TopPeerCategoryCorrespondents, true, nil
	case uint32(0xa0d0744b):
		return SecureValueTypeIdentityCard, true, nil
	case uint32(0x69ec56a3):
		return PrivacyKeyForwards, true, nil
	case uint32(0xfbeec0f0):
		return TopPeerCategoryForwardChats, true, nil
	case uint32(0xbdfb0426):
		return InputPrivacyKeyChatInvite, true, nil
	case uint32(0x40bc6f52):
		return StorageFilePartial, true, nil
	case uint32(0xcbe31e26):
		return SecureValueTypeAddress, true, nil
	case uint32(0x99e3806a):
		return SecureValueTypePassportRegistration, true, nil
	case uint32(0xb3cea0e4):
		return StorageFileMp4, true, nil
	case uint32(0x5719bacc):
		return InputPrivacyKeyProfilePhoto, true, nil
	case uint32(0x8e3ca7ee):
		return SecureValueTypeEmail, true, nil
	case uint32(0xa4f63c0):
		return StorageFilePng, true, nil
	case uint32(0xfabadc5f):
		return InputPrivacyKeyPhoneCall, true, nil
	case uint32(0x1081464c):
		return StorageFileWebp, true, nil
	case uint32(0x500e6dfa):
		return PrivacyKeyChatInvite, true, nil
	case uint32(0xbd17a14a):
		return TopPeerCategoryGroups, true, nil
	case uint32(0x6e425c4):
		return SecureValueTypeDriverLicense, true, nil
	case uint32(0x8b883488):
		return SecureValueTypeRentalAgreement, true, nil
	case uint32(0xcae1aadf):
		return StorageFileGif, true, nil
	case uint32(0xd1219bdd):
		return InputPrivacyKeyAddedByPhone, true, nil
	case uint32(0xdb9e70d2):
		return InputPrivacyKeyPhoneP2P, true, nil
	case uint32(0xe095c1a0):
		return PhoneCallDiscardReasonDisconnect, true, nil
	case uint32(0x5b11125a):
		return BaseThemeArctic, true, nil
	case uint32(0xbc2eab30):
		return PrivacyKeyStatusTimestamp, true, nil
	case uint32(0xd19ae46d):
		return PrivacyKeyPhoneNumber, true, nil
	case uint32(0x4f96cb18):
		return InputPrivacyKeyStatusTimestamp, true, nil
	case uint32(0x3dac6a00):
		return SecureValueTypePassport, true, nil
	case uint32(0xae1e508d):
		return StorageFilePdf, true, nil
	case uint32(0x42ffd42b):
		return PrivacyKeyAddedByPhone, true, nil
	case uint32(0x352dafa):
		return InputPrivacyKeyPhoneNumber, true, nil
	case uint32(0xb320aadb):
		return SecureValueTypePhone, true, nil
	case uint32(0xab661b5b):
		return TopPeerCategoryBotsPM, true, nil
	case uint32(0x161d9628):
		return TopPeerCategoryChannels, true, nil
	case uint32(0x1e76a78c):
		return TopPeerCategoryPhoneCalls, true, nil
	case uint32(0xa4dd4c08):
		return InputPrivacyKeyForwards, true, nil
	default:
		panic("unknown constructor id: " + fmt.Sprintf("%#v", constructorID))
	}
}
