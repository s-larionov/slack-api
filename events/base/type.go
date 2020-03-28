package base

const (
	TypeAppHomeOpened         Type = "app_home_opened"
	TypeAppMention            Type = "app_mention"
	TypeAppRequested          Type = "app_requested"
	TypeAppUninstalled        Type = "app_uninstalled"
	TypeCallRejected          Type = "call_rejected"
	TypeChannelArchive        Type = "channel_archive"
	TypeChannelCreated        Type = "channel_created"
	TypeChannelDeleted        Type = "channel_deleted"
	TypeChannelHistoryChanged Type = "channel_history_changed"
	TypeChannelLeft           Type = "channel_left"
	TypeChannelShared         Type = "channel_shared"
	TypeChannelUnarchive      Type = "channel_unarchive"
	TypeChannelUnshared       Type = "channel_unshared"
	TypeDndUpdated            Type = "dnd_updated"
	TypeDndUpdatedUser        Type = "dnd_updated_user"
	TypeEmailDomainChanged    Type = "email_domain_changed"
	TypeEmojiChanged          Type = "emoji_changed"
	TypeFileChange            Type = "file_change"
	TypeFileCreated           Type = "file_created"
	TypeFileDeleted           Type = "file_deleted"
	TypeFilePublic            Type = "file_public"
	TypeFileShared            Type = "file_shared"
	TypeFileUnshared          Type = "file_unshared"
	TypeGridMigrationFinished Type = "grid_migration_finished"
	TypeGridMigrationStarted  Type = "grid_migration_started"
	TypeGroupArchive          Type = "group_archive"
	TypeGroupClose            Type = "group_close"
	TypeGroupDeleted          Type = "group_deleted"
	TypeGroupHistoryChanged   Type = "group_history_changed"
	TypeGroupLeft             Type = "group_left"
	TypeGroupOpen             Type = "group_open"
	TypeGroupRename           Type = "group_rename"
	TypeGroupUnarchive        Type = "group_unarchive"
	TypeIMClose               Type = "im_close"
	TypeIMCreated             Type = "im_created"
	TypeIMHistoryChanged      Type = "im_history_changed"
	TypeIMOpen                Type = "im_open"
	TypeInviteRequested       Type = "invite_requested"
	TypeLinkShared            Type = "link_shared"
	TypeMemberJoinedChannel   Type = "member_joined_channel"
	TypeMemberLeftChannel     Type = "member_left_channel"
	TypeMessage               Type = "message"
	TypePinAdded              Type = "pin_added"
	TypePinRemoved            Type = "pin_removed"
	TypeReactionAdded         Type = "reaction_added"
	TypeReactionRemoved       Type = "reaction_removed"
	TypeStarAdded             Type = "star_added"
	TypeStarRemoved           Type = "star_removed"
	TypeSubteamCreated        Type = "subteam_created"
	TypeSubteamMembersChanged Type = "subteam_members_changed"
	TypeSubteamSelfAdded      Type = "subteam_self_added"
	TypeSubteamSelfRemoved    Type = "subteam_self_removed"
	TypeSubteamUpdated        Type = "subteam_updated"
	TypeTeamDomainChange      Type = "team_domain_change"
	TypeTeamJoin              Type = "team_join"
	TypeTeamRename            Type = "team_rename"
	TypeTokensRevoked         Type = "tokens_revoked"
)

type Type string
